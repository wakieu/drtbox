package entity

import (
	"strings"
)

type Box struct {
	BoxPath   string `json:"boxpath"`
	Text      string `json:"text"`
}

func NewBox(boxpath string) *Box {
	return &Box{
		BoxPath: boxpath,
	}
}

func NewBoxWithText(boxpath string, text string) *Box {
	return &Box{
		BoxPath: boxpath,
		Text: text,
	}
}

func (b *Box) UpdateText(text string) {
	b.Text = text
}

func (b *Box) IsEmpty() bool {
	return strings.TrimSpace(b.Text) == ""
}