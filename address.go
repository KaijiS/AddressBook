package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

// アドレスの構造
type Address_Person struct {
	Name string
	Mail string
	Tel  string
}

type AdddressList []Address_Person

// 以下3つソート用のインタフェース
func (p AdddressList) Len() int {
	return len(p)
}

func (p AdddressList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p AdddressList) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

// 文字列を1行入力
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

// 編集
func Edit(x *Address_Person) {
	var corNumber string
	for {
		fmt.Println("何を修正しますか")
		fmt.Println("1: 氏名  2: メールアドレス  3: 電話番号 4: 全て")
		fmt.Print("--->")
		corNumber = StrStdin()
		fmt.Println()
		switch corNumber {
		case "1":
			fmt.Println("登録者の氏名を入力してください")
			fmt.Print("--->")
			x.Name = strings.Title(StrStdin())
			fmt.Println()
		case "2":
			fmt.Println("登録者のメールアドレスを登録してください")
			fmt.Print("--->")
			x.Mail = StrStdin()
			fmt.Println()
		case "3":
			fmt.Println("登録者の電話番号を登録してください")
			fmt.Print("--->")
			x.Tel = StrStdin()
			fmt.Println()
		case "4":
			regUser_AllData(x)
		default:
			fmt.Println("正しい値を入力してください")
			fmt.Println()
		}
		if corNumber == "1" || corNumber == "2" || corNumber == "3" || corNumber == "4" {
			break
		}
	}
	return
}

func regUser_AllData(x *Address_Person) {
	fmt.Println("登録者の氏名を入力してください")
	fmt.Print("--->")
	x.Name = strings.Title(StrStdin())
	fmt.Println("登録者のメールアドレスを登録してください")
	fmt.Print("--->")
	x.Mail = StrStdin()
	fmt.Println("登録者の電話番号を登録してください")
	fmt.Print("--->")
	x.Tel = StrStdin()
	fmt.Println()
	return
}

func Conf(x *Address_Person) {
	fmt.Println("以下の情報でよろしいですか?")
	fmt.Printf("氏名:%s\n", x.Name)
	fmt.Printf("メールアドレス:%s\n", x.Mail)
	fmt.Printf("電話番号:%s\n", x.Tel)
	return
}

// ユーザ登録
func regUser(adddress_list AdddressList) AdddressList {
	var user Address_Person
	// ユーザ情報入力
	regUser_AllData(&user)
	for {
		var flag string
		// 入力情報の確認
		Conf(&user)
		fmt.Println("1: 修正  0: 確定")
		fmt.Print("--->")
		flag = StrStdin()
		fmt.Println()
		if flag == "1" {
			// 修正作業
			Edit(&user)
		} else if flag == "0" {
			// 配列に追加
			adddress_list = append(adddress_list, user)
			// 並び替え
			sort.Sort(adddress_list)
		} else {
			fmt.Println("正しい数値を入力してください")
			fmt.Println()
		}

		if flag == "0" {
			break
		}
	}
	return adddress_list
}

// 削除
func remove(adddress_list AdddressList, rmIdx int) AdddressList {
	for {
		fmt.Println("削除してよろしいですか?")
		fmt.Println("1: 削除  2: 戻る")
		fmt.Print("--->")
		removeFlag := StrStdin()
		switch removeFlag {
		case "1":
			var rmAfter AdddressList
			for idx, elem := range adddress_list {
				if idx != rmIdx {
					rmAfter = append(rmAfter, elem)
				}
			}
			return rmAfter
		case "2":
			return adddress_list
		default:
			fmt.Println("正しい数値を入力してくだい")
			fmt.Println()
		}
	}

}

func showAddress(adddress_list AdddressList) AdddressList {
	for idx, adddressData := range adddress_list {
		fmt.Print("No:")
		fmt.Println(idx + 1)
		fmt.Print("氏名: ")
		fmt.Println(adddressData.Name)
		fmt.Print("メールアドレス: ")
		fmt.Println(adddressData.Mail)
		fmt.Print("電話番号: ")
		fmt.Println(adddressData.Tel)
		fmt.Println()
	}
	for {
		fmt.Println("操作を選んでください")
		fmt.Println("1: 編集  2: 削除 3: 更新して表示  4: 戻る")
		fmt.Print("--->")
		ope := StrStdin()
		fmt.Println()

		switch ope {
		case "1":
			fmt.Println("編集する\"No\"を選んでください")
			fmt.Print("--->")
			selectNum := StrStdin()
			i, _ := strconv.Atoi(selectNum)
			Edit(&(adddress_list[i-1]))
			fmt.Println()
			Conf(&(adddress_list[i-1]))

			for {
				fmt.Println("1: 修正  0: 確定")
				fmt.Print("--->")
				flag := StrStdin()
				fmt.Println()
				switch flag {
				case "1":
					Edit(&(adddress_list[i-1]))
				case "0":
					// 並び替え
					sort.Sort(adddress_list)
				default:
					fmt.Println("正しい数値を入力してください")
					fmt.Println()
				}
				if flag == "1" || flag == "0" {
					break
				}
			}

		case "2":
			fmt.Println("削除する\"No\"を選んでください")
			fmt.Print("--->")
			selectNum := StrStdin()
			i, _ := strconv.Atoi(selectNum)
			adddress_list = remove(adddress_list, i-1)
			fmt.Println()

		case "3":
			adddress_list = showAddress(adddress_list)
			return adddress_list

		case "4":
			return adddress_list
		default:
			fmt.Println("正しい数値を入力してください")
			fmt.Println()
		}
	}
}

func Search(adddress_list AdddressList) AdddressList {
	for {
		fmt.Println("検索項目を選んでください")
		fmt.Println("1: 氏名  2: メールアドレス  3: 電話番号  4: 戻る")
		fmt.Print("--->")
		categ := StrStdin()
		fmt.Println()

		switch categ {
		case "1":
			fmt.Println("検索ワードを入力してください")
			fmt.Print("--->")
			word := StrStdin()
			fmt.Println()

			fmt.Println("\"氏名の検索結果\"")
			fmt.Println()
			counter := 0
			for idx, adddressData := range adddress_list {
				if strings.Index(strings.ToLower(adddressData.Name), strings.ToLower(word)) != -1 {
					fmt.Print("No:")
					fmt.Println(idx + 1)
					fmt.Print("氏名: ")
					fmt.Println(adddressData.Name)
					fmt.Print("メールアドレス: ")
					fmt.Println(adddressData.Mail)
					fmt.Print("電話番号: ")
					fmt.Println(adddressData.Tel)
					fmt.Println()
					counter++
				}
			}
			if counter == 0 {
				fmt.Println("該当するデータはありません")
				fmt.Println()
			}

		case "2":
			fmt.Println("検索ワードを入力してください")
			fmt.Print("--->")
			word := StrStdin()
			fmt.Println()

			fmt.Println("\"メールアドレスの検索結果\"")
			fmt.Println()
			counter := 0
			for idx, adddressData := range adddress_list {
				if strings.Index(strings.ToLower(adddressData.Mail), strings.ToLower(word)) != -1 {
					fmt.Print("No:")
					fmt.Println(idx + 1)
					fmt.Print("氏名: ")
					fmt.Println(adddressData.Name)
					fmt.Print("メールアドレス: ")
					fmt.Println(adddressData.Mail)
					fmt.Print("電話番号: ")
					fmt.Println(adddressData.Tel)
					fmt.Println()
					counter++
				}
			}
			if counter == 0 {
				fmt.Println("該当するデータはありません")
				fmt.Println()
			}

		case "3":
			fmt.Println("検索ワードを入力してください")
			fmt.Print("--->")
			word := StrStdin()
			fmt.Println()

			fmt.Println("\"電話番号の検索結果\"")
			fmt.Println()
			counter := 0
			for idx, adddressData := range adddress_list {
				if strings.Index(adddressData.Tel, word) != -1 {
					fmt.Print("No:")
					fmt.Println(idx + 1)
					fmt.Print("氏名: ")
					fmt.Println(adddressData.Name)
					fmt.Print("メールアドレス: ")
					fmt.Println(adddressData.Mail)
					fmt.Print("電話番号: ")
					fmt.Println(adddressData.Tel)
					fmt.Println()
					counter++
				}
			}
			if counter == 0 {
				fmt.Println("該当するデータはありません")
				fmt.Println()
			}

		case "4":
			fmt.Println()
			break

		default:
			fmt.Println("正しい数値を入力してください")
			fmt.Println()
		}
		if categ == "1" || categ == "2" || categ == "3" || categ == "4" {
			break
		}
	}
	return adddress_list
}

func main() {
	var adddress_list AdddressList

	file, err := os.OpenFile("address.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if err := gocsv.UnmarshalFile(file, &adddress_list); err != nil { // Load clients from file
		panic(err)
	}

	for {
		fmt.Println("何を行いますか?")
		fmt.Println("1: 登録データ表示  2: データ登録  3: 検索  4: 終了")
		fmt.Print("--->")
		processNumber := StrStdin()
		fmt.Println()
		switch processNumber {
		case "1":
			if len(adddress_list) == 0 {
				fmt.Println("登録データがありません")
				fmt.Println()
			}
			adddress_list = showAddress(adddress_list)

		case "2":
			if len(adddress_list) < 20 {
				adddress_list = regUser(adddress_list)
			} else {
				fmt.Println("これ以上登録できません")
				fmt.Println()
			}

		case "3":
			adddress_list = Search(adddress_list)

		case "4":
			file, _ := os.OpenFile("address.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
			defer file.Close()
			gocsv.MarshalFile(&adddress_list, file)
			os.Exit(0)

		default:
			fmt.Println("正しい数値を入力してください")
			fmt.Println()
		}
	}
}
