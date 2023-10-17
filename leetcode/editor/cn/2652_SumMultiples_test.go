package lee

import (
	"testing"
)

//给你一个正整数 n ，请你计算在 [1，n] 范围内能被 3、5、7 整除的所有整数之和。
//
// 返回一个整数，用于表示给定范围内所有满足约束条件的数字之和。
//
//
//
// 示例 1：
//
// 输入：n = 7
//输出：21
//解释：在 [1, 7] 范围内能被 3、5、7 整除的所有整数分别是 3、5、6、7 。数字之和为 21 。
//
//
// 示例 2：
//
// 输入：n = 10
//输出：40
//解释：在 [1, 10] 范围内能被 3、5、7 整除的所有整数分别是 3、5、6、7、9、10 。数字之和为 40 。
//
//
// 示例 3：
//
// 输入：n = 9
//输出：30
//解释：在 [1, 9] 范围内能被 3、5、7 整除的所有整数分别是 3、5、6、7、9 。数字之和为 30 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10³
//
//
// Related Topics 数学 👍 48 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
//func sumOfMultiples(n int) int {
//	sum := 0
//	start := 1
//	for {
//		if start <= n {
//			if start%3 == 0 || start%5 == 0 || start%7 == 0 {
//				sum += start
//			}
//			start += 1
//		} else {
//			break
//		}
//	}
//	return sum
//}

func sumOfMultiples(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestSumMultiples(t *testing.T) {
	n := 9
	t.Log(sumOfMultiples(n))
}
