# Name of the binary we build
BINARY_NAME=server
# Quay repo for the project
IMAGE=quay.io/rh_ee_addrew/consoledot-go-starter-app
# Tag for the image
IMAGE_TAG=`git rev-parse --short=7 HEAD`

# Determine the container engine
ifeq ($(shell which podman 2>/dev/null),)
    ifeq ($(shell which docker 2>/dev/null),)
        $(error "No container engine found. Install either podman or docker.")
    else
        CONTAINER_ENGINE = docker
    endif
else
    CONTAINER_ENGINE = podman
endif

# Determine the compose tool
ifeq ($(shell which podman-compose 2>/dev/null),)
    ifeq ($(shell which docker-compose 2>/dev/null),)
        $(error "No compose tool found. Install either podman-compose or docker-compose.")
    else
        COMPOSE_TOOL = docker-compose
    endif
else
    COMPOSE_TOOL = podman-compose
endif

build:
	go build -o bin/${BINARY_NAME} main.go

run: api-docs
	go run main.go

clean:
	go clean
	rm bin/${BINARY_NAME}

test:
	go clean -cache
	go test -v ./...

setup: build
	go install github.com/swaggo/swag/cmd/swag@latest

api-docs:
	swag init

run-ephemeral: api-docs check-image
	oc process -f deploy/clowdapp.yaml -p NAMESPACE=$(NAMESPACE) -p ENV_NAME=env-$(NAMESPACE)  IMAGE=${IMAGE} IMAGE_TAG=${IMAGE_TAG} | oc create -f -

run-local-deps: api-docs
	$(COMPOSE_TOOL) up

build-image:
	$(CONTAINER_ENGINE) build -t ${IMAGE}:${IMAGE_TAG} .

push-image:
	$(CONTAINER_ENGINE) push ${IMAGE}:${IMAGE_TAG}

check-image:
	@if ! $(CONTAINER_ENGINE) images $(IMAGE):$(IMAGE_TAG) --format "{{.Repository}}:{{.Tag}}" | grep -q $(IMAGE):$(IMAGE_TAG); then \
		echo "Image $(IMAGE):$(IMAGE_TAG) not found. Building and pushing..."; \
		$(CONTAINER_ENGINE) build -t $(IMAGE):$(IMAGE_TAG) . ; \
		$(CONTAINER_ENGINE) push $(IMAGE):$(IMAGE_TAG) ; \
	else \
		echo "Image $(IMAGE):$(IMAGE_TAG) already exists. Skipping build and push."; \
	fi