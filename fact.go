package main

import (
    "fmt"
    "log"
	"net/http"
	"strconv"
)
func fact (n int) int {
	if n == 0 {
		return 1
	}
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	} 
	return res
}
func factHandler(w http.ResponseWriter, r *http.Request) {
	in := r.URL.Query().Get("n")
	n, _ := strconv.Atoi(in)

	result := fact(n)
	fmt.Fprintf(w, "%d ! =  %d", n, result);
}

func main() {
	http.HandleFunc("/fact", factHandler)
	fmt.Printf("Starting server at port 8080\n")

    log.Fatal(http.ListenAndServe(":8080", nil))
}