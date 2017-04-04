VERSION=1.0.0

build:
	go build -o bin/egosa cmd/egosa/egosa.go

bundle:
	dep ensure

clean:
	rm -rf bin/*

