build:
	go build

lint:
	golangci-lint run

run:
	./awesome-api > ./awesome.log 2>&1 &

unit-tests:
	-go test -v -short -coverprofile=coverage-units.out
	go tool cover -html=coverage-units.out

integration-tests:
	-go test -v -coverprofile=coverage-integrations.out
	go tool cover -html=coverage-integrations.out

stop:
	kill $$(pidof awesome-api)

clean:
	-rm awesome*
	-rm *.out 

test:
	go test -v -short -coverprofile=coverage-units.out
	go test -v -coverprofile=coverage-integrations.out
 
help:
	$(info help: Show help)
	$(info build: Build golang website)
	$(info clean: Delete executable and log file built by "build")
	$(info run: Start website)
	$(info stop: Shut down website)
