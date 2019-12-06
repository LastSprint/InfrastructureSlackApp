package utils

// JoinByCharacter объединяет массив строк через символ `character`
// и может "окружить" каждый элемент массива значением`surroundBy`
func JoinByCharacter(arr []string, delim string, surroundBy string) string {
	result := ""

	arrLen := len(arr)

	for i, item := range arr {

		result += surroundBy + item + surroundBy

		if i < arrLen-1 {
			result += delim
		}
	}

	return result
}
