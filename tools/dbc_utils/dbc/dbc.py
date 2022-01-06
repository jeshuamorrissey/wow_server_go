import os
import struct
import json

import jinja2

from typing import Any, Text, Type, Iterable


class DBCError(Exception):
    pass


class StringBlock(dict):
    """StringBlock represents a block of strings in a Binary DBC file.

    StringBlocks are made up of 2 dictionaries:
      - One mapping offset --> string (as read from the binary).
      - One mapping string --> offset (for converting back to binary).

    StringBlocks will always have an empty string as the first entry.
    """

    def _MaybeAdd(self, s: str):
        """Maybe add the given string to this block."""
        if s not in self.inverse:
            self[self.next_offset] = s
            self.inverse[s] = self.next_offset
            self.next_offset += len(s) + 1

    def __init__(self, strings: Iterable[str]):
        """Create a new string block from a given list of strings.

        Args:
            strings: A list of strings to add to the block.
        """
        # Setup default state.
        self[0] = ''
        self.inverse = {'': 0}
        self.next_offset = 1

        for s in strings:
            self._MaybeAdd(s.decode())

    def ToBinary(self) -> bytes:
        """Convert this StringBlock to binary.

        Returns:
            An null-terminated list of strings as bytes.
        """
        data = b''
        for _, s in sorted(self.items()):
            data += s.encode() + b'\x00'
        return data

    def OffsetFor(self, s: str) -> int:
        """Get the offset for a given string, adding it if it doesn't exist.

        Args:
            The string to get the offset for.

        Returns:
            The offset to a given string.
        """
        self._MaybeAdd(s)
        return self.inverse[s]


class Record():
    """Records are a representation of a single row within a DBC file.

    Each record corresponds to some binary data, a single JSON object
    or a single Golang struct.
    """

    def __init__(self, id: int):
        self._id = id

    def __str__(self):
        return str(self.__dict__)

    def __repr__(self):
        return str(self)

    def __eq__(self, other: 'Record'):
        return str(self) == str(other)

    def ToBinary(self, string_block: str):
        """Convert the record back to it's binary form."""
        fmt = ''
        data = []
        for field in self.Fields():
            fmt += field.Format()
            data.extend(field.Value(string_block, self))

        return struct.pack(fmt, *data)

    def IndexedFields(self):
        return [f for f in self.Fields() if f.indexed]

    def GoValue(self, field_name: Text) -> Text:
        """Get the Golang representation of the value of some field in this record.

        Args:
            field_name: The name of the field to get the value of.

        Returns:
            A stringified version of the value. For integers/floats, it is just
            the value. For strings, double quotes will be put around it.
        """
        val = getattr(self, field_name)
        if isinstance(val, str):
            val = '"{}"'.format(val)
        return val

    @classmethod
    def Fields(cls):
        raise NotImplementedError()

    @classmethod
    def GoTypeName(cls) -> Text:
        """Get the Golang name for this record type.

        Returns:
            A text version of the name that should be used to refer to this class in Go.
        """
        return cls.__name__

    @classmethod
    def FromBinary(cls, id: int, string_block: StringBlock, args: Iterable[Any]) -> 'Record':
        """Load a new record object from a set of binary arguments."""
        record = cls(id)
        for field in cls.Fields():
            val = field.Load(string_block, args)
            if field.name:
                setattr(record, field.name, val)

        return record

    @classmethod
    def Format(cls) -> Text:
        """Return the struct format string for the whole field."""
        return ''.join((f.Format() for f in cls.Fields()))


class Table():
    """Tables are a representation of a single DBC file.

    Each table is made up of a series of records. Tables typically represent a
    whole file.
    """
    HEADER_SIZE = 20
    MAGIC = 1128416343

    def __init__(self, records: Iterable[Record], record_type: Type):
        self.records = records
        self.record_type = record_type

    def __str__(self):
        return str(self.records)

    def __repr__(self):
        return str(self)

    def ToJSON(self) -> bytes:
        """Convert the table to JSON bytes."""
        result = []
        for record in self.records:
            result.append(
                {k: v for k, v in record.__dict__.items() if k != '_id'})
        return json.dumps(result)

    def ToBinary(self) -> bytes:
        """Convert the table to DBC bytes."""
        string_block = StringBlock([])
        record_data = b''
        for record in self.records:
            record_data += record.ToBinary(string_block)

        # Build the header.
        string_block_data = string_block.ToBinary()
        header_data = struct.pack('IIIII',
                                  Table.MAGIC,
                                  len(self.records),
                                  len(self.records[0].Format()),
                                  struct.calcsize(self.records[0].Format()),
                                  len(string_block_data))

        return header_data + record_data + string_block_data

    def ToGolangPart(self) -> (str, bytes):
        """Convert this table to a Golang partial file.

        This will include the struct defintions and variable definitions and a init function which
        can be called to setup the data.

        Returns:
            A tuple of (file data, init function name).
        """
        dbc_name = self.record_type.__name__
        init_function_name = f'{dbc_name}__init'
        args = {
            'dbc_name': dbc_name,
            'type_name': self.record_type.GoTypeName(),
            'fields': [f for f in self.record_type.Fields() if f.name],
            'index_fields': [f for f in self.record_type.Fields() if f.indexed],
            'num_indexed_fields': sum(1 for f in self.record_type.Fields() if f.indexed),
            'records': self.records,
            'init_function_name': init_function_name,
        }

        # Make a go type for the index.
        index_map_type_list = []
        for f in args['index_fields']:
            index_map_type_list.append('map[{}]'.format(f.GoType()))
        index_map_type_list.append(
            '*{}'.format(self.record_type.GoTypeName()))
        args['index_map_type'] = ''.join(index_map_type_list)
        args['index_map_type_list'] = index_map_type_list

        template_env = jinja2.Environment(
            loader=jinja2.FileSystemLoader(searchpath=os.path.dirname(os.path.realpath(__file__))))
        template_env.trim_blocks = True
        template_env.lstrip_blocks = True
        return template_env.get_template('dbc_golang_part.go.jinja').render(**args), init_function_name

    @classmethod
    def FromJSON(cls, data: bytes, record_type: Type) -> 'Table':
        """Load the table from some JSON data."""
        records = []
        for i, record_data in enumerate(json.loads(data)):
            record = record_type(id=i)
            for k, v in record_data.items():
                setattr(record, k, v)

            records.append(record)

        return Table(records, record_type)

    @classmethod
    def FromBinary(cls, data: bytes, record_type: Type) -> 'Table':
        """Load the table from some binary data."""
        magic, n_records, n_fields, record_size, string_block_size = struct.unpack(
            'IIIII', data[:Table.HEADER_SIZE])
        if magic != Table.MAGIC:
            raise DBCError(
                "Malformed magic value {} (expected {}).".format(magic, Table.MAGIC))

        if n_fields != len(record_type.Format()):
            raise DBCError(
                "DBC has {} fields, but record type {} has {} fields.".format(n_fields, record_type, len(record_type.Format())))

        if record_size != struct.calcsize(record_type.Format()):
            raise DBCError(
                "DBC records are {}b, but record type is {}b.".format(record_size, struct.calcsize(record_type.Format())))

        records_size = n_records * record_size
        record_data = data[Table.HEADER_SIZE:Table.HEADER_SIZE + records_size]
        string_block_data = data[Table.HEADER_SIZE + records_size:]

        if len(string_block_data) != string_block_size:
            raise DBCError("Malformed file: sizes did not match data size.")

        string_block = StringBlock(string_block_data.split(b'\x00')[:-1])

        records = []
        for i, record_args in enumerate(struct.iter_unpack(record_type.Format(), record_data)):
            record_args_iter = iter(record_args)
            records.append(record_type.FromBinary(
                i, string_block, record_args_iter))

            record_args_remaining = list(record_args_iter)
            if len(record_args_remaining) != 0:
                raise DBCError(
                    "Not all fields consumed: {} remaining".format(len(record_args_remaining)))

        return Table(records, record_type)
