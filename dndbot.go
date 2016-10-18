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

func parseMessage(ws *websocket.Conn, m Message) {
	parts := strings.Fields(m.Text)
	if len(parts) < 2 {
		m.Text = fmt.Sprintf("sorry, does not compute\n")
		postMessage(ws, m)
	}
	switch {
  case parts[1] == "roll":
		go roll(ws, m, parts) 
		break
  case parts[1] == "createRandom":
    c := createRandom()  
    s := c.print()
    m.Text = s
    postMessage(ws, m)
    break
  m.Text = fmt.Sprintf("sorry, does not compute\n")
  postMessage(ws, m)
	}
}
