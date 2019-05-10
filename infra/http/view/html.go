package view

import (
	"path/filepath"

	"github.com/tomocy/caster"
)

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
		&caster.TemplateSet{
			Filenames: []string{
				htmlTemplate("master.html"),
				htmlTemplate("header.html"),
				htmlTemplate("error.html"),
			},
		},
	)
	if err != nil {
		panic(err)
	}

	if err := h.caster.ExtendAll(
		map[string]*caster.TemplateSet{},
	); err != nil {
		panic(err)
	}
}

func htmlTemplate(fname string) string {
	projPath, _ := filepath.Abs(".")
	return filepath.Join(projPath, "resource/template", fname)
}
