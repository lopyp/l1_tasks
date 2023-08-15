package main

import "fmt"

// wелевой интерфейс
type Target interface {
	Request() string
}

// адаптируемый класс, который имеет несовместимый интерфейс
type Adaptee struct {
	SpecificRequest string
}

// метод адаптируемого класса
func (a *Adaptee) GetSpecificRequest() string {
	return a.SpecificRequest
}

// класс адаптера, который преобразует интерфейс Adaptee в Target
type Adapter struct {
	Adaptee *Adaptee
}

// метод адаптера, который преобразует вызовы к Adaptee
func (a *Adapter) Request() string {
	return a.Adaptee.GetSpecificRequest()
}

func main() {
	adaptee := &Adaptee{SpecificRequest: "Specific Request"}
	fmt.Println("Adaptee: " + adaptee.GetSpecificRequest())

	adapter := &Adapter{Adaptee: adaptee}
	fmt.Println("Adapter: " + adapter.Request())
}
