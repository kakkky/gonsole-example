package utils

import "fmt"

// 動物関連のユーティリティ関数

// 動物の年齢を人間の年齢に換算
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

// 動物の名前をフォーマット
func FormatAnimalName(name string) string {
	return fmt.Sprintf("🐾 %s", name)
}
