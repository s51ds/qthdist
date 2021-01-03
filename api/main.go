package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"qth/geo"
	"strconv"
	"strings"
)

// next links return QTH
//http://localhost:8080/qth?lat=46.604&lon=15.625 (qtQthPosition)
//http://localhost:8080/qth?jn76to (qtQthLocator)
//
// next links return distance, azimuth, QTH-1 and QTH-2
// http://localhost:8080/qth?jn76to;jn76PO (qtDistLocator)
// http://localhost:8080/qth?lat=46.604&lon=15.625;lat=46.604&lon=15.291 (qtDistPosition)
// http://localhost:8080/qth?jn76to;lat=46.604&lon=15.291 (qtDistLocatorPosition)
// http://localhost:8080/qth?lat=46.604&lon=15.625;jn76PO (qtDistPositionLocator)

func main() {
	http.HandleFunc("/qth", qth)

	addr := "127.0.0.1:8080"
	fmt.Println("Server listen on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func parseReqLatLon(rawRequest string) (latF float64, lonF float64, err error) {
	// lat=46.604&lon=15.625
	ss := strings.Split(rawRequest, "&")
	if len(ss) != 2 {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	}
	ss0 := strings.Split(ss[0], "=")
	if len(ss0) != 2 {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	}
	ss1 := strings.Split(ss[1], "=")
	if len(ss1) != 2 {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	}

	//
	m := make(map[string]string)

	if _, has := m[ss0[0]]; has {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	}
	m[ss0[0]] = ss0[1]

	if _, has := m[ss1[0]]; has {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	}
	m[ss1[0]] = ss1[1]
	//
	//

	if lat, has := m["lat"]; !has {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	} else {
		if latF, err = strconv.ParseFloat(lat, 64); err != nil {
			err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
			return
		}
	}

	if lon, has := m["lon"]; !has {
		err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
		return
	} else {
		if lonF, err = strconv.ParseFloat(lon, 64); err != nil {
			err = errors.New("wrong lat/lon syntax: " + rawRequest + ", example: lat=46.604&lon=15.625")
			return
		}
	}

	// everything is perfect
	return
}

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

// TODO: queryError
func qth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	fmt.Println("----------------")
	query := request.URL.RawQuery
	fmt.Println("RawQuery:", query)

	switch queryType(query) {
	case qtQthPosition: //?lat=46.604&lon=15.625
		{
			if lat, lon, err := parseReqLatLon(query); err != nil {
				s := err.Error() + " " + query
				fmt.Println(s)
				fmt.Fprintf(writer, s)
				return
			} else {
				if qth, err := geo.NewQthFromPosition(lat, lon); err != nil {
					fmt.Println(err.Error())
					fmt.Fprintf(writer, err.Error())
					return
				} else {
					apiResp := ApiResp{
						LocA: qth.Loc,
						LatA: lat,
						LonA: lon,
					}
					json.NewEncoder(writer).Encode(apiResp)
					return
				}
			}
		}
	case qtQthLocator: // ?jn76to
		{
			if qth, err := geo.NewQthFromLocator(query); err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			} else {
				apiResp := ApiResp{
					LocA: query,
					LatA: qth.LatLon.Lat,
					LonA: qth.LatLon.Lon,
				}
				json.NewEncoder(writer).Encode(apiResp)
				return
			}
		}
	case qtDistLocator: // ?jn76to;jn76PO
		{
			ss := strings.Split(query, ";")
			qthA, err := geo.NewQthFromLocator(ss[0])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}
			qthB, err := geo.NewQthFromLocator(ss[1])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}
			dist, azim := qthA.DistanceAndAzimuth(qthB)
			apiResp := ApiResp{
				LocA:     qthA.Loc,
				LocB:     qthB.Loc,
				LatA:     qthA.LatLon.Lat,
				LonA:     qthA.LatLon.Lon,
				LatB:     qthB.LatLon.Lat,
				LonB:     qthB.LatLon.Lon,
				Distance: dist,
				Azimuth:  azim,
			}
			json.NewEncoder(writer).Encode(apiResp)
			return
		}
	case qtDistPosition: // ?lat=46.604&lon=15.625;lat=46.604&lon=15.291
		{
			fmt.Println(query)
			ss := strings.Split(query, ";")

			// QTH A
			lat0, lon0, err := parseReqLatLon(ss[0])
			if err != nil {
				s := err.Error() + " " + query
				fmt.Println(s)
				fmt.Fprintf(writer, s)
				return
			}
			qthA, err := geo.NewQthFromPosition(lat0, lon0)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}
			// QTH B
			lat1, lon1, err := parseReqLatLon(ss[1])
			if err != nil {
				s := err.Error() + " " + query
				fmt.Println(s)
				fmt.Fprintf(writer, s)
				return
			}
			qthB, err := geo.NewQthFromPosition(lat1, lon1)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}

			dist, azim := qthA.DistanceAndAzimuth(qthB)
			apiResp := ApiResp{
				LocA:     qthA.Loc,
				LocB:     qthB.Loc,
				LatA:     qthA.LatLon.Lat,
				LonA:     qthA.LatLon.Lon,
				LatB:     qthB.LatLon.Lat,
				LonB:     qthB.LatLon.Lon,
				Distance: dist,
				Azimuth:  azim,
			}
			json.NewEncoder(writer).Encode(apiResp)
			return
		}
	case qtDistLocatorPosition:
		{ // ?jn76to;lat=46.604&lon=15.291
			fmt.Println(query)
			ss := strings.Split(query, ";")

			// QTH A
			qthA, err := geo.NewQthFromLocator(ss[0])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}

			// QTH B
			lat1, lon1, err := parseReqLatLon(ss[1])
			if err != nil {
				s := err.Error() + " " + query
				fmt.Println(s)
				fmt.Fprintf(writer, s)
				return
			}
			qthB, err := geo.NewQthFromPosition(lat1, lon1)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}

			dist, azim := qthA.DistanceAndAzimuth(qthB)
			apiResp := ApiResp{
				LocA:     qthA.Loc,
				LocB:     qthB.Loc,
				LatA:     qthA.LatLon.Lat,
				LonA:     qthA.LatLon.Lon,
				LatB:     qthB.LatLon.Lat,
				LonB:     qthB.LatLon.Lon,
				Distance: dist,
				Azimuth:  azim,
			}
			json.NewEncoder(writer).Encode(apiResp)
			return
		}
	case qtDistPositionLocator:
		{ //?lat=46.604&lon=15.625;jn76PO
			fmt.Println(query)
			ss := strings.Split(query, ";")

			// QTH A
			lat0, lon0, err := parseReqLatLon(ss[0])
			if err != nil {
				s := err.Error() + " " + query
				fmt.Println(s)
				fmt.Fprintf(writer, s)
				return
			}
			qthA, err := geo.NewQthFromPosition(lat0, lon0)
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}

			// QTH B
			qthB, err := geo.NewQthFromLocator(ss[1])
			if err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			}

			dist, azim := qthA.DistanceAndAzimuth(qthB)
			apiResp := ApiResp{
				LocA:     qthA.Loc,
				LocB:     qthB.Loc,
				LatA:     qthA.LatLon.Lat,
				LonA:     qthA.LatLon.Lon,
				LatB:     qthB.LatLon.Lat,
				LonB:     qthB.LatLon.Lon,
				Distance: dist,
				Azimuth:  azim,
			}
			json.NewEncoder(writer).Encode(apiResp)
			return

		}

	default:
		fmt.Fprintf(writer, "unsupported query %s", query)
		return

	}

}

type ApiResp struct {
	LocA     string  `json:"locA,omitempty"`
	LocB     string  `json:"locB,omitempty"`
	LatA     float64 `json:"latA,omitempty"`
	LonA     float64 `json:"lonA,omitempty"`
	LatB     float64 `json:"latB,omitempty"`
	LonB     float64 `json:"lonB,omitempty"`
	Distance float64 `json:"distance,omitempty"`
	Azimuth  float64 `json:"azimuth,omitempty"`
}
