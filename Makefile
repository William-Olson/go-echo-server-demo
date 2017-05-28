run: kill build
	docker-compose up -d

build:
	docker build -t willko/go-echo-server-demo:latest .

kill:
	docker-compose rm -f || true
	docker-compose kill || true


