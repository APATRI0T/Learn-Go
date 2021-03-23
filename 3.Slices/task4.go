package task4

import (
	"fmt"
)

func Run() {
	fmt.Println("Task4;")
	//arr := [5]float64{1,2,3,4,5}
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	minV := x[0]
	for _, v := range x {
		if v < minV {
			minV = v
		}
	}
	fmt.Println("MinV=", minV)

	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "fff")
	}
	//fmt.Printf("%#v\n",slice)
	//fmt.Println(slice)

	bb := []byte{1, 2, 3}
	fmt.Println("bb=", bb)

	fmt.Println("Append 6,7,8:")
	bb = AppendByte2(bb, 6, 7, 8)
	fmt.Println("bb=", bb)

}

type bs []byte

func AppendByte2(slice bs, data ...byte) bs {
	// https://blog.golang.org/slices-intro
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) { // если емкость меньше чем нужно,  создадим новый срез, опирающийся на в 2 раза больший базовый массив
		newSlice := make(bs, n, n*2+1)
		CopyByte(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	CopyByte(slice[m:n], data)
	return slice
}

func CopyByte(dst bs, src bs) bs {
	for i, v := range src {
		dst[i] = v
	}
	return dst
}

/*
Вопросы.
1. Что за ошибка Warning:(54, 32) Exported function with unexported return type
2.
*/
