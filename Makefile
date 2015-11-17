default: build

clean: 
	rm -f cli/server/server

build: 
	cd service; \
	go build
	cd api; \
	go build
	cd cli/server; \
	go build 
	
deps:
	cd cli/server; \
	go get
