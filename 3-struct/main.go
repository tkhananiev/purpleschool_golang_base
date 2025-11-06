package main

import (
	"3-struct/bins"
	"fmt"
	"time"
)

func main() {

	myBin, err := bins.NewBin("myBin", false, time.Now(), "myBin")
	fmt.Println(myBin, err)
}
