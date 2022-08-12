package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Bukken struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Rent      string `json:"rent"`
	Commonfee string `json:"commonfee"`
	RoomCount int    `json:"roomCount"`
	Skcs      string `json:"skcs"`      // 市区町村
	BukkenUrl string `json:"bukkenUrl"` // 例：/chintai/kanto/kanagawa/40_0660.html
}

type Room struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Rent      string `json:"rent"`
	Commonfee string `json:"commonfee"`
	Status    string `json:"status"`
	Type      string `json:"type"`
}

var rentLow int
var rentHigh int
var roomType string

func main() {
	flag.IntVar(&rentLow, "rent-log", 30000, "家賃FROM。30000")
	flag.IntVar(&rentHigh, "rent-high", 80000, "家賃TO。80000")
	flag.StringVar(&roomType, "room-type", "1R,1K|1DK|1LDK", "間取り。1R,1K|1DK|1LDK")
	flag.Parse()

	fmt.Print("\n☆☆☆ 神奈川 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("14", "01")
	search("14", "02")
	search("14", "03")
	search("14", "04")
	search("14", "05")
	search("14", "06")

	fmt.Print("\n☆☆☆ 東京 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("13", "01")
	search("13", "02")
	search("13", "03")
	search("13", "04")
	search("13", "05")
	search("13", "06")

	fmt.Print("\n☆☆☆ 埼玉 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("11", "01")
	search("11", "02")
	search("11", "03")
	search("11", "04")

	fmt.Print("\n☆☆☆ 千葉 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("12", "01")
	search("12", "02")
	search("12", "03")
	search("12", "04")
	search("12", "05")
	search("12", "06")
	search("12", "07")

	fmt.Print("\n☆☆☆ 京都 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("26", "01")

	fmt.Print("\n☆☆☆ 大阪 ☆☆☆\n\n")
	time.Sleep(5 * time.Second)
	search("27", "01")
	search("27", "02")
}

// tdkf: 都道府県
func search(tdfk string, area string) {
	bukkens := getBukken(tdfk, area)
	for _, bukken := range bukkens {
		if bukken.RoomCount == 0 {
			continue
		}
		fmt.Printf("%s | %s | %s \n", bukken.Name, bukken.Skcs, "https://www.ur-net.go.jp"+bukken.BukkenUrl)
		rooms := getRoom(tdfk, bukken)
		for _, room := range rooms {
			fmt.Printf("\t%s | %s | %s \n", room.Name, room.Rent+room.Commonfee, room.Type)
		}
	}
}

func getBukken(tdfk string, area string) []Bukken {
	formData := url.Values{
		"rent_high": []string{strconv.Itoa(rentHigh)},
		"rent_low":  []string{strconv.Itoa(rentLow)},
		"room":      strings.Split(roomType, "|"),
		"tdfk":      []string{tdfk},
		"area":      []string{area},
	}

	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/bukken/search/list_bukken/", formData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var bukkens []Bukken
	json.NewDecoder(resp.Body).Decode(&bukkens)

	return bukkens
}

func getRoom(tdfk string, bukken Bukken) []Room {
	formData := url.Values{
		"rent_high": []string{strconv.Itoa(rentHigh)},
		"rent_low":  []string{strconv.Itoa(rentLow)},
		"room":      strings.Split(roomType, "|"),
		"tdfk":      []string{tdfk},
		"mode":      []string{"init"},
		"id":        []string{bukken.Id},
	}
	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/room/list/", formData)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var rooms []Room
	json.NewDecoder(resp.Body).Decode(&rooms)

	return rooms
}
