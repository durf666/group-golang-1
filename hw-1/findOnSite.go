package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//findOnSite returns a list of sites where found s string
func findOnSite (s string, searchList []string) (res []string) {
	for _, v := range searchList{
		response, err := http.Get(v)
		if err != nil {
			panic(err) // how to handle the error correctly?
		}
		bytesArray, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if strings.Contains(string(bytesArray), s) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	var s string
	fmt.Println("Insert phrase for search and press enter: ")
	fmt.Scan(&s)

	list := []string{"https://ya.ru"}
	fmt.Println(findOnSite(s, list))
}