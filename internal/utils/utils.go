package utils

import (
	"strconv"
	"strings"
)

// Repeat to repeat string.
func Repeat(str string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(str, n)
}

// PadLeft to pad string to the left.
func PadLeft(str string, l int, p string) string {
	return Repeat(p, l-len([]rune(str))) + str
}

// PadRight to pad string to the right.
func PadRight(str string, l int, p string) string {
	return str + Repeat(p, l-len([]rune(str)))
}

// Ellipsis to truncate string.
func Ellipsis(str string, length int) string {
	r := []rune(str)
	if len(r) > length {
		return string(r[:length-3]) + "..."
	}
	return str
}

// Thousands to format int to thousands string format.
func Thousands(num int) string {
	str := strconv.Itoa(num)
	lStr := len(str)
	digits := lStr
	if num < 0 {
		digits--
	}
	commas := (digits+2)/3 - 1
	lBuf := lStr + commas
	var sbuf [32]byte // pre allocate buffer at stack rather than make([]byte,n)
	buf := sbuf[0:lBuf]
	// copy str from the end
	for si, bi, c3 := lStr-1, lBuf-1, 0; ; {
		buf[bi] = str[si]
		if si == 0 {
			return string(buf)
		}
		si--
		bi--
		// insert comma every 3 chars
		c3++
		if c3 == 3 && (si > 0 || num > 0) {
			buf[bi] = ','
			bi--
			c3 = 0
		}
	}
}

// PtrToStr to convert string pointer to string.
func PtrToStr(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

// PtrToInt to convert int pointer to int.
func PtrToInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

// PtrToBool to convert bool pointer to bool.
func PtrToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
