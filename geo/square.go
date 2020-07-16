package geo

import "fmt"

type square struct {
	// characters {0,1,...9} decoded as
	// longitude {0,2...,18} [degree]
	// latitude {0,1...,9)   [degree]
	decoded LatLonDeg  //characters decoded as longitude and latitude
	encoded latLonChar //latitude and longitude encoded as characters
}

func (a *square) String() string {
	s := ""
	if a.decoded.String() != "" {
		s = fmt.Sprintf("Decoded:%s", a.decoded.String())
	}
	if a.encoded.String() != "" {
		if s == "" {
			s = fmt.Sprintf("Encoded:%s", a.encoded.String())
		} else {
			s += fmt.Sprintf(" Encoded:%s", a.encoded.String())
		}
	}
	return s
}

func (a *square) Equals(b square) bool {
	return a.encoded.Equal(b.encoded) && a.decoded.Equal(b.decoded)
}

func squareEncode(lld LatLonDeg) (field, square) {

	aLat := [...]float64{
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
	mLat := map[int]string{
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
	aLon := [...]float64{
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
	mLon := map[int]string{
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

	s := square{}
	f := fieldEncode(lld)
	iLat, iLon := 0, 0

	fLat := lld.Lat - f.decoded.Lat
	fLon := lld.Lon - f.decoded.Lon

	for _, v := range aLon {
		if fLon >= v && fLon < v+2 {
			iLon = int(v)
			break
		}
	}

	for _, v := range aLat {
		if fLat >= v && fLat < v+1 {
			iLat = int(v)
			break
		}
	}

	s.encoded.setLatChar(mLat[iLat])
	s.encoded.setLonChar(mLon[iLon])
	s.decoded.Lat = float64(iLat)
	s.decoded.Lon = float64(iLon)
	return f, s
}

func squareDecode(llc latLonChar) square {
	s := square{}
	mLat := map[string]float64{
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
	mLon := map[string]float64{
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
	s.decoded.Lat = mLat[llc.getLatChar()]
	s.decoded.Lon = mLon[llc.getLonChar()]
	s.encoded = llc
	return s
}
