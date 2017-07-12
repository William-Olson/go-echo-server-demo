run: kill build
	docker-compose up -d

build:
	docker build -t willko/go-echo-server-demo:latest .

kill:
	docker-compose kill || true
	docker-compose rm -f || true

