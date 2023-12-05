package day_one

import (
  "log"
  "os"
  "strconv"
  "strings"
)

func Answer() (int) {
  var numbers []int
  input, err := os.ReadFile("./input/day_one.txt")

  if err != nil {
    log.Fatal(err)
  }

  lines := strings.Split(string(input), "\n")

  for i := 0; i < len(lines) - 1; i++ {
    first := GetFirstInt(lines[i])
    last := GetLastInt(lines[i])

    number, err := strconv.Atoi(first + last)

    if err != nil {
      log.Fatal(err)
    }

    numbers = append(numbers, number)
  }

  return Reduce(numbers)
}

func GetFirstInt(line string) (string) {
  var firstInt string

  for i := 0; i <= len(line) - 1; i++ {
    character := string(line[i])
    _, err := strconv.Atoi(character)

    if err == nil {
      firstInt = character
      break
    }
  }

  return firstInt
}

func GetLastInt(line string) (string) {
  var firstInt string

  for i := len(line) - 1; i >= 0; i-- {
    character := string(line[i])
    _, err := strconv.Atoi(character)

    if err == nil {
      firstInt = character
      break
    }
  }

  return firstInt
}

func Reduce(numbers []int) (int) {
  sum := 0

  for i := 0; i < len(numbers); i++ {
    sum += numbers[i]
  }

  return sum
}
