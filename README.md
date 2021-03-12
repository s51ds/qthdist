## QTH project provides functionality for distance and azimuth calculation between two positions on the Earth.


### Positions can be provided in two froms:
- latitude/longitude in decimal degrees
- Maidenhead QTH locator (https://en.wikipedia.org/wiki/Maidenhead_Locator_System)

### Conversion between latitude/longitude to Maidenhead QTH locator and vice versa is supported in **package geo**:
- NewQthFromLocator
- NewQthFromPosition

## Restfull API
### Distance and azimuth calculation and qth/lat-lon conversion can be the Network Function "out of the box"
see: ./qth/api/main.go

