dev_envoy:
	(sudo docker build -t uta8a/envoy proxy/;sudo docker run --rm --name envoy -p 9090:9090 uta8a/envoy)
stop_envoy:
	(sudo docker stop envoy)

dev_go:
	(sudo docker build -t uta8a/suburi-server server/;sudo docker run --rm --name suburi-server -p 9090:9090 uta8a/suburi-server)
stop_go:
	(sudo docker stop suburi-server)

up:
	(sudo docker-compose up)
build:
	(sudo docker-compose build)
down:
	(sudo docker-compose down;sudo docker volume rm $$(sudo docker volume ls -qf dangling=true);sudo docker image prune)
delete:
	(sudo docker image prune)
