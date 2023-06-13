package main

//import "fmt"

func main() {
	d := newDeck()
	// fmt.Println(d)
	// d.print()
	d.shuffle1()
	// fmt.Println("###############")
	d.print()
	//fmt.Println("###############")
	d.shuffle2()
	//d.print()
	//fmt.Println("###############")
	d.shuffle3()
	//d.print()
	d.saveToFile()
	d.appendToFile()
	d.readFromFile()
}
