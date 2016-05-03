package polygoncontains

import (
	"encoding/json"
	"testing"
)

var testPolygons = [...]string{
	// 0
	`{
		"type":"Polygon",
		"coordinates":[
			[
				[
					-0.1709747314453125,
					51.519852590742715
				],
				[
					-0.24650573730468753,
					51.471546541834144
				],
				[
					-0.14076232910156253,
					51.46299146603759
				],
				[
					-0.07072448730468751,
					51.54206472590801
				],
				[
					-0.14694213867187503,
					51.57621608189103
				],
				[
					-0.1709747314453125,
					51.519852590742715
				]
			]
		]
	}`,
}

type testPoi struct {
	Lat float64
	Lng float64
}

func (p testPoi) GetLat() float64 {
	return p.Lat
}
func (p testPoi) GetLng() float64 {
	return p.Lng
}

var testPoi1 = testPoi{Lng: -0.14247894287109378, Lat: 51.51173391474148}
var testPoi2 = testPoi{Lng: 0.07553100585937501, Lat: 51.49335472541077}

type Fataler interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func getPolygonCoordinates(f Fataler) PolygonCoordinates {
	var target Geometry
	err := json.Unmarshal([]byte(testPolygons[0]), &target)
	if err != nil {
		f.Fatalf("Error: %v", err)
	}
	var coordinates PolygonCoordinates
	err = json.Unmarshal([]byte(*target.Coordinates), &coordinates)
	if err != nil {
		f.Fatalf("Error: %v", err)
	}
	return coordinates
}

func TestContains(t *testing.T) {
	coordinates := getPolygonCoordinates(t)
	if !coordinates.Contains(testPoi1) {
		t.Errorf("contains failed testPoi1")
	}
	if coordinates.Contains(testPoi2) {
		t.Errorf("contains failed testPoi2")
	}
}

func BenchmarkContains(b *testing.B) {
	coordinates := getPolygonCoordinates(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coordinates.Contains(testPoi1)
		coordinates.Contains(testPoi2)
	}
}

func BenchmarkRayCrossesSegment(b *testing.B) {
	pa := [2]float64{27.01212, 42.212132}
	pb := [2]float64{27.23123, 43.121211}
	for i := 0; i < b.N; i++ {
		rayCrossesSegment(27.123123, 44.23232, pa, pb)
	}
}
