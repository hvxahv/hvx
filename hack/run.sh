#!/bin/bash
svc=$1

if [ "$svc" = "gw" ]; then
  echo "RUN GATEWAY"
  go run ../gateway/main.go run
else
  echo "RUN $svc SERVICE"
  go run ../cmd/"$1"/main.go run
fi
