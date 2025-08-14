package utils

import (
	"strings"
)

func FindBiggestMatchingSuffix(elements []string, target string) (result string) {
	var maxSuffixLength int = len(result)
	for _, element := range elements {
		if strings.HasSuffix(target, element) {
			if len(element) > maxSuffixLength {
				maxSuffixLength = len(element)
				result = element
			}
		}
	}

	InfoLogger.Printf("Best matching suffix: %s for target: %s\n", result, target)
	return result
}

func FindBiggestMatchingPrefix(elements []string, target string) (result string) {
	var maxPrefixLength int = len(result)
	for _, element := range elements {
		if strings.HasPrefix(target, element) {
			if len(element) > maxPrefixLength {
				maxPrefixLength = len(element)
				result = element
			}
		}
	}

	InfoLogger.Printf("Best matching prefix: %s for target: %s\n", result, target)
	return result
}