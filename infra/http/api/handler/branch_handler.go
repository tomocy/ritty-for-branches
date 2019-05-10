package handler

import (
	"net/http"
	"path"

	"github.com/go-chi/chi"
	"github.com/tomocy/ritty-for-branches/config"
	"github.com/tomocy/ritty-for-branches/domain/model"
	"github.com/tomocy/ritty-for-branches/domain/repository"
	"github.com/tomocy/ritty-for-branches/infra/http/view"
)

func newBranchHandler(
	view view.View,
	repo repository.BranchRepository,
) *branchHandler {
	return &branchHandler{
		view: view,
		repo: repo,
	}
}

type branchHandler struct {
	view view.View
	repo repository.BranchRepository
}

func (h *branchHandler) DisposeBranches(w http.ResponseWriter, r *http.Request) {
	branches := h.repo.GetBranches()
	converteds := convertBranchesToDispose(branches)

	if err := h.view.Show(w, "", map[string]interface{}{
		"Branches": converteds,
	}); err != nil {
		logInternalServerError(w, "dispose branches", err)
	}
}

func (h *branchHandler) DisposeBranch(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	branch, err := h.repo.FindBranch(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	converted := convertBranchToDispose(branch)

	if err := h.view.Show(w, "", map[string]interface{}{
		"Branch": converted,
	}); err != nil {
		logInternalServerError(w, "dispose branch", err)
	}
}

func convertBranchesToDispose(branches []*model.Branch) []*branchToDispose {
	converteds := make([]*branchToDispose, len(branches))
	for i, branch := range branches {
		converteds[i] = convertBranchToDispose(branch)
	}

	return converteds
}

func convertBranchToDispose(branch *model.Branch) *branchToDispose {
	return &branchToDispose{
		ID:               branch.ID,
		CategorizedMenus: convertCategorizedMenusToDispose(branch.CategorizedMenus),
	}
}

type branchToDispose struct {
	ID               string
	Name             string
	ImageURL         string
	CategorizedMenus []*categorizedMenusToDispose
}

func convertCategorizedMenusToDispose(categorizedMenu model.CategorizedMenus) []*categorizedMenusToDispose {
	converteds := make([]*categorizedMenusToDispose, len(categorizedMenu))
	var i int
	for category, menus := range categorizedMenu {
		converteds[i] = &categorizedMenusToDispose{
			Category: convertMenuCategoryToDispose(category),
			Menus:    convertMenusToDispose(menus),
		}
		i++
	}

	return converteds
}

type categorizedMenusToDispose struct {
	Category menuCategoryToDispose
	Menus    []*menuToDispose
}

func convertMenuCategoryToDispose(category model.MenuCategory) menuCategoryToDispose {
	return menuCategoryToDispose{
		Name: category.Name,
	}
}

type menuCategoryToDispose struct {
	Name string
}

func convertMenusToDispose(menus []*model.Menu) []*menuToDispose {
	converteds := make([]*menuToDispose, len(menus))
	for i, menu := range menus {
		converteds[i] = convertMenuToDispose(menu)
	}

	return converteds
}

func convertMenuToDispose(menu *model.Menu) *menuToDispose {
	return &menuToDispose{
		Name:              menu.Name,
		Price:             menu.Price,
		Description:       menu.Description,
		ImageURL:          config.Current.Self.Host + path.Join("/", menu.ImagePath),
		Availability:      uint(menu.Availability),
		CategorizedCombos: convertCategorizedCombosToDispose(menu.CategorizedCombos),
	}
}

type menuToDispose struct {
	Name, Description string
	Price             float32
	ImageURL          string
	Availability      uint
	CategorizedCombos []*categorizedCombosToDispose
}

func convertCategorizedCombosToDispose(categorizedCombo model.CategorizedCombos) []*categorizedCombosToDispose {
	converteds := make([]*categorizedCombosToDispose, len(categorizedCombo))
	var i int
	for category, combos := range categorizedCombo {
		converteds[i] = &categorizedCombosToDispose{
			Category: convertComboCategoryToDispose(category),
			Combos:   convertCombosToDispose(combos),
		}
		i++
	}

	return converteds
}

type categorizedCombosToDispose struct {
	Category comboCategoryToDispose
	Combos   []*comboToDispose
}

func convertComboCategoryToDispose(category model.ComboCategory) comboCategoryToDispose {
	return comboCategoryToDispose{
		Name:      category.Name,
		Condition: uint(category.Condition),
	}
}

type comboCategoryToDispose struct {
	Name      string
	Condition uint
}

func convertCombosToDispose(combos []*model.Combo) []*comboToDispose {
	converteds := make([]*comboToDispose, len(combos))
	for i, combo := range combos {
		converteds[i] = &comboToDispose{
			Name:  combo.Name,
			Price: combo.Price,
		}
	}

	return converteds
}

type comboToDispose struct {
	Name  string
	Price float32
}
