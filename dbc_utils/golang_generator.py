"""A utility which generates DBC files.

It can read files in the following formats:
  1. Binary
  2. JSON

... and produce files in the following formats:
  1. Binary
  2. JSON
  3. Go (useful for generating Enums).
"""
from dbc_generator import FILENAME_TO_RECORD_TYPE

import argparse
import os
import subprocess

from dbc import dbc


def main(args: argparse.Namespace):
    for filename in os.listdir(args.in_dir):
        dbc_filename = os.path.splitext(filename)[0]
        if dbc_filename not in FILENAME_TO_RECORD_TYPE:
            print('Unknown DBC {}'.format(dbc_filename))
            continue

        record_type = FILENAME_TO_RECORD_TYPE[dbc_filename]

        with open(os.path.join(args.in_dir, filename), 'r') as f:
            if filename.endswith('.json'):
                table = dbc.Table.FromJSON(f.read(), record_type)
            elif filename.endswith('.dbc'):
                table = dbc.Table.FromBinary(f.read(), record_type)

        out_filename = os.path.join(
            args.out_dir, '{}.go'.format(dbc_filename.lower()))
        with open(out_filename, 'w') as f:
            f.write(table.ToGolang())

        subprocess.run(['gofmt', '-w', out_filename])


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Manage DBC files.')
    parser.add_argument(
        '--in_dir', type=str, help='The directory containing the DBC input files.')
    parser.add_argument(
        '--out_dir', type=str, help='The destination output directory.')
    args = parser.parse_args()
    main(args)
