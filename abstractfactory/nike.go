package abstractfactory

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 22:18
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short{
			logo: "nike",
			size: 14,
		},
	}
}