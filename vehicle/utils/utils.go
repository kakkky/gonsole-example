package utils

import "fmt"

// 乗り物関連のユーティリティ関数

// 年式から車の年数を計算
func CalculateVehicleAge(year int) int {
	currentYear := 2025
	return currentYear - year
}

// 燃費を計算（簡易版）
func CalculateFuelEfficiency(distance, fuel float64) float64 {
	if fuel == 0 {
		return 0
	}
	return distance / fuel
}

// 乗り物の名前をフォーマット
func FormatVehicleName(brand, model string) string {
	return fmt.Sprintf("🚗 %s %s", brand, model)
}
