package dao

import "errors"

var (
	ErrInsertDocument = errors.New("cannot insert document into collection")
	ErrGetDocuments   = errors.New("cannot get documents from collection")
	ErrDeleteDocument = errors.New("cannot delete document from collection")
	ErrUnmarshal      = errors.New("cannot unmarshal values")
)
