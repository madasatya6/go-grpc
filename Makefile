DIRS  = $(shell cd deploy && ls -d */ | grep -v "_output")
ODIR := deploy/_output
IMAGE = registry.app.co.id/vms/be/$(svc)

export VAR_SERVICES       ?= $(DIRS:/=)
export VERSION            ?= $(shell git show -q --format=%h)

pretty:
	gofmt -s -w .

dep:
	go mod download

$(ODIR):
	@mkdir -p $(ODIR)

compile: $(ODIR)
	@$(foreach svc, $(VAR_SERVICES), \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(svc)/$(svc) app/$(svc)/main.go;)

migrate-up:
	go run app/migration/main.go up

migrate-down:
	go run app/migration/main.go down

migrate-status:
	go run app/migration/main.go status

build:
	@$(foreach svc, $(VAR_SERVICES), \
		docker build -t $(IMAGE):latest -f ./deploy/$(svc)/Dockerfile .;)

push:
	@$(foreach svc, $(VAR_SERVICES), \
		docker login registry.app.co.id -u corporate -p SkXkvmX1RmykEVTFgZos && \
		docker push $(IMAGE):latest;)
