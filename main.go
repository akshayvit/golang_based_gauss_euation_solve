package main

import (
	"fmt"
	"goalgos/gauss"
	"net/http"
	"errors"
)


func main() {
	http.HandleFunc("/show3",showfor3)
	http.HandleFunc("/show4",showfor4)
	err:=http.ListenAndServe(":8080",nil)
	if errors.Is(err,http.ErrServerClosed) {
		fmt.Println("Server Closed")
	} else if err!=nil {
		fmt.Println("Server functions handlers some issues")
	} 
	fmt.Println("Listening to port 8080...")
}

func showfor3(w http.ResponseWriter, r *http.Request) {
	gauss.Build(3,w)
}

func showfor4(w http.ResponseWriter, r *http.Request) {
	gauss.Build(4,w)
}

