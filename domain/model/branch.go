package model

type Branch struct {
	ID               string
	Authorization    *Authorization
	CategorizedMenus CategorizedMenus
}

func (b *Branch) RegisterMenu(c MenuCategory, menu *Menu) error {
	if _, err := b.FindMenu(c, menu.Name); err == nil {
		return validationErrorf("register menu", "menu is already registered")
	}

	if b.CategorizedMenus == nil {
		b.CategorizedMenus = make(CategorizedMenus)
	}
	b.CategorizedMenus[c] = append(b.CategorizedMenus[c], menu)

	return nil
}

func (b *Branch) FindMenuCategory(name string) (MenuCategory, error) {
	for stored := range b.CategorizedMenus {
		if stored.Name == name {
			return stored, nil
		}
	}

	return MenuCategory{}, validationErrorf("find menu category", "no such menu")
}

func (b *Branch) UpdateMenu(c MenuCategory, old, new *Menu) error {
	if b.CategorizedMenus == nil {
		b.CategorizedMenus = make(CategorizedMenus)
	}
	storeds := b.CategorizedMenus[c]
	for i, stored := range storeds {
		if stored.Name == old.Name {
			b.CategorizedMenus[c][i] = new
			return nil
		}
	}

	return validationErrorf("register menu", "no such menu")
}

func (b *Branch) FindMenu(c MenuCategory, name string) (*Menu, error) {
	storeds := b.CategorizedMenus[c]
	for _, stored := range storeds {
		if stored.Name == name {
			return stored, nil
		}
	}

	return nil, validationErrorf("register menu", "no such menu")
}
