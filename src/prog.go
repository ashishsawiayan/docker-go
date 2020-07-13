package main

import (
    "fmt"
    "log"
    "net/http"
	"strings"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str1 :=r.URL.Path[6:]
	res1 := strings.Split(str1, "/")
	num1 :=0
	num2 :=0
	err :=""
	ans :=0
	if res1[0]=="sum"{
	num1,err := strconv.ParseFloat(res1[1], 32)
	num2,err := strconv.ParseFloat(res1[2], 32)
	ans :=num1+num2
	fmt.Fprintf(w, "%f!", ans)
		_ = err
		_ = ans
	}
	if res1[0]=="mul"{
	num1,err := strconv.ParseFloat(res1[1], 32)
	num2,err := strconv.ParseFloat(res1[2], 32)
	ans :=num1*num2
	fmt.Fprintf(w, "%f!", ans)
		_ = err
		_ = ans
	}
	if res1[0]=="div"{
	num1,err := strconv.ParseFloat(res1[1], 32)
	num2,err := strconv.ParseFloat(res1[2], 32)
	ans :=num1/num2
	fmt.Fprintf(w, "%f!", ans)
		_ = err
		_ = ans
	}
	if res1[0]=="sub"{
	num1,err := strconv.ParseFloat(res1[1], 32)
	num2,err := strconv.ParseFloat(res1[2], 32)
	ans :=num1-num2
	fmt.Fprintf(w, "%f!", ans)
		_ = err
		_ = ans
	}
    
	
	_ = err
	_ = ans
	_ = num1
	_ = num2
}

func main() {
    http.HandleFunc("/calc/", handler)
    log.Fatal(http.ListenAndServe(":5000", nil))
}
