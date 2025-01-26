package main

/*
In this exercise, you'll build a Computer object that can be customized
with different properties such as CPU, RAM, and storage type. The goal
is to use the Builder pattern to allow for easy construction of a Computer object.
*/

type Computer struct {
	CPU     string
	RAM     int
	Storage string
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

type ComputerBuilder struct {
	Computer
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.RAM = ram
	return b
}

func (b *ComputerBuilder) SetStorage(storage string) *ComputerBuilder {
	b.Storage = storage
	return b
}

func main() {
	builder := NewComputerBuilder()
	computer := builder.SetCPU("Intel Core i7").SetRAM(16).SetStorage("SSD")
	println(computer.CPU, computer.RAM, computer.Storage)
}
