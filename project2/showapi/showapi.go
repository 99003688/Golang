package showapi

import (
	"encoding/json"
	"fmt"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"project2/entity"
	// "text/template"
)

func Index(response http.ResponseWriter, request *http.Request) {

	
	
	req, err := http.Get("https://covidtracking.com/api/states")
	if err != nil {
		log.Fatal(err)
	}


	
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)

	}

	var responseobject entity.Response
	json.Unmarshal(bodyBytes, &responseobject)
	maal:=string(bodyBytes)
	fmt.Fprintf(response,"%+v\n",maal)
	tmp, _ := template.ParseFiles("view/view.html")
	tmp.Execute(response, nil)

}
