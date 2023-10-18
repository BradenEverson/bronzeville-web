package main

import (
    "net/http"
    "log"
)

func main(){
    
    http.HandleFunc("/", mainReq)
    http.HandleFunc("/styles", getStyles)

    log.Println("Now listening on port 8000 B)")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
