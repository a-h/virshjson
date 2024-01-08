{
  description = "virsh-json";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/23.11";
    gitignore = {
      url = "github:hercules-ci/gitignore.nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, gitignore }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {
        inherit system;
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      packages = forAllSystems ({ pkgs, ... }: {
        default = pkgs.buildGo121Module {
          name = "virsh-json";
          src = gitignore.lib.gitignoreSource ./.;
          subPackages = [ "cmd/virsh-json" ];
          vendorHash = "sha256-1wycFQdf6sudxnH10xNz1bppRDCQjCz33n+ugP74SdQ=";
          CGO_ENABLED = 0;
          flags = [
            "-trimpath"
          ];
          ldflags = [
            "-s"
            "-w"
            "-extldflags -static"
          ];
        };
      });

      # `nix develop` provides a shell containing development tools.
      devShell = forAllSystems ({ system, pkgs }:
        pkgs.mkShell {
          buildInputs = with pkgs; [
            go_1_21
          ];
        });

      # This flake outputs an overlay that can be used to add virsh-json to nixpkgs.
      #
      # Example usage:
      #
      # nixpkgs.overlays = [
      #   inputs.virsh-json.overlays.default
      # ];
      overlays.default = final: prev: {
        virsh-json = self.packages.${final.stdenv.system}.virsh-json;
      };
    };
}

