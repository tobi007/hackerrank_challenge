package DrawingBook


func pageCount(n int32, p int32) int32 {

	f := p/2
	var b int32
	if n % 2 == 0{
		b = (n - p+1)/2
	} else {
		b = (n - p)/2
	}

	return min(f,b)

}

func PageCount(n int32, p int32) int32 {
	 return pageCount(n, p)

}

func min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}


