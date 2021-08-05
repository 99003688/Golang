package main

import (
	"log"
	"net/http"
	"project2/showapi"

)

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",showapi.Index)
	log.Fatal(http.ListenAndServe(":8084",mux))
}