build:
	protoc -I. --micro_out=./ proto/vessel/vessel.proto
	protoc -I. --go_out=./ proto/vessel/vessel.proto
	# dep 工具暂不可用，直接手动编译
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	#docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service
	docker run --net=host -e MICRO_SERVER_ADDRESS=:50052 -e MICRO_REGISTRY=etcdv3 -e MICRO_REGISTRY_ADDRESS=10.123.15.136:2379 vessel-service
