package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	a := make(map[string]person)
	a["kevin"] = person{"kevin", 23}
	a["dan"] = person{"dan", 11}

	b := [1]person{
		{"joe", 11},
	}

	var c []person
	c = append(c, person{"bob", 12})

	d := map[string]person{
		"lai": {"lai", 21},
	}

	var e [3]person
	e[0].name = "alen"
	e[0].age = 40

	f := []person{}
	f = append(f, person{"zen", 12})

	aa := make([][]person, 3)
	for i := range aa {
		aa[i] = make([]person, 2)
		aa[i][0] = person{"JJ", 33 + i}
	}

	var bb [][]person
	bb = append(bb, []person{
		{"jason", 1},
		{"frank", 6},
		{"wolf", 6},
	})

	bb = append(bb, []person{
		{"apple", 1},
		{"angle", 6},
	})

	var cc [3][]interface{}
	for i := range cc {
		cc[i] = make([]interface{}, 2)
		cc[i][0] = person{"FF", 33 + i}
	}

	var dd [3][]person
	for i := range dd {
		dd[i] = make([]person, 2)
		dd[i][0] = person{"FF", 33 + i}
	}

	var ff [][]interface{}

	s := make([]interface{}, len(e))
	for i, v := range e {
		s[i] = v
	}

	ff = append(ff, s)

	x := 2
	y := 4
	z := 3
	table := make([][][]int, x)
	for i := range table {
		table[i] = make([][]int, y)
		for j := range table[i] {
			table[i][j] = make([]int, z)
		}
	}

	var fff interface{}
	ddd := 10
	fff = fmt.Sprintf("%v", ddd) + "fbdb"

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)
	fmt.Printf("f: %v\n", f)

	fmt.Printf("aa: %v\n", aa)
	fmt.Printf("bb: %v\n", bb)
	fmt.Printf("cc: %v\n", cc)
	fmt.Printf("dd: %v\n", dd)
	fmt.Printf("ff: %v\n", ff)
	fmt.Printf("table: %v\n", table)

	fmt.Printf("fff: %v\n", fff)
	fmt.Printf("ddd: %v\n", ddd)

}
