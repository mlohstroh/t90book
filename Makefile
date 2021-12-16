EXE=video-viewer

run:
	go build -o $(EXE)
	./$(EXE)

install:
	go build -o $(EXE)
	mv $(EXE) /bin/$(EXE)
docker:
	docker build . -t registry.digitalocean.com/mlohstroh/video-viewer:$(shell git rev-parse --short HEAD)
	docker push registry.digitalocean.com/mlohstroh/video-viewer:$(shell git rev-parse --short HEAD)