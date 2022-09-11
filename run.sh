#!/usr/bin/env bash

#depscan --report_file reports/depscan.json --sync
echo "==== DEPSCAN & SCAN ===="
python3 /usr/local/src/scan
echo "==== FOSSA CLI ===="
fossa analyze -o > reports/logs-output.json
echo "DONE"