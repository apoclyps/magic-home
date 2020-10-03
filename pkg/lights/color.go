package lights

type InvalidColorError struct {
	color string
	msg   string
}

func (err *InvalidColorError) Error() string {
	return err.msg
}

type InvalidColorValue struct {
	value uint8
	msg   string
}

func (err *InvalidColorValue) Error() string {
	return err.msg
}

// Color represents a RGBW color
type Color struct {
	R uint8
	G uint8
	B uint8
	W uint8
}

// NewColor initializes a new Color using valid RGBW values (between 0 - 255).
func NewColor(r uint8, g uint8, b uint8, w uint8) (Color, error) {
	var color Color

	if r > 255 {
		return color, &InvalidColorValue{r, "Red value is not within range. e.g. Value must be between 0 - 255"}
	}
	if g > 255 {
		return color, &InvalidColorValue{r, "Green value is not within range. e.g. Value must be between 0 - 255"}
	}
	if b > 255 {
		return color, &InvalidColorValue{r, "Blue value is not within range. e.g. Value must be between 0 - 255"}
	}
	if w > 255 {
		return color, &InvalidColorValue{r, "White value is not within range. e.g. Value must be between 0 - 255"}
	}

	color = Color{
		R: r,
		G: g,
		B: b,
		W: w,
	}

	return color, nil
}
