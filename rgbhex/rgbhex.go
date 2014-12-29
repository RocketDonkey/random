// Simple command-line tool to convert hex to RGB or vice versa.
package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// An rgbval is a slice containing the RGB values passed to the -rgb flag.
type rgbval RGB

// The flag representing the comma-delimited RGB value.
var rgbflag rgbval

// The -rgb flag needs to satisfy the flag.Value interface, which requires a
// String and Set method.
func (rgbval *rgbval) String() string {
	return fmt.Sprint(*rgbval)
}

// Convert a comma-separated string of RGB values into a value of type RGB.
func (rgbval *rgbval) Set(value string) error {
	rgbSlice := strings.Split(value, ",")
	if len(rgbSlice) != 3 {
		return errors.New(
			"Invalid value for flag rgb. Flag value should be in r,g,b format")
	}

	// Set the values of the flag.
	rgbval.red, _ = strconv.ParseInt(rgbSlice[0], 0, 64)
	rgbval.green, _ = strconv.ParseInt(rgbSlice[1], 0, 64)
	rgbval.blue, _ = strconv.ParseInt(rgbSlice[2], 0, 64)
	return nil
}

// An RGB contains the (R)ed, (G)reen and (B)lue values for a color.
type RGB struct {
	red   int64
	green int64
	blue  int64
}

// Output RGB in a human readable format.
func (rgb *RGB) String() string {
	return fmt.Sprintf("RGB{R:%d, G:%d, B:%d}", rgb.red, rgb.green, rgb.blue)
}

// HexToRGB converts a hex color value to RGB.
func HexToRGB(hexVal int64) *RGB {
	// For a hex value of #AABBCC, the RGB values are:
	//   R: AA
	//   G: BB
	//   B: CC

	// Red: Right-shift the value 16 bits.
	red := hexVal >> 16

	// Green: Right-shift 8 bits and then bitwise AND with 8 bits of 1's.
	green := (hexVal >> 8) & 0xFF

	// Blue: bitwise AND with 8 bits of 1's.
	blue := hexVal & 0xFF
	return &RGB{red, green, blue}
}

// RGBToHex converts an color value to hex.
func RGBToHex(rgbVal *rgbval) string {
	// To convert the RGB value to hex, we can convert each individual value to
	// hex and then concatenate the results together.
	// Example (R, G, B): (10, 15 20) -> (A, F, 14) -> 0A0F14
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%02x", rgbVal.red))
	buffer.WriteString(fmt.Sprintf("%02x", rgbVal.green))
	buffer.WriteString(fmt.Sprintf("%02x", rgbVal.blue))
	return buffer.String()
}

func main() {
	// Parse the command line flags.
	hexPtr := flag.String("hex", "", "a hex value to convert to RGB")
	flag.Var(&rgbflag, "rgb", "an RGB value to convert to hex")
	flag.Parse()

	// If neither flag is provided, print usage and exit.
	if flag.NFlag() == 0 || flag.NFlag() == 2 {
		fmt.Println("Usage: rgbhex [-hex=ABCDEF | -rgb=123,234,100]")
		os.Exit(1)
	}

	if *hexPtr != "" {
		// Convert the provided hex value into an integer.
		hexVal, err := strconv.ParseInt(*hexPtr, 16, 32)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(HexToRGB(hexVal))
	} else {
		// Convert an RGB value to hex.
		fmt.Println(RGBToHex(&rgbflag))
	}
}
