package internal

import (
	"fmt"
	"log"
	"math"
	"strings"
)

//
// Decimal degrees
type LatLonDeg struct {
	Lat, Lon float64
}

func (a *LatLonDeg) String() string {
	return fmt.Sprintf("{%.6f %.6f}", a.Lat, a.Lon)
}

func (a *LatLonDeg) equal(b LatLonDeg) bool {
	return a.String() == b.String()
}

func (a *LatLonDeg) ToLatLonDMS() LatLonDMS {
	dms := LatLonDMS{}

	intLatDeg, fracLatDeg := math.Modf(a.Lat)
	intLatMin, fracLatMin := math.Modf(fracLatDeg * 60)
	intLatSec, _ := math.Modf(fracLatMin * 60)
	dms.latDMS.degrees = intLatDeg
	//	dms.latDMS.minutesDec = math.Abs(fracLatDeg*60)
	dms.latDMS.minutes = int(math.Abs(intLatMin))
	dms.latDMS.seconds = int(math.Abs(intLatSec))

	intLonDeg, fracLonDeg := math.Modf(a.Lon)
	intLonMin, fracLonMin := math.Modf(fracLonDeg * 60)
	intLonSec, _ := math.Modf(fracLonMin * 60)
	dms.lonDMS.degrees = intLonDeg
	//	dms.lonDMS.minutesDec = math.Abs(fracLonDeg * 60)
	dms.lonDMS.minutes = int(math.Abs(intLonMin))
	dms.lonDMS.seconds = int(math.Abs(intLonSec))

	return dms
}

//
// Degrees, Minutes, Seconds
type dms struct {
	degrees float64
	minutes int
	seconds int
}

func (a *dms) String() string {
	return fmt.Sprintf(`%.f°%d'%d"`, a.degrees, a.minutes, a.seconds)
}

type LatLonDMS struct {
	latDMS, lonDMS dms
}

func (a *LatLonDMS) String() string {
	return fmt.Sprintf("Lat=%s, Lon=%s", a.latDMS.String(), a.lonDMS.String())
}

//
// Maidenhead Encoded
type LatLonChar struct {
	LatChar, LonChar byte
}

func isNumber(b byte) bool {
	return b > 47 && b < 58
}

func isLetter(b byte) bool {
	return b > 64 && b < 91
}

func toValidValue(c string) byte {
	if len(c) != 1 {
		log.Printf("Illegal argument, only one character is accepted! argumetn=%s", c)
		return 32 //empty string
	}
	c = strings.ToUpper(c)
	b := byte(c[0])
	if isLetter(b) || isNumber(b) {
		return b
	} else {
		log.Printf("Illegal character:%s", c)
		return 32
	}
}

func (a *LatLonChar) setLatChar(c string) {
	a.LatChar = toValidValue(c)
}

func (a *LatLonChar) setLonChar(c string) {
	a.LonChar = toValidValue(c)
}

func (a *LatLonChar) GetLatChar() string {
	return string(a.LatChar)
}

func (a *LatLonChar) GetLonChar() string {
	return string(a.LonChar)
}

func (a *LatLonChar) isSet() bool {
	return a.LatChar > 0 && a.LonChar > 0
}

func (a *LatLonChar) String() string {
	if a.isSet() {
		return fmt.Sprintf("%s%s", a.GetLonChar(), a.GetLatChar())
	} else {
		return ""
	}
}

func (a *LatLonChar) equal(b LatLonChar) bool {
	return a.String() == b.String()
}
