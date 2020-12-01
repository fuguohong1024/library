package library

import "fmt"


//控制台打印对应字体的颜色
const (
	Black = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func SetBlack(msg string) string {
	return SetColor(msg, 0, 0, Black)
}

func SetRed(msg string) string {
	return SetColor(msg, 0, 0, Red)
}

func SetGreen(msg string) string {
	return SetColor(msg, 0, 0, Green)
}

func SetYellow(msg string) string {
	return SetColor(msg, 0, 0, Yellow)
}

func SetBlue(msg string) string {
	return SetColor(msg, 0, 0, Blue)
}

func SetMagenta(msg string) string {
	return SetColor(msg, 0, 0, Magenta)
}

func SetCyan(msg string) string {
	return SetColor(msg, 0, 0, Cyan)
}

func SetWhite(msg string) string {
	return SetColor(msg, 0, 0, White)
}

func SetColor(msg string, conf, bg,text  int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
