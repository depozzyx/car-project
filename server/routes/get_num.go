package routes

import (
	"encoding/json"
	"net/http"
)

type GetNumResp struct {
	Status string
	Numbers []int64
}

func GetNum(w http.ResponseWriter, req *http.Request){
	status := Status.ok

	w.Header().Set("Content-Type", "application/json")
	p := GetNumResp{status, Nums}
	json.NewEncoder(w).Encode(p)
}