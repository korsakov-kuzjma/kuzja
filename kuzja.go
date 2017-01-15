package kuzja

import (
	"net/http"
	"fmt"
)

func Wrap(fw ...func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, ffw := range fw {
			ffw(w, r)
		}
	}
}
func Tess(){
	fmt.Println("ohoho")
}
