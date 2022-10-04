// Run server:
// 1) Open powershell
// 2) cd go/src/go-server
// 3) go build
// 4) go run
// 5) search engine -> "localhost:8080/" + hello, form, or form.html

// Shutdown server:
// 1) open CMD
// 2) run "netstat -ano | findstr :8080" and find PID of listening TCP
// 3) run "taskkill /PID 20012 /F"
// 4) do step 2 again to double check no more ports are running

package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprint(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
