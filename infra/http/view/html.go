package view

import "github.com/tomocy/caster"

func NewHTML() *HTML {
	h := new(HTML)
	h.mustParseTemplates()

	return h
}

type HTML struct {
	caster caster.Caster
}

func (h *HTML) mustParseTemplates() {
	var err error
	h.caster, err = caster.New(
		&caster.TemplateSet{}
	)
	if err != nil {
		panic(err)
	}

	if err := h.caster.ExtendAll(
		&caster.TemplateSet{},
	); err != nil {
		panic(err)
	}
}
