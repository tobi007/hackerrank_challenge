package CountingValleys


func countingValleys(n int32, s string) int32 {
	step := map[string]int{"D":-1,"U":1}

	valleyNum := 0
	totalStep := 0
	prevTotalStep := 0
	for _, v := range s {
		prevTotalStep = totalStep
		totalStep = totalStep + step[string(v)]

		if totalStep == 0 && prevTotalStep == -1{
			valleyNum++
		}
	}
	return int32(valleyNum)

}

func CountingValleys(n int32, s string) int32 {

	return countingValleys(n , s)

}
