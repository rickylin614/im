.PHONY: default
default:
	echo 'please choose command'

.PHONY: docs
docs:
	swag init -g ./cmd/web/main.go

.PHONY: rundocs
rundocs:
	swag init -g ./cmd/web/main.go
	docker run --rm -it --env GOPATH=/go -v $(shell pwd):/go/src -p 8082:8082 -w /go/src quay.io/goswagger/swagger serve ./docs/swagger.yaml -p 8082 --no-open

model?=""
.PHONY: nunu
nunu:
	nunu create all ${model}
	nunu append ${model}

.PHONY: run
run:
	go run -tags go_json .\cmd\web\main.go

.PHONY: up
up:
	docker-compose -f ./deployments/local/docker-compose.yaml up -d

.PHONY: down
down:
	docker-compose -f ./deployments/local/docker-compose.yaml down

.PHONY: gen
gen:
	go generate ./...