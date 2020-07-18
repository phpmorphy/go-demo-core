package demo

type Demo [150]byte

func (d Demo) Test() {
	d[0] = 1
}
