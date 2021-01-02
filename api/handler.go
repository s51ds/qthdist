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

// TODO: queryError
func qth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	fmt.Println("----------------")
	query := request.URL.RawQuery
	fmt.Println("RawQuery:", query)

	switch queryType(query) {
	case qtQthPosition:
		{
			if lat, lon, err := parseLatLon(request); err != nil {
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
	case qtQthLocator:
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
	case qtDistLocator:
		{
			fmt.Println(query)
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
	case qtDistPosition:
		{
			fmt.Println(query)
		}

	default:
		fmt.Fprintf(writer, "unsupported query %s", query)
		return

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

func parseLatLon(request *http.Request) (latF float64, lonF float64, err error) {
	v := request.URL.Query()
	latS, has := v["lat"]
	if !has {
		err = errors.New("parameter lat is missing in query")
		return
	}
	lonS, has := v["lon"]
	if !has {
		err = errors.New("parameter lon is missing in query")
		return
	}
	if latF, err = strconv.ParseFloat(latS[0], 64); err != nil {
		err = errors.New("parameter lat is not a number")
		return
	}
	if lonF, err = strconv.ParseFloat(lonS[0], 64); err != nil {
		err = errors.New("parameter lon is not a number")
		return
	}
	return
}
