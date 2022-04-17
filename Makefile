SHELL=bash

proto:
	cd ./api/ && buf generate && cd ../
