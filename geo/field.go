package geo

import (
	"fmt"
)

type field struct {
	// characters {A,B,...R} decoded as
	// longitude {-180,-160...,160}
	// latitude {-90,-80...,80)
	decoded LatLonDeg  //characters decoded as longitude and latitude
	encoded latLonChar //latitude and longitude encoded as characters
}

func (a *field) String() string {
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

func (a *field) Equals(b field) bool {
	return a.encoded.Equal(b.encoded) && a.decoded.Equal(b.decoded)
}

func fieldEncode(lld LatLonDeg) field {

	aLat := [...]float64{
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
	mLat := map[int]string{
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

	aLon := [...]float64{
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
	mLon := map[int]string{
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

	a := field{}

	iLat, iLon := 0, 0
	for _, v := range aLon {
		if lld.Lon >= v && lld.Lon < v+20 {
			iLon = int(v)
			//fmt.Printf("lld.Lon=%f iLon=%d \n", lld.Lon, iLon)
			break
		}
	}
	for _, v := range aLat {
		if lld.Lat >= v && lld.Lat < v+10 {
			iLat = int(v)
			//fmt.Printf("lld.Lat=%f iLat=%d \n", lld.Lat, iLat)
			break
		}
	}

	a.encoded.setLatChar(mLat[iLat])
	a.encoded.setLonChar(mLon[iLon])
	a.decoded.Lat = float64(iLat)
	a.decoded.Lon = float64(iLon)
	return a
}

func fieldDecode(llc latLonChar) field {
	a := field{}
	mLat := map[string]float64{
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
	mLon := map[string]float64{
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
	a.decoded.Lat = mLat[llc.getLatChar()]
	a.decoded.Lon = mLon[llc.getLonChar()]
	a.encoded = llc
	return a
}
