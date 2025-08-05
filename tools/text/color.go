/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package text

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
)

func textWithColor(color, text string) string {
	return color + text + reset
}

func RedText(text string) string {
	return textWithColor(red, text)
}

func GreenText(text string) string {
	return textWithColor(green, text)
}

func YellowText(text string) string {
	return textWithColor(yellow, text)
}

func BlueText(text string) string {
	return textWithColor(blue, text)
}

func PurpleText(text string) string {
	return textWithColor(purple, text)
}

func CyanText(text string) string {
	return textWithColor(cyan, text)
}

func WhiteText(text string) string {
	return textWithColor(white, text)
}
