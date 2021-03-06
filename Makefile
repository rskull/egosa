VERSION=2.0.1

build:
	go build -o bin/egosa ./egosa*

build-cross:
	GOOS=linux GOARCH=arm GOARM=6 go build -o bin/linux/amd/egosa-${VERSION}/egosa egosa/*
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/egosa-${VERSION}/egosa egosa/*
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/egosa-${VERSION}/egosa egosa/*

dist: build-cross
	cd bin/linux/arm && tar zcvf egosa-linux-arm-${VERSION}.tar.gz egosa-${VERSION}
	cd bin/linux/amd64 && tar zcvf egosa-linux-amd64-${VERSION}.tar.gz egosa-${VERSION}
	cd bin/darwin/amd64 && tar zcvf egosa-darwin-amd64-${VERSION}.tar.gz egosa-${VERSION}

bundle:
	dep ensure

clean:
	rm -rf bin/*

