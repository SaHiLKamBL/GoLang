package main

import "fmt"

func makemap(id int, name string, age int) map[string]any {
	user := make(map[string]any)

	user["id"] = id
	user["name"] = name
	user["age"] = age

	return user
}

func Q3() {

	// outer map → key = id, value = user map
	users := make(map[int]map[string]any)

	users[123] = makemap(123, "sahil", 23)
	users[145] = makemap(145, "Aum", 22)
	users[156] = makemap(156, "siddhant", 25)
	users[567] = makemap(567, "modi", 27)
	users[231] = makemap(231, "Trump", 67)

	// print full nested map
	fmt.Println(users)
}