VERSION=1.0.0

build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/egosa-${VERSION}/egosa cmd/egosa/egosa.go
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/egosa-${VERSION}/egosa cmd/egosa/egosa.go

dist: build
	cd bin/linux/amd64 && tar zcvf egosa-linux-amd64-${VERSION}.tar.gz egosa-${VERSION}
	cd bin/darwin/amd64 && tar zcvf egosa-darwin-amd64-${VERSION}.tar.gz egosa-${VERSION}

bundle:
	dep ensure

clean:
	rm -rf bin/*

