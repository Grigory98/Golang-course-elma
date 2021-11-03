//Конвертеры
package convert

import (
	"fmt"
	"strings"
)

//Преобразовывает запрос для JSON строки
func ConvertRequest(request []interface{}) string {
	req := fmt.Sprintf("%v", request)
	req = strings.ReplaceAll(req, " ", ",")
	return req
}

//Преобразовывает ответ для JSON строки(двумерный массив)
func ConvertAnswer(answer [][]int) string {
	answ := fmt.Sprintf("%v", answer)
	answ = strings.ReplaceAll(answ, " ", ",")
	return answ
}

//Преобразовывает ответ для JSON строки(массив)
func ConvertAnswer_2(answer []int) string {
	answ := fmt.Sprintf("%v", answer)
	answ = strings.ReplaceAll(answ, " ", ",")
	return answ
}
