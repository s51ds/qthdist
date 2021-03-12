package server

// next links return QTH
//http://localhost:8080/qth?lat=46.604&lon=15.625 (qtQthPosition)
//http://localhost:8080/qth?jn76to (qtQthLocator)
//
// next links return distance, azimuth, QTH-1 and QTH-2
// http://localhost:8080/qth?jn76to;jn76PO (qtDistLocator)
// http://localhost:8080/qth?lat=46.604&lon=15.625;lat=46.604&lon=15.291 (qtDistPosition)
// http://localhost:8080/qth?jn76to;lat=46.604&lon=15.291 (qtDistLocatorPosition)
// http://localhost:8080/qth?lat=46.604&lon=15.625;jn76PO (qtDistPositionLocator)
