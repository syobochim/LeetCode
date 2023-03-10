package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 { // マイナスの数字は回文にならない
		return false
	}
	if x%10 == 0 && x != 0 { // 最後が 0 で終わる 0 以外の数字は回文にならない
		return false
	}
	revertedNum := 0
	for x > revertedNum {
		// 元の数字を10で割った余りを revertNum として登録する
		// その後にxを10で割ることで、1の位を消す
		// 12321の場合
		// 1週目 : x : 1232 , revertNum : 1
		// 2週目 : x : 123 , revertNum : 12
		// 2週目 : x : 12 , revertNum : 123
		// → revertNum が x よりも大きくなったら半分を過ぎたとして処理をストップ
		revertedNum = revertedNum*10 + x%10
		x /= 10
	}
	// x と revertNum を比較する。 奇数桁の場合は revertNum の一の位を切り捨てて比較し、合っていればOK
	return x == revertedNum || x == revertedNum/10
}
