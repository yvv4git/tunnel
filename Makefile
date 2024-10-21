image_build_linux:
	docker buildx build --platform linux/amd64 -t yvv4docker/tunnel-linux .

image_build_macos:
	DOCKER_BUILDKIT=0 docker build --no-cache -t yvv4docker/tunnel-macos -f Dockerfile .


run_server_macos:
	docker run --rm --name tunnel-macos-server --cap-add=NET_ADMIN --device=/dev/net/tun:/dev/net/tun --entrypoint bash -it yvv4docker/tunnel-macos

image_push_macos:
	docker tag yvv4docker/tunnel-macos:latest docker.io/yvv4docker/tunnel-macos:v0.0.1
	docker push docker.io/yvv4docker/tunnel-macos:v0.0.1

image_push_linux:
	docker tag yvv4docker/tunnel-linux:latest docker.io/yvv4docker/tunnel-linux:v0.0.1
	docker push docker.io/yvv4docker/tunnel-linux:v0.0.1


gen_certs:
	./scripts/gen_server_certs.sh
	./scripts/gen_client_certs.sh

compose_up:
	docker compose up

compose_up_only_server:
	docker compose up --scale client=0

compose_down:
	docker compose down

go_build_linux_x64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/tunnel_linux main.go

go_build_linux_386:
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o bin/tunnel_linux_386 main.go