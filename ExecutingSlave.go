package main

import (
	"fmt"
	"net/http"
	//"io/ioutil"
	)


func main () {
	resp,_:= http.Get("http://127.0.0.1:5000/command")
	fmt.Println(resp["Body"])

}



