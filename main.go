package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Bukken struct {
	Place     string `json:"place"` // 住所
	Danchi    string `json:"danchi"`
	DanchiNm  string `json:"danchiNm"`  // 団地名
	RoomCount string `json:"roomCount"` // 空室状況
	Rooms     []Room `json:"room"`
}

type Room struct {
	Id         string `json:"id"`
	RoomNmMain string `json:"roomNmMain"`
	RoomNmSub  string `json:"roomNmSub"`
	Rent       string `json:"rent"`      // 家賃
	Commonfee  string `json:"commonfee"` // 共益費
	Type       string `json:"type"`      // 間取り
}

type Bukken2 struct {
	Name      string `json:"name"`
	Rent      string `json:"rent"`
	Commonfee string `json:"commonfee"`
	RoomCount int    `json:"roomCount"`
	Skcs      string `json:"skcs"`

	BukkenUrl string `json:"bukkenUrl"` // /chintai/kanto/kanagawa/40_0660.html
}

func main() {
	var skcs []string

	fmt.Println("☆☆☆神奈川☆☆☆")
	skcs = []string{
		"104", // 横浜市中区
		"107", // 横浜市磯子区
		"108", // 横浜市金沢区
		"110", // 横浜市戸塚区
		"111", // 横浜市港南区
		"163", // 相模原市南区
		"161", // 相模原市緑区
		"162", // 相模原市中央区
	}
	search("14", skcs)

	fmt.Println("☆☆☆埼玉☆☆☆")
	skcs = []string{
		"201", // 川越市
		"208", // 所沢市
		"227", // 朝霞市
		"235", // 富士見市
		"214", // 春日部市
		"222", // 越谷市
		"221", // 草加市
		"237", // 三郷市
	}
	search("11", skcs)

	// fmt.Println("☆☆☆千葉☆☆☆")
	// skcs = []string{
	// 	"204", // 船橋市
	// 	"203", // 市川市
	// }
	// Search("12", skcs)

	fmt.Println("☆☆☆東京☆☆☆")
	skcs = []string{
		"107", // 墨田区
		"108", // 江東区
		"123", // 江戸川区
		"111", // 大田区
		"112", // 世田谷区
		"120", // 練馬区
		"120", // 練馬区
		"114", // 中野区
		"115", // 杉並区
	}
	search("12", skcs)

	// search2("14", "01")
	// search2("14", "02")
	// search2("14", "03")
	// search2("14", "04")
	// search2("14", "05")
	// search2("14", "06")
}

func search2(tdfk string, area string) {
	formData := url.Values{}
	formData.Add("room", "1R,1K")
	formData.Add("room", "1DK")
	formData.Add("room", "1LDK")
	formData.Add("tdfk", tdfk)
	formData.Add("area", area)
	formData.Add("rent_high", "80000")

	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/bukken/search/list_bukken/", formData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// レスポンス：物件情報
	var results []Bukken2
	// jsonをデコード
	json.NewDecoder(resp.Body).Decode(&results)

	for _, result := range results {
		if result.RoomCount == 0 {
			continue
		}
		fmt.Printf("%s | %d | %s | %s \n", result.Name, result.RoomCount, result.Rent, "https://www.ur-net.go.jp/"+result.BukkenUrl)
	}
}

func search(tdfk string, skcsArr []string) {
	v := url.Values{}
	v.Set("mode", "area")
	v.Set("block", "kanto")
	v.Add("tdfk", tdfk) // 都道府県
	// 市区町村
	for _, skcs := range skcsArr {
		v.Add("skcs", skcs)
	}

	// 間取りタイプ
	v.Add("room", "1R,1K")
	v.Add("room", "1K")
	v.Add("room", "1DK")
	v.Add("room", "1LDK")
	v.Add("room", "2DK")

	// 家賃
	// 共益費を含む
	v.Add("rent_high", "80000")
	v.Add("commonfee", "1")

	v.Add("orderByField", "1")
	v.Add("pageSize", "999")
	v.Add("pageIndex", "0")
	v.Add("pageIndexRoom", "0")
	// fmt.Println(v.Encode())

	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/bukken/result/bukken_result/", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// レスポンス：物件情報
	var bukkens []Bukken
	// jsonをデコード
	json.NewDecoder(resp.Body).Decode(&bukkens)

	for _, bukken := range bukkens {
		if bukken.RoomCount == "0" {
			continue
		}

		fmt.Println("----------------------------------------")
		fmt.Printf("%s  |  %s \n", bukken.DanchiNm, bukken.Place)

		for _, room := range bukken.Rooms {
			fmt.Printf("%10s %10s %10s %10s \n", room.RoomNmMain, room.RoomNmSub, room.Type, room.Rent)
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
}
