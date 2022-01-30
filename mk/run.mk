run:
	go run ./hvx/main.go run

account:
	go run ./cmd/account/main.go run

device:
	go run ./cmd/device/main.go run

notify:
	go run ./cmd/notify/main.go run

articles:
	go run ./app/articles/main.go run

# RUN CHANNEL SERVICES.
chan:
	go run ./app/channel/main.go run