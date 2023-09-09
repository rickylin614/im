package errs

import "fmt"

func ExampleNew() {
	err := New("未知錯誤")
	wErr := fmt.Errorf("封裝一層錯誤 %w", err)
	fmt.Println(wErr)
	fmt.Println(ParseErr(wErr))
	fmt.Println(err.IsErr(wErr))

	// output:
	// 封裝一層錯誤 00-000: 未知錯誤
	// 00-000: 未知錯誤 true
	// true
}
