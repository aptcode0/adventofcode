package solution

import (
	"strconv"
	"strings"
)

type Game struct {
	id    int
	cubes map[string][]int
}

func Parse(data string) []Game {
	lines := strings.Split(data, "\n")
	var games []Game

	for _, l := range lines {
		pos := strings.Index(l, ":")
		id, _ := strconv.Atoi(l[5:pos])
		cubes := countCubesByColor(strings.Split(l[pos+2:], "; "))


		games = append(games, Game{id, cubes})
	}
	return games
}

func countCubesByColor(gameData []string) map[string][]int {
	cubes := make(map[string][]int)
	for _, cubesData := range gameData {
		for _, cubeData := range strings.Split(cubesData, ", ") {
			d := strings.Split(cubeData, " ")
			cnt, _ := strconv.Atoi(d[0])
			color := d[1]
			cubes[color] = append(cubes[color], cnt)
		}
	}
	return cubes
}