package main

type User struct {
	Name string
}

func main() {
	age := 18
	canDrink(&age)
}

func canDrink(age *int) bool {
	return *age >= 18
}

func getAge() *int {
	age := 18
	return &age
}
