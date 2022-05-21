build_debug_server:
	docker-compose -f debug-server.docker-compose.yml up --build

build_debug_client:
	docker-compose -f debug-client.docker-compose.yml up --build

run_debug_client:
	docker-compose -f debug-client.docker-compose.yml up

run_debug_server:
	docker-compose -f debug-server.docker-compose.yml up

build_server:
	docker-compose -f server.docker-compose.yml up --build

build_client:
	docker-compose -f client.docker-compose.yml up --build

run_server:
	docker-compose -f server.docker-compose.yml up

run_client:
	docker-compose -f client.docker-compose.yml up
