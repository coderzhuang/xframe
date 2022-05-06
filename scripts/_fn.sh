realpath() {
  cd "$1" >/dev/null && pwd -P
}


assert_installed() {
  if ! [ -x "$(command -v $1)" ]; then
    echo "Error: $1 is not installed, try following command" >&2
    echo "$2" >&2
    exit 1
  fi
}

panic() {
  echo "$@"
  exit 1
}

list_go_mod() {
  find . -name go.mod ! -path ./vendor/\* -type f
}
