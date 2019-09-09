from record_fields import Int, String, LocalizedString

import dbc


class ChrClasses(dbc.Record):
    @classmethod
    def Fields(cls):
        return [
            Int(name='id'),
            Int(default=1),
            Int(name='primary_stat'),
            Int(name='power_type'),
            String(name='pet_type'),
            LocalizedString(name='name'),
            String(default=lambda r: r.name.upper()),
            Int(default=0),
            Int(default=0),
        ]
