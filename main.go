package main

import (
	"fmt"
	"github.com/XavierCS-dev/GoAdventofCode2023/day1"
	"github.com/XavierCS-dev/GoAdventofCode2023/day2"
	"github.com/XavierCS-dev/GoAdventofCode2023/day3"
	"github.com/XavierCS-dev/GoAdventofCode2023/day4"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("AoC Day 1")
	http.HandleFunc("/", serve)
	socket := "127.0.0.1:8080"
	fmt.Printf("Server: https://%v\n", socket)
	error := http.ListenAndServeTLS(socket, "localhost.crt", "localhost.key", nil)
	if error != nil {
		log.Fatal(error)
	}
}

func serve(write http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/1/1":
		io.WriteString(write, strconv.Itoa(day1.PartOne()))
	case "/1/2":
		io.WriteString(write, strconv.Itoa(day1.PartTwo()))
	case "/2/1":
		io.WriteString(write, strconv.Itoa(day2.PartOne()))
	case "/2/2":
		io.WriteString(write, strconv.Itoa(day2.PartTwo()))
	case "3/1":
		io.WriteString(write, strconv.Itoa(day3.PartOne()))
	case "/4/1":
		io.WriteString(write, strconv.Itoa(day4.PartOne()))
	default:
		io.WriteString(write, "AoC Challenge 2023")
	}
}
