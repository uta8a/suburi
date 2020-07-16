dev_envoy:
	(sudo docker build -t uta8a/envoy proxy/;sudo docker run --rm --name envoy -p 9090:9090 uta8a/envoy)
stop_envoy:
	(sudo docker stop envoy)
up:
	(sudo docker-compose up)
build:
	(sudo docker-compose build)
down:
	(sudo docker-compose down)
