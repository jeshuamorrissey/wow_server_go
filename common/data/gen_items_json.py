import gzip
import json
import sqlite3

ITEM_TEMPLATE_SQL = 'D:\\Users\\Jeshua\\Code\\mangoszero\\database\\World\\Setup\\FullDB\\item_template.sql'
CREATE_TABLES_SQL = '''
CREATE TABLE `item_template` (
    `entry` mediumint(8),
    `class` tinyint(3),
    `subclass` tinyint(3),
    `name` varchar(255),
    `displayid` mediumint(8),
    `Quality` tinyint(3),
    `Flags` int(10),
    `BuyCount` tinyint(3),
    `BuyPrice` int(10),
    `SellPrice` int(10),
    `InventoryType` tinyint(3),
    `AllowableClass` mediumint(9),
    `AllowableRace` mediumint(9),
    `ItemLevel` tinyint(3),
    `RequiredLevel` tinyint(3),
    `RequiredSkill` smallint(5),
    `RequiredSkillRank` smallint(5),
    `requiredspell` mediumint(8),
    `requiredhonorrank` mediumint(8),
    `RequiredCityRank` mediumint(8),
    `RequiredReputationFaction` smallint(5),
    `RequiredReputationRank` smallint(5),
    `maxcount` smallint(5),
    `stackable` smallint(5),
    `ContainerSlots` tinyint(3),
    `stat_type1` tinyint(3),
    `stat_value1` smallint(6),
    `stat_type2` tinyint(3),
    `stat_value2` smallint(6),
    `stat_type3` tinyint(3),
    `stat_value3` smallint(6),
    `stat_type4` tinyint(3),
    `stat_value4` smallint(6),
    `stat_type5` tinyint(3),
    `stat_value5` smallint(6),
    `stat_type6` tinyint(3),
    `stat_value6` smallint(6),
    `stat_type7` tinyint(3),
    `stat_value7` smallint(6),
    `stat_type8` tinyint(3),
    `stat_value8` smallint(6),
    `stat_type9` tinyint(3),
    `stat_value9` smallint(6),
    `stat_type10` tinyint(3),
    `stat_value10` smallint(6),
    `dmg_min1` float,
    `dmg_max1` float,
    `dmg_type1` tinyint(3),
    `dmg_min2` float,
    `dmg_max2` float,
    `dmg_type2` tinyint(3),
    `dmg_min3` float,
    `dmg_max3` float,
    `dmg_type3` tinyint(3),
    `dmg_min4` float,
    `dmg_max4` float,
    `dmg_type4` tinyint(3),
    `dmg_min5` float,
    `dmg_max5` float,
    `dmg_type5` tinyint(3),
    `armor` smallint(5),
    `holy_res` tinyint(3),
    `fire_res` tinyint(3),
    `nature_res` tinyint(3),
    `frost_res` tinyint(3),
    `shadow_res` tinyint(3),
    `arcane_res` tinyint(3),
    `delay` smallint(5),
    `ammo_type` tinyint(3),
    `RangedModRange` float,
    `spellid_1` mediumint(8),
    `spelltrigger_1` tinyint(3),
    `spellcharges_1` tinyint(4),
    `spellppmRate_1` float,
    `spellcooldown_1` int(11),
    `spellcategory_1` smallint(5),
    `spellcategorycooldown_1` int(11),
    `spellid_2` mediumint(8),
    `spelltrigger_2` tinyint(3),
    `spellcharges_2` tinyint(4),
    `spellppmRate_2` float,
    `spellcooldown_2` int(11),
    `spellcategory_2` smallint(5),
    `spellcategorycooldown_2` int(11),
    `spellid_3` mediumint(8),
    `spelltrigger_3` tinyint(3),
    `spellcharges_3` tinyint(4),
    `spellppmRate_3` float,
    `spellcooldown_3` int(11),
    `spellcategory_3` smallint(5),
    `spellcategorycooldown_3` int(11),
    `spellid_4` mediumint(8),
    `spelltrigger_4` tinyint(3),
    `spellcharges_4` tinyint(4),
    `spellppmRate_4` float,
    `spellcooldown_4` int(11),
    `spellcategory_4` smallint(5),
    `spellcategorycooldown_4` int(11),
    `spellid_5` mediumint(8),
    `spelltrigger_5` tinyint(3),
    `spellcharges_5` tinyint(4),
    `spellppmRate_5` float,
    `spellcooldown_5` int(11),
    `spellcategory_5` smallint(5),
    `spellcategorycooldown_5` int(11),
    `bonding` tinyint(3),
    `description` varchar(255),
    `PageText` mediumint(8),
    `LanguageID` tinyint(3),
    `PageMaterial` tinyint(3),
    `startquest` mediumint(8),
    `lockid` mediumint(8),
    `Material` tinyint(4),
    `sheath` tinyint(3),
    `RandomProperty` mediumint(8),
    `block` mediumint(8),
    `itemset` mediumint(8),
    `MaxDurability` smallint(5),
    `area` mediumint(8),
    `Map` smallint(6),
    `BagFamily` mediumint(9),
    `DisenchantID` mediumint(8),
    `FoodType` tinyint(3),
    `minMoneyLoot` int(10),
    `maxMoneyLoot` int(10),
    `Duration` int(11),
    `ExtraFlags` tinyint(1)
);'''


conn = sqlite3.connect(':memory:')
conn.row_factory = sqlite3.Row
conn.execute(CREATE_TABLES_SQL)
conn.executescript(open(ITEM_TEMPLATE_SQL).read())

items = {}


def LoadRowIntoObject(row):
    result = {
        'AllowableClass': row['AllowableClass'],
        'AllowableRace': row['AllowableRace'],
        'AttackRate': row['delay'],
        'BagFamily': row['BagFamily'],
        'Block': row['block'],
        'Bonding': row['bonding'],
        'Class': row['class'],
        'ContainerSlots': row['ContainerSlots'],
        'Description': row['description'],
        'DisenchantID': row['DisenchantID'],
        'DisplayID': row['displayid'],
        'Entry': row['entry'],
        'FoodType': row['FoodType'],
        'InventoryType': row['InventoryType'],
        'ItemLevel': row['ItemLevel'],
        'ItemSet': row['itemset'],
        'LanguageID': row['LanguageID'],
        'LockID': row['lockid'],
        'Material': row['Material'],
        'MaxDurability': row['MaxDurability'],
        'MaxMoneyLoot': row['maxMoneyLoot'],
        'MaxStackSize': row['stackable'],
        'MinMoneyLoot': row['minMoneyLoot'],
        'Name': row['name'],
        'PageMaterial': row['PageMaterial'],
        'PageText': row['PageText'],
        'PerCharacterLimit': row['maxcount'],
        'Quality': row['Quality'],
        'RandomProperty': row['RandomProperty'],
        'RangedModRange': row['RangedModRange'],
        'RequiredAmmoType': row['ammo_type'],
        'RequiredArea': row['area'],
        'RequiredHonorRank': row['requiredhonorrank'],
        'RequiredLevel': row['RequiredLevel'],
        'RequiredMap': row['Map'],
        'RequiredReputationFaction': row['RequiredReputationFaction'],
        'RequiredReputationRank': row['RequiredReputationRank'],
        'RequiredSkillRank': row['RequiredSkillRank'],
        'RequiredSkill': row['RequiredSkill'],
        'RequiredSpell': row['requiredspell'],
        'SheathType': row['sheath'],
        'StartQuest': row['startquest'],
        'SubClass': row['subclass'],
        'TimeUntilDisappear': row['Duration'],
        'VendorBuyPrice': row['BuyPrice'],
        'VendorSellPrice': row['SellPrice'],
        'VendorStackSize': row['BuyCount'],
    }

    # Extra Flags
    result['IsNonConsumable'] = bool(row['ExtraFlags'] & 0x01)
    result['IsRealTimeDuration'] = bool(row['ExtraFlags'] & 0x02)

    # Flags
    result['IsConjured'] = bool(row['Flags'] & 0x0002)
    result['IsLootable'] = bool(row['Flags'] & 0x0004)
    result['IsIndestructible'] = bool(row['Flags'] & 0x0020)
    result['IsUsable'] = bool(row['Flags'] & 0x0040)
    result['IsNoEquipCooldown'] = bool(row['Flags'] & 0x0080)
    result['IsWrapper'] = bool(row['Flags'] & 0x0200)
    result['IsStackable'] = bool(row['Flags'] & 0x0400)
    result['IsPartyLoot'] = bool(row['Flags'] & 0x0800)
    result['IsCharter'] = bool(row['Flags'] & 0x2000)
    result['IsLetter'] = bool(row['Flags'] & 0x4000)
    result['IsPVPReward'] = bool(row['Flags'] & 0x8000)

    SPELL_SCHOOLS = dict(
        normal=0, holy=1, fire=2, nature=3, frost=4, shadow=5, arcane=6,
    )

    result['Resistances'] = {}
    if row['armor']:
        result['Resistances'][SPELL_SCHOOLS['normal']] = row['armor']

    for t in ('arcane', 'fire', 'frost', 'holy', 'nature', 'shadow'):
        if row[t + '_res']:
            result['Resistances'][SPELL_SCHOOLS[t]] = row[t + '_res']

    result['Stats'] = {}
    for i in range(10):
        stype = row['stat_type{}'.format(i + 1)]
        sval = row['stat_value{}'.format(i + 1)]
        if sval > 0:
            result['Stats'][stype] = sval

    result['Damage'] = {}
    for i in range(5):
        dtype = row['dmg_type{}'.format(i + 1)]
        dmin = row['dmg_min{}'.format(i + 1)]
        dmax = row['dmg_max{}'.format(i + 1)]

        if dmin > 0 and dmax > 0:
            result['Damage'][dtype] = dict(
                Min=dmin,
                Max=dmax,
            )

    result['Spells'] = []
    for i in range(5):
        sid = row['spellid_{}'.format(i + 1)]
        strigger = row['spelltrigger_{}'.format(i + 1)]
        scharges = row['spellcharges_{}'.format(i + 1)]
        sppmrate = row['spellppmRate_{}'.format(i + 1)]
        scooldown = row['spellcooldown_{}'.format(i + 1)]
        scategory = row['spellcategory_{}'.format(i + 1)]
        scategory_cooldown = row['spellcategorycooldown_{}'.format(i + 1)]

        if sid:
            result['Spells'].append(dict(
                ID=sid,
                Trigger=strigger,
                Charges=scharges,
                ProcPerMinuteRate=sppmrate,
                Cooldown=scooldown,
                Category=scategory,
                CategoryCooldown=scategory_cooldown,
            ))

    return result


result = {}
cursor = conn.execute("SELECT * FROM item_template")
for row in cursor:
    result[row['entry']] = LoadRowIntoObject(row)


with gzip.open('items.json.gz', 'w') as f:
    f.write(json.dumps(result).encode('utf-8'))
