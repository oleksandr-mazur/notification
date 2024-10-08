package main

import (
	"fmt"
	"notification/lucky"
	"notification/study"
)

type Data struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("Your lucky number is:", lucky.LuckyNumber())
	data := study.Check("Ubuntu")
	fmt.Println(data.Hostname)
	study.Convert()
	return

	var my Data
	Add(5, &my)

	fmt.Println(my)
	m := make(map[string]map[int]int)

	s := m["pi"]
	if s == nil {
		s := map[int]int{}
		m["pi"] = s
		s[1] = 1

		fmt.Println(s)

	}
	fmt.Println(m)

}

func Add(id int, data *Data) {
	data.ID = id
}
