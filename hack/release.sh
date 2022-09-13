#!/bin/bash

#SCRIPT_WD=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

source "./util.sh"

version=$(get_version)
go_version=$(get_go_version)

if [ "$version" == 0 ]; then
  echo "VERSION UNKNOWN"
fi



binary_dir="../build/binary/"
pkg_dir="github.com/hvxahv/hvx/cmd/$1"

if [ $1 == "gw" ]; then
  echo "BUILD GATEWAY"
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-s -w --extldflags "-static -fpic"' \
  -o $binary_dir \
  -ldflags "-X '$pkg_dir/cmd.Version=$version' -X '$pkg_dir/cmd.GoVersion=$go_version'" \
  ../gateway
else
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -ldflags '-s -w --extldflags "-static -fpic"' \
  -o $binary_dir \
  -ldflags "-X '$pkg_dir/cmd.Version=$version' -X '$pkg_dir/cmd.GoVersion=$go_version'" \
  ../cmd/"$1"
fi
