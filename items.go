package main

import (
  "fmt"
)

type Item interface {
  Attack() int
  Block() int
  Use() string
  Print() string
}

type DamageTypes struct {
  Slashing int `json:"slashing"`
  Piercing int `json:"piercing"`
  Bludgeoning int `json:"bludgeoning"`
  Shocking int `json:"shocking"`
  Burning int `json:"burning"`
  Magic int `json:"magic"`
}

type ItemData struct {
  Kills int
  Blocks int
}

type ItemState struct {
  Current int `json:"current"`
  Max int `json:"max"`
}

type Weapon struct {
  Name string `json:"name"`
  DType DamageTypes `json:"dtype"`
  State ItemState `json:"state"`
  Data ItemData `json:"data"`
}

type Armor struct {
  Name string `json:"name"`
  DProtect DamageTypes `json:"dprotect"`
  State ItemState `json:"state"`
  Data ItemData `json:"data"`
}

func CreateItem(parts []string) Item {
  w := Weapon{Name:"Method Not Implemented"}
  return w
}

func (w Weapon) Attack() int {
  return 0
}

func (w Weapon) Block() int {
  return 0
}

func (w Weapon) Use() string {
  return fmt.Sprintf("%v was used", w.Name)
}

func (w Weapon) Print() string {
  return fmt.Sprintf("%v", w.Name)
}

func (a Armor) Attack() int {
  return 0
}

func (a Armor) Block() int {
  return 0
}

func (a Armor) Use() string {
  return fmt.Sprintf("%v was used", a.Name)
}

func (a Armor) Print() string {
  return fmt.Sprintf("%v", a.Name)
}
