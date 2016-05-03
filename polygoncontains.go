package polygoncontains

import "encoding/json"

type Poi interface {
	GetLat() float64
	GetLng() float64
}

// Geometry is the geometry object in the geojson.
type Geometry struct {
	Type        string
	Coordinates *json.RawMessage
}

type PolygonCoordinates [][][2]float64

type PolygonCoordinatesStruct struct {
	coords [][][2]float64
}

// Contains returns true if the polygon contains the poi.
func (c PolygonCoordinates) Contains(p Poi) bool {
	plat := p.GetLat()
	plng := p.GetLng()

	inside := false
	for _, ring := range c {
		l := len(ring)
		for i := range ring {
			j := i + 1
			if j >= l {
				j = 0
			}

			if rayCrossesSegment(plng, plat, ring[i], ring[j]) {
				inside = !inside
			}
		}
	}
	return inside
}

// Contains returns true if the polygon contains the poi.
func (c *PolygonCoordinatesStruct) Contains(p Poi) bool {
	plat := p.GetLat()
	plng := p.GetLng()

	inside := false
	for _, ring := range c.coords {
		l := len(ring)
		for i := range ring {
			j := i + 1
			if j >= l {
				j = 0
			}

			if rayCrossesSegment(plng, plat, ring[i], ring[j]) {
				inside = !inside
			}
		}
	}
	return inside
}

// rayCrossesSegment checks whether a horizontal ray that
// intersects (px, py) crosses the segment formed by a and b.
// coordinates are in the form [longitude, latitude]
func rayCrossesSegment(px, py float64, a, b [2]float64) bool {
	ax := a[0]
	ay := a[1]
	bx := b[0]
	by := b[1]

	return ((ay > py) != (by > py)) &&
		(px < (bx-ax)*(py-ay)/(by-ay)+ax)
}
