package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func DeployHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("this is a error：%s  %s \n", r,time.Now())
		}
	}()
	// todo 执行shell命令
	cmd := exec.Command("/bin/bash", "-c", "sh deploy.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
	fmt.Println("received deploy msg ... ", time.Now())
	return

}

func main() {
	http.HandleFunc("/deploy", DeployHandler)
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
