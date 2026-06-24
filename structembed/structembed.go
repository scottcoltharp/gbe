package structembed

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func (c container) describe() string {
	return fmt.Sprintf("container with num=%v and str=%v", c.num, c.str)
}

type container struct {
	base
	str string
}

func StructEmbed() {

	co := container{
		base: base{
			num: 5,
		},
		str: "some name",
	}

	ba := base{num: 2}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
	d = ba
	fmt.Println("describer:", d.describe())
}
