package document

import "passportreader/lib"

type Document interface {
	Fill(idCard lib.IDCard) error
	GetBody() interface{}
}
