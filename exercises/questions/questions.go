package questions

// QuestionOne performs addition without using arithmetic operators
func QuestionOne(x, y int64) int64 {
	/*
		 No arithmetic, so binary operations required
		 addition:
			 0 + 0 = 00
			 0 + 1 = 01
			 1 + 0 = 01
			 1 + 1 = 10
		can split addition into adding the result without carrying one
		with the result of carrying the one, without the addition
		this is adding an XOR of the two numbers with and AND shifted by one bit
		101 + 011 = 110 + 010 = 100 + 100 = 000 + 1000

		Alternatively you can iterate through the bits in increasing order of significance and perform AND, carrying the one to the next bit if required.
	*/

	if y == 0 {
		return x
	}
	addition := x ^ y
	carry := (x & y) << 1
	return QuestionOne(addition, carry)
}
