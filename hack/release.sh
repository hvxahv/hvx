#!/bin/bash

SCRIPT_WD=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source "${SCRIPT_WD}/lib/util.sh"

version=$(get_version)
go_version=$(get_go_version)

if [ "$version" == 0 ]; then
  echo "VERSION UNKNOWN"
fi

binary_dir="../.release/binary/"
pkg_dir="github.com/hvxahv/hvxahv/cmd/$1"

go build -o $binary_dir -ldflags "-X '$pkg_dir/cmd.Version=$version' -X '$pkg_dir/cmd.GoVersion=$go_version'" ../cmd/"$1"
