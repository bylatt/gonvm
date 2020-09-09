# GoNVM

Manage node version with ease.

## Install

### Download the latest release from [here](https://github.com/clozed2u/gonvm/releases)

### From source

```bash
git clone https://github.com/clozed2u/gonvm.git
cd gonvm
go get -u ./...
go build -o /usr/local/bin/gonvm main.go
```

### Add GoNVM to your path env

```bash
export PATH=$HOME/.gonvm/bin:$PATH
```

Don't forget to reload your shell configuration.

## Usage

```bash
gonvm use [version]
node -v
```
