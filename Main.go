package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"html/template"
	"github.com/waspnesser/sURL/surl_engine"
)

var DomainName = "localhost:8080"
var project_root = os.Getenv("GOPATH")+"/src/github.com/waspnesser/sURL"
var form_template = template.Must(template.ParseFiles(project_root + "/templates/form.html"))

func handler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path == "/"{
		switch r.Method{
			case "GET":
				form_template.Execute(w,nil)
			case "POST":
				if err:= r.ParseForm(); err != nil{
					fmt.Fprintf(w,"ParseForm() err %v",err)
					return
				}
				var url string =  r.FormValue("url")
				surl := surl_engine.AddURL(url)
				surl_st := struct{SURL string;Domain string}{surl.String(),DomainName}
				response_template,_ := template.ParseFiles(project_root+"/templates/response.html")
				response_template.Execute(w,surl_st)

			default:
				fmt.Fprintf(w,"Sorry, only GET and POST methods are supported.")
		}
	} else{
		surl := r.URL.Path[1:]
		url,ok  := surl_engine.GetURL(surl)
		if ok{
			http.Redirect(w,r,url,http.StatusSeeOther)
		} else{
			http.Error(w,"404 Not Found.",http.StatusNotFound)
			return
		}
	}
}

func main(){


	http.HandleFunc("/", handler)


	log.Println("[+] Starting server at localhost:8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
