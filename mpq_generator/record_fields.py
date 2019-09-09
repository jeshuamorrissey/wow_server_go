from typing import Any, Iterable, Text, Optional


class Field:
    def __init__(self, name: Text = None, default: Optional[Any] = None):
        self.name = name
        self.default = default

    def Value(self, string_block, record):
        raise NotImplementedError()

    def _GetValue(self, record):
        if self.name is not None and hasattr(record, self.name):
            return getattr(record, self.name)
        if self.default is not None:
            if callable(self.default):
                return self.default(record)
            return self.default
        raise ValueError('Field should have a name or a default value.')


class Int(Field):
    def Value(self, string_block, record):
        return [super(Int, self)._GetValue(record)]

    @classmethod
    def Format(cls) -> Text:
        return 'I'

    @classmethod
    def Load(cls, string_block, args: Iterable[Any]):
        return next(args)


class String(Field):
    def Value(self, string_block, record):
        return [string_block.OffsetFor(self._GetValue(record))]

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
            0, 0, 0, 0, 0, 0, 0,  # otther locales, unused
            0,  # flags, unused
        ]

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
