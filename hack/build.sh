#!/bin/bash

source ./version.sh

# shellcheck disable=SC2154
if [ "$version" == 0 ]; then
  echo "VERSION UNKNOWN"
fi
  go build -o ../.release -ldflags "-X github.com/hvxahv/hvxahv/hvx/cmd.Version=$version" ../hvx