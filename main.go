package main

import (
	"fmt"
)

func main()  {


	ch := make(chan int,0)
	go func() {
		defer close(ch)
		fmt.Printf("s")
		ch<-1
	}()

	for i:=range ch{
		fmt.Println(i)
	}

	//Commands.Parse()

}
