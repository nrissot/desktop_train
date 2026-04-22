# Tiny train tracks for the bottom of your desktop 🚅

This was a fun weekend+ project (bigger than a "true" weekend project, but most of it was done in a weekend), just for joy and whimsy :)

This project use the excellent [Ebitengine](https://ebitengine.org/), a lightweight game engine for the go programming language. The assets where made using [libresprite](https://libresprite.github.io/)

Both the code and the assets are provided under the [MIT License](/LICENSE.md), meaning you can remix, reuse, modify this project, but a copy of the license must be included with the licensed material.

This project was made without the use of AI.

![](/assets_src/screenshot.png)

## Downloading the executables
You can download the executable from the [latest release](https://github.com/nrissot/desktop_train/releases/latest)

Current supported os/architectures:
- [`darwin/amd64`](https://github.com/nrissot/desktop_train/releases/download/v1.0/desktop_train-darwin-amd64)
- [`darwin/arm64`](https://github.com/nrissot/desktop_train/releases/download/v1.0/desktop_train-darwin-arm64)
- [`darwin/universal`](https://github.com/nrissot/desktop_train/releases/download/v1.0/desktop_train-darwin-universal)
- [`linux/amd64`](https://github.com/nrissot/desktop_train/releases/download/v1.0/desktop_train-linux-amd64)
- [`windows/amd64`](https://github.com/nrissot/desktop_train/releases/download/v1.0/desktop_train-windows-amd64.exe)

> Huge thanks to [@Aminmiri82](https://github.com/Aminmiri82) for providing the darwin versions <3

## Building from source
To build this project from source, you will need a go installation version 1.24.0 or above. you can follow the instruction on  [the official go.dev website](https://go.dev/doc/install) to install go. (if you dont already have it on your machine)

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

> (Notes): (on linux, Ebitengine require the gcc C compiler, and extra dependancy. please check out [this install guide](https://ebitengine.org/en/documents/install.html?os=linux) to learn more)

and yipie, you did it :) double-click on the executable or execute it from your terminal and that's it, enjoy the trains


## Configs / Stats
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