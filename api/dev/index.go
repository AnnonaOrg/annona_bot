package dev

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodGet {
	// 	return
	// }
	// params := r.URL.Query()
	// userId, _ := strconv.Atoi(params.Get("to"))
	// msgText := params.Get("m")
	// groupChannel := r.Param("channel")

	checkToken := r.Header.Get("Apiclient")
	log.Println("checkToken", checkToken)

	// body, err := io.ReadAll(r.Body)
	// common.Must(err)
	// log.Println(string(body))
	// fmt.Fprintf(w, "ok")
}
