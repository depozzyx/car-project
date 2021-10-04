package server

import (
	"fmt"
	"hello/server/routes"
	"net/http"
	"strconv"
)

var Port int64 = 8000

func Server() {
	for {
		http.HandleFunc("/", routes.Index)
		http.HandleFunc("/api/sum", routes.Sum)
		http.HandleFunc("/api/min", routes.Min)
		http.HandleFunc("/api/get_num", routes.GetNum)
		
		PortString := ":" + strconv.FormatInt(Port, 10)

		fmt.Println("Running server on 0.0.0.0" + PortString)
		err := http.ListenAndServe(PortString, nil)

		if err != nil{
			fmt.Println("Error: " + err.Error())
		}
	}
}