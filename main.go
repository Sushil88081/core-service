package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{
		Name: "sushil",
		Age:  20,
	}
	fmt.Println("main func", u.Name)

}
