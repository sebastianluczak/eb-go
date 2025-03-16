{ pkgs, lib, config, inputs, ... }:

{
  env.GREET = "eb";

  packages = [ pkgs.git ];

  languages.go.enable = true;
  languages.typescript.enable = true;
}
