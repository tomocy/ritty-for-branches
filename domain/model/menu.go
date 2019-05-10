package model

type Menu struct {
	Name, Description string
	Price             float32
	ImagePath         string
	Availability      uint
	CategorizedCombos CategorizedCombos
}
