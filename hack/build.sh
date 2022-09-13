#!/bin/bash

if [ $1 == "all" ]; then
  cd ../build && \
  docker-compose down && docker rmi hvx && docker build . -t hvx && docker-compose up -d
fi