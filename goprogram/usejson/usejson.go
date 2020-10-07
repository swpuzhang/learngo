package usejson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type movie struct {
	Title  string
	Year   int
	Actors []string
}
type movie2 struct {
	Title  string
	Year   int      `json:"released,omitempty"` //omitempty 如果为0值, 不解析该字段
	Actors []string `json:"mans,omitempty"`     //逗号后面不能油空格,否则会报错
}

// UseJSON 使用json
func UseJSON() {
	movies := []movie{
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title: "taijiong",
			Year:  0,
		},
	}
	encoded, _ := json.Marshal(movies)
	fmt.Printf("%s\n", encoded)
	encoded, _ = json.MarshalIndent(movies, "", "  ") //第二个参数:每一行的开头设置字符串, 第三个参数:缩进
	fmt.Printf("%s\n", encoded)
	movies2 := []movie2{
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title: "taijiong",
			Year:  0,
		},
		{
			Title:  "taijiong2",
			Year:   0,
			Actors: []string{},
		},
	}
	encoded, _ = json.MarshalIndent(movies2, "", "  ") //第二个参数:每一行的开头设置字符串, 第三个参数:缩进
	fmt.Printf("%s\n", encoded)
	var decodedMovies2 []movie2
	err := json.Unmarshal(encoded, &decodedMovies2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", decodedMovies2)

	var decodedMovies []movie
	err = json.Unmarshal(encoded, &decodedMovies)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", decodedMovies)
	useDecoderEncoder()
}

func useDecoderEncoder() {
	movies2 := []movie2{
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title:  "taijiong",
			Year:   2015,
			Actors: []string{"徐峥", "王宝强", "黄渤"},
		},
		{
			Title: "taijiong",
			Year:  0,
		},
		{
			Title:  "taijiong2",
			Year:   0,
			Actors: []string{},
		},
	}
	fmt.Println("-------------------")
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(movies2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.String())
	fmt.Println("------------------")
	reader := strings.NewReader(buf.String())
	var movs []movie2
	if err := json.NewDecoder(reader).Decode(&movs); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", movs)
}
