package DayOfTheProgrammer

import "fmt"

//import "fmt"

// Complete the solve function below.


func Solve(year int32) string {

	var days int32


	if year >= 1919 {
		days = isLeapGreg(year)
	}else {
		days = isLeapJuli(year)
	}

	rem := 256 - days

	if year == 1918{
		rem = rem + 13
	}
	return fmt.Sprintf("%d.%s.%d", rem, "09", year)
}

func isLeapGreg(year int32) int32  {
	if year % 400 == 0{
		return 244
	}

	if year % 4 == 0 && year % 100 != 0{
		return 244
	}

	return 243
}

func isLeapJuli(year int32) int32  {
	if year % 4 == 0{
		return 244
	}
	return 243
}
