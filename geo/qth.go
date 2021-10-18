package geo

import (
	"fmt"
	"github.com/golang/geo/s2"
	"github.com/s51ds/qthdist/geo/internal"
	"math"
	"regexp"
	"strings"
)

// QTH represents Maidenhead QTH Locator and associated a pair of latitude and longitude
// Use function NewQthFromLocator or NewQthFromPosition to create QTH
type QTH struct {
	Loc    string             // Maidenhead QTH Locator
	LatLon internal.LatLonDeg // LatLon represent a point as a pair of latitude and longitude degrees
	LatLng s2.LatLng          // LatLng represents a point on the unit sphere as a pair of angles.
}

// NewQthFromLocator returns QTH for locator. Locator is case-insensitive 2, 4 or 6 characters; example:
// - JN
// - jN76
// - jn76TO
func NewQthFromLocator(locator string) (*QTH, error) {
	// TODO: maybe pointer to QTH is not good idea
	locator = strings.ToUpper(locator)
	qth := QTH{}
	switch len(locator) {
	case 6:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}[A-X]{2}$`, locator); matched {
				f := internal.FieldDecode(internal.LatLonChar{
					LatChar: locator[1],
					LonChar: locator[0],
				})
				s := internal.SquareDecode(internal.LatLonChar{
					LatChar: locator[3],
					LonChar: locator[2],
				})
				ss := internal.SubsquareDecode(internal.LatLonChar{
					LatChar: locator[5],
					LonChar: locator[4],
				})
				lat := f.Decoded.Lat + s.Decoded.Lat + ss.Decoded.Lat/60 + 0.02083333 // 1.25' / 60
				lon := f.Decoded.Lon + s.Decoded.Lon + ss.Decoded.Lon/60 + 0.04166667 // 2.5' / 60
				return &QTH{
					Loc:    locator,
					LatLon: internal.LatLonDeg{Lat: lat, Lon: lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return &qth, internal.IllegalLocatorError(locator)
			}
		}
	case 4:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}[0-9]{2}$`, locator); matched {
				f := internal.FieldDecode(internal.LatLonChar{
					LatChar: locator[1],
					LonChar: locator[0],
				})
				s := internal.SquareDecode(internal.LatLonChar{
					LatChar: locator[3],
					LonChar: locator[2],
				})
				lat := f.Decoded.Lat + s.Decoded.Lat + 0.5
				lon := f.Decoded.Lon + s.Decoded.Lon + 1
				return &QTH{
					Loc:    locator,
					LatLon: internal.LatLonDeg{Lat: lat, Lon: lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil
			} else {
				return &qth, internal.IllegalLocatorError(locator)
			}
		}
	case 2:
		{
			if matched, _ := regexp.MatchString(`^[A-R]{2}`, locator); matched {
				f := internal.FieldDecode(internal.LatLonChar{
					LatChar: locator[1],
					LonChar: locator[0],
				})
				lat := f.Decoded.Lat + 5
				lon := f.Decoded.Lon + 10

				return &QTH{
					Loc:    locator,
					LatLon: internal.LatLonDeg{Lat: lat, Lon: lon},
					LatLng: s2.LatLngFromDegrees(lat, lon),
				}, nil

			} else {
				return &qth, internal.IllegalLocatorError(locator)
			}
		}

	default:
		return &qth, internal.IllegalLocatorError(locator)
	}
}

// NewQthFromPosition returns QTH for latitude and longitude in decimal degrees
func NewQthFromPosition(latitude, longitude float64) (*QTH, error) {
	lld := internal.LatLonDeg{
		Lat: latitude,
		Lon: longitude,
	}
	if math.Abs(latitude) > 90 || math.Abs(longitude) > 180 {
		return &QTH{}, internal.IllegalLocatorError(lld.String())
	}
	f, s, ss := internal.SubsquareEncode(lld)
	return &QTH{
		Loc:    f.Encoded.GetLonChar() + f.Encoded.GetLatChar() + s.Encoded.GetLonChar() + s.Encoded.GetLatChar() + ss.Encoded.GetLonChar() + ss.Encoded.GetLatChar(),
		LatLon: lld,
		LatLng: s2.LatLngFromDegrees(latitude, longitude),
	}, nil
}

func (a *QTH) String() string {
	return fmt.Sprintf("[ %s %s {%.6f %.6f} ]", a.Loc, a.LatLon.String(), a.LatLng.Lat.Radians(), a.LatLng.Lng.Radians())
}
