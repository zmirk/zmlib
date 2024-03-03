package zmlib

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	println("********** zmlib.go **********")
}

func Percentage(Title string, CharEmpty string, CharFull string, LenBar int, Len int, i int) string {
	CurPercent := math.Floor(100.0 / float64(Len) * float64(i))
	CurLen := int(CurPercent / 100.0 * float64(LenBar))

	result := Title + " ["

	for j := 0; j < CurLen; j++ {
		result = result + CharFull
	}

	for k := 0; k < LenBar-CurLen; k++ {
		result = result + CharEmpty
	}

	result = result + "] "
	result = result + fmt.Sprintf("%.0f", CurPercent) + "%"

	result = result + " (" + fmt.Sprintf("%d", i) + "/" + fmt.Sprintf("%d", Len) + ")"

	return result
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandSleep(min, max int) {
	curSleep := RandInt(min, max)
	time.Sleep(time.Second / time.Duration(curSleep))
}
