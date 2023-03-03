impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        for (i, num1) in nums.iter().enumerate() {
            for (j, num2) in nums.iter().enumerate() {
                if i == j {
                    continue
                }
                if num1 + num2 == target {
                    return vec![i as i32, j as i32]
                }
            }
        }
        vec![]
    }
}

struct Solution;

#[test]
fn test_two_sum() {
    assert_eq!(Solution::two_sum(vec![2,7,11,15], 9), [0,1]);
    assert_eq!(Solution::two_sum(vec![3,2,4], 6), [1,2]);
    assert_eq!(Solution::two_sum(vec![3,3], 6), [0,1]);
}