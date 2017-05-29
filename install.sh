#!/bin/bash
set -euo pipefail

GIT_URL="https://raw.githubusercontent.com/monofuel/monobot/master"

if [[ $EUID == 0 ]]
then
  echo "setting up project"
  id -u monobot &> /dev/null || useradd monobot
  mkdir -p /opt/monobot/
elif [[ $USER != "monobot" ]]
then
  echo "run this as monobot to update, or as root to install"
  exit
fi
echo "updating"
cd /opt/monobot
wget "$GIT_URL/install.sh" -N -O install.sh.tmp
chmod +x install.sh.tmp
wget "$GIT_URL/start.sh" -N
chmod +x start.sh
wget "$GIT_URL/monobot" -N
chmod +x monobot

if [[ $EUID == 0 ]]
then
  echo "fixing permissions"
  chown -R monobot /opt/monobot
fi
mv install.sh.tmp install.sh