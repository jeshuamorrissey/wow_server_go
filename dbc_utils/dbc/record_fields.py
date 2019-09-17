from typing import Any, Iterable, Text, Optional, Type, Tuple


class Field:
    def __init__(self, name: Text = None, default: Optional[Any] = None, indexed: bool = False, foreign_key: Optional[Tuple[Type, Text]] = None):
        self.name = name
        self.default = default
        self.indexed = indexed

        self.foreign_key_field = None
        self.foreign_key_type = None
        if foreign_key:
            record_type, field_name = foreign_key
            for field in record_type.Fields():
                if field.name == field_name:
                    self.foreign_key_field = field
                    self.foreign_key_type = record_type.GoTypeName()

    def Value(self, string_block, record):
        raise NotImplementedError()

    def _GetValue(self, record):
        if self.name is not None and hasattr(record, self.name):
            return getattr(record, self.name)
        if self.default is not None:
            if callable(self.default):
                return self.default(record)
            return self.default
        raise ValueError('Field {} in {} should have a name or a default value.'.format(
            self, type(record)))

    def GoName(self) -> Text:
        if not self.name:
            return None

        name = ''.join(x.title() for x in self.name.split('_'))
        name = name.replace('Id', 'ID')
        name = name.replace('_ID', 'ID')
        return name

    def GoType(self) -> Text:
        if self.foreign_key_type is not None:
            return '*{}'.format(self.foreign_key_type)

        return None

    @classmethod
    def Format(cls) -> Text:
        raise NotImplementedError()

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        raise NotImplementedError()


class Int(Field):
    def Value(self, string_block, record):
        return [super(Int, self)._GetValue(record)]

    def GoType(self) -> Text:
        go_type = super(Int, self).GoType()
        if go_type:
            return go_type

        return 'int'

    @classmethod
    def Format(cls) -> Text:
        return 'I'

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        return next(args)


class ID(Int):
    def __init__(self, name: Text = None, default: Optional[Any] = None, indexed: bool = False):
        super(ID, self).__init__(name='_id', default=default, indexed=indexed)


class Float(Field):
    def GoType(self) -> Text:
        go_type = super(Float, self).GoType()
        if go_type:
            return go_type

        return 'float32'

    def Value(self, string_block, record):
        return [super(Float, self)._GetValue(record)]

    @classmethod
    def Format(cls) -> Text:
        return 'f'

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        return next(args)


class String(Field):
    def Value(self, string_block, record):
        return [string_block.OffsetFor(self._GetValue(record))]

    def GoType(self) -> Text:
        go_type = super(String, self).GoType()
        if go_type:
            return go_type

        return 'string'

    @classmethod
    def Format(cls) -> Text:
        return 'I'

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        return string_block[next(args)]


class LocalizedString(Field):
    def Value(self, string_block, record):
        return [
            string_block.OffsetFor(self._GetValue(record)),  # enUS
            0, 0, 0, 0, 0, 0, 0,  # other locales, unused
            0,  # flags, unused
        ]

    def GoType(self) -> Text:
        go_type = super(LocalizedString, self).GoType()
        if go_type:
            return go_type

        return 'string'

    @classmethod
    def Format(cls) -> Text:
        return 'IIIIIIIII'

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        locales = dict(
            enUS=next(args),
            koKR=next(args),  # unused
            frFR=next(args),  # unused
            deDE=next(args),  # unused
            enCN=next(args),  # unused
            enTW=next(args),  # unused
            esES=next(args),  # unused
            esMX=next(args),  # unused
        )

        next(args)  # flags, unused

        return string_block[locales['enUS']]
