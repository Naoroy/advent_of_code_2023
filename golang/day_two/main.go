package day_two

import (
  "log"
  "os"
  "strings"
)

func Answer() (string) {
  input, err := os.ReadFile("./input/day_two.txt")

  if err != nil {
    log.Fatal(err)
  }
  games := GetGames(string(input))

  log.Println(games)
  throws := strings.Split(string(games[0]), ";")
  log.Println(throws)

  return "Day two"
}

func GetGames(input string) ([]string) {
  lines := strings.Split(string(input), "\n")
  log.Println(lines)

  for i := 0; i < len(lines) - 1; i++ {
    lines[i] = strings.Split(lines[i], ":")[1]
  }

  return lines
}
