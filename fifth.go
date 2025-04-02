package main

import "fmt"

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("делитель равен 0")
	}
	return a / b, nil
}

func main() {
	d, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("d = %d", d)
	}
}
