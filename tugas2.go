package main

import (
	"log"
)

func main() {
	var triangle pattern

	triangle = star{5}
	log.Println(triangle.	rataKiri())
	log.Println(triangle.	rataKanan())
}

type pattern interface {
	rataKiri() string
	rataKanan() string
}

type star struct {
	total int
}

func (n star) rataKiri() string {
	str := ""
	for i := 0; i < n.total; i++ {
		strl := ""
		for j := 0; j <= i; j++ {
			strl += "*"
		}
		log.Println(strl)
	}
	return str
}

func (n star) rataKanan() string {
	str := ""
	for i := 0; i < n.total; i++ {
		strr := ""
		for j := 0; j < n.total; j++ {
			if j >= i {
				strr += "*"
			} else {
				strr += " "
			}
		}
		log.Println(strr)
	}
	return str
}