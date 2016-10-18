package main

import (
  "bytes"
  "fmt"
  "math/rand"
  "reflect"
)

type Stats struct {
  Strength int
  Dexterity int
  Intelligence int
  Charisma int
  Wisdom int
  Luck int
  Constitution int
}

type Character struct {
  Name string
  Stats Stats
  // to come:
  // StatusEffects
  // Description
  // Skills
}

func createRandom() Character {
  stats := Stats{
    Strength: rand.Intn(100),
    Dexterity: rand.Intn(100),
    Intelligence: rand.Intn(100),
    Charisma: rand.Intn(100),
    Wisdom: rand.Intn(100),
    Luck: rand.Intn(100),
    Constitution: rand.Intn(100),
  }
  c := Character{"RandomName", stats}
  return c
}

func (s Stats) print() string {
  var buffer bytes.Buffer
  r := reflect.ValueOf(&s).Elem()
  typeOfS := r.Type()
  for i := 0; i < r.NumField(); i++ {
    f := r.Field(i)
    stat_str := fmt.Sprintf("%s: %v\n", typeOfS.Field(i).Name, f.Interface())
    buffer.WriteString(stat_str)
  }
  return buffer.String()
}

func (c Character) print() string {
  var buffer bytes.Buffer
  buffer.WriteString(c.Name + "\n")
  buffer.WriteString(c.Stats.print())
  return buffer.String()
}
