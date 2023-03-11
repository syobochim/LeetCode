package roman;

func romanToInt2(s string) int {
    romanToIntMap := map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    result := 0
    prev := 0

    for i := len(s) - 1; i >= 0; i-- {
        current := romanToIntMap[s[i]]  // 文字を取り出してMapから数値を取り出す
        if current < prev {
            // IV など、prev より current の方が小さい場合は、 5 - 1 のように数値を引く
            result -= current
        } else {
            result += current  // 通常は出てきた文字に合致する数値を加算する
        }
        prev = current
    }
    return result
}