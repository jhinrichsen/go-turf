package main

import (
	"log"
)

const strike = 10

// Game class is obviously needed for the bowling kata, says Uncle bob.
type Game struct {
	// Now take a deep breath, and enjoy the absence of constructors, new,
	// memory allocation, static initializers e.a.
	pins []int
}

func main() {
}

// byte should do it, but spec requires an int
// By allowing a variadic number of pins, we get rid of rollMany()
func (g *Game) roll(pins ...int) {
	// Frances Campoy on 'append to nil slices',
	// https://www.youtube.com/watch?v=ynoY2xz-F8s at 19:10
	g.pins = append(g.pins, pins...)
}

// Anything but the last roll
func (g Game) isMiddleRoll(idx int) bool {
	return idx+1 < len(g.pins)
}

func (g Game) isSpare(idx int) bool {
	// Cannot roll a spare in tail position
	return (g.isMiddleRoll(idx)) &&
		g.pins[idx]+g.pins[idx+1] == strike
}

func (g Game) isStrike(idx int) bool {
	return g.pins[idx] == strike
}

func (g Game) traditionalScore() int {
	score := 0
	roll := 0
	frame := 0
	secondRoll := false
	for frame < 10 {
		log.Printf("roll %v, frame %v\n", roll, frame)
		switch {
		case g.isSpare(roll):
			log.Printf("#%v is a spare\n", roll)
			score += g.pins[roll]
			roll++
			score += g.pins[roll]
			frame++
		case g.isStrike(roll):
			log.Printf("#%v is a strike\n", roll)
			score += g.pins[roll]
			if g.isMiddleRoll(roll) {
				score += g.pins[roll+1]
			}
			if g.isMiddleRoll(roll + 1) {
				score += g.pins[roll+2]
			}
			frame++
		default:
			score += g.pins[roll]
			if secondRoll {
				frame++
			}
			secondRoll = !secondRoll
		}
		log.Printf("New score after %v rolls is %v\n", roll, score)
		roll++
	}
	log.Printf("score for %v pins %v is %v\n", len(g.pins), g.pins, score)
	return score
}
