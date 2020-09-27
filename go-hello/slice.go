package main

import "fmt"

func main() {
	data := [] byte { 'h' , 'e' , 'l' , 'l' , 'o' ,
		',' , ' ' , ' ' , 'o' , 'r' , 'l' , ' ' }
	data = Filter(data, isSpace)
	fmt.Println(data)
}

func isSpace(b byte) bool {
	if b == ' ' {
		return true
	}

	return false
}

func Filter(s []byte, fn func (x byte) bool) [] byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}

	return b
}