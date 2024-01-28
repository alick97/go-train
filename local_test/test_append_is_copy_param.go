package main

import "fmt"

type s struct {
	Name string
}

func show1(sList []s) {
	fmt.Printf("address in show1, addr: %p\n", &sList[0])
}

func show2(sParam ...s) {
	fmt.Printf("address in show2, addr: %p\n", &sParam[0])
}

func test_call_append_param_by_coping() {
	// Does ... param is trans by copy when call func?
	// conclusion: no, by reference, like trans param as slice
	sList := []s{{Name: "n1"}}
	fmt.Printf("address in test_call, addr: %p\n", &sList[0])
	show1(sList)
	show2(sList...)
}

func main() {
	test_call_append_param_by_coping()
}
