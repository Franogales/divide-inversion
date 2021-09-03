package main

import (
	"errors"
	"fmt"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}
type CresitType struct {
	Credit300 int32
	Credit500 int32
	Credit700 int32
}

func (cr *CresitType) Assign(investment int32) (int32, int32, int32, error) {
	const (
		creditBase int32 = 100
		creditA    int32 = 300
		creditB    int32 = 500
		creditC    int32 = 700
	)
	if investment < creditA || investment%creditBase != 0 {
		cr.Credit300, cr.Credit500, cr.Credit700 = 0, 0, 0
		return cr.Credit300, cr.Credit500, cr.Credit700, errors.New("does not meet minimum requirements")
	}
	var residuo700, residuo500, residuo300 int32
	first700, first500 := true, true
	for {
		if first700 {
			cr.Credit700 = investment / creditC
		}
		if investment >= creditC && cr.Credit700 > 0 {
			residuo700 = investment - (creditC * cr.Credit700)
			//credito700=1
			//residuo700=500
			if residuo700 == 0 && investment == creditC {
				return cr.Credit300, cr.Credit500, cr.Credit700, nil
			}
			if residuo700/creditB > 0 {
				if first500 {
					cr.Credit500 = residuo700 / creditB
				}
				// fmt.Printf("credit300 = %d, credit500 = %d, credit700 = %d \n", cr.Credit300, cr.Credit500, cr.Credit700)
				residuo500 = residuo700 - (creditB * cr.Credit500)
				if residuo500 == 0 {
					return cr.Credit300, cr.Credit500, cr.Credit700, nil
				}
				if residuo500/creditA > 0 {
					cr.Credit300 = residuo500 / creditA
					residuo300 = residuo500 - (creditA * cr.Credit300)
					if residuo300 > 0 {
						cr.Credit700 -= 1
						first700 = false
						first500 = true
						continue
					}
					return cr.Credit300, cr.Credit500, cr.Credit700, nil
				} else {
					cr.Credit500 -= 1
					first500 = false
					continue
				}
			} else {
				cr.Credit700 -= 1
				first700 = false
				continue
			}
		}
		if first500 {
			cr.Credit500 = investment / creditB
		}
		if investment >= creditB && cr.Credit700 == 0 && cr.Credit500 > 0 {
			// fmt.Printf("credit300 = %d, credit500 = %d, credit700 = %d \n", cr.Credit300, cr.Credit500, cr.Credit700)
			residuo500 = investment - (creditB * cr.Credit500)
			if residuo500 == 0 && investment == creditB {
				return cr.Credit300, cr.Credit500, cr.Credit700, nil
			}
			if residuo500/creditA > 0 {
				cr.Credit300 = residuo500 / creditA
				residuo300 = residuo500 - (creditA * cr.Credit300)
				if residuo300 > 0 {
					cr.Credit500 -= 1
					first500 = false
					continue
				}
				fmt.Println(cr.Credit300)
				fmt.Println(cr.Credit500)
				fmt.Println(cr.Credit700)
				return cr.Credit300, cr.Credit500, cr.Credit700, nil
			} else {
				cr.Credit500 -= 1
				first500 = false
				continue
			}
		}
		cr.Credit300 = investment / creditA
		if investment >= creditA && cr.Credit700 == 0 && cr.Credit500 == 0 && cr.Credit300 > 0 {
			residuo300 = investment - (creditA * cr.Credit300)
			if residuo300 > 0 {
				return cr.Credit300, cr.Credit500, cr.Credit700, errors.New("the investment could not be assigned in any credit")
			}
			return cr.Credit300, cr.Credit500, cr.Credit700, nil
		}
	}
}
