"""A utility which generates DBC files.

It can read files in the following formats:
  1. Binary
  2. JSON

... and produce files in the following formats:
  1. Binary
  2. JSON
  3. Go (useful for generating Enums).
"""
from typing import List
from dbc_generator import FILENAME_TO_RECORD_TYPE

import argparse
import os
import subprocess
import jinja2

from dbc import dbc

def gen_golang_file(chunks: List[bytes], init_function_names: List[str]) -> bytes:
    """Convert the table to a Golang file."""
    args = {
        'package': 'dbc',
        'chunks': chunks,
        'init_function_names': init_function_names,
    }

    template_env = jinja2.Environment(
        loader=jinja2.FileSystemLoader(searchpath=os.path.join(os.path.dirname(os.path.realpath(__file__)), 'dbc')))
    template_env.trim_blocks = True
    template_env.lstrip_blocks = True
    return template_env.get_template('dbc_golang.go.jinja').render(**args)


def main(args: argparse.Namespace):
    golang_chunks = []
    init_function_names = []

    for filename in args.DATA_FILES:
        dbc_filename = os.path.splitext(os.path.basename(filename))[0]
        if dbc_filename not in FILENAME_TO_RECORD_TYPE:
            print('Unknown DBC {}'.format(dbc_filename))
            continue

        record_type = FILENAME_TO_RECORD_TYPE[dbc_filename]

        with open(filename, 'r') as f:
            if filename.endswith('.json'):
                table = dbc.Table.FromJSON(f.read(), record_type)
            elif filename.endswith('.dbc'):
                table = dbc.Table.FromBinary(f.read(), record_type)

        
        golang_chunk, init_function_name = table.ToGolangPart()

        golang_chunks.append(golang_chunk)
        init_function_names.append(init_function_name)

    with open(args.output_file, 'w') as f:
        f.write(gen_golang_file(golang_chunks, init_function_names))

    subprocess.run(['gofmt', '-w', args.output_file])


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Manage DBC files.')
    parser.add_argument(
        'DATA_FILES', type=str, nargs='+', help='The directory containing the DBC input files.')
    parser.add_argument(
        '--output_file', type=str, help='The destination output go file.')
    args = parser.parse_args()
    main(args)
