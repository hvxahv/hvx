#!/bin/bash

if [ "$1" == "up" ]; then
	cd ../build && \
	docker-compose down && docker rmi hvx && docker build . -t hvx && docker-compose up -d
elif [ "$1" == "down" ]; then
	cd ../build && \
	docker-compose down
fi
