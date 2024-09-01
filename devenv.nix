{
  pkgs,
  lib,
  config,
  inputs,
  ...
}:
{
  env.GREET = "devenv";

  packages = with pkgs; [
    git
    docker
    docker-compose
    go-migrate
  ];

  # Main service language
  languages.go = {
    enable = true;
    package = pkgs.go_1_21;
  };

  # Main database
  # services.postgres = {
  #   enable = true;
  #   package = pkgs.postgresql_15;
  #   initialDatabases = [ { name = "LibreLib"; } ];
  # };

  # Web-proxy
  # services.nginx.enable = true;

  # Cache for service
  # services.redis.enable = true;

  scripts.hello.exec = ''
    echo hello from $GREET
  '';

  enterShell = # bash
    ''
      source ./.env
      Backend for LibreLib
      go version
    '';

  enterTest = ''
    echo "Running tests"
    git --version | grep --color=auto "${pkgs.git.version}"
  '';

}
