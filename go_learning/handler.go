package main

import(
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"hello")
}

func byehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"bye\n")
	method := r.Method
	path := r.URL.Path
	fmt.Fprintf(w,"Method :%s , Path :%s",method,path)
}

func main() {
	http.HandleFunc("/hello",hellohandler)
	http.HandleFunc("/bye",byehandler)
	http.ListenAndServe(":8080",nil)
}