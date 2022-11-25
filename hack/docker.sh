#!/bin/bash

if [ "$1" == "dup" ]; then
	cd ../build && \
	docker-compose down && docker rmi hvturingga/hvxahv && docker build . -t hvturingga/hvxahv && docker-compose up -d
elif [ "$1" == "down" ]; then
	cd ../build && \
	docker-compose down
elif [ "$1" == "up" ]; then
	cd ../build && \
  	docker build . -t hvturingga/hvxahv && docker-compose up -d
elif [ "$1" == "push" ]; then
  	docker push hvturingga/hvxahv:latest
fi
