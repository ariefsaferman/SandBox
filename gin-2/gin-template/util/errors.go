package util

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrBookSameTitle = errors.New("cannot add book with same title")
var ErrQtyBook = errors.New("quantity error cannot below 0")
var ErrQtyBookNumeric = errors.New("quantity must be number")
var ErrBadRequest = errors.New("bad request")
var ErrInternalServer = errors.New("internal server error")
