bin:
	go build -o ./bin/synk

build:
	cd frontend && npm install && npm run build

run:
	./bin/synk