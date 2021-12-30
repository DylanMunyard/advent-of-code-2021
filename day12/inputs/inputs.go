package inputs

import (
	"regexp"
	"strings"
)

var Arr = []string{
	"vp-BY",
	"ui-oo",
	"kk-IY",
	"ij-vp",
	"oo-start",
	"SP-ij",
	"kg-uj",
	"ij-UH",
	"SP-end",
	"oo-IY",
	"SP-kk",
	"SP-vp",
	"ui-ij",
	"UH-ui",
	"ij-IY",
	"start-ui",
	"IY-ui",
	"uj-ui",
	"kk-oo",
	"IY-start",
	"end-vp",
	"uj-UH",
	"ij-kk",
	"UH-end",
	"UH-kk",
}

type Cave struct {
	Big      bool
	End      bool
	Start    bool
	Connects []string
}

var rIsUpper, _ = regexp.Compile("[^a-z]")

func Map() map[string]Cave {
	paths := make(map[string]Cave)
	for _, p := range Arr {
		pieces := strings.SplitN(p, "-", 2)
		from := pieces[0]
		to := pieces[1]
		fCave := parseCave(from, to)
		tCave := parseCave(to, from)
		if path, key := paths[from]; key {
			fCave.Connects = append(path.Connects, fCave.Connects...)
		}
		paths[from] = fCave
		if path, key := paths[to]; key {
			tCave.Connects = append(path.Connects, tCave.Connects...)
		}
		paths[to] = tCave
	}

	return paths
}

func IsBig(cave string) bool {
	return rIsUpper.MatchString(cave)
}

func parseCave(sCave string, connect string) Cave {
	var cave Cave
	cave.Start = sCave == "start"
	cave.End = sCave == "end"
	cave.Big = rIsUpper.MatchString(sCave)
	cave.Connects = []string{connect}
	return cave
}
