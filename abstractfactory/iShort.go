package abstractfactory

/**
 *
 * @author: hoangtq
 * @timeCreate: 27/08/2020 22:02
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) getSize() int {
	return s.size
}
