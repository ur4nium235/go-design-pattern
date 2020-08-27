package abstractfactory

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 21:53
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short{
			logo: "adidas",
			size: 14,
		},
	}
}
