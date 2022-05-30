DOCKERHUB_USER_NAME=lptonussi

##############
# playground #
##############






build_debug_server:
	docker-compose -f debug-server.docker-compose.yml up --build

build_debug_client:
	docker-compose -f debug-client.docker-compose.yml up --build

build_debug_hermes:
	docker-compose -f debug-hermes.docker-compose.yml up --build

build_server:
	docker-compose -f server.docker-compose.yml up --build

build_hermes:
	docker-compose -f hermes.docker-compose.yml up --build

build_client:
	docker-compose -f client.docker-compose.yml up --build

build_client_with_python_server:
	docker-compose -f go-client-python-server.docker-compose.yml up --build

build_mock:
	docker-compose -f mock-hashicorp-raft-join-server.docker-compose.yml up --build




run_debug_client:
	docker-compose -f debug-client.docker-compose.yml up

run_debug_hermes:
	docker-compose -f debug-hermes.docker-compose.yml up

run_debug_server:
	docker-compose -f debug-server.docker-compose.yml up

run_server:
	docker-compose -f server.docker-compose.yml up

run_client:
	docker-compose -f client.docker-compose.yml up

run_mock:
	docker-compose -f mock-hashicorp-raft-join-server.docker-compose.yml up






run_python_http_log_server:
	docker-compose -f go-client-python-server.docker-compose.yml up http-log-server

run_go_http_log_client:
	docker-compose -f go-client-python-server.docker-compose.yml up http-log-client

run_hermes:
	docker-compose -f hermes.docker-compose.yml up







docker_build_client:
	docker build -t ${DOCKERHUB_USER_NAME}/public:go-http-log-client -f client.dockerfile .

docker_run_client:
	docker run --name go-http-log-client ${DOCKERHUB_USER_NAME}/public:go-http-log-client







docker_down:
	docker-compose -f go-client-python-server.docker-compose.yml down
	docker-compose -f client.docker-compose.yml down
	docker-compose -f server.docker-compose.yml down
	docker-compose -f hermes.docker-compose.yml down
	docker-compose -f mock-hashicorp-raft-join-server.docker-compose.yml

