package server

import (
	"fmt"
	"net/http"
)

func Http(addr string) error {
	fmt.Println("Server listen on", addr)
	http.HandleFunc("/qth", qth)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil // never happens
}
