.PHONY: clean
clean:
	rm -rf ./build
	rm ./app

.PHONY: build
build:
	./front-end/node_modules/.bin/esbuild --bundle ./front-end/src/app.jsx --outfile=./build/app.js && cp ./front-end/src/index.html ./build
	go build -o ./app ./main.go
