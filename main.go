package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type BukkenItem struct {
	Place     string     `json:"place"` // 住所
	Danchi    string     `json:"danchi"`
	DanchiNm  string     `json:"danchiNm"`  // 団地名
	RoomCount string     `json:"roomCount"` // 空室状況
	Rooms     []RoomItem `json:"room"`
}

type RoomItem struct {
	Id         string `json:"id"`
	RoomNmMain string `json:"roomNmMain"`
	RoomNmSub  string `json:"roomNmSub"`
	Rent       string `json:"rent"`      // 家賃
	Commonfee  string `json:"commonfee"` // 共益費
	Type       string `json:"type"`      // 間取り
}

func main() {
	// 104 横浜市中区
	// 107 横浜市磯子区
	// 163 相模原市南区
	// 161 相模原市緑区
	// 162 相模原市中央区
	Search("14", []string{"104", "107", "163", "162", "161"})
}

func Search(tdfk string, skcsArr []string) {
	v := url.Values{}
	v.Set("mode", "area")
	v.Add("block", "kanto")
	v.Add("tdfk", tdfk) // 都道府県
	// 市区町村
	for _, skcs := range skcsArr {
		v.Add("skcs", skcs)
	}
	// 間取りタイプ
	v.Add("room", "1K")
	v.Add("room", "1DK")
	v.Add("room", "1LDK")
	// v.Add("room", "2DK")

	v.Add("orderByField", "1")
	v.Add("pageSize", "10")
	v.Add("pageIndex", "0")
	v.Add("pageIndexRoom", "0")
	fmt.Println(v.Encode())

	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/bukken/result/bukken_result/", v)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// レスポンス：物件情報
	var bukkenArray []BukkenItem
	// jsonをデコード
	json.NewDecoder(resp.Body).Decode(&bukkenArray)

	for _, bukkenItem := range bukkenArray {
		// if bukkenItem.RoomCount == "0" {
		// 	continue
		// }
		fmt.Println("--------------------")
		fmt.Println(bukkenItem.DanchiNm)
		fmt.Println(bukkenItem.Place)

		for _, roomItem := range bukkenItem.Rooms {
			fmt.Printf("%10s %10s %10s %10s \n", roomItem.RoomNmMain, roomItem.RoomNmSub, roomItem.Type, roomItem.Rent)
		}
	}
}
