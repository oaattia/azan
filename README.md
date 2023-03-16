# Azan Player

This repository contains go file to fetch the prayers times from my-masjid.com and then extract the salah times from the api response so it can be played on micro computers like rasperrypi.

The problem that I wanted to solve, there is no a good app to play the real times of the prayers in Berlin (as I'm living in berlin), I tried to use many mobile apps that at the end found they are not accurate, and also wanted to have something that keep running in my home and not consuming so much power.

my-masjid.com has prayers for most of the mosquees and it added by the mosquee itself and make it easy to follow a unified way to have azan.

# Prerequisites
- If you are using it on rasperry pi like me, you can use the generated binary at the [release](https://github.com/oaattia/azan/releases) page.
- Any rasperry-pi model, but also you can try to use it on mini PC if you can't find one to buy online.
- Speakers you can connect to rasperry pi.
- install `sox` on you rasperry pi so you can play the wav files from terminal.

# Getting Started 

```bash
sudo apt-get update
sudo apt-get install sox -y
```

# Usage
I use crontab to schedule running every minute.
```
* * * * * /home/pi/salah
```

- Create a directory on rasperry pi with directory media under this directory `azan` and `duaa` directory is placed.
- `azan` is the azan dir contain all the wav files for the azans which can played randomly everytime if there is many files.
- `duaa` is the duaa dir contain all the wav files for the duaa which can played randomly after azan.
- Download the binary file and place it in the same directory.

In this repo two wav samples you can directly use.

# License
[MIT license](LICENSE)
