package main

type Location struct {
	Start    int
	End      int
	Filename string
}

type File struct {
	Name       string
	Expression Term
	Location   Location
}

type Term interface{}

func (i *Int) isTerm() {}

func main() {

}
