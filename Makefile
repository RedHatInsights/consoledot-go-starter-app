BINARY_NAME=server
IMAGE=quay.io/rh_ee_addrew/consoledot-go-starter-app
IMAGE_TAG=latest

build:
	go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	go test -v ./...

setup: build
	go install github.com/swaggo/swag/cmd/swag@latest

api-docs:
	swag init

run-minikube:
	oc create namespace go-starter-app
	oc process --local -f deploy/clowdenv.yaml -p NAMESPACE=go-starter-app | oc create -f -
	oc process --local -f deploy/clowdapp.yaml -p NAMESPACE=go-starter-app -p ENV_NAME=env-starter-app IMAGE=${IMAGE} IMAGE_TAG=${IMAGE_TAG} | oc create -f -

run-ephemeral:
	oc process -f deploy/clowdapp.yaml -p NAMESPACE=$(NAMESPACE) -p ENV_NAME=env-$(NAMESPACE)  IMAGE=${IMAGE} IMAGE_TAG=${IMAGE_TAG} | oc create -f -
