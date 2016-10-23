package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

var charMap = map[string]Character{}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: dndbot slack-bot-token\n")
		os.Exit(1)
	}

	ws, id := slackConnect(os.Args[1])
	fmt.Println("dndbot ready, ^C exits")

	for {
		m, err := getMessage(ws)
		if err != nil {
			log.Fatal(err)
		}

		if m.Type == "message" && strings.HasPrefix(m.Text, "<@"+id+">") {
			parseMessage(ws, m)
		}
	}
}

type Roll struct {
	dice int
	size int
}

func parseRoll(s string) (Roll, error) {
	nums := strings.Split(s, "d")
	if len(nums) < 1 {
		return Roll{}, fmt.Errorf("Incorrect format: %v", s)
	} else if len(nums) == 1 {
		size, err := strconv.Atoi(nums[0])
		if err != nil {
			return Roll{}, fmt.Errorf("Could not convert numbers: %v", nums)
		}
		return Roll{1, size}, nil
	} else {
		d, err := strconv.Atoi(nums[0])
		size, err := strconv.Atoi(nums[1])
		if err != nil {
			return Roll{}, fmt.Errorf("Could not convert numbers: %v", nums)
		}
		return Roll{d, size}, nil
	}
}

func getRollResult(s string) (string, error) {
	roll, err := parseRoll(s)
	if err != nil {
		return "", err
	}
	result := 0
	for i := 0; i < roll.dice; i++ {
		result = result + rand.Intn(roll.size+1)
	}
	return strconv.Itoa(result), nil
}

func roll(ws *websocket.Conn, m Message, parts []string) {
  text, err := getRollResult(parts[2])
  m.Text = text
  if err != nil {
    m.Text = fmt.Sprintf("incorrect format: %v", err)
  }
  postMessage(ws, m)
}

func errorMessage(ws *websocket.Conn, m Message, err error) {
  m.Text = fmt.Sprintf("sorry, does not compute: %v\n", err)
  postMessage(ws, m)
}

func getChar(parts []string) (Character, error) {
  if len(parts) < 3 {
    return Character{}, fmt.Errorf("not enough args")
  }
  name := parts[2]
  c := charMap[name]
  emptyChar := Character{}
  if c == emptyChar {
    return c, fmt.Errorf("character %v doesn't exist", name)
  }
  return c, nil
}

func printHelp(ws *websocket.Conn, m Message, parts []string) {
  m.Text = `General usage: @testbot <command> args
  
  Available commands:
     help:           prints this help message
     roll:           rolls dice specified by <#1>d<#2> where #1 is
                     the number of dice and #2 is the size of dice
     createRandom:   creates a random character with random stats
                     Note: not balanced at all, for testing
     printChar:      prints everything about a character`
  postMessage(ws, m)
}

func parseMessage(ws *websocket.Conn, m Message) {
	log.Printf("Got message from channel: %v\n%v", m.Channel, m.Text)
	parts := strings.Fields(m.Text)
	if len(parts) < 2 {
		m.Text = fmt.Sprintf("sorry, does not compute\n")
		postMessage(ws, m)
	}
	switch {
	case parts[1] == "help":
	  go printHelp(ws, m, parts)
	  break
  case parts[1] == "roll":
		go roll(ws, m, parts) 
		break
  case parts[1] == "createRandom":
    c := createRandom()  
    charMap[c.Name] = c
    s := c.print()
    m.Text = s
    postMessage(ws, m)
    break
  case parts[1] == "printChar":
    c, err := getChar(parts)
    if err != nil {
      errorMessage(ws, m, err)  
    }
    m.Text = c.print()
    postMessage(ws, m)
    break
  err = fmt.Errorf("Not a valid command")
  errorMessage(ws, m, err)
	}
}
