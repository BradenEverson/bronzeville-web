package main

import (
    "net/http"
    "log"
)

func main(){
    //Http route section
    //Each http.HandleFunc method takes in the route along with the function used to return the htmx needed
    //Handler functions are found at the handlers.go file on the root directory
    http.HandleFunc("/", mainReq)
    http.HandleFunc("/styles", getStyles)

    //Set up the actual http server and listen on a specified port
    //TODO dynamically change port so that heroku can automatically hand one over :)
    log.Println("Now listening on port 8000 B)")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
