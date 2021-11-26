package main

import "fmt"

func main() {
	var hello string
	var a, b, c string
	var (
        name string = "Tom"
        age int = 27
    )
    name2 := "Jeck"
	hello = "string1"
	a,b,c = "string2","string3","string4"
	fmt.Println(hello,a,b,c)
	fmt.Println(name,age)
	fmt.Println(name2)
}
