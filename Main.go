package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"net/http"
	"html/template"
	"github.com/waspnesser/sURL/surl_engine"
)

var DomainName = "localhost:8080"
var project_root = os.Getenv("GOPATH")+"/src/github.com/waspnesser/sURL"
var form_template = template.Must(template.ParseFiles(project_root + "/templates/form.html"))


func getTimeCoefficient(unitTime string) int{
	switch unitTime{
	case "Seconds":
		return 1
	case "Minutes":
		return 60
	case "Hours":
		return 60*60
	default :
		return 1
	}

}

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
				var url string		=  r.FormValue("url")
				ttl,_			:=  strconv.Atoi(r.FormValue("ttl"))
				var unitTime string	=  r.FormValue("unit-time")
				ttl *= getTimeCoefficient(unitTime)
				surl := surl_engine.AddURL(url,ttl)
				surl_st := struct{SURL string;Domain string}{surl.String(),DomainName}
				response_template,_ := template.ParseFiles(project_root+"/templates/response.html")
				response_template.Execute(w,surl_st)

			default:
				fmt.Fprintf(w,"Sorry, only GET and POST methods are supported.")
		}
	} else{
		surl := r.URL.Path[1:]
		url,err  := surl_engine.GetURL(surl)
		if err==false{
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
