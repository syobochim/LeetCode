package ransomNote

func canConstruct2(ransomNote string, magazine string) bool {
    // アルファベット格納用の26文字の配列を作る
    // アルファベットそれぞれに対して、使われている個数を保持し、利用するごとに引いていく
    magazineLetters := [26]int{0}
	for _, char := range magazine {
        // magazineに利用されているアルファベットに対して、カウントアップしていく
        magazineLetters[char - 'a']++  // a は 97 の char になるので、その分を引いて 0 始まりの数値に変換する
    }

    for _, char := range ransomNote {
        // 配列入っているアルファベットのストックを確認し、ストックがない場合はfalseを返す
        if magazineLetters[char - 'a'] == 0 {
            return false
        }
        // ストックが見つかったら１つ消費する
        magazineLetters[char - 'a']--
    }
    return true
}
