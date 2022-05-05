package main

// func SearchBukkenRoom() {
// 	v := url.Values{}
// 	v.Set("mode", "area")
// 	v.Add("skcs", "104")
// 	v.Add("skcs", "106")
// 	v.Add("block", "kanto")
// 	v.Add("tdfk", "14")
// 	v.Add("orderByField", "1")
// 	v.Add("pageSize", "10")
// 	v.Add("pageIndex", "0")
// 	v.Add("shisya", "40")
// 	v.Add("danchi", "142") // 団地ID
// 	v.Add("shikibetu", "0")
// 	v.Add("pageIndexRoom", "0")

// 	fmt.Println(v.Encode())

// 	resp, err := http.PostForm("https://chintai.sumai.ur-net.go.jp/chintai/api/bukken/result/bukken_result_room/", v)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	var rooms []RoomItem

// 	json.NewDecoder(resp.Body).Decode(&rooms)

// 	// data, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// json.Unmarshal(data, &rooms)

// 	fmt.Println(rooms)
// }
