package main

import (
	"github.com/XavierCS-dev/GoAdventofCode2023/day_1"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1")
	http.HandleFunc("/", serve)
	http.HandleFunc("/day_one/part_one", serve)
	http.HandleFunc("/day_one/part_two", serve)
	socket := "127.0.0.1:8080"
	fmt.Printf("Server: http://%v\n", socket)
	error := http.ListenAndServe(socket, nil)
	if error != nil {
		log.Fatal(error)
	}
}

func serve(write http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/day_one/part_one" {
		io.WriteString(write, strconv.Itoa(day_1.Part_one()))
	} else if request.URL.Path == "/day_one/part_two" {
        io.WriteString(write, strconv.Itoa(day_1.Part_two()))
    } else {
		io.WriteString(write, "AoC Challenge 2023")
	}
}


