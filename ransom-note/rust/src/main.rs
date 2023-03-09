impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut magazine_letters: [i32; 26] = [0;26];
        for char in magazine.chars() {
            let code = char as u8 - 'a' as u8;
            magazine_letters[code as usize] += 1;
        }

        for char in ransom_note.chars() {
            let code = char as u8 - 'a' as u8;
            if magazine_letters[code as usize] == 0 {
                return false;
            }
            magazine_letters[code as usize] -= 1;
        }
        true
    }
}

struct Solution;

#[test]
fn test_true_case1() {
    let ransom_note = "aa".to_string();
    let magazine = "aab".to_string();
    assert!(Solution::can_construct(ransom_note, magazine));
}

#[test]
fn test_true_case2() {
    let ransom_note = "aab".to_string();
    let magazine = "baa".to_string();
    assert!(Solution::can_construct(ransom_note, magazine));
}

#[test]
fn test_false_case1() {
    let ransom_note = "a".to_string();
    let magazine = "b".to_string();
    assert!(!Solution::can_construct(ransom_note, magazine));
}

#[test]
fn test_false_case2() {
    let ransom_note = "aa".to_string();
    let magazine = "ab".to_string();
    assert!(!Solution::can_construct(ransom_note, magazine));
}