package models

type MenuOptionEnums struct {
	Luminance  string
	PixelSort  string
	RedShift   string
	BlueShift  string
	GreenShift string
	AlphaShift string
}

var MenuOptions = MenuOptionEnums{
	Luminance:  "l",
	PixelSort:  "s",
	RedShift:   "r",
	BlueShift:  "b",
	GreenShift: "g",
	AlphaShift: "a",
}
