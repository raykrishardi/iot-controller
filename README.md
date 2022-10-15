# IoT Controller
IoT controller server and client (cli) built using Go gRPC

## Getting Started

## Supported Feature
- LED (on/off)

## Local Deployment

### Run on local machine

#### Server (Raspberry pi)
```
docker run -d -p 50001:50001 -e GRPC_PORT=50001 raylayadi/iot-controller-grpc:rpi-latest

git clone https://github.com/raykrishardi/iot-controller.git
cd iot-controller
make build-iot-grpc-rpi
cd iot-controller-grpc
env GRPC_PORT=50001 ./iotGRPC
```

#### Client (MacOS)
```
docker run -e GRPC_HOST="<rpi IP>" -e GRPC_PORT="50001" -it raylayadi/iot-controller-cli:latest led --pin 17 --active=true
```
