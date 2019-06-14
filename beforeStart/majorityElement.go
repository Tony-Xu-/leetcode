func majorityElement(nums []int) int {
    countMap := map[int]int{}
    for _, n := range nums {
        if _, ok := countMap[n]; !ok {
            countMap[n]=0
        }
        countMap[n]=countMap[n]+1
        if countMap[n]>len(nums)/2 {
            return n
        }
    }
    return 0
}
