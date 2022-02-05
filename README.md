# ConfeTTY

![build workflow](https://github.com/maaslalani/confetty/actions/workflows/build.yml/badge.svg)

<https://user-images.githubusercontent.com/42545625/128612977-5d6e0321-9584-48b5-8ff8-dd5b811211d3.mov>

Confetti (or fireworks) inside your terminal.

## Preview

You can quickly preview `confetty` through `ssh` (thanks to [charmbracelet/wish](https://github.com/charmbracelet/wish))

```bash
# Confetti
ssh -p 2222 ssh.caarlos0.dev
# Fireworks
ssh -p 2223 ssh.caarlos0.dev
```

## Installation

### Using go toolchain

```bash
go install github.com/maaslalani/confetty
```

### Using homebrew

```bash
brew install maaslalani/tap/confetty
```

### Using yum

```bash
yum install -y <<latest rpm url from releases section>>
```

### Using apt

```bash
apt install -y <<latest deb url from releases section>>
```

### Other platforms

Head over to the [releases section](https://github.com/maaslalani/confetty/releases) and download the binary for your platform.

## Usage

```bash
confetty
```

```bash
confetty fireworks
```

Press any key to cause more confetti / fireworks to appear.
`Ctrl-C` or `q` to exit.

## Why?

¯\\\_(ツ)\_/¯
