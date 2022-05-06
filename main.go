package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Bukken struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Rent      string `json:"rent"`
	Commonfee string `json:"commonfee"`
	RoomCount int    `json:"roomCount"`
	Skcs      string `json:"skcs"`
	BukkenUrl string `json:"bukkenUrl"` // /chintai/kanto/kanagawa/40_0660.html
}

type Room struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Rent      string `json:"rent"`
	Commonfee string `json:"commonfee"`
	Status    string `json:"status"`
	Type      string `json:"type"`
}

func main() {
	fmt.Print("\n☆☆☆神奈川☆☆☆\n\n")
	search("14", "01")
	search("14", "02")
	search("14", "03")
	search("14", "04")
	search("14", "05")
	search("14", "06")

	fmt.Print("\n☆☆☆東京☆☆☆\n\n")
	search("13", "01")
	search("13", "02")
	search("13", "03")
	search("13", "04")
	search("13", "05")
	search("13", "06")

	fmt.Print("\n☆☆☆埼玉☆☆☆\n\n")
	search("11", "01")
	search("11", "02")
	search("11", "03")
	search("11", "04")

	fmt.Print("\n☆☆☆千葉☆☆☆\n\n")
	search("12", "01")
	search("12", "02")
	search("12", "03")
	search("12", "04")
	search("12", "05")
	search("12", "06")
	search("12", "07")
}

func search(tdfk string, area string) {
	bukkens := fetchBukken(tdfk, area)
	for _, bukken := range bukkens {
		if bukken.RoomCount == 0 {
			continue
		}
		fmt.Printf("%s | %s | %s \n", bukken.Name, bukken.Skcs, "https://www.ur-net.go.jp"+bukken.BukkenUrl)
		rooms := fetchRoom(tdfk, bukken)
		for _, room := range rooms {
			fmt.Printf("\t%s | %s | %s \n", room.Name, room.Rent+room.Commonfee, room.Type)
		}
	}

	time.Sleep(5 * time.Second)
}

func fetchBukken(tdfk string, area string) []Bukken {
	formData := url.Values{
		"rent_high": []string{"80000"},
		"rent_low":  []string{"40000"},
		"room":      []string{"1R,1K", "1DK", "1LDK"},
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

func fetchRoom(tdfk string, bukken Bukken) []Room {
	formData := url.Values{
		"rent_high": []string{"80000"},
		"rent_low":  []string{"40000"},
		"room":      []string{"1R,1K", "1DK", "1LDK"},
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
