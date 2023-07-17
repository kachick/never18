# never18

Happy happy 17th birthday to everyone, I hope you all had a safe and happy 17th birthday this year!

## Installation

`go install` is also okay, or use [prebuilt binaries](https://github.com/kachick/never18/releases)

```console
> curl -L https://github.com/kachick/never18/releases/latest/download/never18_Linux_x86_64.tar.gz | tar xvz -C ./ never18
> ./never18 --version
never18 0.0.1 (70f68fa) # 2023-07-17T21:58:05Z
```

## Usage

```console
> date
Mon Jul 17 22:40:49 JST 2023

> never18 --from "-552-09-28"
e57b65abbbf7a2d5786acc86fdf56cde060ed026

> never18 --to $(date)

> never18 bump && git commit -m 'Bump nixpkgs to latest' *.nix
[main 213d1bf] Bump nixpkgs to latest
 1 file changed, 1 insertion(+), 1 deletion(-)
```
