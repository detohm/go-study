package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("."))
	if err := http.ListenAndServe("localhost:8080", fileServer); err != nil {
		fmt.Println(err)
		return
	}
}

/* command
cd app && GOOS=js GOARCH=wasm go build -o main.wasm && cd -
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
*/
