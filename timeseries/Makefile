
all: binaries

binaries: bin/influxdb-queries bin/influxdb-updates

bin/influxdb-queries:
	test -d bin || mkdir bin
	go build -o bin/influxdb-queries cmd/influxdb-queries.go

bin/influxdb-updates:
	test -d bin || mkdir bin
	go build -o bin/influxdb-updates cmd/influxdb-updates.go

deps:
	go mod download | true

clean:
	rm -f bin/influxdb-queries
	rm -f bin/influxdb-updates
	rm -d bin

.PHONY: all binaries deps clean

