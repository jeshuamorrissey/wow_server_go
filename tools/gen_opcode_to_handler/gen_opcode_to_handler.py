import argparse
import os
import re
import sys
from typing import Dict, List

_TYPE_REGEXP = re.compile(r'type (Client[a-zA-Z]+) struct')
_OPCODE_METHOD_REGEXP = 'func \(.*\*{}\) OpCode.*return static\.([a-zA-Z]+)'

GOLANG_FILE_TEMPLATE = '''
package world

import (
    "github.com/jeshuamorrissey/wow_server_go/server/world/data/dynamic/interfaces"
    "github.com/jeshuamorrissey/wow_server_go/server/world/data/static"
    "github.com/jeshuamorrissey/wow_server_go/server/world/packet"
    "github.com/jeshuamorrissey/wow_server_go/server/world/packet/handlers"
    "github.com/jeshuamorrissey/wow_server_go/server/world/system"
)


var opCodeToHandler = map[static.OpCode]func(interfaces.ClientPacket, *system.State) ([]interfaces.ServerPacket, error){
%(generated_functions)s
}
'''

GOLANG_GENERATED_FUNCTION_TEMPLATE = '''    static.%(op_code)s: func(pkt interfaces.ClientPacket, state *system.State) ([]interfaces.ServerPacket, error) {
        return handlers.Handle%(packet_name)s(pkt.(*packet.%(packet_name)s), state)
    },'''

MOVE_OPCODES = [
    'OpCodeClientMoveHeartbeat',
    'OpCodeClientMoveSetFacing',
    'OpCodeClientMoveStartBackward',
    'OpCodeClientMoveStartForward',
    'OpCodeClientMoveStartStrafeLeft',
    'OpCodeClientMoveStartStrafeRight',
    'OpCodeClientMoveStartTurnLeft',
    'OpCodeClientMoveStartTurnRight',
    'OpCodeClientMoveStop',
    'OpCodeClientMoveStopStrafe',
    'OpCodeClientMoveStopTurn',
]


def find_packet_types(file_content: str) -> List[str]:
    return _TYPE_REGEXP.findall(file_content)


def find_opcode_names(file_content: str,
                      packet_types: List[str]) -> Dict[str, str]:
    packet_to_opcode = {}
    for packet_type in packet_types:
        # Special case: ClientMove can handle multiple opcodes.
        if packet_type == 'ClientMove':
            packet_to_opcode[packet_type] = MOVE_OPCODES
            continue

        regexp = re.compile(_OPCODE_METHOD_REGEXP.format(packet_type),
                            re.MULTILINE | re.DOTALL)
        opcodes = regexp.findall(file_content)
        if not opcodes:
            print(f'Could not find opcode for {packet_type}', file=sys.stderr)
            continue

        packet_to_opcode[packet_type] = opcodes[0]

    return packet_to_opcode


def main(package_path: str):
    packet_to_opcode = {}
    for file in sorted(os.listdir(package_path)):
        filepath = os.path.join(package_path, file)
        if os.path.splitext(filepath)[1] == '.go':
            with open(filepath) as f:
                file_content = f.read()
                packet_types = find_packet_types(file_content)
                packet_to_opcode.update(
                    find_opcode_names(file_content, packet_types))

    generated_functions = []
    for packet_type_name, opcode_name in sorted(packet_to_opcode.items()):
        if isinstance(opcode_name, list):
            for op in opcode_name:
                generated_functions.append(
                    GOLANG_GENERATED_FUNCTION_TEMPLATE %
                    dict(op_code=op, packet_name=packet_type_name))
        else:
            generated_functions.append(
                GOLANG_GENERATED_FUNCTION_TEMPLATE %
                dict(op_code=opcode_name, packet_name=packet_type_name))

    print(GOLANG_FILE_TEMPLATE %
          dict(generated_functions='\n'.join(generated_functions)))


if __name__ == '__main__':
    argument_parser = argparse.ArgumentParser()
    argument_parser.add_argument(
        '--package_path',
        type=str,
        required=True,
        help='The path to the golang package which contains the packet types.')
    main(**argument_parser.parse_args().__dict__)
