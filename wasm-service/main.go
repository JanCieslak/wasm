package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Request struct {
	LangType string `json:"langType"`
	Code     string `json:"code"`
}

func getCommand(langType string) *exec.Cmd {
	switch langType {
	case "AssemblyScript":
		return exec.Command("asc", "../code/asc/test.ts", "--outFile", "../code/asc/test.wasm", "--optimize")
	case "Rust":
		return exec.Command("") // TODO
	case "Go":
		return exec.Command("") // TODO
	default:
		panic("Given language type is not supported")
	}
}

func getRequest(r *http.Request) Request {
	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("Received request: %s \n", string(input))

	var request Request
	err = json.Unmarshal(input, &request)
	if err != nil {
		panic(err)
	}

	return request
}

func saveFile(request Request) {
	log.Printf("Saving %s code", request.LangType)

  err := ioutil.WriteFile("../code/asc/test.ts", []byte(request.Code), 0644)
  if err != nil {
    panic(err)
  }
}

func compile(langType string) {
	log.Printf("Compiling %s to wasm", langType)

	cmd := getCommand(langType)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	log.Printf("Sucessfully compiled %s \n", langType)
}

func readWasm(langType string) []byte {
	fileBytes, err := ioutil.ReadFile("../code/asc/test.wasm")
	if err != nil {
		panic(err)
	}

	return fileBytes
}

func handleCompile(w http.ResponseWriter, r *http.Request) {
	request := getRequest(r)
	saveFile(request)
	compile(request.LangType)
	wasmBytes := readWasm(request.LangType)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(wasmBytes)
}

func main() {
	http.HandleFunc("/compile", handleCompile)
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    panic(err)
  }
}
