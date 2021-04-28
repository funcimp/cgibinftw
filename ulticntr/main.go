package main

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"
	"html/template"
)

var (
	//go:embed assets/html.go.html
	htmlTemplate string
)

var (
	//go:embed assets/images/starbg.gif
	starBG []byte
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
		fmt.Println(err)
	}
}

type img struct {
	data   []byte
	Width  uint
	Height uint
}

func (i img) URI() template.URL {
	encoded := base64.StdEncoding.EncodeToString(i.data)
	return template.URL(fmt.Sprintf("data:image/gif;base64,%v", encoded))
}

type assets struct {
	Background img
	Images     map[string]img
	Alien      img
	Disk       img
	Spidey     img
	Yinyang    img
	Peace      img
	Counter    uint64
}

func newAssets() assets {
	images := make(map[string]img)
	images["alien"] = img{data: alien, Width: 82, Height: 90}
	images["disk"] = img{data: disk, Width: 65, Height: 80}
	images["peace"] = img{data: peace, Width: 74, Height: 75}
	images["yinyang"] = img{data: yinyang, Width: 65, Height: 65}
	images["spidey"] = img{data: spidey, Width: 241, Height: 124}

	return assets{
		Images:     images,
		Background: img{data: starBG}}
}

func render() error {
	a := newAssets()
	counter, err := logVisit()
	if err != nil {
		return err
	}
	a.Counter = counter

	t, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if err := t.Execute(&b, a); err != nil {
		return err
	}
	fmt.Println(b.String())
	return nil
}
