package main

import "fmt"

func main() {
	opaPDP := &OpaPDP{policy: "test"}
	fmt.Println("OpaPDP policy:", opaPDP.policy)
}
