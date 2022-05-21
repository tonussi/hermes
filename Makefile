build_debug:
	docker-compose -f debug.docker-compose.yml up --build

build_server:
	docker-compose -f server.docker-compose.yml up --build

build_client:
	docker-compose -f client.docker-compose.yml up --build

run_debugger:
	docker-compose -f debug.docker-compose.yml up

run_server:
	docker-compose -f server.docker-compose.yml up

run_client:
	docker-compose -f client.docker-compose.yml up
