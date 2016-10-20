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

type DerivedStat struct {
  CurrentValue int
  MaxValue int
}

type DerivedStats struct {
  Health DerivedStat
  Armor DerivedStat
  Stamina DerivedStat
  CarryWeight DerivedStat
}

type Character struct {
  Name string
  Stats Stats
  DStats DerivedStats
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

func (c Character) print() string {
  var buffer bytes.Buffer
  buffer.WriteString(c.Name + "\n")
  buffer.WriteString(c.DStats.print())
  buffer.WriteString(c.Stats.print())
  return buffer.String()
}
