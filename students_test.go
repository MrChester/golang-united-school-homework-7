package coverage

import (
	"os"
	"time"
	"testing"
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
	}{
		"empty": 
		{
			People: People{},
		},
		"one value": 
		{
			People: People{
				{
				firstName: "Ivan",
				lastName: "Petrov",
				birthDay: time.Date(1987, 07, 06, 21, 34, 0, 0, time.UTC),
				},
			},
		},
		"several value": 
		{
			People: People{
				{
					firstName: "Ivan",
					lastName:  "Bobr",
					birthDay:  time.Date(1999, time.November, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Ivan",
					lastName:  "Petrov",
					birthDay:  time.Date(1999, time.November, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Petr",
					lastName:  "Kotov",
					birthDay:  time.Date(1987, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Lidia",
					lastName:  "Azamatova",
					birthDay:  time.Date(1999, time.November, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Ivan",
					lastName:  "Bodrov",
					birthDay:  time.Date(1987, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Petro",
					lastName:  "Petrovskiy",
					birthDay:  time.Date(1987, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
				{
					firstName: "Petro",
					lastName:  "Petrovskiy",
					birthDay:  time.Date(1987, time.July, 06, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			v := tt.People.Len()

			if v > 1 {
				lv1 := tt.People.Less(0, 1)
				if lv1 {
					tt.People.Swap(0, 1)

				}
				lv2 := tt.People.Less(4, 5)
				if lv2 {
					tt.People.Swap(4, 5)

				}
				lv3 := tt.People.Less(3, 4)
				if lv3 {
					tt.People.Swap(3, 4)

				}
			}
		})
	}
}

