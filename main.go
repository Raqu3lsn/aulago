package main

import "fmt"
//map podemos escolher a chave e o valor
func main(){
	menu := map [string] float64 {
		"pizza": 40.00,
		"suco": 12.50,
		"x-tudo": 22.90, 
	}
	fmt.Println(menu["pizza"])
	for k,v := range menu{
	fmt.Println(k,"R$",v)
	}
	contanova := novaConta("Astrubal")
	fmt.Println(contanova)

	contanova2 := novaConta("Astrubal")
	fmt.Println(contanova2)
}