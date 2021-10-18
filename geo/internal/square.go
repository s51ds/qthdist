package internal

import (
	"fmt"
)

var (
	squareDigitToLetterLat map[int]string
	squareLetterToDigitLat map[string]float64

	squareDigitToLetterLon map[int]string
	squareLetterToDigitLon map[string]float64

	squareDegLatitudes = [...]float64{
		0,
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
	}

	squareDegLongitudes = [...]float64{
		0,
		2,
		4,
		6,
		8,
		10,
		12,
		14,
		16,
		18,
	}
)

func init() {

	squareDigitToLetterLat = map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "3",
		4: "4",
		5: "5",
		6: "6",
		7: "7",
		8: "8",
		9: "9",
	}
	squareLetterToDigitLat = map[string]float64{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	squareDigitToLetterLon = map[int]string{
		0:  "0",
		1:  "0",
		2:  "1",
		3:  "1",
		4:  "2",
		5:  "2",
		6:  "3",
		7:  "3",
		8:  "4",
		9:  "4",
		10: "5",
		11: "5",
		12: "6",
		13: "6",
		14: "7",
		15: "7",
		16: "8",
		17: "8",
		18: "9",
		19: "9",
	}

	squareLetterToDigitLon = map[string]float64{
		"0": 0,
		"1": 2,
		"2": 4,
		"3": 6,
		"4": 8,
		"5": 10,
		"6": 12,
		"7": 14,
		"8": 16,
		"9": 18,
	}

}

type Square struct {
	// characters {0,1,...9} Decoded as
	// longitude {0,2...,18} [degree]
	// latitude {0,1...,9)   [degree]
	Decoded LatLonDeg  //characters Decoded as longitude and latitude
	Encoded LatLonChar //latitude and longitude Encoded as characters
}

func (a Square) String() string {
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

func SquareEncode(lld LatLonDeg) (Field, Square) {

	s := Square{}
	f := FieldEncode(lld)
	iLat, iLon := 0, 0

	fLat := lld.Lat - f.Decoded.Lat
	fLon := lld.Lon - f.Decoded.Lon

	for _, v := range squareDegLongitudes {
		if fLon >= v && fLon < v+2 {
			iLon = int(v)
			break
		}
	}

	for _, v := range squareDegLatitudes {
		if fLat >= v && fLat < v+1 {
			iLat = int(v)
			break
		}
	}

	s.Encoded.setLatChar(squareDigitToLetterLat[iLat])
	s.Encoded.setLonChar(squareDigitToLetterLon[iLon])
	s.Decoded.Lat = float64(iLat)
	s.Decoded.Lon = float64(iLon)
	return f, s
}

func SquareDecode(llc LatLonChar) Square {
	s := Square{}
	s.Decoded.Lat = squareLetterToDigitLat[llc.GetLatChar()]
	s.Decoded.Lon = squareLetterToDigitLon[llc.GetLonChar()]
	s.Encoded = llc
	return s
}
