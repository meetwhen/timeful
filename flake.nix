{
  description = "Timeful development shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/c8f90650c15282fa8656a041bfbbd2403997a9a7";
    systems.url = "github:nix-systems/default/future-26.11";
    flake-parts = {
      url = "github:hercules-ci/flake-parts/17c9d6cdfc60c64f4ee8d306f9bc0b4ccb51481e";
      inputs.nixpkgs-lib.url = "github:nix-community/nixpkgs.lib";
    };
    backlog-md.url = "github:MrLesk/Backlog.md/583f928dfa65266df994a4323566eb426446ad55";
  };

  outputs = inputs@{ flake-parts, systems, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import systems;

      perSystem = { pkgs, system, ... }:
        let
          graphify-cli = pkgs.writeShellScriptBin "graphify" ''
            exec ${pkgs.graphify}/bin/graphify "$@"
          '';
        in {
        apps.graphify-mcp = {
          type = "app";
          program = "${pkgs.writeShellScriptBin "graphify-mcp" ''
            export PYTHONPATH="${pkgs.python3Packages.makePythonPath [ pkgs.graphify pkgs.python3Packages.mcp ]}''${PYTHONPATH:+:$PYTHONPATH}"
            exec ${pkgs.python3}/bin/python3 -m graphify.serve graphify-out/graph.json
          ''}/bin/graphify-mcp";
        };

        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.nodejs_26
            pkgs.python3
            pkgs.go
            pkgs.playwright-driver.browsers
            graphify-cli
            inputs.backlog-md.packages.${system}.default
            pkgs.ripgrep
          ];
          shellHook = ''
             export BACKLOG_CWD="$PWD"
             export PLAYWRIGHT_BROWSERS_PATH="${pkgs.playwright-driver.browsers}"
          '';
        };
      };
    };
}
