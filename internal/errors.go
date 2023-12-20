package internal

import "matryer/pkg/joker"

var (
	ErrNoContent   = joker.NewErrorf(joker.CodeNoContent, "no content")
	ErrSqlQueryRow = joker.NewErrorf(joker.CodeSqlQueryRow, "")
	ErrSqlSelect   = joker.NewErrorf(joker.CodeSqlSelect, "")
)
