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
	render()
}

type img struct {
	data  []byte
	alt   string
	class string
}

func (i img) URI() template.URL {
	encoded := base64.StdEncoding.EncodeToString(i.data)
	return template.URL(fmt.Sprintf("data:image/gif;base64,%v", encoded))
}

type assets struct {
	Background img
	Alien      img
	Disk       img
	Spidey     img
	Yinyang    img
	Peace      img
	Counter    uint64
}

func render() {
	a := assets{
		Background: img{
			data: starBG,
		},
		Alien: img{
			data:  alien,
			class: "alien",
		},
		Disk: img{
			data:  disk,
			class: "disk",
		},
		Spidey: img{
			data:  spidey,
			class: "spidey",
		},
		Yinyang: img{
			data:  yinyang,
			class: "yinyang",
		},
		Peace: img{
			data:  peace,
			class: "peace",
		},
		Counter: getCount(),
	}

	t, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
		return
	}
	var b bytes.Buffer
	if err := t.Execute(&b, a); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b.String())
	fmt.Println(getTable())
}
