package routes

type status_type struct {
	ok string
	err string
}
var Status = status_type{"Ok", "Error"}

var Nums = make([]int64, 0)

