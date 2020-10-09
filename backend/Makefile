PROJECT_NAME=soap
PROJECT_PATH=github.com/soap/moment-save

BRANCH=`git rev-parse --abbrev-ref HEAD`
SHORT_SHA=`git rev-parse --short=7 HEAD`
MEDIUM_SHA=`git rev-parse --short=20 HEAD`
GOPATH=$(abspath $(subst $(PROJECT_PATH),,$(CURDIR)))
COGNITO_POOL=test

$(eval SHA = `git rev-parse HEAD`)

LDFLAGS=-ldflags "-X github.com/soap/moment-save/internal/version.Sha=$(SHA) -X github.com/soap/moment-save/internal/version.Branch=$(BRANCH)"

AWS_REGION=us-east-1
AWS_PROFILE=soap
DOCKER_REPO=
DOCKER_PASSWORD=$(DOCKER_REPO)
DOCKER_SOAPCALC_REPO=$(DOCKER_REPO)/soap
DOCKER_ADMIN_REPO=$(DOCKER_REPO)/db-admin

PROTO_PATH=$(abspath $(subst $(PROJECT_PATH),,$(CURDIR)))
PROTO_PATH=$(abspath $(subst $(PROJECT_PATH),,$(CURDIR)))

PROTOC_ZIP=protoc-3.12.3-linux-x86_64.zip
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3

DOCKER_DB_NAME=soap-db
DOCKER_DB_TAG=10.6
DOCKER_DB_REDIS=soap-redis

DOCKER_REACT_NAME=soap-react
DOCKER_REACT_TAG=1.0

DOCKER_NGINX_NAME=soap-nginx
DOCKER_NGINX_TAG=1.0

BIN_DIR="./bin"
SERVICE_MAIN="."
SERVICE_BINARY_NAME="service"
SERVICE_BINARY64_NAME="service.amd64"
EXPORT_DIR="docker/postgres"

TERRAFORM_DIR=$(GOPATH)/github.com/soap/terraform
DEPLOY_SOAPCALC_DIR=$(TERRAFORM_DIR)
DEPLOY_ADMIN_DIR=$(TERRAFORM_DIR)
ifndef DEPLOY
	override DEPLOY = plan
endif

ifndef DEPLOY_SHA
	override DEPLOY_SHA = $(SHA)
endif


GOFILES := $(wildcard *.go)

export GO111MODULE=on

default: all

all: build test 

# basic initialization commands
init: get_geoip get_cognito

build_pkg: generate
	cd pkg; go build ./...

generate: protobufs
	cd pkg; go generate ./...

gqlgen:
	go run github.com/99designs/gqlgen generate

protobufs:
	test -n "$(PROTO_PATH)"
	protoc -I=$(PROTO_PATH) --go_out=plugins=grpc:$(PROTO_PATH) $(PROTO_PATH)/$(PROJECT_PATH)/pkg/domain/reward/soap.proto

# build client and service
# build: build_pkg build_service
build: build_service

# build service code
build_service:
	go build $(LDFLAGS) -o $(BIN_DIR)/$(SERVICE_BINARY_NAME) $(SERVICE_MAIN)

deploy_soap_list:
	AWS_PROFILE=$(AWS_PROFILE) aws ecr list-images --repository-name soap | jq '.imageIds[].imageTag' | sort

deploy_soap:
	cd $(DEPLOY_SOAPCALC_DIR); AWS_PROFILE=$(AWS_PROFILE) terraform12 $(DEPLOY) -var="image_id=$(DEPLOY_SHA)"

deploy_admin:
	cd $(DEPLOY_ADMIN_DIR); AWS_PROFILE=$(AWS_PROFILE) terraform12 $(DEPLOY) -var="image_id=$(DEPLOY_SHA)"

# build docker files for all containers (db, admin, soap)
docker_build: docker_db docker_admin docker_soap docker_react docker_nginx

# start docker compose setup
docker_up: 
	docker-compose -f ../docker/docker-compose.yml up -d --remove-orphans

# shutdown docker compose setup
docker_down:
	docker-compose -f ../docker/docker-compose.yml down

# push all docker files to ECR
docker_push: docker_login docker_soap_push docker_admin_push

# login to the local soap db (outdated)
db:
	# psql -h localhost -U soap -p 15432 soap
	docker exec -it $(DOCKER_DB_NAME) psql -U soap soap

# login to the local soap db
db_soap:
	# psql -h localhost -U soap -p 15432 soap
	docker exec -it $(DOCKER_DB_NAME) psql -U soap soap

# List all tables in the database, useful for testing migrations were successful
db_tables:
	docker exec -it $(DOCKER_DB_NAME) psql -U soap django -c "\dt"
	docker exec -it $(DOCKER_DB_NAME) psql -U soap soap -c "\dt"
	docker exec -it $(DOCKER_DB_NAME) psql -U soap survey -c "\dt"

db_redis:
	docker exec -it $(DOCKER_DB_REDIS) redis-cli -h localhost -p 6379

# Export local database data to EXPORT_DIR
db_export:
	docker exec -it $(DOCKER_DB_NAME) pg_dump -U soap --schema-only django > $(EXPORT_DIR)/django-schema.sql
	docker exec -it $(DOCKER_DB_NAME) pg_dump -U soap --data-only django > $(EXPORT_DIR)/django-data.sql
	docker exec -it $(DOCKER_DB_NAME) pg_dump -U soap --schema-only soap > $(EXPORT_DIR)/soap-schema.sql
	docker exec -it $(DOCKER_DB_NAME) pg_dump -U soap --data-only soap > $(EXPORT_DIR)/soap-data.sql

# import test data previously exported to db
db_import: 
	docker exec -it $(DOCKER_DB_NAME) psql -U soap django -f /backup/django-data.sql
	docker exec -it $(DOCKER_DB_NAME) psql -U soap soap -f /backup/soap-data.sql

# build docker stuff
docker: docker_soap docker_up

# build binary needed for moment save docker container
docker_soap_build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build $(LDFLAGS) -o $(BIN_DIR)/$(SERVICE_BINARY64_NAME) $(SERVICE_MAIN)

# build soap docker container
docker_soap: #docker_soap_build
	docker build -t $(DOCKER_SOAP_REPO):$(SHA) .
	docker tag $(DOCKER_SOAP_REPO):$(SHA) $(DOCKER_SOAPCALC_REPO):$(SHA)

docker_soap_restart: docker_soap docker_up

# push moment save container to ECR
docker_soap_push:
	docker push $(DOCKER_SOAP_REPO):$(SHA)

# dump soap docker logs
docker_soap_logs:
	docker logs -f soap

# build admin docker
docker_admin:
	cd ../admin; docker build -t $(DOCKER_ADMIN_REPO):$(SHA) --build-arg "VERSION=$(SHA)" .
	docker tag $(DOCKER_ADMIN_REPO):$(SHA) $(DOCKER_ADMIN_REPO):0389-$(SHORT_SHA)
	docker tag $(DOCKER_ADMIN_REPO):$(SHA) $(DOCKER_ADMIN_REPO):latest

# push admin container to ecr
docker_admin_push:
	docker push $(DOCKER_ADMIN_REPO):$(SHA)
	docker push $(DOCKER_ADMIN_REPO):0389-$(SHORT_SHA)

# initial local admin with superuser
docker_admin_init:
	docker exec -it admin /venv/bin/python manage.py createsuperuser

# dump admin docker logs
docker_admin_logs:
	docker logs -f admin

# run ecr login
docker_login:
	aws ecr get-login-password \
        --region $(AWS_REGION) | docker login \
        --username AWS \
        --password-stdin $(DOCKER_REPO)

# build local db docker container
docker_db:
	cd ../docker/postgres; docker build -t $(DOCKER_DB_NAME):$(DOCKER_DB_TAG) .

# dump local db docker logs
docker_db_logs:
	docker logs -f $(DOCKER_DB_NAME)

# blow away local db volume useful for testing
docker_db_rm:
	docker volume rm docker_$(DOCKER_DB_NAME)

docker_react:
# We run yarn install inside the container and on the host
# We run it on the host so that eslint will run correctly in our editors such as VS Code
	cd rewardunit; \
	docker build -t $(DOCKER_REACT_NAME):$(DOCKER_REACT_TAG) .; \
	docker run --rm $(DOCKER_REACT_NAME):$(DOCKER_REACT_TAG) cat yarn.lock > yarn.lock;

# Use this command to bring the Docker containers up after you have rebuilt the reward unit image.
# In the docker-compose files, the reward unit image will ignore the node_module folder
# on the host and use the one in the container.
# However, when there are changes to node_modules after a change to package.json
# and a yarn execution, docker-compose will still cache the anonymous volume.
# We must add a --renew-anon-volumes to tell it to re-sync the node_modules folder.
# Thus the command below will tell docker-compose to re-sync.
docker_node_modules: 
	docker-compose -f docker-compose.yml up -d --remove-orphans --renew-anon-volumes

docker_react_logs:
	docker logs -f --tail 100 $(DOCKER_REACT_NAME)

docker_react_restart:
	docker-compose restart react

# Run webpack to compile the reward unit for prod deployment.
docker_react_webpack:
	docker-compose exec react npx webpack -p

# Run webpack to compile what we would see in production.
# It does not perform a real deploy
docker_react_prod:
	docker-compose exec react npx webpack -p --env.GIT_SHA=$(MEDIUM_SHA)

docker_nginx:
# We run yarn install inside the container and on the host
# We run it on the host so that eslint will run correctly in our editors such as VS Code
	cd nginx; \
	docker build -t $(DOCKER_NGINX_NAME):$(DOCKER_NGINX_TAG) .

docker_pull: docker_login
	docker-compose pull

docker_logs:
	docker-compose logs -f --tail=100

# run go vet
vet:
	go vet ./...

# run tests
test: vet
	GOPATH=$(GOPATH) go test -v -cover ./...

get_cognito:
	cd etc; wget -O cognito-jwks.json https://cognito-idp.us-east-1.amazonaws.com/$(COGNITO_POOL)/.well-known/jwks.json

# clean up old stuff
clean:
	rm -f $(BIN_DIR)/$(SERVICE_BINARY_NAME)
	rm -f $(BIN_DIR)/$(CLIENT_BINARY_NAME)
	rm -f $(BIN_DIR)/$(SERVICE_BINARY64_NAME)
	go clean ./...

# run soap service
run: build_service
	source source.env; $(BIN_DIR)/$(SERVICE_BINARY_NAME)

upload:
	curl http://localhost:8080/3.0/graphql \
		-F operations='{ "query": "mutation DoUpload($file: Upload!, $title: String!) { upload(file: $file, title: $title) }", "variables": { "file": null, "title": null } }' \
		-F map='{ "file": ["variables.file"], "title": ["variables.title"] }' \
		-F file=@myfile.txt \
		-F title="My content title"

# run the client
client: build_client
	$(BIN_DIR)/$(CLIENT_BINARY_NAME)

sha:
	echo $(SHA)

.PHONY: all build test clean run
