// This is a port of the original lunar lander written in FOCAL
// Source: http://www.vintage-basic.net/bcg/lunar.bas
// It is ported as close as possible, so don't take this as idiomatic code :-)

package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat(" ", 33), "LUNAR")
	fmt.Println(strings.Repeat(" ", 15), "CREATIVE COMPUTING MORRISTOWN, NEW JERSEY")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("THIS IS A COMPUTER SIMULATION OF AN APOLLO LUNAR")
	fmt.Println("LANDING CAPSULE.")
	fmt.Println()
	fmt.Println()
	fmt.Println("THE ON-BOARD COMPUTER HAS FAILED (IT WAS MADE BY")
	fmt.Println("XEROX) SO YOU HAVE TO LAND THE CAPSULE MANUALLY.")
line70:
	fmt.Println()
	fmt.Println("SET BURN RATE OF RETRO ROCKETS TO ANY VALUE BETWEEN")
	fmt.Println("0 (FREE FALL) AND 200 (MAXIMUM BURN) POUNDS PER SECOND.")
	fmt.Println("SET NEW BURN RATE EVERY 10 SECONDS.")
	fmt.Println()
	fmt.Println("CAPSULE WEIGHT 32,500 LBS; FUEL WEIGHT 16,500 LBS.")
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("GOOD LUCK")
	L := 0.0
	fmt.Println()
	fmt.Println("SEC", "MI + FT", "MPH", "LB FUEL", "BURN RATE")
	A := 120.0
	V := 1.0
	M := 33000.0
	N := 16500.0
	G := 1E-03
	Z := 1.8
line150:
	fmt.Println(L, int(A), int(5280*(A-float64(int(A)))), 3600*V, M-N)
	K := 0.0
	fmt.Scanf("%f", &K)
	T := 10.0
	D, I, J, S, W := 0.0, 0.0, 0.0, 0.0, 0.0

	GOSUB330 := func() {
		L = L + S
		T = T - S
		M = M - S*K
		A = I
		V = J
	}

	GOSUB420 := func() {
		Q := S * K / M
		J = V + G*S + Z*(-Q-Q*Q/2-math.Pow(Q, 3)/3-math.Pow(Q, 4)/4-math.Pow(Q, 5)/5)
		I = A - G*S*S/2 - V*S + Z*S*(Q/2+math.Pow(Q, 2)/6+math.Pow(Q, 3)/12+math.Pow(Q, 4)/20+math.Pow(Q, 5)/30)
	}

line160:
	if M-N < 1E-03 {
		goto line240
	}
	if T < 1E-03 {
		goto line150
	}
	S = T
	if M >= N+S*K {
		goto line200
	}
	S = (M - N) / K
	I = 0.0
	J = 0.0
line200:
	GOSUB420()
	if I <= 0 {
		goto line340
	}
	if V <= 0 {
		goto line230
	}
	if J < 0 {
		goto line370
	}
line230:
	GOSUB330()
	goto line160
line240:
	fmt.Println("FUEL OUT AT", L, "SECONDS")
	S = (-V + math.Sqrt(V*V+2*A*G)) / G
	V = V + G*S
	L = L + S
line260:
	W = 3600 * V
	fmt.Println("ON MOON AT", L, "SECONDS - IMPACT VELOCITY", W, "MPH")
	if W <= 1.2 {
		fmt.Println("PERFECT LANDING!")
		goto line440
	}
	if W <= 10 {
		fmt.Println("GOOD LANDING (COULD RE BETTER)")
		goto line440
	}
	if W > 60 {
		goto line300
	}
	fmt.Println("CRAFT DAMAGE... YOU'RE STRANDED HERE UNTIL A RESCUE")
	fmt.Println("PARTY ARRIVES. HOPE YOU HAVE ENOUGH OXYGEN!")
	goto line440
line300:
	fmt.Println("SORRY THERE NERE NO SURVIVORS. YOU BLOW IT!")
	fmt.Println("IN FACT, YOU BLASTED A NEW LUNAR CRATER", W*.227, "FEET DEEP!")
	goto line440

line340:
	if S < 5E-03 {
		goto line260
	}
	D = V + math.Sqrt(V*V+2*A*(G-Z*K/M))
	S = 2 * A / D
	GOSUB420()
	GOSUB330()
	goto line340
line370:
	W = (1 - M*G/(Z*K)) / 2
	S = M*V/(Z*K*(W+math.Sqrt(W*W+V/Z))) + .05
	GOSUB420()
	if I <= 0 {
		goto line340
	}
	GOSUB330()
	if J > 0 {
		goto line160
	}
	if V > 0 {
		goto line370
	}
	goto line160
line440:
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("TRY AGAIN??")
	goto line70
}
