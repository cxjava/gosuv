#!/bin/bash

set -euo pipefail

find dist/ -type f -iname 'gosuv*' | xargs -I{} -n1 -P 4 upx --best "{}"
