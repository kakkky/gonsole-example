package utils

import "fmt"

// 植物関連のユーティリティ関数

// 植物の成長段階を判定
func GetGrowthStage(age int) string {
	switch {
	case age < 1:
		return "苗"
	case age < 3:
		return "若木"
	case age < 10:
		return "成木"
	default:
		return "古木"
	}
}

// 植物の名前をフォーマット
func FormatPlantName(name string) string {
	return fmt.Sprintf("🌱 %s", name)
}
