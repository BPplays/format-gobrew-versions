package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type semver struct {
	major int
	minor int
	patch int
}

func containsNumeric(input string) bool {
	// Define a regular expression pattern that matches anything other than 0-9 or .
	pattern := "^[0-9]+\\.[0-9]+\\.[0-9]+$"

	// Compile the regular expression
	regexp := regexp.MustCompile(pattern)

	// Use the regular expression to check if the input contains any non-numeric characters
	return regexp.MatchString(input)
}

func string_to_semver(s string) (semver) {
	if !containsNumeric(s) {
		log.Fatal("non semver", s)
	}

	// output := semver{0, 0, 0}

	str_sl := strings.Split(s, ".")
	int_sl := [3]int{0, 0, 0}

    for i := 0; i <= 2; i++ {
		in, err := strconv.Atoi(str_sl[i])
		if err != nil {
			log.Fatal("non semver", err)
		}
		int_sl[i] = in
    }



	// fmt.Println(int_sl[0], int_sl[1], int_sl[2])
	return semver{int_sl[0], int_sl[1], int_sl[2]}
}


func gobrew_lr() (string) {


	cmd := exec.Command("gobrew", "ls-remote", ) // "sed \"s,\\x1B\\[[0-9;]*[a-zA-Z],,g\"", "|", "tac"


	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), err)
		log.Fatal(string(output), err)
	}

	escaped := string(output)
	regex := regexp.MustCompile("\x1b[^m]*m")
	cleaned := regex.ReplaceAllString(escaped, "")

	
	// fmt.Println(string(output))
	return cleaned
}


func reverseSlice(inputSlice interface{}) interface{} {
	sliceValue := reflect.ValueOf(inputSlice)

	if sliceValue.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	length := sliceValue.Len()
	reversed := reflect.MakeSlice(sliceValue.Type(), length, length)

	for i := 0; i < length; i++ {
		reversed.Index(length-i-1).Set(sliceValue.Index(i))
	}

	return reversed.Interface()
}


func gobrew_parse(s string) ([]semver) {
	// Define a regular expression for splitting by comma
	// regex := regexp.MustCompile("(?m)^[A-Za-z0-9.]")

	// Use the Split function to split the string
	result := []semver{}
	fields := strings.Fields(s)

	for _, field := range fields {
		if field != "" {
			if containsNumeric(field) {

				result = append(result, string_to_semver(field))
				// fmt.Printf("-%v_\n", field)
			} // else {
			// 	// fmt.Printf("-%v_\n", field)
			// }
		}

	}


	result = reverseSlice(result).([]semver)
	return result

}




func gobrew_filter(sem []semver) ([]semver) {
	// Define a regular expression for splitting by comma
	// regex := regexp.MustCompile("(?m)^[A-Za-z0-9.]")

	// Use the Split function to split the string
	result := []semver{}


	last_maj := -1

	last_min := -1

	for _, s := range sem {

		if last_maj != s.major {
			result = append(result, s)
			last_maj = s.major
			last_min = s.minor
		} else {
			if last_min != s.minor {
				result = append(result, s)
				last_min = s.minor
			}
		}





	}


	result = reverseSlice(result).([]semver)
	return result

}



func main() {

	fmt.Println(gobrew_filter(gobrew_parse(gobrew_lr())))
	
}