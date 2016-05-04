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
	// 1
	`{
		"type":"Polygon",
		"coordinates":[[[21.84585,47.33609],[21.83778,47.33245],[21.82821,47.3317],[21.8147,47.33654],[21.81492,47.34647],[21.81327,47.35176],[21.81468,47.35346],[21.81444,47.35667],[21.81669,47.36426],[21.81655,47.36898],[21.80972,47.3769],[21.8088,47.38007],[21.80576,47.38416],[21.80598,47.39546],[21.80406,47.40303],[21.80611,47.40671],[21.80772,47.4153],[21.82223,47.4241],[21.82266,47.43042],[21.82874,47.44071],[21.8408,47.44058],[21.8403,47.44187],[21.85007,47.44382],[21.85508,47.43932],[21.85853,47.43068],[21.86656,47.43716],[21.86619,47.43901],[21.87814,47.44973],[21.88391,47.44687],[21.89054,47.45319],[21.90242,47.4472],[21.90583,47.45009],[21.9085,47.45512],[21.91146,47.45829],[21.93384,47.44474],[21.92845,47.43533],[21.92546,47.43397],[21.92519,47.43221],[21.92112,47.42843],[21.91622,47.4256],[21.91581,47.42173],[21.9143,47.42025],[21.90859,47.41708],[21.90999,47.41312],[21.90647,47.4017],[21.91372,47.39078],[21.92619,47.39363],[21.93091,47.38804],[21.93832,47.39027],[21.94295,47.38198],[21.95231,47.38168],[21.95591,47.3761],[21.95257,47.3746],[21.95042,47.37524],[21.94523,47.37515],[21.9385,47.37298],[21.93968,47.36761],[21.93767,47.36612],[21.93467,47.36027],[21.92524,47.35353],[21.9236,47.34981],[21.92327,47.34413],[21.92158,47.34143],[21.91352,47.33242],[21.90796,47.32876],[21.90739,47.32622],[21.90451,47.32221],[21.89613,47.31627],[21.89324,47.30631],[21.88817,47.29473],[21.88875,47.29681],[21.88169,47.30041],[21.87703,47.30035],[21.8778,47.30354],[21.87386,47.30365],[21.87351,47.30443],[21.86807,47.30607],[21.86829,47.30506],[21.86663,47.30519],[21.86573,47.30671],[21.86233,47.30849],[21.86086,47.30996],[21.86305,47.31033],[21.84585,47.33609]]]
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

var testPois = [][]testPoi{
	{
		testPoi{Lng: -0.14247894287109378, Lat: 51.51173391474148},
		testPoi{Lng: 0.07553100585937501, Lat: 51.49335472541077},
	},
	{
		testPoi{Lng: 21.891628, Lat: 47.380651},
		testPoi{Lng: 18.528550, Lat: 47.033121},
	},
}

/*
var testPoi1 = testPoi{Lng: -0.14247894287109378, Lat: 51.51173391474148}
var testPoi2 = testPoi{Lng: 0.07553100585937501, Lat: 51.49335472541077}

var testPoi2_1 = testPoi{Lng: 21.891628, Lat: 47.380651}
var testPoi2_2 = testPoi{Lng: 18.528550, Lat: 47.033121}
*/

type Fataler interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func getPolygonCoordinates(f Fataler, index int) PolygonCoordinates {
	var target Geometry
	err := json.Unmarshal([]byte(testPolygons[index]), &target)
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

func doTestContains(t *testing.T, index int) {
	coordinates := getPolygonCoordinates(t, index)
	if !coordinates.Contains(testPois[index][0]) {
		t.Errorf("contains failed for polygon %v", index)
	}
	if coordinates.Contains(testPois[index][1]) {
		t.Errorf("not contains failed for polygon %v", index)
	}
}

func doTestContainsStruct(t *testing.T, index int) {
	coordinates := getPolygonCoordinates(t, index)
	coordinatesStruct := &PolygonCoordinatesStruct{
		coords: coordinates,
	}
	if !coordinatesStruct.Contains(testPois[index][0]) {
		t.Errorf("contains failed for polygon %v", index)
	}
	if coordinatesStruct.Contains(testPois[index][1]) {
		t.Errorf("not contains failed for polygon %v", index)
	}
}

func doBenchmarkContains(b *testing.B, index int) {
	coordinates := getPolygonCoordinates(b, index)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coordinates.Contains(testPois[index][0])
		coordinates.Contains(testPois[index][1])
	}
}

func doBenchmarkStructContains(b *testing.B, index int) {
	coordinates := getPolygonCoordinates(b, index)
	coordinatesStruct := &PolygonCoordinatesStruct{
		coords: coordinates,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coordinatesStruct.Contains(testPois[index][0])
		coordinatesStruct.Contains(testPois[index][1])
	}
}

func TestContains(t *testing.T) {
	doTestContains(t, 0)
}

func TestStructContains(t *testing.T) {
	doTestContainsStruct(t, 0)
}

func BenchmarkContains(b *testing.B) {
	doBenchmarkContains(b, 0)
}

func BenchmarkStructContains(b *testing.B) {
	doBenchmarkStructContains(b, 0)
}

func TestContains2(t *testing.T) {
	doTestContains(t, 1)
}

func TestStructContains2(t *testing.T) {
	doTestContainsStruct(t, 1)
}

func BenchmarkContains2(b *testing.B) {
	doBenchmarkContains(b, 1)
}

func BenchmarkStructContains2(b *testing.B) {
	doBenchmarkStructContains(b, 1)
}

func BenchmarkRayCrossesSegment(b *testing.B) {
	pa := [2]float64{27.01212, 42.212132}
	pb := [2]float64{27.23123, 43.121211}
	for i := 0; i < b.N; i++ {
		rayCrossesSegment(27.123123, 44.23232, pa, pb)
	}
}

// manually inlined

func doTestContainsInline(t *testing.T, index int) {
	coordinates := getPolygonCoordinates(t, index)
	if !coordinates.ContainsInline(testPois[index][0]) {
		t.Errorf("contains failed for polygon %v", index)
	}
	if coordinates.ContainsInline(testPois[index][1]) {
		t.Errorf("not contains failed for polygon %v", index)
	}
}

func doTestContainsStructInline(t *testing.T, index int) {
	coordinates := getPolygonCoordinates(t, index)
	coordinatesStruct := &PolygonCoordinatesStruct{
		coords: coordinates,
	}
	if !coordinatesStruct.ContainsInline(testPois[index][0]) {
		t.Errorf("contains failed for polygon %v", index)
	}
	if coordinatesStruct.ContainsInline(testPois[index][1]) {
		t.Errorf("not contains failed for polygon %v", index)
	}
}

func doBenchmarkContainsInline(b *testing.B, index int) {
	coordinates := getPolygonCoordinates(b, index)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coordinates.ContainsInline(testPois[index][0])
		coordinates.ContainsInline(testPois[index][1])
	}
}

func doBenchmarkStructContainsInline(b *testing.B, index int) {
	coordinates := getPolygonCoordinates(b, index)
	coordinatesStruct := &PolygonCoordinatesStruct{
		coords: coordinates,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coordinatesStruct.ContainsInline(testPois[index][0])
		coordinatesStruct.ContainsInline(testPois[index][1])
	}
}

func TestContainsInline(t *testing.T) {
	doTestContainsInline(t, 0)
}

func TestStructContainsInline(t *testing.T) {
	doTestContainsStructInline(t, 0)
}

func BenchmarkContainsInline(b *testing.B) {
	doBenchmarkContainsInline(b, 0)
}

func BenchmarkStructContainsInline(b *testing.B) {
	doBenchmarkStructContainsInline(b, 0)
}

func TestContains2Inline(t *testing.T) {
	doTestContainsInline(t, 1)
}

func TestStructContains2Inline(t *testing.T) {
	doTestContainsStructInline(t, 1)
}

func BenchmarkContains2Inline(b *testing.B) {
	doBenchmarkContainsInline(b, 1)
}

func BenchmarkStructContains2Inline(b *testing.B) {
	doBenchmarkStructContainsInline(b, 1)
}
