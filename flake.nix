{
  description = "Timeful development shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/f4b6996c4e8b9ee06ce147ec344c885f51071b14";
    systems.url = "github:nix-systems/default/future-26.11";
    flake-parts = {
      url = "github:hercules-ci/flake-parts/17c9d6cdfc60c64f4ee8d306f9bc0b4ccb51481e";
      inputs.nixpkgs-lib.url = "github:nix-community/nixpkgs.lib";
    };
  };

  outputs = inputs@{ flake-parts, systems, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import systems;

      perSystem = { pkgs, ... }: {
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.nodejs_26
            pkgs.python3
          ];
        };
      };
    };
}
