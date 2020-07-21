package internal

import (
	"fmt"
)

var (
	fieldDigitToLetterLat map[int]string
	fieldDigitToLetterLon map[int]string

	fieldLetterToDigitLat map[string]float64
	fieldLetterToDigitLon map[string]float64

	fieldDegLatitudes = [...]float64{
		-90,
		-80,
		-70,
		-60,
		-50,
		-40,
		-30,
		-20,
		-10,
		0,
		10,
		20,
		30,
		40,
		50,
		60,
		70,
		80,
	}
	fieldDegLongitudes = [...]float64{
		-180,
		-160,
		-140,
		-120,
		-100,
		-80,
		-60,
		-40,
		-20,
		0,
		20,
		40,
		60,
		80,
		100,
		120,
		140,
		160,
	}
)

func init() {
	fieldDigitToLetterLat = map[int]string{
		-90: "A",
		-80: "B",
		-70: "C",
		-60: "D",
		-50: "E",
		-40: "F",
		-30: "G",
		-20: "H",
		-10: "I",
		0:   "J",
		10:  "K",
		20:  "L",
		30:  "M",
		40:  "N",
		50:  "O",
		60:  "P",
		70:  "Q",
		80:  "R",
	}

	fieldDigitToLetterLon = map[int]string{
		-180: "A",
		-170: "A",
		-160: "B",
		-150: "B",
		-140: "C",
		-130: "C",
		-120: "D",
		-110: "D",
		-100: "E",
		-90:  "E",
		-80:  "F",
		-70:  "F",
		-60:  "G",
		-50:  "G",
		-40:  "H",
		-30:  "H",
		-20:  "I",
		-10:  "I",
		0:    "J",
		10:   "J",
		20:   "K",
		30:   "K",
		40:   "L",
		50:   "L",
		60:   "M",
		70:   "M",
		80:   "N",
		90:   "N",
		100:  "O",
		110:  "O",
		120:  "P",
		130:  "P",
		140:  "Q",
		150:  "Q",
		160:  "R",
		170:  "R",
	}

	fieldLetterToDigitLat = map[string]float64{
		"A": -90,
		"B": -80,
		"C": -70,
		"D": -60,
		"E": -50,
		"F": -40,
		"G": -30,
		"H": -20,
		"I": -10,
		"J": 0,
		"K": 10,
		"L": 20,
		"M": 30,
		"N": 40,
		"O": 50,
		"P": 60,
		"Q": 70,
		"R": 80,
	}

	fieldLetterToDigitLon = map[string]float64{
		"A": -180,
		"B": -160,
		"C": -140,
		"D": -120,
		"E": -100,
		"F": -80,
		"G": -60,
		"H": -40,
		"I": -20,
		"J": 0,
		"K": 20,
		"L": 40,
		"M": 60,
		"N": 80,
		"O": 100,
		"P": 120,
		"Q": 140,
		"R": 160,
	}
}

type Field struct {
	// characters {A,B,...R} Decoded as
	// longitude {-180,-160...,160}
	// latitude {-90,-80...,80)
	Decoded LatLonDeg  //characters Decoded as longitude and latitude
	Encoded LatLonChar //latitude and longitude Encoded as characters
}

func (a *Field) String() string {
	s := ""
	if a.Decoded.String() != "" {
		s = fmt.Sprintf("Decoded:%s", a.Decoded.String())
	}
	if a.Encoded.String() != "" {
		if s == "" {
			s = fmt.Sprintf("Encoded:%s", a.Encoded.String())
		} else {
			s += fmt.Sprintf(" Encoded:%s", a.Encoded.String())
		}
	}
	return s
}

func FieldEncode(lld LatLonDeg) Field {
	a := Field{}

	iLat, iLon := 0, 0
	for _, v := range fieldDegLongitudes {
		if lld.Lon >= v && lld.Lon < v+20 {
			iLon = int(v)
			break
		}
	}
	for _, v := range fieldDegLatitudes {
		if lld.Lat >= v && lld.Lat < v+10 {
			iLat = int(v)
			break
		}
	}

	a.Encoded.setLatChar(fieldDigitToLetterLat[iLat])
	a.Encoded.setLonChar(fieldDigitToLetterLon[iLon])
	a.Decoded.Lat = float64(iLat)
	a.Decoded.Lon = float64(iLon)
	return a
}

func FieldDecode(llc LatLonChar) Field {
	a := Field{}
	a.Decoded.Lat = fieldLetterToDigitLat[llc.GetLatChar()]
	a.Decoded.Lon = fieldLetterToDigitLon[llc.GetLonChar()]
	a.Encoded = llc
	return a
}
