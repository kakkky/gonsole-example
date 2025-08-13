package utils

import "fmt"

// Vehicle-related utility functions

// CalculateVehicleAge calculates vehicle age from model year
func CalculateVehicleAge(year int) int {
	currentYear := 2025
	return currentYear - year
}

// CalculateFuelEfficiency calculates fuel efficiency (simple version)
func CalculateFuelEfficiency(distance, fuel float64) float64 {
	if fuel == 0 {
		return 0
	}
	return distance / fuel
}

// FormatVehicleName formats vehicle name with emoji
func FormatVehicleName(brand, model string) string {
	return fmt.Sprintf("🚗 %s %s", brand, model)
}
