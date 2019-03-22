/* ************************************************************************** */
/*                                                                            */
/*  main.go                                                                   */
/*                                                                            */
/*   By: elhmn <www.elhmn.com>                                                */
/*             <nleme@live.fr>                                                */
/*                                                                            */
/*   Created:                                                 by elhmn        */
/*   Updated: Fri Mar 22 13:48:38 2019                        by bmbarga      */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	// 	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	cPort      = "4390"
	cBlockPath = "./res/block.json"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am root"))
}

func analytics(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	text := r.FormValue("text")
	if text == "" {
		text = "No data sent !"
	}

	block, err := ioutil.ReadFile(cBlockPath)
	if err != nil {
		http.Error(w, "Error : "+err.Error(), http.StatusBadRequest)
		return
	}

	data := "{" + "\"blocks\"" + ":" + string(block) + "}"

	w.Write([]byte(data))
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
