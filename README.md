# MonoBot

monofuel's tweakable bot. Can be either ran on it's own as app/main.go, or included into other go projects. The goal is to be an easy way to add 'bot-like' functionality to projects. The basic abilities are to broadcast a message, and to handle commands.

# easy install
- as root: `curl https://raw.githubusercontent.com/monofuel/monobot/master/install.sh | bash`
- add `@reboot su monobot -c 'cd /opt/monobot && ./start.sh'` to your crontab

# working
- irc integration
- discord integration
- command handling
- pushbullet integration

# future goals
- provide api
- cooperate with other APIs
- nagios integration?
- self-updating
- installable as a daemon
- cross platform
- slack
- matrix