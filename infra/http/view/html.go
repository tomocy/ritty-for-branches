package view

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/tomocy/caster"
	"github.com/tomocy/ritty-for-branches/domain/model"
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
			"profile.index": &caster.TemplateSet{
				Filenames: []string{
					htmlTemplate("profile/index.html"),
				},
				FuncMap: template.FuncMap{
					"from": func(b *model.Branch) string {
						if 1 <= len(b.OpeningDates) {
							return b.OpeningDates[0].From.Format("15:04")
						}

						return ""
					},
					"to": func(b *model.Branch) string {
						if 1 <= len(b.OpeningDates) {
							return b.OpeningDates[0].To.Format("15:04")
						}

						return ""
					},
					"isOpen": func(b *model.Branch, day uint) bool {
						for _, date := range b.OpeningDates {
							if date.Day == day {
								return true
							}
						}

						return false
					},
				},
			},
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
