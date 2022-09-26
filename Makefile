run:
	go run cmd/server/main.go

build-docker:
	docker build -t "volume-assignment" -f Dockerfile .

run-docker:
	docker run -p "8080:8080" rollee