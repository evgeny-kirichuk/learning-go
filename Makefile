basics:
		@go run ./cmd/basics/main.go

serveBooking:
		@go run ./cmd/bookingApi/main.go

buildBooking:
		@go build -o ./bin/bookingApi ./cmd/bookingApi/main.go

serveProdBooking: buildBooking
		@./bin/bookingApi

testBooking:
		@go test -v ./bookingApi/...