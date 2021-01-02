package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qth/geo"
	"strconv"
	"strings"
)

// next links return QTH
//http://localhost:8080/qth?lat=46.604&lon=15.625
//http://localhost:8080/qth?jn76to
//
// next links return distance, azimuth, QTH-1 and QTH-2
// http://localhost:8080/qth?jn76to;jn76PO
// http://localhost:8080/qth?lat=46.604&lon=15.625;lat=46.604&lon=15.291
// http://localhost:8080/qth?jn76to;lat=46.604&lon=15.291
// http://localhost:8080/qth?lat=46.604&lon=15.625;jn76PO

func main() {
	http.HandleFunc("/qth", qth)

	addr := "127.0.0.1:8080"
	fmt.Println("Server listen on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println(err.Error())
	}
}

func api(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	fmt.Println("----------------")

	query := request.URL.RawQuery
	fmt.Println("RawQuery:", query)
	ss := strings.Split(query, ";")
	for _, v := range ss {
		fmt.Println(v)
	}
	switch len(ss) {
	case 1:
		{
			// query can be be:
			// - locator -> returns: lat,lon (http://localhost:8080/?jn76to)
			// - latitude, longitude -> returns: qth locator (http://localhost:8080/?lat=46.604&lon=15.625)

			v := request.URL.Query()
			if len(v) == 2 { //latitude, longitude -> returns: qth locator (http://localhost:8080/?lat=46.604&lon=15.625)
				v := request.URL.Query()
				lat, has := v["lat"]
				if !has {
					fmt.Fprintf(writer, "parameter lat is missing in query! %s", query)
					return
				}
				lon, has := v["lon"]
				if !has {
					fmt.Fprintf(writer, "parameter lon is missing in query! %s", query)
					return
				}

				//
				var err error
				var latf, lonf float64
				if latf, err = strconv.ParseFloat(lat[0], 64); err != nil {
					fmt.Fprintf(writer, "parameter lat is not a number! %s", query)
					return
				}
				if lonf, err = strconv.ParseFloat(lon[0], 64); err != nil {
					fmt.Fprintf(writer, "parameter lon is not a number! %s", query)
					return
				}
				if qth, err := geo.NewQthFromPosition(latf, lonf); err != nil {
					fmt.Println(err.Error())
					fmt.Fprintf(writer, err.Error())
					return
				} else {
					apiResp := ApiResp{
						LocA: qth.Loc,
						LatA: latf,
						LonA: lonf,
					}
					json.NewEncoder(writer).Encode(apiResp)
					return
				}

			} else { // locator -> returns: lat,lon (http://localhost:8080/?jn76to)
				if qth, err := geo.NewQthFromLocator(ss[0]); err != nil {
					fmt.Println(err.Error())
					fmt.Fprintf(writer, err.Error())
					return
				} else {
					apiResp := ApiResp{
						LocA: ss[0],
						LatA: qth.LatLon.Lat,
						LonA: qth.LatLon.Lon,
					}
					json.NewEncoder(writer).Encode(apiResp)
					return
				}
			}

		}
	case 2:
		{
			// query can be be:
			// - locatorA, locatorB -> distance and azimuth: http://localhost:8080/?jn76to;jn76PO
			if qth, err := geo.NewQthFromLocator(ss[0]); err != nil {
				fmt.Println(err.Error())
				fmt.Fprintf(writer, err.Error())
				return
			} else {
				apiResp := ApiResp{
					LocA: ss[0],
					LatA: qth.LatLon.Lat,
					LonA: qth.LatLon.Lon,
				}
				json.NewEncoder(writer).Encode(apiResp)
				return
			}
		}

	}

	fmt.Fprintf(writer, query)
}

func distanceAndAzimuth(w http.ResponseWriter, r *http.Request) {
	//for name, headers := range r.Header {
	//	for _, h := range headers {
	//		fmt.Fprintf(w, "%v: %v\n", name, h)
	//	}
	//}

	var err error
	var aQTH, bQTH *geo.QTH

	if aQTH, err = geo.NewQthFromLocator("JN76TO"); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if bQTH, err = geo.NewQthFromLocator("JN76PO"); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	dist, azim := aQTH.DistanceAndAzimuth(bQTH)

	qthResp := ApiResp{
		LocA:     aQTH.Loc,
		LocB:     bQTH.Loc,
		LatA:     aQTH.LatLon.Lat,
		LonA:     aQTH.LatLon.Lon,
		LatB:     bQTH.LatLon.Lat,
		LonB:     bQTH.LatLon.Lon,
		Distance: dist,
		Azimuth:  azim,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(qthResp)

	//	fmt.Fprintf(w, "dist: %.1f km, azim: %.1f deg", dist, azim)
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
