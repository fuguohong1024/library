package library

import "fmt"

// 字体颜色
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

// 背景颜色
const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// 字体格式
const (
	None = iota
	// 加粗
	Bold
	// 斜体
	Italic = iota + 1
	// 下划线
	Underscore
	// 闪烁
	Flash
	// 删除线
	Delete = iota + 4
)

type Font struct {
	Format int
	Bg     int
	Color  int
}

//  "\033[字背景颜色;文字颜色m  你要显示的内容  \033[0m"
//                    |                         |
//                 控制颜色                  控制其他属性
// ESC的ascii码 0x1B
// bg 背景色  text 字体颜色  conf

// 传参字体   背景色   字体色
func Newcolor(format, bg, color int) *Font {
	return &Font{Format: format,
		Color: color,
		Bg:    bg,
	}
}

func (f *Font) String(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, f.Format, f.Bg, f.Color, msg, 0x1B)
}
