package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/govalidator"
	derr "github.com/tomocy/ritty-for-branches/domain/error"
	"github.com/tomocy/ritty-for-branches/domain/model"
	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/domain/service"
	"github.com/tomocy/ritty-for-branches/infra/http/route"
	"github.com/tomocy/ritty-for-branches/infra/http/session"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func newBranchHandler(
	view view.View,
	repo repository.BranchRepository,
	storageServ service.StorageService,
) *branchHandler {
	return &branchHandler{
		view:        view,
		repo:        repo,
		storageServ: storageServ,
	}
}

type branchHandler struct {
	view        view.View
	repo        repository.BranchRepository
	storageServ service.StorageService
}

func (h *branchHandler) ShowMenus(w http.ResponseWriter, r *http.Request) {
	id, err := session.FindAuthenticBranch(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.view.Show(w, "menu.index", map[string]interface{}{
		"Branch": branch,
	}); err != nil {
		logInternalServerError(w, "show restaurant menus", err)
	}
}

func (h *branchHandler) ShowMenuRegistrationForm(w http.ResponseWriter, r *http.Request) {
	cname := r.URL.Query().Get("category_name")
	category := model.MenuCategory{
		Name: cname,
	}

	if err := h.view.Show(w, "menu.new", map[string]interface{}{
		"Category": category,
		"Error":    session.GetErrorMessage(w, r),
	}); err != nil {
		logInternalServerError(w, "show branch menu registration form", err)
	}
}

func (h *branchHandler) RegisterMenu(w http.ResponseWriter, r *http.Request) {
	id, err := session.FindAuthenticBranch(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := validateRequestToRegisterMenu(r); err != nil {
		session.SetErrorMessage(w, r, err)
		redirect(w, r, route.Web.Route("menu.new").String())
		return
	}
	category, menu, err := h.convertRequestToRegisterMenu(r)
	if err != nil {
		if derr.InInput(err) {
			session.SetErrorMessage(w, r, err)
			redirect(w, r, route.Web.Route("menu.new").String())
			return
		}

		logInternalServerError(w, "register menu", err)
		return
	}

	if err := branch.RegisterMenu(category, menu); err != nil {
		if derr.InInput(err) {
			session.SetErrorMessage(w, r, err)
			redirect(w, r, route.Web.Route("menu.new").String())
			return
		}

		logInternalServerError(w, "register menu", err)
		return
	}

	if err := h.repo.SaveBranch(branch); err != nil {
		logInternalServerError(w, "register menu", err)
		return
	}

	redirect(w, r, route.Web.Route("menu.index").String())
}

func (h *branchHandler) ShowMenu(w http.ResponseWriter, r *http.Request) {
	id, err := session.FindAuthenticBranch(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cname := chi.URLParam(r, "category_name")
	category, err := branch.FindMenuCategory(cname)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	name := chi.URLParam(r, "name")
	menu, err := branch.FindMenu(category, name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := h.view.Show(w, "menu.show", map[string]interface{}{
		"Category": category,
		"Menu":     menu,
		"Error":    session.GetErrorMessage(w, r),
		"Action":   fmt.Sprintf("/menus/%s/%s", category.Name, menu.Name),
		"Method":   "PUT",
	}); err != nil {
		logInternalServerError(w, "show menu", err)
		return
	}
}

func (h *branchHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	id, err := session.FindAuthenticBranch(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cname := chi.URLParam(r, "category_name")
	category, err := branch.FindMenuCategory(cname)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	name := chi.URLParam(r, "name")
	menu, err := branch.FindMenu(category, name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := validateRequestToUpdateMenu(r); err != nil {
		session.SetErrorMessage(w, r, err)
		redirect(w, r, fmt.Sprintf("%s/%s/%s", route.Web.Route("menu.show"), category.Name, menu.Name))
		return
	}
	newCategory, newMenu, err := h.convertRequestToRegisterMenu(r)
	if err != nil {
		if derr.InInput(err) {
			session.SetErrorMessage(w, r, err)
			redirect(w, r, fmt.Sprintf("%s/%s/%s", route.Web.Route("menu.show"), category.Name, menu.Name))
			return
		}

		logInternalServerError(w, "register menu", err)
		return
	}
	if newMenu.ImagePath == "" {
		newMenu.ImagePath = menu.ImagePath
	} else {
		h.storageServ.DeleteImage(menu.ImagePath)
	}

	if err := branch.UpdateMenu(newCategory, menu, newMenu); err != nil {
		if derr.InInput(err) {
			session.SetErrorMessage(w, r, err)
			redirect(w, r, fmt.Sprintf("%s/%s/%s", route.Web.Route("menu.show"), category.Name, menu.Name))
			return
		}

		logInternalServerError(w, "register menu", err)
		return
	}

	if err := h.repo.SaveBranch(branch); err != nil {
		logInternalServerError(w, "register menu", err)
		return
	}

	redirect(w, r, route.Web.Route("menu.index").String())
}

func (h *branchHandler) convertRequestToRegisterMenu(r *http.Request) (model.MenuCategory, *model.Menu, error) {
	category := model.MenuCategory{}

	price, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		return category, nil, validationErrorf("convert request to register menu", err)
	}
	var imagePath string
	image, header, err := r.FormFile("image")
	if err == nil {
		imagePath, err = h.storageServ.SaveImage(image, filepath.Ext(header.Filename))
		if err != nil {
			return category, nil, devErrorf("convert request to register menu", err)
		}
	}
	availability, err := strconv.ParseUint(r.FormValue("availability"), 10, 8)
	if err != nil {
		return category, nil, validationErrorf("convert request to register menu", err)
	}
	ccombos, err := h.convertRequestToCategorizedCombos(r)
	if err != nil {
		return category, nil, err
	}

	category.Name = r.FormValue("category_name")
	return category, &model.Menu{
		Name:              r.FormValue("name"),
		Description:       r.FormValue("description"),
		Price:             float32(price),
		ImagePath:         imagePath,
		Availability:      uint(availability),
		CategorizedCombos: ccombos,
	}, nil
}

func validateRequestToRegisterMenu(r *http.Request) error {
	names := []string{
		"category_name", "price", "image", "availability",
	}
	options := govalidator.Options{
		Request: r,
		Rules: govalidator.MapData{
			"category_name": []string{"required"},
			"price":         []string{"required", "float"},
			"file:image":    []string{"required", "mime:image/jpeg,image/jpg,image/png"},
			"availability":  []string{"required"},
		},
	}

	if err := validate(names, options); err != nil {
		return validationErrorf("validate request to register menu", err)
	}

	return nil
}

func validateRequestToUpdateMenu(r *http.Request) error {
	names := []string{
		"category_name", "price", "image", "availability",
	}
	options := govalidator.Options{
		Request: r,
		Rules: govalidator.MapData{
			"category_name": []string{"required"},
			"price":         []string{"required", "float"},
			"file:image":    []string{"mime:image/jpeg,image/jpg,image/png"},
			"availability":  []string{"required"},
		},
	}

	if err := validate(names, options); err != nil {
		return validationErrorf("validate request to register menu", err)
	}

	return nil
}

func (h *branchHandler) convertRequestToCategorizedCombos(r *http.Request) (model.CategorizedCombos, error) {
	is := r.Form["combo_category[]"]
	categorizedCombos := make(map[model.ComboCategory][]*model.Combo)
	for i := range is {
		key := fmt.Sprintf("combo_category[%d]", i)
		name := r.FormValue(key + "[name]")
		cond, err := strconv.ParseUint(r.FormValue(key+"[condition]"), 10, 8)
		if err != nil {
			return nil, validationErrorf("convert reqeust to categorized combos", err)
		}
		category := model.ComboCategory{
			Name:      name,
			Condition: uint(cond),
		}

		combos, err := convertRequestToCombos(r, key)
		if err != nil {
			return nil, err
		}
		categorizedCombos[category] = combos
	}

	return categorizedCombos, nil
}

func convertRequestToCombos(r *http.Request, key string) ([]*model.Combo, error) {
	is := r.Form[key+"[combo][]"]
	combos := make([]*model.Combo, len(is))
	for j := range is {
		comboKey := fmt.Sprintf("%s[combo][%d]", key, j)
		comboName := r.FormValue(comboKey + "[name]")
		comboPrice, err := strconv.ParseFloat(r.FormValue(comboKey+"[price]"), 32)
		if err != nil {
			break
		}
		combos[j] = &model.Combo{
			Name:  comboName,
			Price: float32(comboPrice),
		}
	}

	return combos, nil
}
