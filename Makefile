run_app:
	go run main.go server

.PHONY: build_linux
build_linux:
	docker buildx build --platform linux/amd64 -t yvv4git/tunnel-linux .

.PHONY: build_macos
build_macos:
	DOCKER_BUILDKIT=0 docker build --no-cache -t yvv4git/tunnel-macos -f Dockerfile .

run_linux:
	#docker run --rm --entrypoint "/app/tunnel server" yvv4git/tunnel-linux
	docker run --rm --entrypoint bash -it yvv4git/tunnel-linux

run_macos:
	#docker run --rm --name tunnel-macos --cap-add=NET_ADMIN --device=/dev/net/tun:/dev/net/tun -it yvv4git/tunnel-macos
	docker run --rm --name tunnel-macos --cap-add=NET_ADMIN --device=/dev/net/tun:/dev/net/tun --entrypoint bash -it yvv4git/tunnel-macos