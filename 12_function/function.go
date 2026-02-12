package main

// func add(a int, b int) int {
// 	return a + b
// }
// func main() {
// 	result := add(3, 5)
// 	fmt.Println(result)

// }

//
// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }
// func main() {
// 	// lang1, lang2, lang3 := getLanguages()
// 	// fmt.Println(lang1, lang2, lang3)

// 	//or
// 	lang1, lang2, _ := getLanguages() //_ we passed it if value unused
// 	fmt.Println(lang1, lang2)
// }

//
// func processIt(fn func(a int) int) {
// 	fn(1)
// }

//
func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	// fn := func(a int) int {
	// 	return 2
	// }
	fn := processIt()
	fn(6)
}
