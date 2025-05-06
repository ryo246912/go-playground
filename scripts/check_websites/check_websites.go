package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// 新しいタイプのresultは、WebsiteCheckerの戻り値をチェック対象のURLに関連付けるために作成されました
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// for _, url := range urls {
	// 	results[url] = wc(url)
	// }

	// Goに新しいgoroutineを開始するよう指示するには、キーワードの前にキーワードgoを置くことで、関数呼び出しをgoステートメントに変換します
	// 「go doSomething()」。
	// 匿名関数には、それらを便利にするいくつかの機能があり、そのうち2つは上記で使用しています。
	// まず、宣言と同時に実行できます-これは、無名関数の最後にある()が行っていることです。
	// 次に、定義されている字句スコープへのアクセスを維持します

	// for _, url := range urls {
	// 	// 変数urlがforループの反復ごとに再利用されることです。 毎回urlsから新しい値を取得します。
	// 	// しかし、それぞれのゴルーチンはurl変数への参照を持っています。それらは独自の独立したコピーを持っていません。
	// 	// go func() {
	// 	// 	results[url] = wc(url)
	// 	// }()
	// 	go func(u string) {
	// 		results[u] = wc(u)
	// 	}(url)
	// }

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {

		// 高速化したいコードの部分を並列化し、同時に実行できない部分は線形的に発生するようにしました。
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
