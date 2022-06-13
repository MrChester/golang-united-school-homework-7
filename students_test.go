package coverage

import (
	"os"
	"time"
	"testing"
	"reflect"	
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func TestPerson(t *testing.T) {
	cases := map[string]struct {
		People People
		expLen int
		expLess bool
	}{
		"len test case": {
			People: People{
				{
					firstName: "Obi Van",
					lastName:  "Kenobi",
					birthDay:  time.Date(1987, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Darth",
					lastName:  "Vader",
					birthDay:  time.Date(2000, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
			},
			expLen: 2,
		},
		"equal birth date": {
			People: People{
				{
					firstName: "Darth",
					lastName:  "Vader",
					birthDay:  time.Date(1989, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Obi Van",
					lastName:  "Kenobi",
					birthDay:  time.Date(1989, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
			},
			expLen: 2,
			expLess: true,
		},
		"equal names": {
			People: People{
				{
					firstName: "Darth",
					lastName:  "Maul",
					birthDay:  time.Date(2000, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Darth",
					lastName:  "Sidius",
					birthDay:  time.Date(2000, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
			},
			expLen: 2,
			expLess: true,

		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			v := tt.People

			len := v.Len()
			less1 := v.Less(0, 1)

			if !reflect.DeepEqual(len, tt.expLen) {
				t.Errorf("the value of len is incorrect exp: %v, got %v", tt.expLen, len)
			}
			if less1 != tt.expLess {
				t.Errorf("the result of Less function is incorrect exp: %v, got %v", tt.expLess, less1)
			} else {
				v.Swap(0, 1)
			}
		})
	}	
}

func TestMatrix(t *testing.T) {
	cases := map[string]struct {
		matrix string
		expRows [][]int
		expCols [][]int
	}{
		"empty matrix": {
			matrix: "",
		},
		"number of rows and cols equal": {
			matrix: "3\n 3 3",
		},
		"valid matrix": {
			matrix: "3 3\n3 3 \n 3 3",
			expRows: [][]int{{3, 3}, {3, 3}, {3, 3}},
			expCols: [][]int{{3, 3, 3}, {3, 3, 3}},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {			
			m, err := New(tt.matrix)

			if err != nil{
				t.Errorf("Error - %v", err)		
			} else {
				r := m.Rows()
				c := m.Cols()
				m.Set(1, 1, 5)
				m.Set(1, 5, 5)
				m.Set(-1, 0, 5)				
				
				if !reflect.DeepEqual(r, tt.expRows) {
					t.Errorf("incorect number of rows - exp: %v, got %v", tt.expRows, r)
				}
				if !reflect.DeepEqual(c, tt.expCols) {
					t.Errorf("incorect number of cols - exp: %v, got %v", tt.expCols, c)
				}	
			}
		})
	}
}

