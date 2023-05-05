package data

type User struct {
	id   int
	name string
	age  int
}

var users = []User{
	{id: 1, name: "user1", age: 20},
	{id: 2, name: "user2", age: 20},
	{id: 3, name: "user3", age: 20},
}
