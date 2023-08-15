package main

type Human struct {
	Name string
	Age  int
}

func (h Human) getAge() {
	print(h.Age)
}

type Action struct {
	Human
}

func main() {
	age := Action{Human: Human{Name: "Sergey", Age: 42}}
	age.getAge()
}
