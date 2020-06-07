func change(amount int, coins []int) int {
    if amount == 0 {
        return 1
    }
    
    if len(coins) == 0 {
        return 0
    }
    
    opt := make([]int, amount+1)
    opt[0] = 1
    
    for _, c := range coins {
        for i := 1; i <= amount; i++ {
            if i - c >= 0 {
                opt[i] += opt[i-c]
            }
        }
    }
    
    return opt[amount]
}
