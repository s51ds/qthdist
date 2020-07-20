package geo

const earthRadiusKm = 6371.0088 // mean Earth radius in km

// DistanceLocator returns distance between two Maidenhead locators in km
// and any error encountered
func DistanceLocator(locatorA string, locatorB string) (float64, error) {
	a, err := MakeQthFromLOC(locatorA)
	b, err := MakeQthFromLOC(locatorB)
	if err != nil {
		return 0, err
	} else {
		d := a.LatLng.Distance(b.LatLng)
		return d.Radians() * earthRadiusKm, nil
	}
}

// DistanceQTH returns distance between a and b in km
func DistanceQTH(a, b QTH) float64 {
	d := a.LatLng.Distance(b.LatLng)
	return d.Radians() * earthRadiusKm
}

// Distance returns distance between a and b in km
func (a *QTH) Distance(b *QTH) float64 {
	d := a.LatLng.Distance(b.LatLng)
	return d.Radians() * earthRadiusKm
}
