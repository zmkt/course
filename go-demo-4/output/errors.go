package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	intValue, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intValue)
		return
	}

	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
		return
	}

	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
		return
	}

	color.Red("Неизвестный тип ошибки")
}

// func sum[T int | string](a, b T) T {
// 	switch d := any(a).(type) {
// 	case string:
// 		fmt.Println(d)
// 	}
// 	return a + b
// }

// func sumInt(a, b int) int {
// 	return a + b
// }

// func sumFloat32(a, b float32) float32 {
// 	return a + b
// }

// type List[T any] struct {
// 	elements []T
// }

// func (l *List[T]) addElement () {

// }
