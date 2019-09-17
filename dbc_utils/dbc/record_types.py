from dbc import dbc
from dbc.record_fields import Int, ID, String, LocalizedString, Float


class ChrClasses(dbc.Record):
    @classmethod
    def GoTypeName(cls):
        return "Class"

    @classmethod
    def Fields(cls):
        return [
            Int(name='id'),
            Int(default=1),
            Int(name='primary_stat'),
            Int(name='power_type'),
            String(name='pet_type'),
            LocalizedString(name='name', indexed=True),
            String(default=lambda r: r.name.upper()),
            Int(default=0),
            Int(default=0),
        ]


class ChrRaces(dbc.Record):
    @classmethod
    def GoTypeName(cls):
        return 'Race'

    @classmethod
    def Fields(cls):
        return [
            Int(name='id'),
            Int(name='flags'),
            Int(name='faction_id'),  # faction_id
            Int(name='unk'),  # unk
            Int(name='male_display_id'),
            Int(name='female_display_id'),
            String(name='client_prefix'),
            Float(name='mount_scale'),
            Int(name='base_language'),
            Int(name='creature_type'),
            Int(name='login_effect_spell_id'),
            Int(name='combat_stun_spell_id'),
            Int(name='res_sickness_spell_id'),
            Int(name='splash_sound_id'),
            Int(name='starting_taxi_nodes'),
            String(name='client_file_string'),
            Int(name='cinematic_sequence_id'),
            LocalizedString(name='name', indexed=True),
            String(name='male_feature_name'),
            String(name='female_feature_name'),
            String(name='hair_customization_name'),
        ]


class ChrStartingStats(dbc.Record):
    @classmethod
    def GoTypeName(cls):
        return 'StartingStats'

    @classmethod
    def Fields(cls):
        return [
            ID(),
            String(name='class', indexed=True,
                   foreign_key=(ChrClasses, 'name')),
            String(name='race', indexed=True, foreign_key=(ChrRaces, 'name')),
            Int(name='strength'),
        ]


class ChrStartingLocations(dbc.Record):
    @classmethod
    def GoTypeName(cls):
        return 'StartingLocations'

    @classmethod
    def Fields(cls):
        return [
            ID(),
            String(name='race', indexed=True, foreign_key=(ChrRaces, 'name')),
            Int(name='map'),
            Int(name='zone'),
            Float(name='x'),
            Float(name='y'),
            Float(name='z'),
            Float(name='o'),
        ]
