package main

import (
	"testing"
)

func TestCobranza(t *testing.T) {
	setTests := []struct {
		Name        string
		Input       int32
		Expected1   int32
		Expected2   int32
		Expected3   int32
		ExpectedErr string
	}{
		{
			Name:        "Igual que 50",
			Input:       50,
			Expected1:   0,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "does not meet minimum requirements",
		},
		{
			Name:        "Igual que 200",
			Input:       200,
			Expected1:   0,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "does not meet minimum requirements",
		},
		{
			Name:        "Igual que 300",
			Input:       300,
			Expected1:   1,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 400",
			Input:       400,
			Expected1:   0,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "the investment could not be assigned in any credit",
		},
		{
			Name:        "Igual que 500",
			Input:       500,
			Expected1:   0,
			Expected2:   1,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 600",
			Input:       600,
			Expected1:   2,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 700",
			Input:       700,
			Expected1:   1,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 800",
			Input:       800,
			Expected1:   1,
			Expected2:   1,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 900",
			Input:       900,
			Expected1:   3,
			Expected2:   0,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 1000",
			Input:       1000,
			Expected1:   0,
			Expected2:   2,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 1100",
			Input:       1100,
			Expected1:   2,
			Expected2:   1,
			Expected3:   0,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 1200",
			Input:       1200,
			Expected1:   0,
			Expected2:   1,
			Expected3:   1,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 1300",
			Input:       1300,
			Expected1:   2,
			Expected2:   0,
			Expected3:   1,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 3000",
			Input:       3000,
			Expected1:   2,
			Expected2:   2,
			Expected3:   2,
			ExpectedErr: "",
		},
		{
			Name:        "Igual que 6700",
			Input:       6700,
			Expected1:   2,
			Expected2:   1,
			Expected3:   8,
			ExpectedErr: "",
		},
	}

	for _, tt := range setTests {
		t.Run(tt.Name, func(t *testing.T) {
			credit := CresitType{}
			a, b, c, err := credit.Assign(tt.Input)
			if a != tt.Expected1 && b != tt.Expected2 && c != tt.Expected3 && err.Error() != tt.ExpectedErr {
				t.Errorf("Failed %s : expected a=%d, b=%d, c=%d, error=%s. you got %d, %d, %d,%s", tt.Name, tt.Expected1, tt.Expected2, tt.Expected3, tt.ExpectedErr, a, b, c, err.Error())
			}
		})
	}
}
