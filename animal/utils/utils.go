package utils

import "fmt"

// Animal-related utility functions

// ConvertToHumanAge converts animal age to human equivalent age
func ConvertToHumanAge(animalAge int, animalType string) int {
	switch animalType {
	case "dog":
		return animalAge * 7
	case "cat":
		return animalAge * 6
	default:
		return animalAge * 5
	}
}

// FormatAnimalName formats animal name with emoji
func FormatAnimalName(name string) string {
	return fmt.Sprintf("🐾 %s", name)
}
