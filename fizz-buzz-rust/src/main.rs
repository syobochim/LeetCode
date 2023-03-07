impl Solution {
    pub fn fizz_buzz(n: i32) -> Vec<String> {
        let mut result: Vec<String> = vec![];
        for i in 1..n+1 {
            if i % 15 == 0 {
                result.push("FizzBuzz".to_string());
            } else if i % 3 == 0 {
                result.push("Fizz".to_string());
            } else if i % 5 == 0 {
                result.push("Buzz".to_string());
            } else {
                result.push(i.to_string());
            }
        }
        return result;
    }
}

struct Solution;

#[test]
fn test_fizz_buzz() {
    assert_eq!(Solution::fizz_buzz(3), ["1","2","Fizz"]);
    assert_eq!(Solution::fizz_buzz(5), ["1","2","Fizz","4","Buzz"]);
    assert_eq!(Solution::fizz_buzz(15), ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"])
}