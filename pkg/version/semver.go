package version

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

type Semver struct {
	Major string
	Minor string
	Patch string
	Other string
}

func NewVersion(in string) Semver {
	parts := strings.Split(in, ".")
	if len(parts) == 1 {
		return Semver{Major: parts[0]}
	} else if len(parts) == 2 {
		return Semver{Major: parts[0], Minor: parts[1]}
	} else if len(parts) == 3 {
		return Semver{Major: parts[0], Minor: parts[1], Patch: parts[2]}
	} else {
		return Semver{Major: parts[0], Minor: parts[1], Patch: parts[2], Other: parts[3]}
	}
}

func toInt(in string) (int, error) {
	if in == "" {
		return 0, nil
	}
	s, err := strconv.Atoi(in)
	if err != nil {
		return -1, errors.Wrap(err, "unable to parse as int: "+in)
	}
	return s, nil
}

func (v *Semver) String() string {
	if v.Minor == "" {
		return v.Major
	} else if v.Patch == "" {
		return fmt.Sprintf("%s.%s", v.Major, v.Minor)
	} else if v.Other == "" {
		return fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Patch)
	} else {
		return fmt.Sprintf("%s.%s.%s.%s", v.Major, v.Minor, v.Patch, v.Other)
	}
}

func (v Semver) LessThan(o Semver) bool {
	if v.Major != o.Major {
		val, err := v.lessThan(v.Major, o.Major)
		if err != nil {
			logrus.Warnf("unable to compare '%s' & '%s'", v.String(), o.String())
			return false
		}
		return val
	}

	if v.Minor != o.Minor {
		val, err := v.lessThan(v.Minor, o.Minor)
		if err != nil {
			logrus.Warnf("unable to compare '%s' & '%s'", v.String(), o.String())
			return false
		}
		return val
	}

	if v.Patch != o.Patch {
		val, err := v.lessThan(v.Patch, o.Patch)
		if err != nil {
			logrus.Warnf("unable to compare '%s' & '%s'", v.String(), o.String())
			return false
		}
		return val
	}

	return strings.Compare(v.Other, o.Other) < 0
}

func (v Semver) lessThan(v1 string, v2 string) (bool, error) {
	i1, err := toInt(v1)
	if err != nil {
		return false, err
	}
	i2, err := toInt(v2)
	if err != nil {
		return false, err
	}
	return i1 < i2, nil
}

type bySemVer []string

func (s bySemVer) Len() int {
	return len(s)
}
func (s bySemVer) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s bySemVer) Less(i, j int) bool {
	v1 := NewVersion(s[i])
	v2 := NewVersion(s[j])

	return v1.LessThan(v2)
}