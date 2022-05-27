package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var carSelected string

	if option1 < option2 {
		carSelected = option1
	} else {
		carSelected = option2
	}

	return carSelected + " is clearly the better choice."

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var costPercentage float64

	if int(age) < 3 {
		costPercentage = float64(80) / float64(100)
	} else if int(age) >= 10 {
		costPercentage = float64(50) / float64(100)
	} else {
		costPercentage = float64(70) / float64(100)
	}

	return costPercentage * originalPrice
}
