/**
 * 	除使用第三方包外，有时还可能需要创建包。
 */
package temperature

func CtoF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
