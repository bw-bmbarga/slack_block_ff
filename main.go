/* ************************************************************************** */
/*                                                                            */
/*  main.go                                                                   */
/*                                                                            */
/*   By: elhmn <www.elhmn.com>                                                */
/*             <nleme@live.fr>                                                */
/*                                                                            */
/*   Created:                                                 by elhmn        */
/*   Updated: Fri Mar 22 20:37:52 2019                        by bmbarga      */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	cPort      = "4390"
	cBlockPath = "./res/button_block.json"
)

type sAction struct {
	ActionID int    `json:"action_id"`
	BlockID  string `json:"block_id"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am root"))
}

func answerInteractive(actions interface{}, url interface{}) {
	time.Sleep(5 * time.Second)
	var buf bytes.Buffer
	message := struct {
		Text string `json:"text"`
	}{Text: "Message received"}

	json.NewEncoder(&buf).Encode(message)
	http.Post(url.(string), "application/json", &buf)
}

func interactions(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	text := r.FormValue("text")
	if text == "" {
		text = "No data sent !"
	}

	payload := r.Form.Get("payload")

	var tmp map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &tmp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//start response go routine
	{
		actions := tmp["actions"]
		url := tmp["response_url"]
		go answerInteractive(actions, url)
	}
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
	http.HandleFunc("/interactions", interactions)
	if err := http.ListenAndServe(":"+cPort, nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
