package main

import ( 
 "fmt" 
 "log" 
 "net/http"
 "strings" 
) 

func main() {
	// each request contains /hello calls handler
	http.HandleFunc("/hello", handler)  
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) 
} 

// handler echoes the Path component of the requested URL 
func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	switch r.Method {
		case "GET":
			get(w, r)
		default:
			fmt.Fprintf(w, "Incorrect response: GET method response is waited\n") 
	}
}

func get (w http.ResponseWriter, r *http.Request) {
	var name string = r.FormValue("name")
	// defense of blank lines
	strings.Replace(name, " ", "", -1)
	if strings.EqualFold(name, "") {
		http.Error(w, "Please, enter the name in request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %v!\n", name) 
}
