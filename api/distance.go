package api

import "github.com/s51ds/qthdist/geo"

func Distance(locatorA, locatorB string) (distance, azimuth float64, err error) {

	var qthA, qthB geo.QTH
	qthA, err = geo.NewQthFromLocator(locatorA)
	if err != nil {
		return
	}
	qthB, err = geo.NewQthFromLocator(locatorB)
	if err != nil {
		return
	}
	distance, azimuth = qthA.DistanceAndAzimuth(qthB)

	return
}
