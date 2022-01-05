#!/bin/bash

set -euo pipefail

bazel run //:gazelle-update-repos && bazel run //:gazelle
