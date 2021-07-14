package shortpaste

import "gopkg.in/go-playground/validator.v9"

func (link *Link) validate() error {
	return validator.New().Struct(link)
}

func (file *File) validate() error {
	return validator.New().Struct(file)
}

func (text *Text) validate() error {
	return validator.New().Struct(text)
}
