package geo

import (
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

// MakeQthFromLOC returns QTH for qthLocator
func MakeQthFromLOC(qthLocator string) (QTH, error) {
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
				return qth, illegalLocatorError(qthLocator)
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
				return qth, illegalLocatorError(qthLocator)
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
				return qth, illegalLocatorError(qthLocator)
			}
		}

	default:
		return qth, illegalLocatorError(qthLocator)
	}
}

// MakeQthFromLatLon returns QTH for latitude, longitude
func MakeQthFromLatLon(latitude, longitude float64) (QTH, error) {
	lld := LatLonDeg{
		Lat: latitude,
		Lon: longitude,
	}
	if math.Abs(latitude) > 90 || math.Abs(longitude) > 180 {
		return QTH{}, illegalLocatorError(lld.String())
	}
	f, s, ss := subsquareEncode(lld)
	return QTH{
		Loc:    f.encoded.getLonChar() + f.encoded.getLatChar() + s.encoded.getLonChar() + s.encoded.getLatChar() + ss.encoded.getLonChar() + ss.encoded.getLatChar(),
		LatLon: lld,
		LatLng: s2.LatLngFromDegrees(latitude, longitude),
	}, nil
}

func (a QTH) String() string {
	return fmt.Sprintf("[ %s %s {%.6f %.6f} ]", a.Loc, a.LatLon.String(), a.LatLng.Lat.Radians(), a.LatLng.Lng.Radians())
}
