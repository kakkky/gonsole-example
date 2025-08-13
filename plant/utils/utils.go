package utils

import "fmt"

// Plant-related utility functions

// GetGrowthStage determines the plant's growth stage based on age
func GetGrowthStage(age int) string {
	switch {
	case age < 1:
		return "Seedling"
	case age < 3:
		return "Young Plant"
	case age < 10:
		return "Mature Plant"
	default:
		return "Old Plant"
	}
}

// FormatPlantName formats plant name with emoji
func FormatPlantName(name string) string {
	return fmt.Sprintf("🌱 %s", name)
}
