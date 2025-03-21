package utils

import "slices"

func IsExist(arr []string, elm string) bool {
	return slices.Contains(arr, elm)
}
