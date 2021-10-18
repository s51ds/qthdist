package api

import (
	"github.com/s51ds/qthdist/geo"
	"sync"
)

var (
	qthCache    = make(map[string]geo.QTH)
	qthCacheMux sync.RWMutex
)

func Distance(locatorA, locatorB string) (distance, azimuth float64, err error) {

	var (
		qthA, qthB geo.QTH
		has        bool
	)

	// is QTH for locatorA in the cache
	qthCacheMux.RLock()
	qthA, has = qthCache[locatorA]
	qthCacheMux.RUnlock()
	if !has { // it is not yet in the cache
		qthA, err = geo.NewQthFromLocator(locatorA)
		if err != nil {
			return
		}
		qthCacheMux.Lock()
		qthCache[locatorA] = qthA // now it is in the cache
		qthCacheMux.Unlock()
	}

	// is QTH for locatorB in the cache
	qthCacheMux.RLock()
	qthB, has = qthCache[locatorB]
	qthCacheMux.RUnlock()
	if !has { // it is not yet in the cache
		qthB, err = geo.NewQthFromLocator(locatorB)
		if err != nil {
			return
		}
		qthCacheMux.Lock()
		qthCache[locatorB] = qthB // now it is in the cache
		qthCacheMux.Unlock()
	}

	distance, azimuth = qthA.DistanceAndAzimuth(qthB)

	return
}
