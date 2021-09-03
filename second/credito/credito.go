package credito

import (
	"errors"
	"fmt"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditType struct {
	Credit300 int32 `json:"credit_type_300"`
	Credit500 int32 `json:"credit_type_500"`
	Credit700 int32 `json:"credit_type_700"`
}

func New() *CreditType {
	return &CreditType{}
}
func (cr *CreditType) Assign(investment int32) (int32, int32, int32, error) {
	const (
		creditBase int32 = 100 /* solo usado para marcar la base de los creditos */
		creditA    int32 = 300 /* El tipo de credito A es de 300 */
		creditB    int32 = 500 /* El tipo de credito B es de 500 */
		creditC    int32 = 700 /* El tipo de credito C es de 700 */
	)
	/* Validamos que se cumpla lo siguiente. Inversion sea divisible entre 100 sin residuo. Que la inversion sea al menos del primer tipo de credito(300)  */
	if investment < creditA || investment%creditBase != 0 {
		cr.Credit300, cr.Credit500, cr.Credit700 = 0, 0, 0
		return cr.Credit300, cr.Credit500, cr.Credit700, errors.New("does not meet minimum requirements")
	}
	var residuo700, residuo500, residuo300 int32
	/* first700 y first500 son solo un flag para identificar la primera iteracion por el tipo de credito */
	first700, first500 := true, true
	for {
		if first700 {
			cr.Credit700 = investment / creditC
		}
		/* si la inversion es mayor/igual que el tipo de credito de 700 y al menos se ha podido otorgar un credito de 700  */

		if investment >= creditC && cr.Credit700 > 0 {
			/* el residuo es igual a la diferencia entre la inversion menos la mutiplicacion de el tipo de credito (700) por la cantidad de credito otorgado del tipo C(700) */
			residuo700 = investment - (creditC * cr.Credit700)
			/* aqui valide la inversion por el tipo de credito para poder otorgar otos tipos de credito mas que solos de 700 */
			if residuo700 == 0 && investment == creditC {
				return cr.Credit300, cr.Credit500, cr.Credit700, nil
			}
			/* si el residuo es al menos divisible una vez entre el tipo de credito B(500), buscamos los posibles creditos entregables de 500 */
			if residuo700/creditB > 0 {
				if first500 {
					cr.Credit500 = residuo700 / creditB
				}
				residuo500 = residuo700 - (creditB * cr.Credit500)
				if residuo500 == 0 {
					return cr.Credit300, cr.Credit500, cr.Credit700, nil
				}
				/* si el residuo de la divicion entre residuo500/300 es al menos 1 , quiere decir que hay mas creditos que otorgar de 300   */
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
		/* si la inversion es menor que 700, empezamos a sacar los posibles creditos de tipo B y A (500 y 300) */
		if first500 {
			cr.Credit500 = investment / creditB
		}
		if investment >= creditB && cr.Credit700 == 0 && cr.Credit500 > 0 {
			fmt.Printf("%d,%d,%d \n", cr.Credit300, cr.Credit500, cr.Credit700)
			residuo500 = investment - (creditB * cr.Credit500)
			if residuo500 == 0 {
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
		/* si la inversion es menor que 500, sacamos los posibles creditos del tipo A */
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
