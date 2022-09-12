#!/bin/bash

cd ../build && \
docker-compose down && docker rmi hvx && docker build . -t hvx