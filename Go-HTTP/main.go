package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Website is runing....")

	http.HandleFunc("/", PageLogin)
	http.HandleFunc("/my/about", PageAbout)
	http.HandleFunc("/my/admin", PageAdmin)
	http.HandleFunc("/my/struct", PageStruct)
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":3342", nil))

}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println("ctx: ", ctx)
	fmt.Fprintln(w, ctx)
}

func PageLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func PageAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>ABOUT</h1>")
}

func PageAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Admin</h1>")
}

func PageStruct(w http.ResponseWriter, r *http.Request) {
	var data = map[string]interface{}{
		"name":   "DLU",
		"adress": "Da Lat",
	}
	json.NewEncoder(w).Encode(data)
}
