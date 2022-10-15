IOT_CONTROLLER_CLI_BINARY=iot
IOT_CONTROLLER_GRPC_BINARY=iotGRPC
IOT_CONTROLLER_CLI_DEST_PATH="webshell-ws"

build-iot-cli:
	@echo "Building ${IOT_CONTROLLER_CLI_BINARY} binary..."
	cd iot-controller-cli && env GOOS=linux CGO_ENABLED=0 go build -o ${IOT_CONTROLLER_CLI_BINARY} ./cmd/iot
	@echo "Done!"

build-iot-cli-rpi:
	@echo "Building ${IOT_CONTROLLER_CLI_BINARY} binary..."
	cd iot-controller-cli && env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o ${IOT_CONTROLLER_CLI_BINARY} ./cmd/iot
	@echo "Done!"

add-iot-cli: build-iot-cli
	@echo "Adding ${IOT_CONTROLLER_CLI_BINARY} binary to ${IOT_CONTROLLER_CLI_DEST_PATH}..."
	cd iot-controller-cli && mv ${IOT_CONTROLLER_CLI_BINARY} ../${IOT_CONTROLLER_CLI_DEST_PATH}
	@echo "Done!"

build-iot-grpc:
	@echo "Building ${IOT_CONTROLLER_GRPC_BINARY} binary..."
	cd iot-controller-grpc && env GOOS=linux CGO_ENABLED=0 go build -o ${IOT_CONTROLLER_GRPC_BINARY} ./cmd/api
	@echo "Done!"

build-iot-grpc-rpi:
	@echo "Building ${IOT_CONTROLLER_GRPC_BINARY} binary..."
	cd iot-controller-grpc && env GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -o ${IOT_CONTROLLER_GRPC_BINARY} ./cmd/api
	@echo "Done!"