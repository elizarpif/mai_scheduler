package cmd

import "errors"


func NewErrNotFound()error{
	return  errors.New("not found")
}

func NewErrParseHtml()error{
	return  errors.New("parse html")
}
