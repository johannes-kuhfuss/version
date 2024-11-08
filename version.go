// The package version handles version numbers in the form "major.minor.patch"
package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major uint64
	minor uint64
	patch uint64
}

// New creates a new version based on three seperate numbers for major, minor and patch
func New(major uint64, minor uint64, patch uint64) Version {
	return Version{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// parseTwo parses the first two parts of the version, major and minor
func (v *Version) parseTwo(parts []string) error {
	if len(parts) < 2 || len(parts) > 3 {
		return errors.New("expected version string as \"major.minor\" or \"major.minor.patch\"")
	}
	major, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		return fmt.Errorf("could not parse %v as major version", parts[0])
	}
	minor, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return fmt.Errorf("could not parse %v as minor version", parts[1])
	}
	v.major = major
	v.minor = minor
	return nil
}

// parseThree parses all three parts of the version, major, minor and patch
func (v *Version) parseThree(parts []string) error {
	if len(parts) != 3 {
		return errors.New("expected version string as \"major.minor.patch\"")
	}
	err := v.parseTwo(parts)
	if err != nil {
		return err
	}
	patch, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		v.major = 0
		v.minor = 0
		return fmt.Errorf("could not parse %v as patch version", parts[2])
	}
	v.patch = patch
	return nil
}

// Parse takes a string and tries to parse it as a version. It accepts either the form "major.minor" or "major.minor.patch"
// For versions given as "major.minor", patch will be set to 0
func (v *Version) Parse(vStr string) error {
	parts := strings.Split(vStr, ".")
	switch {
	case len(parts) < 2 || len(parts) > 3:
		return errors.New("expected version string as \"major.minor\" or \"major.minor.patch\"")
	case len(parts) == 2:
		err := v.parseTwo(parts)
		if err != nil {
			return err
		}
	case len(parts) == 3:
		err := v.parseThree(parts)
		if err != nil {
			return err
		}
	}
	return nil
}

// Prints a version in its canonical form "major.minor.patch"
func (v *Version) Print() string {
	vStr := []string{
		strconv.FormatUint(v.major, 10),
		strconv.FormatUint(v.minor, 10),
		strconv.FormatUint(v.patch, 10),
	}
	return strings.Join(vStr, ".")
}

// IsNull returns true, if a version is zero in all components
func (v *Version) IsNull() bool {
	return (v.major == 0 && v.minor == 0 && v.patch == 0)
}

// IsEqual compares two versions and returns true if they are equal
func (v *Version) IsEqual(v2 Version) bool {
	return *v == v2
}

// IsBigger compares two versions and returns true if the first version is bigger than the second
func (v *Version) IsBigger(v2 Version) bool {
	if v.major > v2.major {
		return true
	}
	if v.minor > v2.minor {
		return true
	}
	if v.patch > v2.patch {
		return true
	}
	return false
}
