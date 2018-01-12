package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func getNumber(r *bufio.Reader, question string) int {
	fmt.Println(question)
	carte1s, _, err := r.ReadLine()
	if err != nil {
		log.Println(err)
		return getNumber(r, question)
	}

	carte1, err := strconv.Atoi(string(carte1s))

	if err != nil {
		log.Println("Put a number in between 1-10.")
		return getNumber(r, question)
	}

	switch {
	case carte1 < 1:
		log.Println("Number is to small. Must be between 1-10.")
		return getNumber(r, question)
	case carte1 > 10:
		log.Println("Number is to big. Must be between 1-10.")
		return getNumber(r, question)
	}

	return carte1
}

type action int

func (a action) String() string {
	switch a {
	case SPLIT:
		return "Split your hand."
	case TAP:
		return "You should tap."
	case STAND:
		return "You should stand."
	case BLACKJACK:
		return "Stand you have blackjack!"
	default:
		return "Invalid action."
	}
}

const (
	SPLIT action = iota
	TAP
	STAND
	BLACKJACK
)

var states = [10][10]action{
	// {ACE, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	{SPLIT, TAP, TAP, TAP, TAP, TAP, TAP, TAP, TAP, BLACKJACK},           //1
	{TAP, SPLIT, TAP, TAP, TAP, TAP, TAP, TAP, TAP, TAP},                 //2
	{TAP, TAP, SPLIT, TAP, TAP, TAP, TAP, TAP, TAP, TAP},                 //3
	{TAP, TAP, TAP, SPLIT, TAP, TAP, TAP, TAP, TAP, TAP},                 //4
	{TAP, TAP, TAP, TAP, SPLIT, TAP, TAP, TAP, TAP, STAND},               //5
	{TAP, TAP, TAP, TAP, TAP, SPLIT, TAP, TAP, STAND, STAND},             //6
	{TAP, TAP, TAP, TAP, TAP, TAP, SPLIT, STAND, STAND, STAND},           //7
	{TAP, TAP, TAP, TAP, TAP, TAP, TAP, SPLIT, STAND, STAND},             //8
	{TAP, TAP, TAP, TAP, TAP, STAND, STAND, STAND, SPLIT, STAND},         //9
	{BLACKJACK, TAP, TAP, TAP, STAND, STAND, STAND, STAND, STAND, SPLIT}, //10
}
