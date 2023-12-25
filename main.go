package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func gobrew_lr() (string) {


	cmd := exec.Command("gobrew", "ls-remote")


	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), err)
		log.Fatal(string(output), err)
	}

	// fmt.Println(string(output))
	return string(output)
}


func gobrew_parse(s string) ([]string) {
	// Define a regular expression for splitting by comma
	regex := regexp.MustCompile("(?m)^[A-Za-z0-9.]")

	// Use the Split function to split the string
	result := regex.Split(s, -1)

	return result

}




func main() {

	fmt.Println(gobrew_parse(gobrew_lr()))
	
}