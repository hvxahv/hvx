run:
	go run ./hvx/main.go run

acct:
	go run ./app/accounts/main.go run

articles:
	go run ./app/articles/main.go run

# RUN CHANNEL SERVICES.
chan:
	go run ./app/channel/main.go run