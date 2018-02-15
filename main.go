package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func doSomethingHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var localStr string
	json.NewDecoder(r.Body).Decode(&struct {
		*string `json:"jsonParam"`
	}{&localStr})

	fmt.Printf("localStr: %s", localStr)

	if localStr == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("localStr empty"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("localStr not empty"))
	return
}
func main() {
	router := httprouter.New()
	router.POST("/dosomething", doSomethingHandler)
	http.ListenAndServe(":8000", router)
}
