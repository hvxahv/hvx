#!/bin/bash

SCRIPT_WD=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source "${SCRIPT_WD}/lib/util.sh"

version=$(get_version)

if [ "$version" == 0 ]; then
  echo "VERSION UNKNOWN"
fi
  go build -o ../.release -ldflags "-X github.com/hvxahv/hvxahv/cmd/account/cmd.Version=$version" ../cmd/account