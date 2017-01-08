package main

import ( 
 "fmt" 
 "log" 
 "net/http" 
) 

func main() { 
	var name string = "test"
	http.HandleFunc("/hello", func (w http.ResponseWriter, r *http.Request) {
		handler(w, r, name)
	}) // each request contains /hello calls handler 
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) 
} 

// handler echoes the Path component of the requested URL. 
func handler(w http.ResponseWriter, r *http.Request, name string) {
	fmt.Fprintf(w, "Hello, %q\n", name) 
}
