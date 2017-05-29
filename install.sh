#!/bin/bash
set -euo pipefail

GIT_URL="https://raw.githubusercontent.com/monofuel/monobot/master"

if [[ $EUID -e 0 ]]
then
  id -u monobot &> /dev/null || useradd monobot
  mkdir -p /opt/monobot/
  chown -R monobot /opt/monobot
  su monobot
fi
# assuming we are running as monobot
cd /opt/monobot
wget "$GIT_URL/install.sh"
wget "$GIT_URL/start.sh"
