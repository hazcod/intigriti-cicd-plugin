all: clean build run

build:
	mkdir -p dist/
	CGO_ENABLED=0 go build -ldflags '-w -s -extldflags "-static"' -o dist/icp ./cmd

run:
	./dist/icp --loglevel=debug --conf=test/icp.yml

clean:
	rm -r ./dist || true
