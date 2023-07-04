package mens

var people = man{}

type man struct {
	name     string
	lastName string
	age      int
	gender   string
	crimes   int
}

func NewMan(name string, lastName string, age int, gender string, crimes int) man {
	people.name = name
	people.lastName = lastName
	people.age = age
	people.gender = gender
	people.crimes = crimes
	return people
}

func (man *man) Name() string {
	return man.name
}

func (man *man) LastName() string {
	return man.lastName
}

func (man *man) Age() int {
	return man.age
}

func (man *man) Gender() string {
	return man.gender
}

func (man *man) Crimes() int {
	return man.crimes
}

type Niggers interface {
	Name() string
	LastName() string
	Age() int
	Gender() string
	Crimes() int
}
