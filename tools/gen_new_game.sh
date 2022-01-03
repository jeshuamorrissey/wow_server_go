#!/bin/bash

set -euo pipefail

SAVE_NAME="jeshua"
REPO_ROOT=$(git rev-parse --show-toplevel)

rm -rf "example_data/${SAVE_NAME?}"
bazel run //tools/new_game -- -name="${SAVE_NAME?}" -save-dir="${REPO_ROOT?}/example_data"