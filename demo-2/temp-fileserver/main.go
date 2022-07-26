package main

import (
	"log"
	"net/http"
)

func cors(fs http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Access-Control-Allow-Origin", "*")
    fs.ServeHTTP(w, r)
  }
}

func main() {
  fs := http.FileServer(http.Dir("../pkg/"))
	http.Handle("/", cors(fs))
  log.Fatal(http.ListenAndServe(":1234", nil))
}
