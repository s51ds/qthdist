package geo

import (
	"math"
)

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

// Distance returns distance in km and azimuth in decimal degrees from a to b
func (a *QTH) DistanceAndAzimuth(b *QTH) (dist, azimuth float64) {
	d := a.LatLng.Distance(b.LatLng)
	return d.Radians() * earthRadiusKm, a.AzimuthTo(b)
}

const (
	rad = 180.0 / math.Pi
)

// AzimuthTo Calculates forward azimuth in decimal degrees from a to b
// Original Implementation from: http://www.movable-type.co.uk/scripts/latlong.html
func (a *QTH) AzimuthTo(b *QTH) float64 {

	diffInLongitude := (b.LatLng.Lng - a.LatLng.Lng).Radians()

	lat1 := a.LatLng.Lat.Radians()
	lat2 := b.LatLng.Lat.Radians()

	y := math.Sin(diffInLongitude) * math.Cos(lat2)
	x := math.Cos(lat1)*math.Sin(lat2) - math.Sin(lat1)*math.Cos(lat2)*math.Cos(diffInLongitude)

	azimDeg := math.Atan2(y, x) * rad
	if azimDeg < 0 {
		return 360.0 + azimDeg
	} else {
		return azimDeg
	}
}
