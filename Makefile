.PHONY: default
default:
	echo 'please choose command'

.PHONY: docs
docs:
	swag init  -g ./cmd/apis/main.go

.PHONY: rundocs
rundocs:
	docker run --rm -it --env GOPATH=/go -v $(shell pwd):/go/src -p 8082:8082 -w /go/src quay.io/goswagger/swagger serve ./docs/swagger.yaml -p 8082 --no-open
