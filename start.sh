#!/bin/bash
set -euo pipefail

while [ false ]
do

  ./install.sh || true

  ./monobot
  printf "\nrestarting\n"
  sleep 2s
done