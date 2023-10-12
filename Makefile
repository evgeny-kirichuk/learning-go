basics:
		@go run ./cmd/basics/main.go

booking:
		@go run ./cmd/bookingApi/main.go

buildBooking:
		@go build -o ./bin/bookingApi ./cmd/bookingApi/main.go

testBooking:
		@go test -v ./cmd/bookingApi/...