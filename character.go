package main

import (
  "bytes"
  "fmt"
  "math/rand"
  "reflect"
)

type Stats struct {
  Strength int `json:"strength"`
  Dexterity int `json:"dexterity"`
  Intelligence int `json:"intelligence"`
  Charisma int `json:"charisma"`
  Wisdom int `json:"wisdom"`
  Luck int `json:"luck"`
  Constitution int `json:"constitution"`
}

type DerivedStat struct {
  CurrentValue int `json:"current"`
  MaxValue int `json:"max"`
}

type DerivedStats struct {
  Health DerivedStat `json:"health"`
  Armor DerivedStat `json:"armor"`
  Stamina DerivedStat `json:"stamina"`
  CarryWeight DerivedStat `json:"carry"`
}

type Character struct {
  Name string `json:"name"`
  Stats Stats `json:"stats"`
  DStats DerivedStats `json:"dstats"`
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
  dstats := DerivedStats{}
  c := Character{"RandomName", stats, dstats}
  return c
}

func CreateChar(parts []string) Character {
  if len(parts) < 1 {
    return createRandom()
  }
  switch {
  case  parts[0] == "random":
    return createRandom()
  }
  return createRandom()
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

func (d DerivedStat) print() string {
  return fmt.Sprintf("%v/%v", d.CurrentValue, d.MaxValue)
}

func (d DerivedStats) print() string {
  var buffer bytes.Buffer
  buffer.WriteString("Health: " + d.Health.print() + "\n")
  return buffer.String() 
}

func (c Character) Print() string {
  var buffer bytes.Buffer
  buffer.WriteString(c.Name + "\n")
  buffer.WriteString(c.DStats.print())
  buffer.WriteString(c.Stats.print())
  return buffer.String()
}
