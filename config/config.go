package config

import "github.com/BurntSushi/toml"

type PointList struct {
	ANNIVERSARY Anniversary `toml:"anniversary`
}

type Anniversary struct {
	LIVE []Live
}

type Live struct {
	NAME  string `toml: name`
	POINT int    `toml: point`
	NEED  int    `toml: need`
}

func ParseToml(fileName string) (PointList, error) {
	var pointlist PointList
	_, err := toml.DecodeFile(fileName, &pointlist)
	return pointlist, err
}
