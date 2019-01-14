package main

import(
	"fmt"
)

var pipe chan int

func add(a int, b int)  { 
	sum := a + b
	pipe <- sum
}

func main()  {
	pipe = make(chan int, 3)
	go add(100, 200) 
	fmt.Println(<- pipe)
	fmt.Println("hello world")
}