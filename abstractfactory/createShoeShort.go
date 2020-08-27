package abstractfactory

import "fmt"

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 22:27
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func InitShoeShort()  {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShoe.setSize(20)
	nikeShort := nikeFactory.makeShort()
	nikeShort.setSize(22)

	adidasShoe := adidasFactory.makeShoe()
	adidasShoe.setSize(21)
	adidasShort := adidasFactory.makeShort()
	adidasShort.setSize(23)

	showShoeDetails(nikeShoe)
	showShoeDetails(adidasShoe)
	showShortDetails(nikeShort)
	showShortDetails(adidasShort)
}

func showShoeDetails(s iShoe) {
	fmt.Println("---SHOE---")
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func showShortDetails(s iShort) {
	fmt.Println("---SHORT---")
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

