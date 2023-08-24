package users

import (
	"fmt"
)

func InitUser() {
	var frank Cashier = NewEmployee("frank")
	var fab Admin = NewEmployee("fabricio")
	total := frank.CalcTotal(10, 20, 30, 40)

	fmt.Printf("%+v", frank)
	fmt.Println("")
	fmt.Println(total)
	fab.DesactivarEmployeer(frank)
	total = frank.CalcTotal(20, 30, 40)

	fab.ActivarEmployeer(frank)

	total = frank.CalcTotal(10, 20, 30, 40)
	fmt.Println("")
	fmt.Println(total)

}