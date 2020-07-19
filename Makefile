test-python:
	go build --buildmode=c-shared -o ./test/python/libprime.so ./main.go
	docker build -t ffipython -f test/python/Dockerfile .
test-node:
	go build --buildmode=c-shared -o ./test/nodejs/libprime.so ./main.go
	docker build -t ffinode -f test/nodejs/Dockerfile .
test-numbers:
	go test --cover
run-server:
	go run ./main.go
