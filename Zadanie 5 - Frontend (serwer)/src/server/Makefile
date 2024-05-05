include Makefile.defs

main:
	$(GO_ENVS) go build -a -installsuffix cgo -mod readonly -ldflags="$(GO_BUILD_DFLAGS)" ./cmd/$(PROJECT_NAME)/main.go

./openapi/swagger.json: ./openapi/spec/$(PROJECT_NAME)*.yaml
	swagger flatten ./openapi/spec/$(PROJECT_NAME)-openapi.yaml $^ -o $@

.PHONY: all
all: clean swagger build

.PHONY: clean
clean:
	rm ./main
	rm ./openapi/swagger.json
	rm -Rf ./openapi/gen

.PHONY: gen-openapi
gen-openapi:
	@make gen-openapi-server

.PHONY: gen-openapi-server
gen-openapi-server:
	mkdir -p $(OPENAPI_GEN_DIR)
	go-swagger.exe generate server -t $(OPENAPI_GEN_DIR) -f $(OPENAPI_SPEC_FILE) -s server --exclude-main -A $(PROJECT_CAPITAL_NAME)

.PHONY: swagger
swagger: gen-openapi

.PHONY: build
build: main

.PHONY: mongo
mongo:
	killall mongod
	rm -rf /tmp/mongodb
	mkdir -p /tmp/mongodb
	mongod --dbpath /tmp/mongodb --replSet rs0 > /dev/null &
	sleep 3
	mongo --eval 'rs.initiate()'
	sleep 3
	mongo ClientsService --eval 'db.createUser({user:"mongo_user",pwd:"mongo_pass",roles:["readWrite","dbAdmin"]})'
