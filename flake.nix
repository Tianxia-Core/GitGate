# ffp: FFP-NIX-FLAKE declarative nix
{
  description = "GitGate v1.0.0 — GitGate Zero-Token GitHub Commit, 14-Gate Differential PR Reviewer & Clean-Room Auto-Merge CLI";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  };

  outputs = { self, nixpkgs, ... }:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      artifactMeta = {
        "x86_64-linux" = {
          url = "https://github.com/Tianxia-Core/GitGate/releases/download/v1.0.0/gitgate-v1.0.0-linux-x86_64";
          sha256 = "5a7dda9c7c8c694380af8db279fd088a91a307ecd35f09d83dda6a337bdb4c8a";
        };
        "aarch64-linux" = {
          url = "https://github.com/Tianxia-Core/GitGate/releases/download/v1.0.0/gitgate-v1.0.0-linux-aarch64";
          sha256 = "23097709caec65f22e0c217f2e1f48a9b85ed894be4139258f3c301b72147be3";
        };
      };
      mkSovereignPkg = pkgs: system:
        let
          metaInfo = artifactMeta.${system};
        in
        pkgs.stdenvNoCC.mkDerivation {
          pname = "gitgate";
          version = "1.0.0";
          src = pkgs.fetchurl {
            inherit (metaInfo) url sha256;
          };
          dontUnpack = true;
          installPhase = ''
            runHook preInstall
            mkdir -p $out/bin
            install -m 0755 $src $out/bin/gitgate
            runHook postInstall
          '';
          meta = with pkgs.lib; {
            description = "GitGate Zero-Token GitHub Commit, 14-Gate Differential PR Reviewer & Clean-Room Auto-Merge CLI";
            homepage = "https://rasaxmachina.com";
            platforms = supportedSystems;
            mainProgram = "gitgate";
          };
        };
    in {
      packages = forAllSystems (system:
        let
          pkgs = import nixpkgs { inherit system; };
          pkg = mkSovereignPkg pkgs system;
        in {
          default = pkg;
          "gitgate" = pkg;
        }
      );

      apps = forAllSystems (system: {
        default = {
          type = "app";
          program = "${self.packages.${system}.default}/bin/gitgate";
        };
      });

      overlays.default = final: prev: {
        "gitgate" = mkSovereignPkg final final.stdenv.hostPlatform.system;
      };

      nixosModules.gitgate = { config, lib, pkgs, ... }:
        let
          cfg = config.services.gitgate;
        in {
          options.services.gitgate = {
            enable = lib.mkEnableOption "GitGate Sovereign Service";
            package = lib.mkOption {
              type = lib.types.package;
              default = self.packages.${pkgs.stdenv.hostPlatform.system}.default;
              description = "Package providing gitgate.";
            };
          };
          config = lib.mkIf cfg.enable {
            environment.systemPackages = [ cfg.package ];
          };
        };
    };
}
