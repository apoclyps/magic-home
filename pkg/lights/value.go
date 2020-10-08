package lights

// Value represents the string value used to present a color.
type Value struct {
	value string
}

// NewValue initialize a string or hex representation of a Color
func NewValue(value string) (Value, error) {
	return Value{value: value}, nil
}

// GetColorByHex attempt's to find a matching Color using the value as a hex value.
func (c *Value) GetColorByHex() (Color, error) {
	var color Color

	if c.value != "" {
		color, err := HexToColor(c.value)
		if err == nil {
			return color, nil
		}
	}

	return color, &InvalidColorError{c.value, "Value is invalid"}
}

// GetColorByName attempt's to find a matching Color using the value as a color preset.
func (c *Value) GetColorByName() (Color, error) {
	var color Color

	if c.value == "" {
		return color, &InvalidColorError{c.value, "Value is invalid"}
	}

	cn, _ := lookupColorByName(c.value)
	if cn != nil {
		return *cn, nil
	}

	cl, err := lookupPrimaryColor(c.value)
	if err != nil {
		return *cl, &InvalidColorError{c.value, "Value is invalid"}
	}
	if cl != nil {
		return *cl, nil
	}

	return color, &InvalidColorError{c.value, "Value is invalid"}
}
