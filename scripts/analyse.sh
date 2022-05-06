#!/usr/bin/env bash

source scripts/_fn.sh

assert_installed golangci-lint "brew install golangci/tap/golangci-lint"

while read -r line
do
  pushd "$(dirname "$line")" > /dev/null || exit 1
    echo "linting mod $line"
    golangci-lint run --fix --timeout 3m
    status=$?
    [[ $status -eq 0 ]] || exit $status
  popd > /dev/null || exit 1
done < <(list_go_mod)
