{
  inputs = {
    # Candidate channels
    #   - https://github.com/kachick/anylang-template/issues/17
    #   - https://discourse.nixos.org/t/differences-between-nix-channels/13998
    # How to update the revision
    #   - `nix flake update --commit-lock-file` # https://nixos.org/manual/nix/stable/command-ref/new-cli/nix3-flake-update.html
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      rec {
        devShells.default = with pkgs;
          mkShell {
            buildInputs = [
              # https://github.com/NixOS/nix/issues/730#issuecomment-162323824
              bashInteractive

              go_1_20
              nil
              nixpkgs-fmt
              dprint
              actionlint
              go-task
              goreleaser
              typos
              go-tools
            ];
          };

        packages.never18 = pkgs.stdenv.mkDerivation
          {
            name = "never18";
            src = self;
            buildInputs = with pkgs; [
              go_1_20
              go-task
            ];
            buildPhase = ''
              # https://github.com/NixOS/nix/issues/670#issuecomment-1211700127
              export HOME=$(pwd)
              task build
            '';
            installPhase = ''
              mkdir -p $out/bin
              install -t $out/bin dist/never18
            '';
          };

        packages.default = packages.never18;

        # `nix run`
        apps.default = {
          type = "app";
          program = "${packages.never18}/bin/never18";
        };
      }
    );
}
