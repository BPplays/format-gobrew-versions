package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type semver struct {
	major int
	minor int
	patch int
}

func containsNonNumeric(input string) bool {
	// Define a regular expression pattern that matches anything other than 0-9 or .
	pattern := "[^0-9.]"

	// Compile the regular expression
	regexp := regexp.MustCompile(pattern)

	// Use the regular expression to check if the input contains any non-numeric characters
	return regexp.MatchString(input)
}

func string_to_semver(s string) (semver) {
	if containsNonNumeric(s) {
		log.Fatal("non semver")
	}

	// output := semver{0, 0, 0}

	str_sl := strings.Split(s, ".")
	int_sl := []int{0, 0, 0}

    for i := 0; i < 2; i++ {
		in, err := strconv.Atoi(str_sl[i])
		if err != nil {
			log.Fatal("non semver")
		}
		int_sl[in] = in
    }



	// fmt.Println(string(output))
	return semver{int_sl[0], int_sl[1], int_sl[2]}
}


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


func gobrew_parse(s string) ([]semver) {
	// Define a regular expression for splitting by comma
	// regex := regexp.MustCompile("(?m)^[A-Za-z0-9.]")

	// Use the Split function to split the string
	result := []semver{}
	fields := strings.Fields(s)

	for _, field := range fields {
		result = append(result, string_to_semver(field)) 
	}

	return result

}




func main() {

	fmt.Println(gobrew_parse(gobrew_lr()))
	
}