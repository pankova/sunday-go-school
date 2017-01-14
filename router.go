package main

import ( 
 "fmt" 
 "github.com/julienschmidt/httprouter"
 "github.com/sirupsen/logrus" 
 "log" 
 "net/http"
 "strings"
 "io/ioutil"
) 

func main() {
	logrus.Info("Hi. I am working!")
	router := httprouter.New()
	router.GET("/hello", getSimple)
	router.GET("/hello/:name", getParam)
	router.POST("/hello", post)
	log.Fatal(http.ListenAndServe(":8000", router)) 
} 

func getSimple (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query().Get("name")
	outName(w, name)
	
}

func getParam (w http.ResponseWriter, _ *http.Request, param httprouter.Params) {
	name := param.ByName("name")
	outName(w, name)
}

func post (w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Please, enter the name in request", http.StatusBadRequest)
		logrus.Info("Got request without name!")
		return
	}
	name := string(body)
	outName(w, name)
}

func outName (w http.ResponseWriter, name string){
	// defense of blank lines
	strings.Replace(name, " ", "", -1)
	if strings.EqualFold(name, "") {
		http.Error(w, "Please, enter the name in request", http.StatusBadRequest)
		logrus.Info("Got request without name!")
		return
	}
	fmt.Fprintf(w, "Hello, %v!\n", name) 
	logrus.Infof("Got name, %v", name)
	

}

