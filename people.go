package main

type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Parent_id int    `json:"parent_id"`
}

var people []Person = []Person{
	{
		Id:        1,
		Name:      "zaidy",
		Parent_id: 0,
	},
	{
		Id:        2,
		Name:      "dad",
		Parent_id: 1,
	},
	{
		Id:        3,
		Name:      "mum",
		Parent_id: 1,
	},
	{
		Id:        4,
		Name:      "me",
		Parent_id: 1,
	},
	{
		Id:        5,
		Name:      "berel",
		Parent_id: 1,
	},
	{
		Id:        6,
		Name:      "baila",
		Parent_id: 1,
	},
	{
		Id:        7,
		Name:      "motty",
		Parent_id: 1,
	},
	{
		Id:        8,
		Name:      "yehudis hecht",
		Parent_id: 5,
	},
}

func (p *Person) GetParent() *Person {
	for _, person := range people {
		if person.Id == p.Parent_id {
			return &person
		}
	}
	return nil
}

func (p *Person) GetChildren() []Person {
	var children []Person
	for _, person := range people {
		if person.Parent_id == p.Id {
			children = append(children, person)
		}
	}
	return children
}
