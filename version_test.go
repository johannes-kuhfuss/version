package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTwoZeroLengthReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{}
	err := ver.parseTwo(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor\" or \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoWrongLength1ReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1"}
	err := ver.parseTwo(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor\" or \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoWrongLength4ReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2", "3", "4"}
	err := ver.parseTwo(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor\" or \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoMajorNotANumberReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"A", "2"}
	err := ver.parseTwo(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse A as major version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoMinorNotANumberReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "B"}
	err := ver.parseTwo(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse B as minor version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoReturnsNil(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2"}
	err := ver.parseTwo(parts)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, ver.major)
	assert.EqualValues(t, 2, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoIgnoresThirdReturnsNil(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2", "3"}
	err := ver.parseTwo(parts)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, ver.major)
	assert.EqualValues(t, 2, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreeZeroLengthReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{}
	err := ver.parseThree(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreeTwoPartsReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2"}
	err := ver.parseThree(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreePatchNotANumberReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2", "C"}
	err := ver.parseThree(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse C as patch version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreeMinorNotANumberReturnsError(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "B", "3"}
	err := ver.parseThree(parts)
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse B as minor version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreeReturnsNil(t *testing.T) {
	var (
		ver Version
	)
	parts := []string{"1", "2", "3"}
	err := ver.parseThree(parts)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, ver.major)
	assert.EqualValues(t, 2, ver.minor)
	assert.EqualValues(t, 3, ver.patch)
}

func TestParseNoversionStringReturnsError(t *testing.T) {
	var (
		ver Version
	)
	err := ver.Parse("no version string")
	assert.NotNil(t, err)
	assert.EqualValues(t, "expected version string as \"major.minor\" or \"major.minor.patch\"", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoPartsCannotParseReturnsError(t *testing.T) {
	var (
		ver Version
	)
	err := ver.Parse("2024.A")
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse A as minor version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreePartsCannotParseReturnsError(t *testing.T) {
	var (
		ver Version
	)
	err := ver.Parse("2024.10.C")
	assert.NotNil(t, err)
	assert.EqualValues(t, "could not parse C as patch version", err.Error())
	assert.EqualValues(t, 0, ver.major)
	assert.EqualValues(t, 0, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseTwoPartsReturnsNil(t *testing.T) {
	var (
		ver Version
	)
	err := ver.Parse("2024.10")
	assert.Nil(t, err)
	assert.EqualValues(t, 2024, ver.major)
	assert.EqualValues(t, 10, ver.minor)
	assert.EqualValues(t, 0, ver.patch)
}

func TestParseThreePartsReturnsNil(t *testing.T) {
	var (
		ver Version
	)
	err := ver.Parse("2024.10.2")
	assert.Nil(t, err)
	assert.EqualValues(t, 2024, ver.major)
	assert.EqualValues(t, 10, ver.minor)
	assert.EqualValues(t, 2, ver.patch)
}

func TestNew(t *testing.T) {
	ver := New(1, 2, 3)
	assert.EqualValues(t, 1, ver.major)
	assert.EqualValues(t, 2, ver.minor)
	assert.EqualValues(t, 3, ver.patch)
}

func TestPrint(t *testing.T) {
	ver := New(1, 2, 3)
	vStr := ver.Print()
	assert.EqualValues(t, "1.2.3", vStr)
}

func TestIsNullIsTrue(t *testing.T) {
	ver := New(0, 0, 0)
	n := ver.IsNull()
	assert.True(t, n)
}

func TestIsNullIsFalse(t *testing.T) {
	ver := New(1, 2, 3)
	n := ver.IsNull()
	assert.False(t, n)
}

func TestIsEqualIsEqual(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(1, 2, 3)
	e := v1.IsEqual(v2)
	assert.True(t, e)
}

func TestIsEqualIsNotEqual(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(1, 2, 4)
	e := v1.IsEqual(v2)
	assert.False(t, e)
}

func TestIsBiggerIsNotBiggerPatch(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(1, 2, 4)
	b := v1.IsBigger(v2)
	assert.False(t, b)
}

func TestIsBiggerIsNotBiggerMinor(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(1, 3, 3)
	b := v1.IsBigger(v2)
	assert.False(t, b)
}

func TestIsBiggerIsNotBiggerMajor(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(2, 2, 3)
	b := v1.IsBigger(v2)
	assert.False(t, b)
}

func TestIsBiggerIsEqual(t *testing.T) {
	v1 := New(1, 2, 3)
	v2 := New(1, 2, 3)
	b := v1.IsBigger(v2)
	assert.False(t, b)
}

func TestIsBiggerIsBiggerPatch(t *testing.T) {
	v1 := New(1, 2, 4)
	v2 := New(1, 2, 3)
	b := v1.IsBigger(v2)
	assert.True(t, b)
}

func TestIsBiggerIsBiggerMinor(t *testing.T) {
	v1 := New(1, 3, 3)
	v2 := New(1, 2, 3)
	b := v1.IsBigger(v2)
	assert.True(t, b)
}

func TestIsBiggerIsBiggerMajor(t *testing.T) {
	v1 := New(2, 2, 3)
	v2 := New(1, 2, 3)
	b := v1.IsBigger(v2)
	assert.True(t, b)
}
