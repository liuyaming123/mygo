package myInterface

import "fmt"

type code string

type Programmer interface {
	WriteHelloWorld() code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() code {
	return "fmt.Println(\"Hello World!\")"
}

func (j *JavaProgrammer) WriteHelloWorld() code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}
func MyIf() {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
