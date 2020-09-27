package main

import "fmt"


type duration int

func (d duration) pretty() string {
	return fmt.Sprintf("Duration: %d", d)
}

func main() {
	fmt.Println("Hello World!")
	fmt.Println("dd")








	duration(42).pretty()




	colors := map[string]string{}

	colors["red"] = "#da1337"
	fmt.Println(colors)
}
