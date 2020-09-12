package lights

// A Candle light source with a Kelvin temperature of 1900
func Candle() Color {
	return Color{R: 255, G: 147, B: 41, W: 0}
}

// A 40W Tungeson light source with a Kelvin temperature of 2600
func Tungsen40W() Color {
	return Color{R: 255, G: 197, B: 143, W: 0}
}

// A 100W Tungeson light source with a Kelvin temperature of 2850
func Tungsen100W() Color {
	return Color{R: 255, G: 214, B: 170, W: 0}
}

// A Halogen light source with a Kelvin temperature of 3200
func Halogen() Color {
	return Color{R: 255, G: 241, B: 224, W: 0}
}

// A Carbon Arch light source with a Kelvin temperature of 5200
func CarbonArc() Color {
	return Color{R: 255, G: 250, B: 244, W: 0}
}

// A Highnoon Sun light source with a Kelvin temperature of 5400
func HighNoonSun() Color {
	return Color{R: 255, G: 255, B: 251, W: 0}
}

// A direct Sun light source with a Kelvin temperature of 6000
func DirectSunlight() Color {
	return Color{R: 255, G: 255, B: 251, W: 0}
}

// A overcast sky light source with a Kelvin temperature of 7000
func OvercastSky() Color {
	return Color{R: 255, G: 255, B: 251, W: 0}
}

// A clear blue sky light source with a Kelvin temperature of 20000
func ClearBlueSky() Color {
	return Color{R: 64, G: 156, B: 255, W: 0}
}
