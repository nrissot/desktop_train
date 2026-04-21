# Tiny train tracks for the bottom of your desktop 🚅

This was a weekend project made entirely without AI, just for fun :)

This project use the excellent [Ebitengine](https://ebitengine.org/), a lightweight game engine for the go programming language. The assets where made using [libresprite](https://libresprite.github.io/)

Both the code and the assets are provided under the [MIT License](/LICENSE.md), meaning you can remix, reuse, modify this project, but a copy of the license must be included with the licensed material.

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

and yipie, you did it :)

## TODO

- [ ] Embed executable icon
- [ ] FPS / TPS in config
- [ ] More graffiti