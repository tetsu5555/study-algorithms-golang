package quick

func QuickSort(values []int) (ret []int) {
	// 要素数が1以下の配列はそれ以上ソートする必要がない
	if len(values) < 2 {
		return values
	}

	// 配列の先頭をピボットに選ぶ
	pivot := values[0]

	// ピボットを基準にしてあたいの代償で配列を二つに分割する
	left := []int{}
	right := []int{}

	for _, v := range values[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	// 分割した配列をそれぞれ再帰的にソートする
	left = QuickSort(left)
	right = QuickSort(right)

	// left(小さい) + pivot(基準値) + right(大きい)の順番で配列を組み立てる
	// ここが実際のソート処理
	ret = append(left, pivot)
	ret = append(ret, right...)

	return
}
