// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"fmt"
)

const (
	// PlanetMercury is a Planet of type Mercury
	PlanetMercury Planet = iota
	// PlanetVenus is a Planet of type Venus
	PlanetVenus
	// PlanetEarth is a Planet of type Earth
	PlanetEarth
	// PlanetMars is a Planet of type Mars
	PlanetMars
	// PlanetJupiter is a Planet of type Jupiter
	PlanetJupiter
	// PlanetSaturn is a Planet of type Saturn
	PlanetSaturn
	// PlanetUranus is a Planet of type Uranus
	PlanetUranus
	// PlanetNeptune is a Planet of type Neptune
	PlanetNeptune
)

const _PlanetName = "MercuryVenusEarthMarsJupiterSaturnUranusNeptune"

var _PlanetMap = map[Planet]string{
	0: _PlanetName[0:7],
	1: _PlanetName[7:12],
	2: _PlanetName[12:17],
	3: _PlanetName[17:21],
	4: _PlanetName[21:28],
	5: _PlanetName[28:34],
	6: _PlanetName[34:40],
	7: _PlanetName[40:47],
}

// String implements the Stringer interface.
func (x Planet) String() string {
	if str, ok := _PlanetMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Planet(%d)", x)
}

var _PlanetValue = map[string]Planet{
	_PlanetName[0:7]:   0,
	_PlanetName[7:12]:  1,
	_PlanetName[12:17]: 2,
	_PlanetName[17:21]: 3,
	_PlanetName[21:28]: 4,
	_PlanetName[28:34]: 5,
	_PlanetName[34:40]: 6,
	_PlanetName[40:47]: 7,
}

// ParsePlanet attempts to convert a string to a Planet
func ParsePlanet(name string) (Planet, error) {
	if x, ok := _PlanetValue[name]; ok {
		return x, nil
	}
	return Planet(0), fmt.Errorf("%s is not a valid Planet", name)
}

// MarshalJSON implements the json.Marshaler method
func (x *Planet) MarshalJSON() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the json.Unmarshaler method
func (x *Planet) UnmarshalJSON(text []byte) error {
	name := string(text)
	tmp, err := ParsePlanet(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
