package geo

const earthRadiusKm = 6371.0088 // mean Earth radius in km

// LocatorDistance returns distance between two Maidenhead locators in km
// and any error encountered
// LocatorDistance is mainly used in tests. Preferred function for distance
// calculation is Distance
func LocatorDistance(locatorA string, locatorB string) (float64, error) {
	var err error
	qthA, err := NewQthFromLocator(locatorA)
	if err != nil {
		return 0, err
	}
	qthB, err := NewQthFromLocator(locatorB)
	if err != nil {
		return 0, err
	}
	d := qthA.LatLng.Distance(qthB.LatLng)
	return d.Radians() * earthRadiusKm, nil
}

// Distance returns distance between a and b in km
func (a *QTH) Distance(b *QTH) float64 {
	d := a.LatLng.Distance(b.LatLng)
	return d.Radians() * earthRadiusKm
}
