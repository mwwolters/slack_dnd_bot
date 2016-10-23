package main

import (
  "fmt"
  "encoding/json"
  "os"
  "io/ioutil"
)

type GameState struct {
  CharMap map[string]Character
}

func (g GameState) GetChar(name string) (Character, error) {
  c := gs.CharMap[name]
  emptyChar := Character{}
  if c == emptyChar {
    return c, fmt.Errorf("character %v doesn't exist", name)
  }
  return c, nil
}

func (g GameState) LoadChar(name string) error {
  c := Character{}
  charFile, err := os.Open(name + ".json")
  if err != nil {
    return fmt.Errorf("Could not open character file: %v", err)
  }
  jsonParser := json.NewDecoder(charFile)
  if err = jsonParser.Decode(&c); err != nil {
    return fmt.Errorf("could not load character %v: %v", name, err)
  }
  g.CharMap[name] = c
  return nil
}

func (g GameState) SaveChar(name string) error {
  c := g.CharMap[name]
  emptyChar := Character{}
  if c == emptyChar {
    return fmt.Errorf("character not found")
  }
  out, err := json.Marshal(c)
  if err != nil {
    return fmt.Errorf("failed to make json: %v", err)
  }
  ioutil.WriteFile(name + ".json", out, 0644)
  return nil
}

func (g GameState) AddChar(c Character) {
  g.CharMap[c.Name] = c
}

func (g GameState) Load(parts []string) (string, error) {
  if len(parts) < 4 {
    return "", fmt.Errorf("too few arguments for load")
  }
  switch {
  case parts[2] == "character":
    err := g.LoadChar(parts[3])
    if err != nil {
      return "", fmt.Errorf("load error: %v", err)
    }
    return fmt.Sprintf("Loaded character: %v", parts[3]), nil
  }
  return "", fmt.Errorf("Invalid load command: %v", parts[2])
}

func (g GameState) Save(parts []string) (string, error) {
  if len(parts) < 4 {
    return "", fmt.Errorf("too few arguments for load")
  }
  switch {
  case parts[2] == "character":
    err := g.SaveChar(parts[3])
    if err != nil {
      return "", fmt.Errorf("save error: %v", err)
    }
    return fmt.Sprintf("Saved character: %v", parts[3]), nil
  }
  return "", fmt.Errorf("Invalid save command: %v", parts[2])
}
