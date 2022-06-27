build:
	go build

lint:
	golangci-lint run

run:
	./awesome-api > ./awesome.log 2>&1 &

unit-test:
	go test -v -short -coverprofile=coverage-units.out

stop:
	kill $$(pidof awesome-api)

clean:
	rm awesome*

test:
	curl http://localhost:9999
	curl http://localhost:9999/health

help:
	$(info help: Show help)
	$(info build: Build golang website)
	$(info clean: Delete executable and log file built by "build")
	$(info run: Start website)
	$(info stop: Shut down website)
