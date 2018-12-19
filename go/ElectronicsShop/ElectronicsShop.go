package ElectronicsShop


func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	*/
	var budget int32 = 0
	for _, k := range keyboards  {
		for _, d := range drives{
			if (k + d) <= b && (k + d) > budget{
				budget = k + d
			}
		}
	}

	if budget == 0{
		return -1
	}
	return budget

}

func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	return getMoneySpent(keyboards, drives, b)
}


