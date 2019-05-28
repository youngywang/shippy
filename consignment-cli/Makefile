build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run --net=host -e MICRO_REGISTRY=etcdv3 -e MICRO_REGISTRY_ADDRESS=10.123.15.136:2379 consignment-cli
