{ pkgs ? import (fetchTarball {
  name = "nixos-unstable-2022-12-07";
  url = "https://github.com/NixOS/nixpkgs/archive/e62dbca3771f42c81dec4e704c09ef110ef3a3f8.tar.gz";
  sha256 = "sha256-jj3Pq8Co7lRem8lF9vqeBTTTs+rX3UjSuebAF+QGezw=";
}) { } }:

with pkgs;

let
  olaf = buildGoModule rec {
    pname = "olaf";
    version = "v1.0.0-beta1";
    vendorSha256 = "sha256-Z5iIG0xXSb5iFbwWbDH7YD5hoX2Pe3TfJ103JX5bXOE=";

    src = ./.;

    ldflags = [ "-s" "-w" "-X main.Version=${version}" ];

    buildInputs = [ xorg.libX11.dev ];

    meta = with lib; {
      homepage = "https://interrato.dev/olaf";
      description = "A simple multi-clipboard system.";
      license = licenses.bsd3;
    };
  };
in {
  olaf = olaf;
}
