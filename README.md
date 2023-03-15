# Azan timings

This repository contains of a go file to fetch the salah times from my-masjid.com and then extract the salah times from the
api response so it can be played on small computers like rasperrypi 

The problem that I wanted to solve, there is no a good app to show the real times of islamic prayers in Berlin (as I'm living in berlin) 
I tried to use many mobile apps that at the end found they are not accurate, and also wanted to have something that keep running in my home and not consuming so much power

I tried also to use azan timers that are can be found on amazon, but the problem again they are limited to some standard azan sound that I didn't like and I wanted to use custom azan اذان and custom duaa(دعاء)

# Prerequisites
- If you are using it on rasperry pi like me, you can use the generated binary at the release page.
- Any rasperry-pi model, but also you can try to use it on mini PC if you can't find one to buy online.
- Speakers you can connect to rasperry pi.
- install `sox` on you rasperry pi so you can play the wav files from terminal.

# Getting Started 

```bash
sudo apt-get install sox -y
sudo apt-get update
sudo apt-get install -y supervisor 
sudo service supervisor start
```

# Usage
- I use supervisor to monitor the binary so it restart everytime it fails
- sudo vim /etc/supervisor/conf.d/azan.conf (so we can add configs for supervisor)

Configs should be like that:
```bash
[program:azan]
directory=path/to/your/binary/directory
command=path/to/your/binary/directory/azan
autostart=true
autorestart=true
stderr_logfile=/var/log/salah.err
stdout_logfile=/var/log/salah.log
```
- Create a directory on rasperry pi with directory `azan` and `duaa` and the binary file downloaded.
- Place in azan اذان dir any wav file you would like to play and also put in `duaa` دعاء another wav file for the duaa دعاء you want to play after azan.
- I have in this repo two wav samples you can directly use.

# License
[MIT license](LICENSE)
