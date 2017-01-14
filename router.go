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

// обработка запроса с параметрами в строке, например
// GET http://127.0.0.1:8000/hello\?name\=Def
func getSimple (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.URL.Query().Get("name")
	outName(w, name)
	
}

// обработка запроса с параметрами в параметрах роутера, например
// GET http://127.0.0.1:8000/hello/Piu
func getParam (w http.ResponseWriter, _ *http.Request, param httprouter.Params) {
	name := param.ByName("name")
	outName(w, name)
}

// обработка запроса с параметрами в его теле, например
// POST -d "" http://127.0.0.1:8000/hello
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
	// избавляемся от пустых строк
	strings.Replace(name, " ", "", -1)
	if strings.EqualFold(name, "") {
		http.Error(w, "Please, enter the name in request", http.StatusBadRequest)
		logrus.Info("Got request without name!")
		return
	}
	fmt.Fprintf(w, "Hello, %v!\n", name) 
	logrus.Infof("Got name, %v", name)
	

}

