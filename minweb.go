package main

import ( 
 "fmt" 
 "log" 
 "net/http" 
) 

func main() { 
	http.HandleFunc("/hello", handler) // each request contains /hello calls handler 
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) 
} 

// handler echoes the Path component of the requested URL 
func handler(w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	fmt.Fprintf(w, "Hello, %v!\n", name) 
}
