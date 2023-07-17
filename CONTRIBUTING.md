# How to develop

## Setup

1. Install [Nix](https://nixos.org/) package manager
2. Run `nix develop` or `direnv allow` in project root
3. You can use development tools

```console
> nix develop
(prepared shell)

> task fmt
task: [fmt] dprint fmt
task: [fmt] go fmt

> task
task: [build] ..."
task: [test] go test
task: [lint] dprint check
task: [lint] go vet
PASS
ok      never18    0.313s

> find dist
dist
dist/metadata.json
dist/config.yaml
dist/never18_linux_amd64_v1
dist/never18_linux_amd64_v1/never18
dist/artifacts.json

> ./dist/never18_linux_amd64_v1/never18 --version
never18 0.1.1-next (906924b) # 2023-06-19T09:33:14Z
```
