package services

import "strings"

const (
    alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    length   = uint64(len(alphabet))
)

func encode(number uint64) string {
    if number == 0 {
        return string(alphabet[0])
    }
    result := make([]byte, 0)
    for ; number > 0; number = number / length {
        result = append(result, alphabet[number%length])
    }
    reverse(result)
    return string(result)
}

func decode(str string) uint64 {
    var result uint64
    for _, c := range []byte(str) {
        index := strings.IndexByte(alphabet, c)
        if index < 0 {
            return 0
        }
        result = result*length + uint64(index)
    }
    return result
}

func reverse(arr []byte) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
