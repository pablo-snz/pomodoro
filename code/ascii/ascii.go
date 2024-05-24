package ascii

import (
	"fmt"
	"strings"
)

// Example 10: 30
//   ██  ██████	     ██████  ██████
// ████  ██  ██  ██      ██  ██  ██
//   ██  ██  ██      ██████  ██  ██
//   ██  ██  ██  ██      ██  ██  ██
//   ██  ██████	     ██████  ██████
// Ascii code idea from: https://github.com/zs5460/art/tree/master

func getArt(char string) []string {
	font := make(map[string][]string)
	font["0"] = []string{
		" ██████",
		" ██  ██",
		" ██  ██",
		" ██  ██",
		" ██████",
	}
	font["1"] = []string{
		"     ██",
		"   ████",
		"     ██",
		"     ██",
		"     ██",
	}
	font["2"] = []string{
		" ██████",
		"     ██",
		" ██████",
		" ██    ",
		" ██████",
	}
	font["3"] = []string{
		" ██████",
		"     ██",
		" ██████",
		"     ██",
		" ██████",
	}

	font["4"] = []string{
		" ██  ██",
		" ██  ██",
		" ██████",
		"     ██",
		"     ██",
	}

	font["5"] = []string{
		" ██████",
		" ██    ",
		" ██████",
		"     ██",
		" ██████",
	}

	font["6"] = []string{
		" ██████",
		" ██    ",
		" ██████",
		" ██  ██",
		" ██████",
	}

	font["7"] = []string{
		" ██████",
		"     ██",
		"     ██",
		"     ██",
		"     ██",
	}

	font["8"] = []string{
		" ██████",
		" ██  ██",
		" ██████",
		" ██  ██",
		" ██████",
	}

	font["9"] = []string{
		" ██████",
		" ██  ██",
		" ██████",
		"     ██",
		" ██████",
	}

	font[":"] = []string{
		"   ",
		" ██",
		"   ",
		" ██",
		"   ",
	}

	return font[char]

}

// String returns an art string
func Asciify(text string, states []string) string {
	if len(text) == 0 {
		return ""
	}

	if len(states) < 6 {
		for i := len(states); i < 6; i++ {
			states = append(states, "")
		}
	}
	chars := strings.Split(text, "")
	s := getArt(chars[0])
	for i := 1; i < len(chars); i++ {
		s = fuse(s, getArt(chars[i]))
	}

	s = fuse(s, states)

	return strings.Replace(strings.Join(s, "\n"), "$", " ", -1)
}

func fuse(left, right []string) []string {
	for i := 0; i < len(left); i++ {
		left[i] = left[i] + " " + right[i]
	}
	return left

}

func GenerateProgressBar(percentage int) string {
	var progressBar string

	// Asegúrate de que el porcentaje esté en el rango correcto
	if percentage < 0 {
		percentage = 0
	} else if percentage > 100 {
		percentage = 100
	}

	// Calcula la cantidad de "=" y " " en la barra de progreso
	equalsCount := percentage / 5
	spacesCount := 20 - equalsCount

	// Construye la barra de progreso
	progressBar += " ["
	for i := 0; i < equalsCount; i++ {
		progressBar += "="
	}
	progressBar += ">"
	for i := 0; i < spacesCount; i++ {
		progressBar += " "
	}
	progressBar += "] " + fmt.Sprintf("%d%%", percentage)

	return progressBar
}
