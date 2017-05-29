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
curl "$GIT_URL/install.sh" > install.sh
chmod +x install.sh
curl "$GIT_URL/start.sh" > start.sh
chmod +x start.sh
curl "$GIT_URL/monobot" > monobot
chmod +x monobot

if [[ $EUID == 0 ]]
then
  echo "fixing permissions"
  chown -R monobot /opt/monobot
fi