/* ************************************************************************** */
/*                                                                            */
/*  main.go                                                                   */
/*                                                                            */
/*   By: elhmn <www.elhmn.com>                                                */
/*             <nleme@live.fr>                                                */
/*                                                                            */
/*   Created:                                                 by elhmn        */
/*   Updated: Fri Mar 22 12:38:50 2019                        by bmbarga      */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	cPort = "4390"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am root"))
}

func analytics(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	if text == "" {
		text = "No data sent !"
	}

	w.Write([]byte(text))
}

func run() {
	fmt.Println("Server started at port :", cPort)
	http.HandleFunc("/", root)
	http.HandleFunc("/analytics", analytics)
	if err := http.ListenAndServe(":"+cPort, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
