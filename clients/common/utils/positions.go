package utils

import (
	"github.com/defiants-co/perpstream-go-2/clients/common/models"
)

func FuturePositionSetsAreEqual(oldPositions []models.Future, newPositions []models.Future) bool {
	if len(oldPositions) != len(newPositions) {
		return false
	}

	for _, o := range oldPositions {
		hasMatch := false
		for _, n := range newPositions {
			if o.Equal(&n) {
				hasMatch = true
			}
		}

		if !hasMatch {
			return false
		}
	}

	return true
}

func OptionPositionSetsAreEqual(oldPositions []models.Option, newPositions []models.Option) bool {
	if len(oldPositions) != len(newPositions) {
		return false
	}

	for _, o := range oldPositions {
		hasMatch := false
		for _, n := range newPositions {
			if o.Equal(&n) {
				hasMatch = true
			}
		}

		if !hasMatch {
			return false
		}
	}

	return true
}
