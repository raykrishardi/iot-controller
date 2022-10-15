# IoT Controller
IoT controller server and client (cli) built using Go gRPC

## Getting Started

## Supported Feature
- LED (on/off)

## Local Deployment

### Run on local machine

#### Server (e.g. raspberry pi)
```
docker run -d -p 50001:50001 -e GRPC_PORT=50001 raylayadi/iot-controller-grpc:latest
```
#### Client
```
git clone https://github.com/raykrishardi/iot-controller.git
cd iot-controller
make build-iot-cli
cd iot-controller-cli
./iot led --pin 17 --active=true # Turn LED on pin 17
./iot led --pin 17 --active=false # Turn LED off pin 17
```
