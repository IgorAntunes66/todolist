package todo

func IsLetter(texto string) bool {
	for i := range len(texto) {
		if int(texto[i]) >= 65 && int(texto[i]) <= 90 || int(texto[i]) >= 97 && int(texto[i]) <= 122 || int(texto[1]) == 32 {
			continue
		} else {
			return false
		}
	}
	return true
}
