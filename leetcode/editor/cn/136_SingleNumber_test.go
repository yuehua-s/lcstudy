package lee

import (
	"testing"
)

//给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
//
//
//
//
//
//
//
// 示例 1 ：
//
//
//输入：nums = [2,2,1]
//输出：1
//
//
// 示例 2 ：
//
//
//输入：nums = [4,1,2,1,2]
//输出：4
//
//
// 示例 3 ：
//
//
//输入：nums = [1]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
// 除了某个元素只出现一次以外，其余每个元素均出现两次。
//
//
// Related Topics 位运算 数组 👍 3030 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//func singleNumber(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	if len(nums) == 1 {
//		return nums[0]
//	}
//	m := make(map[int]int)
//	for _, n := range nums {
//		if _, ok := m[n]; !ok {
//			m[n] = 1
//		} else {
//			m[n] += 1
//		}
//	}
//	for k, v := range m {
//		if v == 1 {
//			return k
//		}
//	}
//	return 0
//}

func singleNumber(nums []int) int {
	single := 0
	for _, n := range nums {
		single ^= n
	}
	return single
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSingleNumber(t *testing.T) {
	//nums := []int{2, 2, 1}
	nums := []int{4, 1, 2, 1, 2}
	t.Log(singleNumber(nums))
}
