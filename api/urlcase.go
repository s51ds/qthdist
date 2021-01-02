package main

// next links return QTH

//
// next links return distance, azimuth, QTH-1 and QTH-2

const (
	queryUnsupported         = iota
	queryQthPosition         // http://localhost:8080/qth?lat=46.604&lon=15.625
	queryQthLocator          // http://localhost:8080/qth?jn76to
	queryDistLocator         // http://localhost:8080/qth?jn76to;jn76PO
	queryDistPosition        // http://localhost:8080/qth?lat=46.604&lon=15.625;lat=46.604&lon=15.291
	queryDistLocatorPosition // http://localhost:8080/qth?jn76to&lon=15.625;lat=46.604&lon=15.291
	queryDistPositionLocator // http://localhost:8080/qth?lat=46.604&lon=15.625;jn76PO
)

func queryType() int {

	return queryUnsupported
}
