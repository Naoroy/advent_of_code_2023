package day_two

import (
  "log"
  "os"
  "strings"
)

func Answer() (int) {
  var idSum int
  input, err := os.ReadFile("./input/day_two.txt")

//   redCubes := 12
//   greenCubes := 13
//   blueCubes := 14
  if err != nil {
    log.Fatal(err)
  }
  games := GetGames(string(input))

  for i := 0; i <= len(games) -1; i++ {
    a := FormatThrows(games[i])
    log.Println(games[i])
    log.Println(a)
  }
  // Iterate over games and validate each (ValidateGame)
    // split by comma and find <int> <color> then validate against total available
  return idSum
}

func FormatThrows(throws []string) ([]string) {
  var formatedThrows []string
  for j := 0; j <= len(throws) -1; j++ {
    throws := strings.Split(throws[j], ",")
    n := strings.Split(throws[0], ",")[0]
    formatedThrows = strings.Fields(n)
  }
  return formatedThrows
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
