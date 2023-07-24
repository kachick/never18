{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/7fd307937db70af23b956c4539033542809ae263";
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
              goreleaser
              # To embed commit ref for revision
              git
            ];
            buildPhase = ''
              task build
            '';
            installPhase = ''
              mkdir -p $out/bin
              install -t $out/bin dist/never18
            '';
          };

        defaultPackage = packages.never18;
      }
    );
}
