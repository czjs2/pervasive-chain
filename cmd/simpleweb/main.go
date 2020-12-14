package main

import (
	"fmt"
	"net/http"
	"time"
)

func DeployHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("this is a error：%s  %s \n", r, time.Now())
		}
	}()
	fmt.Printf("incoming  %v  \n", time.Now())
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
	return

}

func main() {
	http.HandleFunc("/xjrwTest", DeployHandler)
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
