#!/bin/bash
set -euo pipefail

while [ false ]
do

  ./install.sh || true

  ./monobot || true
  printf "\nrestarting\n"
  sleep 2s
done