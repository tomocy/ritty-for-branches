package view

import (
	"html/template"
	"net/http"
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

func (h *HTML) Show(w http.ResponseWriter, name string, data interface{}) error {
	return h.caster.Cast(w, name, data)
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
		map[string]*caster.TemplateSet{
			"menu.index": &caster.TemplateSet{
				Filenames: []string{
					htmlTemplate("menu/index.html"),
				},
			},
			"menu.new": &caster.TemplateSet{
				Filenames: []string{
					htmlTemplate("menu/single.html"),
				},
				FuncMap: template.FuncMap{
					"add": func(a, b int) int {
						return a + b
					},
				},
			},
			"menu.show": &caster.TemplateSet{
				Filenames: []string{
					htmlTemplate("menu/single.html"),
				},
				FuncMap: template.FuncMap{
					"add": func(a, b int) int {
						return a + b
					},
				},
			},
		},
	); err != nil {
		panic(err)
	}
}

func htmlTemplate(fname string) string {
	projPath, _ := filepath.Abs(".")
	return filepath.Join(projPath, "resource/template", fname)
}
