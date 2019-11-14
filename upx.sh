#!/bin/bash

set -euo pipefail
binaryName="$1"
find dist/ -type f -iname "${binaryName}*" | xargs -I{} -n1 -P 4 upx --best "{}"
