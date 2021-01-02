package main

import "strings"

// next links return QTH

//
// next links return distance, azimuth, QTH-1 and QTH-2

type qt int

const (
	qtUnsupported         qt = iota
	qtQthPosition            // http://localhost:8080/qth?lat=46.604&lon=15.625
	qtQthLocator             // http://localhost:8080/qth?jn76to
	qtDistLocator            // http://localhost:8080/qth?jn76to;jn76PO
	qtDistPosition           // http://localhost:8080/qth?lat=46.604&lon=15.625;lat=46.604&lon=15.291
	qtDistLocatorPosition    // http://localhost:8080/qth?jn76to;lat=46.604&lon=15.291
	qtDistPositionLocator    // http://localhost:8080/qth?lat=46.604&lon=15.625;jn76PO
)

// queryType returns qt based on separators (& and ;) in the query
func queryType(query string) qt {

	if strings.Contains(query, ";") {
		// distance, azimuth, qth1, qth2
		if strings.Contains(query, "&") {
			switch strings.Count(query, "&") {
			case 1:
				{
					if strings.Index(query, ";") < strings.Index(query, "&") {
						return qtDistLocatorPosition
					} else {
						return qtDistPositionLocator
					}
				}
			case 2:
				return qtDistPosition
			}

		} else {
			return qtDistLocator
		}

	} else {
		// qth
		if strings.Contains(query, "&") {
			return qtQthPosition
		} else {
			if len(query) == 0 {
				return qtUnsupported
			}
			return qtQthLocator
		}
	}

	return qtUnsupported
}
