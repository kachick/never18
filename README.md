# never18

Happy happy 17th birthday to everyone, I hope you all had a safe and happy 17th birthday this year!

## Installation

[Prebuilt binaries](https://github.com/kachick/never18/releases)

```console
> curl -L https://github.com/kachick/never18/releases/latest/download/never18_Linux_x86_64.tar.gz | tar xvz -C ./ never18
> ./never18 --version
never18 0.0.1 (70f68fa) # 2023-07-18T21:58:05Z
```

`go install`

```console
> go install github.com/kachick/never18/cmd/never18@latest
go: downloading...
> ${GOPATH:-"$HOME/go"}/bin/never18 --version
never18 dev (rev)
```

## Usage

```console
> never18 --birth=1962-08-07
17 years, 527 months, 11 days

> never18 --birth=1962-08-07 --limit=12
12 years, 587 months, 11 days

> never18 --birth=1962-08-07 --limit=12 --moment=2112-09-03
12 years, 1656 months, 27 days

> never18 --birth=1962-08-07 --nominally
60 years, 11 months, 11 days
```

You can check other options via `--help`.

- [1962-08-07](https://en.wikipedia.org/wiki/Nobita_Nobi) # There are several theories.
- [2112-09-03](https://en.wikipedia.org/wiki/Doraemon_(character))

## Bug Report

Add the `--doctor` option, and report your terminal screenshot and output from GitHub issues.
