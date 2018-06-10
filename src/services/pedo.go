package services

import (
	"strconv"
	"github.com/rafaelcam/go-pedo-api/src/validators"
)

// Calls the validator and if the parameters are valid
// prints the numbers from `start` to` end` with the following rules:
//
// For multiples of 3, print the word 'Pé' instead of the number.
// For multiples of 5, print the word 'Do' instead of the number.
// For multiples of both 3 and 5, print 'PéDo'.
func GetPeDoByRange(start int, end int) ([]string, error) {
	if err := validators.ValidateRange(start, end); err != nil {
		return nil, err
	}

	var result []string
	for number := start; number <= end; number++ {
		result = append(result, getPeDoByNumber(number))
	}

	return result, nil
}

func getPeDoByNumber(number int) string {
	if number % 15 == 0 {
		return "PéDo"
	} else if number % 3 == 0 {
		return "Pé"
	} else if number % 5 == 0 {
		return "Do"
	} else {
		return strconv.Itoa(number)
	}
}