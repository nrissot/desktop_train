# Tiny train tracks for the bottom of your desktop 🚅

This was a fun weekend+ project (bigger than a "true" weekend project, but most of it was done in a weekend), just for joy and whimsy :)

This project use the excellent [Ebitengine](https://ebitengine.org/), a lightweight game engine for the go programming language. The assets where made using [libresprite](https://libresprite.github.io/)

Both the code and the assets are provided under the [MIT License](/LICENSE.md), meaning you can remix, reuse, modify this project, but a copy of the license must be included with the licensed material.

This project was made without the use of AI.

![](/assets_src/screenshot.png)

## Building from source
This project is not compiled, you need to build it yourself.

To use this project, you will need a go installation version 1.24.0 or above. you can follow the instruction on  [the official go.dev website](https://go.dev/doc/install) to install go. (if you dont already have it on your machine)

> you can type
> ```go version```
> to check you have the correct version of go installed

Now that you have go installed, lets get started:

First, Clone this repo

```sh
git clone https://github.com/nrissot/desktop_train.git
```

Then, you can simply run

```sh
go build
```

and yipie, you did it :) double-click on the executable or execute it from your terminal and that's it, enjoy the trains


## Config
upon the first launch of the app, it'll create a `.desktop_train/` directory. This directory contains two files `config.json` and `stats.json`.

the `config.json` file contains 4 fields:
```json
{
	"Window_y_offset": 39, 		// offset in pixel from the bottom of the monitor screen 
	"Only_Clean_Tracks": false, // should the tracks use only the plain version of the tile ? (no variants)
	"Tracks_Visible": true,		// should the tracks be visible ?
	"Log_Statistics": true		// should desktop train log it's statistics in the stats.json file ?
}
```

the `stat.json` store the amount of trains that passed on your screen (purely informative, doesnt serve any purpose (kinda like this project actually, it just sit there and look fun))
```json
{
	"Train_Counts": 1408
}
```