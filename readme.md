# GoNVM

Manage node version with ease.

## Install

### From source

```bash
git clone https://github.com/clozed2u/gonvm.git
cd gonvm/cmd
go build main.go -o /usr/local/bin/gonvm
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
