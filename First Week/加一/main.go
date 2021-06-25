func plusOne(digits []int) []int { 
    add := 0
    for i := len(digits) - 1; i >= 0; i-- {
        var tempResult int
        if i == len(digits) - 1 {
            tempResult = (digits[i] + 1) % 10
            add = (digits[i] + 1) / 10
            digits[i] = tempResult
            continue
        }
        tempResult = (digits[i] + add) % 10
        add = (digits[i] + add) / 10
        digits[i] = tempResult
    }
    if add != 0 {
        digits = append([]int{add}, digits...)
    }
    return digits
}