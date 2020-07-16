//Package geo supports distance calculation between two positions on the Earth.
//Positions can be provided in two froms: latitude/longitude or as Maidenhead QTH locator
package geo

import (
	"errors"
	"fmt"
	"github.com/golang/geo/s2"
	"math"
	"regexp"
	"strings"
)

type QTH struct {
	Loc    string    // Maidenhead QTH Locator
	LatLon LatLonDeg // LatLon represent a point as a pair of latitude and longitude degrees
	LatLng s2.LatLng // LatLng represents a point on the unit sphere as a pair of angles.
}

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.00001
}

func (a QTH) String() string {
	return fmt.Sprintf("[ %s %s {%.6f %.6f} ]", a.Loc, a.LatLon.String(), a.LatLng.Lat.Radians(), a.LatLng.Lng.Radians())
}

func QthEqual(a, b QTH) bool {
	eq := a.Loc == b.Loc
	if eq {
		eq = almostEqual(a.LatLon.Lat, b.LatLon.Lat)
	} else {
		return false
	}
	if eq {
		eq = almostEqual(a.LatLon.Lon, b.LatLon.Lon)
	} else {
		return false
	}
	if eq {
		eq = almostEqual(a.LatLng.Lat.Radians(), b.LatLng.Lat.Radians())
	} else {
		return false
	}
	if eq {
		eq = almostEqual(a.LatLng.Lng.Radians(), b.LatLng.Lng.Radians())
	} else {
		return false
	}
	return eq
}

const earthRadiusKm = 6371.0088 // mean Earth radius in km

// returns distance between two Maindenhead locators in km
func DistanceLocator(locatorA string, locatorB string) (float64, error) {
	a, err := NewQthFromLOC(locatorA)
	b, err := NewQthFromLOC(locatorB)
	if err != nil {
		return 0, err
	} else {
		d := a.LatLng.Distance(b.LatLng)
		return d.Radians() * earthRadiusKm, nil
	}
}

// returns Distance between two QTH variables in km
func DistanceQTH(a, b QTH) float64 {
	d := a.LatLng.Distance(b.LatLng)
	return d.Radians() * earthRadiusKm
}

func illegalArgumentError(arg string) error {
	return errors.New(fmt.Sprintf("Illegal argumet value! qthLocator=%s", arg))
}

// This function creates new QTH variable from provided QTH locator
// or returns error if QTH locator is wrong formatted
func NewQthFromLOC(qthLocator string) (QTH, error) {
	qthLocator = strings.ToUpper(qthLocator)
	qth := QTH{}
	switch len(qthLocator) {
	case 6:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}[A-X]{2}$`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				s := squareDecode(latLonChar{
					latChar: qthLocator[3],
					lonChar: qthLocator[2],
				})
				ss := subsquareDecode(latLonChar{
					latChar: qthLocator[5],
					lonChar: qthLocator[4],
				})
				lat := f.decoded.Lat + s.decoded.Lat + ss.decoded.Lat/60 + 0.02083333 // 1.25' / 60
				lon := f.decoded.Lon + s.decoded.Lon + ss.decoded.Lon/60 + 0.04166667 // 2.5' / 60
				return QTH{
					Loc:    qthLocator,
					LatLon: LatLonDeg{lat, lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}
	case 4:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}$`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				s := squareDecode(latLonChar{
					latChar: qthLocator[3],
					lonChar: qthLocator[2],
				})
				lat := f.decoded.Lat + s.decoded.Lat + 0.5
				lon := f.decoded.Lon + s.decoded.Lon + 1
				return QTH{
					Loc:    qthLocator,
					LatLon: LatLonDeg{lat, lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}
	case 2:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}`, qthLocator); matched {
				f := fieldDecode(latLonChar{
					latChar: qthLocator[1],
					lonChar: qthLocator[0],
				})
				lat := f.decoded.Lat + 5
				lon := f.decoded.Lon + 10

				return QTH{
					Loc:    qthLocator,
					LatLon: LatLonDeg{lat, lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil

			} else {
				return qth, illegalArgumentError(qthLocator)
			}
		}

	default:
		return qth, illegalArgumentError(qthLocator)
	}
}

// This function creates new QTH variable from provided latitude, longitude
// or returns error if QTH locator is wrong formatted
func NewQthFromLatLon(latitude, longitude float64) (QTH, error) {
	lld := LatLonDeg{
		Lat: latitude,
		Lon: longitude,
	}
	if math.Abs(latitude) > 90 || math.Abs(longitude) > 180 {
		return QTH{}, illegalArgumentError(lld.String())
	}
	f, s, ss := subsquareEncode(lld)
	return QTH{
		Loc:    f.encoded.getLonChar() + f.encoded.getLatChar() + s.encoded.getLonChar() + s.encoded.getLatChar() + ss.encoded.getLonChar() + ss.encoded.getLatChar(),
		LatLon: lld,
		LatLng: s2.LatLngFromDegrees(latitude, longitude),
	}, nil
}
