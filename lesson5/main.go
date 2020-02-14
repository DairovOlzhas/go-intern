package main


import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"

)

var (
	port = "8080"
)

//func createFileHandler(w http.ResponseWriter, r *http.Request) {
//
//	out, err := exec.Command("ls", "-a").Output()
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(string(out))
//
//	jsonAnswer, err := json.Marshal(out)
//	if err != nil {
//		panic(err)
//	}
//
//	w.Write(jsonAnswer)
//
//}


func main(){
	r := mux.NewRouter()

	r.HandleFunc("/createFile", createFileHandler)
	http.Handle("/", r)
	fmt.Println("Server started")
	http.ListenAndServe(":"+port, r)
}
