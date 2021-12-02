gen:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=import \
        proto/*.proto
clean:
	rm proto/*.go

run:
	go run main.go