package main

type Person struct {
	Name     string
	Age      int
	Location string
	Wealth   float64
	Career   string
}
type PersonBuilder interface {
	SetName(name string) PersonBuilder
	SetAge(n int) PersonBuilder
	SetLocation(location string) PersonBuilder
	SetWealth(wealth float64) PersonBuilder
	SetCareer(career string) PersonBuilder
	Build() *Person
}

func (p *Person) SetName(name string) PersonBuilder {
	p.Name = name
	return p
}

func (p *Person) SetAge(age int) PersonBuilder {
	p.Age = age
	return p
}
func (p *Person) SetLocation(name string) PersonBuilder {
	p.Location = name
	return p
}
func (p *Person) SetWealth(amount float64) PersonBuilder {
	p.Wealth = amount
	return p
}
func (p *Person) SetCareer(name string) PersonBuilder {
	p.Career = name
	return p
}

func (p *Person) Build() *Person {
	return p
}
func (p *Person) NewPerson(builder PersonBuilder) Person {
	return *builder.SetName("Felicity Smoke").SetAge(28).SetLocation("Carifonia").SetWealth(50000000).Build()
}
func main() {
	// person := Person{}
	// builder := PersonBuilder{}

	// newperson =
}
