package main

import (
	"fmt"
)

func main() {
	str := "Ну начнём!"
	fmt.Scan(str)
	pat, chetire, tri, dva, kol, kolvo := OsenkiF(str)
	sredball := SredBallF(pat, chetire, tri, dva, kol, kolvo)
	if sredball < 1.5 {
		d44, d45, d454, d445, d5 := DoChetverkiIPaterki(pat, chetire, tri, dva, kol, kolvo, sredball)
		fmt.Printf("Надеюсь это шутка но всё же: /n Средний балл: &f /n Всего оценок &f /n До четвёрки надо ещё: &d по <4> или &d по <5> /n Оптимальный вариант: &d по <5> и &d по <4> /n До пятёрки надо: &d по <5>", sredball, kolvo, d44, d45, d454, d445, d5)
	} else if sredball < 2.5 {
		d44, d45, d454, d445, d5 := DoChetverkiIPaterki(pat, chetire, tri, dva, kol, kolvo, sredball)
		fmt.Printf("КАК??? /n Средний балл: &f /n Всего оценок &f /n До четвёрки надо ещё: &d по <4> или &d по <5> /n Оптимальный вариант: &d по <5> и &d по <4> /n До пятёрки надо: &d по <5>", sredball, kolvo, d44, d45, d454, d445, d5)
	} else if sredball < 3.5 {
		d44, d45, d454, d445, d5 := DoChetverkiIPaterki(pat, chetire, tri, dva, kol, kolvo, sredball)
		fmt.Printf("Средний балл: &f /n Всего оценок &f /n До четвёрки надо ещё: &d по <4> или &d по <5> /n Оптимальный вариант: &d по <5> и &d по <4> /n До пятёрки надо: &d по <5>", sredball, kolvo, d44, d45, d454, d445, d5)
	} else if sredball < 4.5 {
		d5 := DoPaterki(pat, chetire, tri, dva, kol, kolvo, sredball)
		fmt.Printf("Средний балл: &f /n Всего оценок &f /n До пятёрки надо: &d по <5>", sredball, kolvo, d5)
	} else {
		fmt.Printf("КАК??? /n Надеюсь это не шутка: /n Средний балл: &f /n Всего оценок &f", sredball, kolvo)
	}
}

func OsenkiF(str string) (float64, float64, float64, float64, float64, float64) {
	pat, chetire, tri, dva, kol, kolvo := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0
	for _, i := range str {
		if i == '5' {
			pat++
			kolvo++
		} else if i == '4' {
			chetire++
			kolvo++
		} else if i == '3' {
			tri++
			kolvo++
		} else if i == '2' {
			dva++
			kolvo++
		} else if i == '1' {
			pat++
			kolvo++
		} else {
			fmt.Errorf("Ошибка ввода")
		}
	}
	return pat, chetire, tri, dva, kol, kolvo
}

func SredBallF(pat, chetire, tri, dva, kol, kolvo float64) float64 {
	return (pat*5 + chetire*4 + tri*3 + dva*2 + kol*1) / kolvo
}

func DoChetverkiIPaterki(pat, chetire, tri, dva, kol, kolvo, sredball float64) (int, int, int, int, int) {
	sredballPr, patPr, chetPr, kolvoPr := sredball, pat, chetire, kolvo
	i := 0
	for sredballPr < 3.5 {
		i++
		chetPr++
		kolvoPr++
		sredballPr = SredBallF(pat, chetPr, tri, dva, kol, kolvoPr)
	}
	sredballPr, patPr, chetPr, kolvoPr = sredball, pat, chetire, kolvo
	o := 0
	for sredballPr < 3.5 {
		o++
		patPr++
		kolvoPr++
		sredballPr = SredBallF(patPr, chetire, tri, dva, kol, kolvoPr)
	}
	x := o - 1
	y := 0
	for sredballPr < 3.5 {
		y++
		chetPr++
		kolvoPr++
		sredballPr = SredBallF(patPr, chetire, tri, dva, kol, kolvoPr)
	}
	sredballPr, patPr, chetPr, kolvoPr = sredball, pat, chetire, kolvo
	r := 0
	for sredballPr < 4.5 {
		r++
		patPr++
		kolvoPr++
		sredballPr = SredBallF(patPr, chetire, tri, dva, kol, kolvoPr)
	}
	return i, o, x, y, r
}

func DoPaterki(pat, chetire, tri, dva, kol, kolvo, sredball float64) int {
	sredballPr, patPr, kolvoPr := sredball, pat, kolvo
	r := 0
	for sredballPr < 4.5 {
		r++
		patPr++
		kolvoPr++
		sredballPr = SredBallF(patPr, chetire, tri, dva, kol, kolvoPr)
	}
	return r
}
