package abstractfactory

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 21:58
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) setSize(size int){
	s.size = size
}

func (s *shoe) getLogo() string  {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}