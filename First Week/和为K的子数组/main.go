func subarraySum(nums []int, k int) int {
    hash := make(map[int]int)
    add := 0
    count := 0
    hash[0] = 1
    for i := 0; i < len(nums); i++ {
        add += nums[i]
        if _, ok := hash[add - k]; ok {
            count += hash[add - k]
        }
        hash[add] += 1
    }
    return count
    // hash := make(map[int]int)
    // // 获得前缀和，存入哈希表
    // sum := 0
    // count := 0
    // hash[0] = 0
    // for i := 0; i < len(nums); i++ {
    //     sum += nums[i]
    //     hash[i+1] = sum
    // }
    // for i := 0; i < len(nums); i++ {
    //     for j := i + 1; j <= len(nums); j++ {
    //         if hash[j] - hash[i] == k {
    //             count ++
    //         }
    //     }
    // }
    // return count
}