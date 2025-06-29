// 198. House Robber
package dynamicprogramming

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums), len(nums))
	dp[0] = nums[0]
	dp[1] = max(dp[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(dp)-1]
}

//runtime 0ms
//Memory 4.00MB Beats 76.39%
