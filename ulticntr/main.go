package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"
)

var (
	//go:embed assets/html.go.html
	htmlTemplate string
	//go:embed assets/images/starbg.gif
	starBG []byte
	//go:embed assets/images/fireworksbg.gif
	fireworksBG []byte
	//go:embed assets/images/alien.gif
	alien []byte
	//go:embed assets/images/disk.gif
	disk []byte
	//go:embed assets/images/spidey.gif
	spidey []byte
	//go:embed assets/images/yinyang.gif
	yinyang []byte
	//go:embed assets/images/peace.gif
	peace []byte
)

func main() {
	if err := render(); err != nil {
		log.Println(err)
	}
}

type img struct {
	data   []byte
	Class  string
	Width  uint
	Height uint
}

func (i img) URI() template.URL {
	uri := fmt.Sprintf("data:image/gif;base64,%v", base64.StdEncoding.EncodeToString(i.data))
	return template.URL(uri)
}

func (a assets) GetClassName(i int) string {
	return a.Images[a.Arrangement[i]].Class
}

type assets struct {
	Background  img
	Images      []img
	Arrangement []uint
	Counter     uint64
}

type option func(*assets)

func withFireworks() option {
	return func(a *assets) {
		a.Background = img{data: fireworksBG}
	}
}

func withRandomImages() option {
	return func(a *assets) {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 4
		arr := []uint{0}
		for i := 1; i <= 12; i++ {
			arr = append(arr, uint(rand.Intn(max-min+1)+min))
		}
		a.Arrangement = arr
	}
}

func newAssets(options ...option) assets {
	images := []img{
		{data: spidey, Class: "spidey", Width: 241, Height: 124},
		{data: alien, Class: "alien", Width: 82, Height: 90},
		{data: disk, Class: "disk", Width: 65, Height: 80},
		{data: peace, Class: "peace", Width: 74, Height: 75},
		{data: yinyang, Class: "yinyang", Width: 65, Height: 65},
	}
	arr := []uint{0, 1, 3, 2, 1, 4, 1, 1, 4, 1, 2, 3, 1}
	a := assets{
		Images:      images,
		Arrangement: arr,
		Background:  img{data: starBG},
	}
	for _, opt := range options {
		opt(&a)
	}
	return a
}

func parseOpts(queryString string) (o []option) {
	values, _ := url.ParseQuery(queryString)
	if values.Get("random") == "true" {
		o = append(o, withRandomImages())
	}
	if values.Get("bg") == "fireworks" {
		o = append(o, withFireworks())
	}
	return o
}

func render() error {
	t := template.Must(template.New("page").Parse(htmlTemplate))
	a := newAssets(parseOpts(os.Getenv("QUERY_STRING"))...)
	counter, err := logVisit()
	if err != nil {
		return err
	}
	a.Counter = counter
	var b bytes.Buffer
	if err := t.Execute(&b, a); err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}
