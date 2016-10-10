PWD = `pwd`

clean:
				go clean ./...

build: 	clean
			 	GOPATH=$(PWD)/vendor CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extld ld -extldflags -static' -a -x -o ola .

docker: build
			  docker build --no-cache --pull -t symfoni/ola:latest .
