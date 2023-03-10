impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 {
            return false
        }
        if x % 10 == 0 && x != 0 {
            return false
        }
        let mut num = x;
        let mut reverted_num: i32 = 0;
        while num > reverted_num {
            reverted_num = reverted_num * 10 + num % 10;
            num = num / 10;
        }
        num == reverted_num || num == reverted_num / 10
    }
}

struct Solution;

#[test]
fn test_is_palindrome_number_1221() {
    assert_eq!(Solution::is_palindrome(1221), true);
}

#[test]
fn test_is_palindrome_number_12321() {
    assert_eq!(Solution::is_palindrome(12321), true);
}

#[test]
fn test_is_palindrome_number_123() {
    assert_eq!(Solution::is_palindrome(123), false);
}

#[test]
fn test_is_palindrome_number_12345() {
    assert_eq!(Solution::is_palindrome(12345), false);
}

#[test]
fn test_is_palindrome_number_10() {
    assert_eq!(Solution::is_palindrome(10), false);
}
