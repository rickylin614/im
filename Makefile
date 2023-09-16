.PHONY: default
default:
	echo 'please choose command'

.PHONY: swag
docs:
	go install github.com/swaggo/swag/cmd/swag@v1.16.1

.PHONY: docs
docs:
	swag init -g ./cmd/web/main.go

.PHONY: rundocs
rundocs:
	swag init -g ./cmd/web/main.go
	docker run --rm -it --env GOPATH=/go -v $(shell pwd):/go/src -p 8082:8082 -w /go/src quay.io/goswagger/swagger serve ./docs/swagger.yaml -p 8082 --no-open

.PHONY: rundocs2
rundocs2:
	docker run --rm -it --env GOPATH=/go -v C:/code/golang/im:/go/src -p 8082:8082 -w /go/src quay.io/goswagger/swagger serve ./docs/swagger.yaml -p 8082 --no-open

	