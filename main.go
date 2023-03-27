package main

import (
	"fmt"
	
	"net/http"
	"os"
	"strings"
)
func read_file(file_path string) []string {
  content, err := os.ReadFile(file_path)

  if( err != nil ) {
    panic("File does not exist")
  }

  return strings.Split(string(content), "\n")
}


func main() {
  words := read_file("./words.txt")
  var url string = "http://localhost:3333/"
  
  conn, err := http.Get(url)
  if( err != nil )  {
    fmt.Println(conn)
    os.Exit(1)
  } else {

    tr := &http.Transport{
      MaxIdleConns: 10,
      IdleConnTimeout: 30 * 1000,
      DisableCompression: true,
    }
    client:= &http.Client{Transport: tr}



    for _, value := range(words) {
      resp , erro  := client.Get(url + string(value))
      if ( erro != nil ) {
        fmt.Println("Error")
      } else {
        if( resp.StatusCode != 404 ) {
          fmt.Printf("/%v => (%v code)\n", string(value), resp.StatusCode)
        }
      }
    }
  }
}
