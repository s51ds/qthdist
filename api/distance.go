package api

import (
	"github.com/s51ds/qthdist/geo"
	"sync"
)

type distValueCache struct {
	distance, azimuth float64
}

var (
	qthCache    = make(map[string]geo.QTH) // key is locator
	qthCacheMux sync.RWMutex

	distCache    = make(map[string]distValueCache) // key is locatorA+locatorB
	distCacheMux sync.RWMutex
)

func Distance(locatorA, locatorB string) (distance, azimuth float64, err error) {
	var (
		qthA, qthB   geo.QTH
		hasCachedQth bool

		cachedDistAndAzimuth    distValueCache
		hasCachedDistAndAzimuth bool
	)
	// if distance has been calculated before then it is cached
	distCacheKey := locatorA + locatorB
	distCacheMux.RLock()
	cachedDistAndAzimuth, hasCachedDistAndAzimuth = distCache[distCacheKey]
	distCacheMux.RUnlock()
	if hasCachedDistAndAzimuth {
		return cachedDistAndAzimuth.distance, cachedDistAndAzimuth.azimuth, nil
	}

	// is QTH for locatorA in the cache
	qthCacheMux.RLock()
	qthA, hasCachedQth = qthCache[locatorA]
	qthCacheMux.RUnlock()
	if !hasCachedQth { // it is not yet in the cache
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
	qthB, hasCachedQth = qthCache[locatorB]
	qthCacheMux.RUnlock()
	if !hasCachedQth { // it is not yet in the cache
		qthB, err = geo.NewQthFromLocator(locatorB)
		if err != nil {
			return
		}
		qthCacheMux.Lock()
		qthCache[locatorB] = qthB // now it is in the cache
		qthCacheMux.Unlock()
	}

	distance, azimuth = qthA.DistanceAndAzimuth(qthB)

	if !hasCachedDistAndAzimuth {
		distCacheMux.Lock()
		distCache[distCacheKey] = distValueCache{
			distance: distance,
			azimuth:  azimuth,
		}
		distCacheMux.Unlock()
	}

	return
}
