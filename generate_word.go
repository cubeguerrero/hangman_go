package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func GenerateWord() string {
	contents, _ := ioutil.ReadFile("words.txt")
	c := string(contents)
	words := strings.Split(c, "\n")
	s := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(s)
	return words[r.Intn(len(words)-1)]
}
