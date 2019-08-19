import gzip
import json
import sqlite3

CREATURE_TEMPLATE_SQL = 'D:\\Users\\Jeshua\\Code\\mangoszero\\database\\World\\Setup\\FullDB\\creature_template.sql'
CREATE_TABLES_SQL = '''
CREATE TABLE `creature_template` (
    `Entry` mediumint(8),
    `Name` char(100),
    `SubName` char(100),
    `MinLevel` tinyint(3),
    `MaxLevel` tinyint(3),
    `ModelId1` mediumint(8),
    `ModelId2` mediumint(8),
    `ModelId3` mediumint(8),
    `ModelId4` mediumint(8),
    `FactionAlliance` smallint(5),
    `FactionHorde` smallint(5),
    `Scale` float,
    `Family` tinyint(4),
    `CreatureType` tinyint(3),
    `InhabitType` tinyint(3),
    `RegenerateStats` tinyint(3),
    `RacialLeader` tinyint(3),
    `NpcFlags` int(10),
    `UnitFlags` int(10),
    `DynamicFlags` int(10),
    `ExtraFlags` int(10),
    `CreatureTypeFlags` int(10),
    `SpeedWalk` float,
    `SpeedRun` float,
    `UnitClass` tinyint(3),
    `Rank` tinyint(3),
    `HealthMultiplier` float,
    `PowerMultiplier` float,
    `DamageMultiplier` float,
    `DamageVariance` float,
    `ArmorMultiplier` float,
    `ExperienceMultiplier` float,
    `MinLevelHealth` int(10),
    `MaxLevelHealth` int(10),
    `MinLevelMana` int(10),
    `MaxLevelMana` int(10),
    `MinMeleeDmg` float,
    `MaxMeleeDmg` float,
    `MinRangedDmg` float,
    `MaxRangedDmg` float,
    `Armor` mediumint(8),
    `MeleeAttackPower` int(10),
    `RangedAttackPower` smallint(5),
    `MeleeBaseAttackTime` int(10),
    `RangedBaseAttackTime` int(10),
    `DamageSchool` tinyint(4),
    `MinLootGold` mediumint(8),
    `MaxLootGold` mediumint(8),
    `LootId` mediumint(8),
    `PickpocketLootId` mediumint(8),
    `SkinningLootId` mediumint(8),
    `KillCredit1` int(11),
    `KillCredit2` int(11),
    `MechanicImmuneMask` int(10),
    `SchoolImmuneMask` int(10),
    `ResistanceHoly` smallint(5),
    `ResistanceFire` smallint(5),
    `ResistanceNature` smallint(5),
    `ResistanceFrost` smallint(5),
    `ResistanceShadow` smallint(5),
    `ResistanceArcane` smallint(5),
    `PetSpellDataId` mediumint(8),
    `MovementType` tinyint(3),
    `TrainerType` tinyint(4),
    `TrainerSpell` mediumint(8),
    `TrainerClass` tinyint(3),
    `TrainerRace` tinyint(3),
    `TrainerTemplateId` mediumint(8),
    `VendorTemplateId` mediumint(8),
    `GossipMenuId` mediumint(8),
    `EquipmentTemplateId` mediumint(8),
    `Civilian` tinyint(3),
    `AIName` char(64)
);

CREATE TABLE `creature_model_info` (
  `modelid` mediumint(8),
  `bounding_radius` float,
  `combat_reach` float,
  `gender` tinyint(3),
  `modelid_other_gender` mediumint(8),
  `modelid_other_team` mediumint(8)
);
'''


conn = sqlite3.connect(':memory:')
conn.row_factory = sqlite3.Row
conn.execute(CREATE_TABLES_SQL)
conn.executescript(open(CREATURE_TEMPLATE_SQL).read())

items = {}


def LoadRowIntoObject(row):
    result = {
        'Entry': row['Entry'],
        'Name': row['Name'],
        'SubName': row['SubName'],
        'MinLevel': row['MinLevel'],
        'MaxLevel': row['MaxLevel'],
        'FactionAlliance': row['FactionAlliance'],
        'FactionHorde': row['FactionHorde'],
        'Scale': row['Scale'],
        'Family': row['Family'],
        'CreatureType': row['CreatureType'],
        'InhabitType': row['InhabitType'],
        'RegenerateStats': row['RegenerateStats'],
        'RacialLeader': row['RacialLeader'],
        'DynamicFlags': row['DynamicFlags'],
        'SpeedWalk': row['SpeedWalk'],
        'SpeedRun': row['SpeedRun'],
        'UnitClass': row['UnitClass'],
        'Rank': row['Rank'],
        'HealthMultiplier': row['HealthMultiplier'],
        'PowerMultiplier': row['PowerMultiplier'],
        'DamageMultiplier': row['DamageMultiplier'],
        'DamageVariance': row['DamageVariance'],
        'ArmorMultiplier': row['ArmorMultiplier'],
        'ExperienceMultiplier': row['ExperienceMultiplier'],
        'MinLevelHealth': row['MinLevelHealth'],
        'MaxLevelHealth': row['MaxLevelHealth'],
        'MinLevelMana': row['MinLevelMana'],
        'MaxLevelMana': row['MaxLevelMana'],
        'MinMeleeDmg': row['MinMeleeDmg'],
        'MaxMeleeDmg': row['MaxMeleeDmg'],
        'MinRangedDmg': row['MinRangedDmg'],
        'MaxRangedDmg': row['MaxRangedDmg'],
        'Armor': row['Armor'],
        'MeleeAttackPower': row['MeleeAttackPower'],
        'RangedAttackPower': row['RangedAttackPower'],
        'MeleeBaseAttackTime': row['MeleeBaseAttackTime'],
        'RangedBaseAttackTime': row['RangedBaseAttackTime'],
        'DamageSchool': row['DamageSchool'],
        'MinLootGold': row['MinLootGold'],
        'MaxLootGold': row['MaxLootGold'],
        'LootID': row['LootId'],
        'PickpocketLootID': row['PickpocketLootId'],
        'SkinningLootID': row['SkinningLootId'],
        'KillCredit1': row['KillCredit1'],
        'KillCredit2': row['KillCredit2'],
        'MechanicImmuneMask': row['MechanicImmuneMask'],
        'SchoolImmuneMask': row['SchoolImmuneMask'],
        'ResistanceHoly': row['ResistanceHoly'],
        'ResistanceFire': row['ResistanceFire'],
        'ResistanceNature': row['ResistanceNature'],
        'ResistanceFrost': row['ResistanceFrost'],
        'ResistanceShadow': row['ResistanceShadow'],
        'ResistanceArcane': row['ResistanceArcane'],
        'PetSpellDataID': row['PetSpellDataId'],
        'MovementType': row['MovementType'],
        'TrainerType': row['TrainerType'],
        'TrainerSpell': row['TrainerSpell'],
        'TrainerClass': row['TrainerClass'],
        'TrainerRace': row['TrainerRace'],
        'TrainerTemplateID': row['TrainerTemplateId'],
        'VendorTemplateID': row['VendorTemplateId'],
        'GossipMenuID': row['GossipMenuId'],
        'EquipmentTemplateID': row['EquipmentTemplateId'],
        'Civilian': row['Civilian'],
        # 'AIName': row['AIName'],
    }

    result['Models'] = {}
    for i in range(4):
        model_id = row['ModelId{}'.format(i+1)]
        if model_id != 0:
            model_info = conn.execute(
                "SELECT * FROM create_model_info WHERE modelid = {}".format(model_id))[0]
            result['Models'].append({
                'ID': model_id,
                'BoundingRadius': model_info['bounding_radius'],
                'CombadReach': model_info['combat_reach'],
                'Gender': model_info['gender'],
            })

    flags = row['NpcFlags']
    result['HasGossip'] = bool(flags & 0x00000001)
    result['IsQuestgiver'] = bool(flags & 0x00000002)
    result['IsVendor'] = bool(flags & 0x00000004)
    result['IsFlightmaster'] = bool(flags & 0x00000008)
    result['IsTrainer'] = bool(flags & 0x00000010)
    result['IsSpirithealer'] = bool(flags & 0x00000020)
    result['IsSpiritguide'] = bool(flags & 0x00000040)
    result['IsInnkeeper'] = bool(flags & 0x00000080)
    result['IsBanker'] = bool(flags & 0x00000100)
    result['IsPetitioner'] = bool(flags & 0x00000200)
    result['IsTabarddesigner'] = bool(flags & 0x00000400)
    result['IsBattlemaster'] = bool(flags & 0x00000800)
    result['IsAuctioneer'] = bool(flags & 0x00001000)
    result['IsStablemaster'] = bool(flags & 0x00002000)
    result['CanRepair'] = bool(flags & 0x00004000)

    flags = row['ExtraFlags']
    result['IsInstanceBound'] = bool(flags & 0x00000001)
    result['NoAggro'] = bool(flags & 0x00000002)
    result['NoParry'] = bool(flags & 0x00000004)
    result['NoParryHasten'] = bool(flags & 0x00000008)
    result['NoBlock'] = bool(flags & 0x00000010)
    result['NoCrush'] = bool(flags & 0x00000020)
    result['NoXPAtKill'] = bool(flags & 0x00000040)
    result['IsInvisible'] = bool(flags & 0x00000080)
    result['IsNotTauntable'] = bool(flags & 0x00000100)
    result['HasAggroZone'] = bool(flags & 0x00000200)
    result['IsGuard'] = bool(flags & 0x00000400)
    result['NoCallAssist'] = bool(flags & 0x00000800)
    result['IsActive'] = bool(flags & 0x00001000)
    result['IsMMapForceEnable'] = bool(flags & 0x00002000)
    result['IsMMapForceDisable'] = bool(flags & 0x00004000)
    result['WalksInWater'] = bool(flags & 0x00008000)
    result['HasNoSwimAnimation'] = bool(flags & 0x00010000)

    flags = row['CreatureTypeFlags']
    result['IsTameable'] = bool(flags & 0x00000001)
    result['IsGhostVisible'] = bool(flags & 0x00000002)
    result['IsHerbLoot'] = bool(flags & 0x00000100)
    result['IsMiningLoot'] = bool(flags & 0x00000200)
    result['CanAssist'] = bool(flags & 0x00001000)
    result['IsEngineerLoot'] = bool(flags & 0x00008000)

    return result


result = {}
cursor = conn.execute("SELECT * FROM creature_template")
for row in cursor:
    result[row['entry']] = LoadRowIntoObject(row)


with gzip.open('units.json.gz', 'w') as f:
    f.write(json.dumps(result).encode('utf-8'))
