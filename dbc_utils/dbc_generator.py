"""A utility which generates DBC files.

It can read files in the following formats:
  1. Binary
  2. JSON

... and produce files in the following formats:
  1. Binary
  2. JSON
  3. Go (useful for generating Enums).
"""
import argparse
import os

from dbc import record_types, dbc

FILENAME_TO_RECORD_TYPE = dict(
    ChrClasses=record_types.ChrClasses,
    ChrRaces=record_types.ChrRaces,
    ChrStartingLocations=record_types.ChrStartingLocations,
    ChrStartingStats=record_types.ChrStartingStats,
    CharBaseInfo=record_types.CharBaseInfo,
)


def main(args: argparse.Namespace):
    record_type = FILENAME_TO_RECORD_TYPE[os.path.splitext(
        os.path.basename(args.src))[0]]

    if args.src.endswith('.dbc'):
        with open(args.src, 'rb') as f:
            table = dbc.Table.FromBinary(f.read(), record_type)
    elif args.src.endswith('.json'):
        with open(args.src, 'r') as f:
            table = dbc.Table.FromJSON(f.read(), record_type)
    else:
        raise ValueError('Unknown input file type')

    if args.dst.endswith('.dbc'):
        output = table.ToBinary()
    elif args.dst.endswith('.json'):
        output = table.ToJSON()
    else:
        raise ValueError('Unknown output file type')

    with open(args.dst, 'w') as f:
        f.write(output)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Manage DBC files.')
    parser.add_argument(
        'src', type=str, help='The source file to use for the conversion.')
    parser.add_argument(
        'dst', type=str, help='The destination file for the conversion.')
    args = parser.parse_args()
    main(args)
