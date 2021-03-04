example-app:
	rm -rf example/html
	mkdir example/html
	tinygo build -o example/html/wasm.wasm -target wasm -no-debug example/wasm.go
	cp example/wasm_exec.js example/html/
	cp example/wasm.js example/html/
	cp example/index.html example/html/
	go run example/wasm-server/main.go