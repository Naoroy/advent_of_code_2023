package day_two

import (
  "log"
  "os"
  "strings"
)

func Answer() (int) {
  var idSum int
  input, err := os.ReadFile("./input/day_two.txt")

  redCubes := 12
  greenCubes := 13
  blueCubes := 14
  if err != nil {
    log.Fatal(err)
  }
  games := GetGames(string(input))

  // Iterate over games and validate each (ValidateGame)
    // split by comma and find <int> <color> then validate against total available
  return idSum
}

func GetGames(input string) ([][]string) {
  lines := strings.Split(string(input), "\n")
  var throws [][]string

  for i := 0; i <= len(lines) - 1; i++ {
    if lines[i] == "" {
      continue
    }
    lines[i] = strings.Split(lines[i], ":")[1]
    throws = append(throws, strings.Split(lines[i], ";"))
  }

  return throws
}
