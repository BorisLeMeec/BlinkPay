package main

import "net/http"

func main() {
	//TO DO Implement handler
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8001", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "Hello world !"}`))
}