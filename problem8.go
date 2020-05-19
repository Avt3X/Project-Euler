package main

import (
	"strconv"
	"strings"
)

var str string = "73167176531330624919225119674426574742355349194934969835203127745063262395783180169848018694788518438586156078911294949545950173795833195285320880551112540698747158523863050715693290963295227443043557668966489504452445231617318564030987111217223831136222989342338030813533627661428280644448664523874930358907296290491560440772390713810515859307960866701724271218839987979087922749219016997208880937766572733300105336788122023542180975125454059475224352584907711670556013604839586446706324415722155397"

func problem8_2() int {
	res := 1
	for i := 0; i < len(str)-13; i++ {
		p := 1
		for j := 0; j < 13; j++ {
			re, _ := strconv.Atoi(string(str[i+j]))
			p *= re
		}

		if res < p {
			res = p
		}
	}

	return res
}

func problem8() int {
	res := 1
	split := strings.Split(str, "0")
	for i := 0; i < len(split); i++ {
		r := split[i]
		if len(r) >= 13 {
			last := 13
			for i2 := range r {
				if last > len(r) {
					break
				}

				subStr := r[i2:last]
				p := 1
				for i3 := range subStr {
					re, _ := strconv.Atoi(string(subStr[i3]))
					p *= re
				}
				if res < p {
					res = p
				}
				last++
			}
		}
	}

	return res
}
