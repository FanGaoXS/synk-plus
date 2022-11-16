bin:
	go build -o ./bin/synk

build:
	cd frontend && npm run build

run:
	./bin/synk