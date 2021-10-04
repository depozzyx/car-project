package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
)


type MinResponse struct {
	Minus int64
}

func Min(w http.ResponseWriter, req *http.Request){
	query := req.URL.Query()

	num1, _ := strconv.ParseInt(query.Get("n1"), 10, 64)
	num2, _ := strconv.ParseInt(query.Get("n2"), 10, 64)

	n := num1 - num2

	w.Header().Set("Content-Type", "application/json")
	p := MinResponse{n}
	json.NewEncoder(w).Encode(p)
}

