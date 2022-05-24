// Code generated by "enumer -type=N76E616LDROMSize -trimprefix=N76E616LDROM -transform=snake -json -text"; DO NOT EDIT

package n76

import (
	"encoding/json"
	"fmt"
)

const _N76E616LDROMSizeName = "0kb1kb2kb3kb4kb"

var _N76E616LDROMSizeIndex = [...]uint8{0, 3, 6, 9, 12, 15}

func (i N76E616LDROMSize) String() string {
	if i >= N76E616LDROMSize(len(_N76E616LDROMSizeIndex)-1) {
		return fmt.Sprintf("N76E616LDROMSize(%d)", i)
	}
	return _N76E616LDROMSizeName[_N76E616LDROMSizeIndex[i]:_N76E616LDROMSizeIndex[i+1]]
}

var _N76E616LDROMSizeValues = []N76E616LDROMSize{0, 1, 2, 3, 4}

var _N76E616LDROMSizeNameToValueMap = map[string]N76E616LDROMSize{
	_N76E616LDROMSizeName[0:3]:   0,
	_N76E616LDROMSizeName[3:6]:   1,
	_N76E616LDROMSizeName[6:9]:   2,
	_N76E616LDROMSizeName[9:12]:  3,
	_N76E616LDROMSizeName[12:15]: 4,
}

// N76E616LDROMSizeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func N76E616LDROMSizeString(s string) (N76E616LDROMSize, error) {
	if val, ok := _N76E616LDROMSizeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to N76E616LDROMSize values", s)
}

// N76E616LDROMSizeValues returns all values of the enum
func N76E616LDROMSizeValues() []N76E616LDROMSize {
	return _N76E616LDROMSizeValues
}

// IsAN76E616LDROMSize returns "true" if the value is listed in the enum definition. "false" otherwise
func (i N76E616LDROMSize) IsAN76E616LDROMSize() bool {
	for _, v := range _N76E616LDROMSizeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for N76E616LDROMSize
func (i N76E616LDROMSize) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for N76E616LDROMSize
func (i *N76E616LDROMSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("N76E616LDROMSize should be a string, got %s", data)
	}

	var err error
	*i, err = N76E616LDROMSizeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for N76E616LDROMSize
func (i N76E616LDROMSize) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for N76E616LDROMSize
func (i *N76E616LDROMSize) UnmarshalText(text []byte) error {
	var err error
	*i, err = N76E616LDROMSizeString(string(text))
	return err
}
