package main

import "fmt"

type User struct {
	Name string
}

func changeName(u User) {
	u.Name = "Mateus"
}

func changeName2(u *User) {
	u.Name = "Mateus"
}

type User2 struct {
	Scores []int
}

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value++
}

func main() {
	fmt.Println("Hello")

	u1 := User{Name: "Hiraeth"}
	changeName(u1)
	fmt.Println(u1)

	changeName2(&u1)
	fmt.Println(u1)

	fmt.Println("Slices")
	u2 := User2{Scores: []int{10, 20}}
	u3 := u2
	u3.Scores[0] = 999
	fmt.Println(u2.Scores)

	fmt.Println("Mais exemplos")
	c := Counter{}
	c.Increment()
	fmt.Println(c)
}

/*
	Pointers em Go

	FUNDAMENTO 1: TUDO EM GO É PASSADO POR VALOR

	- Go sempre passa cópia
	- Se é pra compartilhar, você passa um pointer

	----------------- Exemplo em Java

	User u1 = new User("Ana");
	User u2 = u1;
	u2.name = "Carlos";

	(O nome de "u1" VAI SER SOBRESCRITO)

	----------------- Exemplo em Go

	type User struct {
    	Name string
	}

	u1 := User{Name: "Ana"}
	u2 := u1
	u2.Name = "Carlos" ()

	("u2" vira UMA CÓPIA. Nada é sobrescrito)

	FUNDAMENTO 2: SLICES SÃO DIFERENTES...


*/
