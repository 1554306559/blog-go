package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "go"
	indexData.Desc = "hello world"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func main() {
	//程序入口，一个程序只有一个入口
	server := http.Server{
		Addr: "127.0.0.1:8888",
	}
	http.HandleFunc("/", indexHandler)
	if err := server.ListenAndServe(); err != nil {
		log.Panicln(err)
	}
}
