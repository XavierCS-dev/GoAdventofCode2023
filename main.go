package main

import (
	"github.com/XavierCS-dev/GoAdventofCode2023/day1"
    "github.com/XavierCS-dev/GoAdventofCode2023/day2"
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
	http.HandleFunc("/day_two/part_one", serve)
	socket := "127.0.0.1:8080"
	fmt.Printf("Server: http://%v\n", socket)
	error := http.ListenAndServe(socket, nil)
	if error != nil {
		log.Fatal(error)
	}
}

func serve(write http.ResponseWriter, request *http.Request) {
    switch request.URL.Path {
        case "/day_one/part_one" :
            io.WriteString(write, strconv.Itoa(day1.PartOne()))
        case "/day_one/part_two":
            io.WriteString(write, strconv.Itoa(day1.PartTwo()))
        case "/day_two/part_one":
            io.WriteString(write, strconv.Itoa(day2.PartOne()))
        default:
            io.WriteString(write, "AoC Challenge 2023")
    }

}


