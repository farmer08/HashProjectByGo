package main

import (
	"HashProjectByGo/HMap"
	"fmt"
)

func main() {
	HMap.InitBuckets()

	HMap.AddKeyValye("go","welcome")
	HMap.AddKeyValye("node","wow")

	fmt.Println(HMap.GetValueByKey("go"))
	fmt.Println(HMap.GetValueByKey("node"))

}
