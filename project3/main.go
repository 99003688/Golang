package main

import (
	handlers "project3/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/addnewuser/", handlers.AddNewUserFunc)
	mux.HandleFunc("/notsucceded", handlers.NotSucceded)

	mux.HandleFunc("/deleted", handlers.DeletedFunc)
	mux.HandleFunc("/deleteuser/deleted", handlers.DeleteUserFunc)
	mux.HandleFunc("/deleteuser/", handlers.DeleteUserServe)
	mux.HandleFunc("/deleteuser/notsuccededdelete", handlers.NotSuccededDelete)

	mux.HandleFunc("/", handlers.IndexFunc)

	mux.HandleFunc("/showuser/show", handlers.ShowUserFunc)
	mux.HandleFunc("/showuser/", handlers.ShowUser)
	mux.HandleFunc("/showuser/notsuccededshow/", handlers.NotSuccededShow)

	log.Fatal(http.ListenAndServe(":8080",mux))
}
}