package user

import (
	"github.com/robertobouses/api_ejercicio1/http"
)

type CustomRectangle struct {
	http.RegRectangle
}

func (r CustomRectangle) Area() int {
	return r.Long * r.Short
}
