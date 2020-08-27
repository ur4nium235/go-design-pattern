package abstractfactory

import "fmt"

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 21:51
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}
