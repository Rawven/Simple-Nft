package util

import "github.com/valyala/fastjson"

var p fastjson.Parser

func GetFastJson() *fastjson.Parser {
	return &p
}
