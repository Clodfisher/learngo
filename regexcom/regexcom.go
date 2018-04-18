package main

import (
	"fmt"
	"regexp"
)

const (
	email = "My email is HunterYuan@gmail.com" +
		" My email is juju@gamil.com" +
		" My email is HunterYuan@gmail.cm" +
		" My email is HunterYuan@gmail.com.cn"
	regexpFormat = "([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\\.([a-zA-Z0-9.]+)"
)

func main() {
	re := regexp.MustCompile(regexpFormat)
	amatch := re.FindAllStringSubmatch(email, -1)

	for _, e := range amatch {
		fmt.Printf("%v\n", e)
	}
}
