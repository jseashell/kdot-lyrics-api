.PHONY: build clean gosumgen

build: gosumgen
	export GO111MODULE=on
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

clean:
	rm -rf ./bootstrap ./vendor go.sum
	
gosumgen: clean
	go mod tidy

deploy: clean build
	sls deploy --verbose
