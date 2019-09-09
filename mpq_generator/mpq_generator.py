"""A utility which generate a MPQ patch file.

It reads JSON DBC files and produces a single MPQ in the requested location.

It also takes any data within "other_data" and copies it into the MPQ.
"""
import argparse
import os
import tempfile
import subprocess

import dbc
import record_types
import dbc_generator

FILENAME_TO_RECORD_TYPE = dict(
    ChrClasses=record_types.ChrClasses,
)


def main(args: argparse.Namespace):
    my_dir = os.path.dirname(os.path.realpath(__file__))

    # Build script using MPQEditor.
    script = [
        'new "{}"'.format(args.out),
        'add "{}" "{}" /r'.format(args.out,
                                  os.path.join(my_dir, 'other_data')),
    ]

    # Make DBC files in a temporary directory.
    with tempfile.TemporaryDirectory() as d:
        os.mkdir(os.path.join(d, 'DBFilesClient'))

        for filename, record_type in dbc_generator.FILENAME_TO_RECORD_TYPE.items():
            with open('dbc_data/{}.json'.format(filename)) as f:
                table = dbc.Table.FromJSON(f.read(), record_type)

            output_filename = os.path.join(
                d, 'DBFilesClient', '{}.dbc'.format(filename))
            with open(output_filename, 'wb') as f:
                f.write(table.ToBinary())

        script.append('add "{}" "{}" /r'.format(args.out, d))

        with tempfile.NamedTemporaryFile(delete=False) as script_file:
            script_file.write('\n'.join(script).encode())
            script_file.close()

            if os.path.exists(args.out):
                os.remove(args.out)

            subprocess.check_output(
                (args.mpqeditor_bin, '-console', script_file.name))

            os.remove(script_file.name)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Manage DBC files.')
    parser.add_argument(
        '--mpqeditor_bin',
        type=str,
        help='The location of the MPQEditor.exe binary.',
        default='D:\\Games\\World of Warcraft (Vanilla) - Tools\\MPQEditor\\MPQEditor.exe')
    parser.add_argument(
        'out', type=str, help='The output MPQ filepath.')
    args = parser.parse_args()
    main(args)
