package lights

// A White light
func White() Color {
	return Color{R: 255, G: 255, B: 255, W: 0}
}

// A Yellow light
func Black() Color {
	return Color{R: 0, G: 0, B: 0, W: 0}
}

// A Red light
func Red() Color {
	return Color{R: 255, G: 0, B: 0, W: 0}
}

// A Blue light
func Blue() Color {
	return Color{R: 0, G: 0, B: 255, W: 0}
}

// A Green light
func Green() Color {
	return Color{R: 0, G: 255, B: 0, W: 0}
}

// A Yellow light
func Yellow() Color {
	return Color{R: 255, G: 255, B: 0, W: 0}
}

// A Pink light
func Pink() Color {
	return Color{R: 255, G: 0, B: 255, W: 0}
}

// A Cyan light
func Cyan() Color {
	return Color{R: 0, G: 255, B: 255, W: 0}
}

// A Silver light
func Silver() Color {
	return Color{R: 192, G: 192, B: 192, W: 0}
}

// A Gray light
func Gray() Color {
	return Color{R: 128, G: 128, B: 128, W: 0}
}

// A Maroon light
func Maroon() Color {
	return Color{R: 128, G: 255, B: 255, W: 0}
}

// A Olive light
func Olive() Color {
	return Color{R: 128, G: 128, B: 0, W: 0}
}

// A Purple light
func Purple() Color {
	return Color{R: 128, G: 0, B: 128, W: 0}
}

// A Teal light
func Teal() Color {
	return Color{R: 0, G: 128, B: 128, W: 0}
}

// A Navy light
func Navy() Color {
	return Color{R: 0, G: 0, B: 128, W: 0}
}

// A Orange light
func Orange() Color {
	return Color{R: 255, G: 165, B: 0, W: 0}
}

var ColorPresets = map[string]Color{
	"white":  White(),
	"black":  Black(),
	"red":    Red(),
	"blue":   Blue(),
	"green":  Green(),
	"yellow": Yellow(),
	"pink":   Pink(),
	"cyan":   Cyan(),
	"silver": Silver(),
	"gray":   Gray(),
	"maroon": Maroon(),
	"olive":  Olive(),
	"purple": Purple(),
	"teal":   Teal(),
	"navy":   Navy(),
	"orange": Orange(),
}
