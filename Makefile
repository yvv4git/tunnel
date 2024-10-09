build_linux:
	docker buildx build --platform linux/amd64 -t yvv4git/tunnel-linux .

build_macos:
	DOCKER_BUILDKIT=0 docker build --no-cache -t yvv4git/tunnel-macos -f Dockerfile .


run_server_macos:
	docker run --rm --name tunnel-macos-server --cap-add=NET_ADMIN --device=/dev/net/tun:/dev/net/tun --entrypoint bash -it yvv4git/tunnel-macos


compose_up:
	docker compose up --build

compose_down:
	docker compose down