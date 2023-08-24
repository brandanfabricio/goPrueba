package exponses

func Average(expns ...float32) float32 {

	return Sum(expns...) / float32(len(expns))

}

func Sum(expns ...float32) float32 {

	var sum float32
	for _, expn := range expns {
		sum += expn
	}
	return sum
}

func Max(expns ...float32) float32 {
	if len(expns) ==0{
		return 0
	}
	var max float32 = expns[0]
	for _, expn := range expns {
		if expn > max {
			max = expn
		}


	}

	return max
	
}


func Min(expns ...float32) float32{
	if len(expns) ==0{
		return 0
	}

	var min float32 = expns[0]

	for _, expn := range expns {
		if expn < min {
			min = expn
		}
	}

	return min
}