{
  description = "lcsc api wrapper";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    { nixpkgs
    , flake-utils
    , goflake
    , ...
    }:
    flake-utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
        inherit system;

        overlays = [ goflake.overlay ];
      };
      buildDeps = with pkgs; [ git go_1_23 gnumake ];
      devDeps = with pkgs;
        buildDeps
        ++ [
          golangci-lint
          goreleaser
        ];
    in
    { devShell = pkgs.mkShell { buildInputs = devDeps; }; });
}
