package lights

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/go-playground/colors.v1"
)

type InvalidHexValue struct {
	value string
	msg   string
}

func (err *InvalidHexValue) Error() string {
	return fmt.Sprintf("%s : %s", err.msg, err.value)
}

func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func validHexPrefix(hex string) bool {
	return hex[0] == '#'
}

func HexToColor(hex string) (Color, error) {
	var color Color

	if !validHexPrefix(hex) {
		hex = "#" + hex
	} else {
		return color, &InvalidHexValue{hex, "Hex Value is invalid"}
	}

	hexColor, _ := colors.ParseHEX(hex)
	rgba := hexColor.ToRGBA()

	color, err := NewColor(
		rgba.R,
		rgba.G,
		rgba.B,
		uint8(rgba.A*255),
	)
	if err != nil {
		return color, &InvalidHexValue{hex, "Hex Value is invalid"}
	}

	return color, nil
}

func lookupColorByName(color string) (*Color, *InvalidColorError) {
	// TODO: refactor file read to allow it to be mocked easily
	content, err := ioutil.ReadFile("colors.min.json")

	if err != nil {
		fmt.Printf("Error while reading a file %v", err)
	}
	var hexes map[string]string
	_ = json.Unmarshal(content, &hexes)

	names := reverseMap(hexes)
	hex, ok := names[string(color)]
	if !ok {
		return nil, &InvalidColorError{color, "Color is invalid"}
	}
	h, err := HexToColor(hex)
	if err != nil {
		return nil, &InvalidColorError{color, "Color is invalid"}
	}
	return &h, nil
}

func lookupPrimaryColor(color string) (*Color, *InvalidColorError) {
	c, ok := ColorPresets[color]
	if !ok {
		return nil, &InvalidColorError{color, "Color is invalid"}
	}
	return &c, nil
}
