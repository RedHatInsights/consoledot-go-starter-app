# Name of the binary we build
BINARY_NAME=server
# Quay repo for the project
IMAGE=quay.io/rh_ee_addrew/consoledot-go-starter-app
# Tag for the image
IMAGE_TAG=`git rev-parse --short=7 HEAD`

# If GOPATH isn't set, set it
ifndef GOPATH
export GOPATH := $(HOME)/go
endif
# Check if GOPATH/bin is in PATH
ifneq (,$(findstring $(GOPATH)/bin,$(PATH)))
# NOP - GOPATH/bin is already in PATH
else
# Add GOPATH's bin directory to the PATH
export PATH := $(GOPATH)/bin:$(PATH)
endif

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

# Builds the application binary
build:
	go build -o bin/${BINARY_NAME} main.go

# Runs the application binary
run: generate-api-docs
	go run main.go

# Cleans our project: deletes binaries
clean:
	go clean
	rm bin/${BINARY_NAME}

# Runs unit tests
test:
	go clean -cache
	go test -v ./...

# Builds the binary and installs the api docs binary
setup: build
	go install github.com/swaggo/swag/cmd/swag@latest

# Generates the api docs
generate-api-docs:
	swag init

# Deploys the application to ephemeral
deploy: generate-api-docs build-image push-image
	oc process -f deploy/clowdapp.yaml -p NAMESPACE=$(NAMESPACE) -p ENV_NAME=env-$(NAMESPACE)  IMAGE=${IMAGE} IMAGE_TAG=${IMAGE_TAG} | oc create -f -

# Runs the application's dependencies locally
run-local-deps:
	$(COMPOSE_TOOL) up

# Checks if an image exists in the repo that corresponds to the git SHA at head
# If it does not, it builds it
build-image:
	@if ! $(CONTAINER_ENGINE) images $(IMAGE):$(IMAGE_TAG) --format "{{.Repository}}:{{.Tag}}" | grep -q $(IMAGE):$(IMAGE_TAG); then \
		echo "Image $(IMAGE):$(IMAGE_TAG) not found. Building and pushing..."; \
		$(CONTAINER_ENGINE) build -t $(IMAGE):$(IMAGE_TAG) . ; \
	else \
		echo "Image $(IMAGE):$(IMAGE_TAG) already exists. Skipping build."; \
	fi

# Checks if an image exists in the repo that corresponds to the git SHA at head
# If it does not, it builds it
push-image:
	@if ! $(CONTAINER_ENGINE) images $(IMAGE):$(IMAGE_TAG) --format "{{.Repository}}:{{.Tag}}" | grep -q $(IMAGE):$(IMAGE_TAG); then \
		echo "Image $(IMAGE):$(IMAGE_TAG) not found. Building and pushing..."; \
		$(CONTAINER_ENGINE) build -t $(IMAGE):$(IMAGE_TAG) . ; \
		$(CONTAINER_ENGINE) push $(IMAGE):$(IMAGE_TAG) ; \
	else \
		echo "Image $(IMAGE):$(IMAGE_TAG) already exists. Skipping push."; \
	fi