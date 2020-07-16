package geo

import (
	"fmt"
	"math"
)

type subsquare struct {
	// characters {A,B,...,X} decoded as
	// longitude {0,5,10...,55} [minute]
	// latitude {0.0, 2.5, 5.0,..., 110, 115)  [minute]
	decoded LatLonDeg  //characters decoded as longitude and latitude
	encoded latLonChar //latitude and longitude encoded as characters
}

func (a *subsquare) String() string {
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

func (a *subsquare) Equals(b square) bool {
	return a.encoded.Equal(b.encoded) && a.decoded.Equal(b.decoded)
}

func subsquareEncode(lld LatLonDeg) (field, square, subsquare) {
	fld, sqr := squareEncode(lld)
	subsqr := subsquare{}

	latMinutes := math.Abs(fld.decoded.Lat+sqr.decoded.Lat-lld.Lat) * 60
	lonMinutes := math.Abs(fld.decoded.Lon+sqr.decoded.Lon-lld.Lon) * 60

	subsquareLetters := [24]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X"}

	//SUBSQUARE LAT
	// 65  66    67   68   69   70    71     72    73    74    75    76    77  78     79     80    81    82   83    84     85    86    87   88    [ASCII]
	//
	// 0    1    2    3    4     5     6     7     8     9     10    11    12  13     14     15    16    17   18    19     20    21    22   23   [IDX]
	//"A", "B", "C", "D", "E",  "F",  "G",  "H",  "I",  "J",  "K",   "L", "M", "N",   "O",  "P",  "Q",  "R",  "S",  "T",  "U",  "V",  "W", "X"
	//0.0, 2.5, 5.0, 7.5, 10.0, 12.5, 15.0, 17.5, 20.0, 22.5, 25.0, 27.5, 30.0, 32.5, 35.0, 37.5, 40.0, 42.5, 45.0, 47.5, 50.0, 52.5, 55.0, 57.5 [MINUTES]
	//
	//SQUARE LAT
	//Y                                                                                                                                           Y+1 [DEG]
	//Y:{0,1,2,....9}
	//
	//FIELD LAT: {-90,-80,-70,...,80}

	aLat := [24]float64{0.0, 2.5, 5.0, 7.5, 10.0, 12.5, 15.0, 17.5, 20.0, 22.5, 25.0, 27.5, 30.0, 32.5, 35.0, 37.5, 40.0, 42.5, 45.0, 47.5, 50.0, 52.5, 55.0, 57.5}
	for i, v := range aLat {
		if latMinutes >= v && latMinutes < v+2.5 {
			subsqr.encoded.latChar = byte(subsquareLetters[i][0])
			subsqr.decoded.Lat = v
			break
		}
	}

	//SUBSQUARE LON
	// 65  66    67   68   69   70    71     72    73    74    75    76    77  78     79     80    81    82   83    84     85    86    87   88    [ASCII]
	//
	// 0    1    2    3    4     5     6     7     8     9     10    11    12  13     14     15    16    17   18    19     20    21    22   23   [IDX]
	//"A", "B", "C", "D", "E",  "F",  "G",  "H",  "I",  "J",  "K",   "L", "M", "N",   "O",  "P",  "Q",  "R",  "S",  "T",  "U",  "V",  "W", "X"
	// 0,   5,  10,   15, 20,   25,   30,    35,  40,    45,   50,   55,   60, 65,    70,   75,   80,    85,  90,   95,   100,  105,  110, 115   [MINUTES]
	//
	//SQUARE LON
	//X                                                                                                                                           X+2 [DEG]
	//X:{0,2,4,...,18}
	//
	//FIELD LON: {-180,-160,-140,...,160}

	aLon := [24]float64{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115}
	for i, v := range aLon {
		if lonMinutes >= v && lonMinutes < v+5 {
			subsqr.encoded.lonChar = byte(subsquareLetters[i][0])
			subsqr.decoded.Lon = v
			break
		}
	}

	return fld, sqr, subsqr
}

func subsquareDecode(llc latLonChar) subsquare {
	s := subsquare{}
	mLat := map[string]float64{
		"A": 0.0,
		"B": 2.5,
		"C": 5.0,
		"D": 7.5,
		"E": 10.0,
		"F": 12.5,
		"G": 15.0,
		"H": 17.5,
		"I": 20.0,
		"J": 22.5,
		"K": 25.0,
		"L": 27.5,
		"M": 30.0,
		"N": 32.5,
		"O": 35.0,
		"P": 37.5,
		"Q": 40.0,
		"R": 42.5,
		"S": 45.0,
		"T": 47.5,
		"U": 50.0,
		"V": 52.5,
		"W": 55.0,
		"X": 57.5,
	}
	mLon := map[string]float64{
		"A": 0,
		"B": 5,
		"C": 10,
		"D": 15,
		"E": 20,
		"F": 25,
		"G": 30,
		"H": 35,
		"I": 40,
		"J": 45,
		"K": 50,
		"L": 55,
		"M": 60,
		"N": 65,
		"O": 70,
		"P": 75,
		"Q": 80,
		"R": 85,
		"S": 90,
		"T": 95,
		"U": 100,
		"V": 105,
		"W": 110,
		"X": 115,
	}
	s.decoded.Lat = mLat[llc.getLatChar()]
	s.decoded.Lon = mLon[llc.getLonChar()]
	s.encoded = llc
	return s
}
