package commands

import (
	"github.com/crcms/blog/web/internal/domain/document/models"
	"github.com/jinzhu/gorm"
	"time"

	//"github.com/crcms/blog/web/internal/domain/document/models"
	//"github.com/firmeve/firmeve/database"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/support/path"
	//"log"

	//"github.com/jinzhu/gorm"

	//"github.com/jinzhu/gorm"
	//"github.com/russross/blackfriday"
	"regexp"
	"strings"

	//"github.com/gomarkdown/markdown/html"

	//"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/html"
	"io/ioutil"
	//"strings"

	//"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/html"

	//"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/html"

	//"github.com/gomarkdown/markdown"
	//"github.com/gomarkdown/markdown/html"
	"github.com/spf13/cobra"
)

type MarkdownCommand struct {
}

func (m MarkdownCommand) CobraCmd() *cobra.Command {
	command := new(cobra.Command)
	command.Use = "markdown"
	command.Short = "Parse all markdown file"
	command.Flags().StringP("path", "", path.RunRelative("../../../../sources"), "file path")

	return command
}

func (m MarkdownCommand) Run(root contract.BaseCommand, cmd *cobra.Command, args []string) {
	logger := root.Application().Resolve(`logger`).(contract.Loggable)
	path := cmd.Flag(`path`).Value.String()

	if path == `` {
		logger.Error("path error")
	}

	db := root.Resolve(`db.connection`).(*gorm.DB)
	//fmt.Printf("%#v",db)
	//db.AutoMigrate(&models.Document{})
	//
	//d := &models.Document{
	//	Uuid:    "fdsa",
	//	Title:   "fdsafdsafdsafdsa",
	//	Content: "fdsa",
	//}
	//
	//db.Create(d)
	//
	//v := new(models.Document)
	//e := db.Where("title=?","fdsafdsafdsafdsa").Find(v).Error
	//if e != nil {
	//	log.Println(e.Error())
	//}
	//fmt.Println("=============")
	//fmt.Println(v.Title)
	//
	////db.
	//
	//fmt.Println("=============")
	////
	if err := getAllFile(db,path); err != nil {
		logger.Error("%#v", err)
	}

	// 先解析一个单文件
	//md := []byte("## markdown document")
	//output := markdown.ToHTML(md, nil, nil)
	//fmt.Println(cmd.Flag(`path`).Value.String())
	//fmt.Printf("%s", "output")
}

// return (ast.GoToNext, true) to tell html renderer to skip rendering this node
// (because you've rendered it)
//func renderHookDropCodeBlock(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
//	fmt.Printf("%#v",node.)
//	// skip all nodes that are not CodeBlock nodes
//	if _, ok := node.(*ast.CodeBlock); !ok {
//		return ast.GoToNext, false
//	}
//	// custom rendering logic for ast.CodeBlock. By doing nothing it won't be
//	// present in the output
//	return ast.GoToNext, true
//}

//md := "test\n```\nthis code block will be dropped from output\n```\ntext"
//html := markdown.ToHTML([]byte(s), nil, renderer)

func getAllFile(db *gorm.DB,filepath string) error {
	// 遍历所有目录和文件
	resources, err := ioutil.ReadDir(filepath)
	if err != nil {
		return err
	}

	for _, file := range resources {
		if file.IsDir() {
			//getAllFile()
		} else {
			data := parseFile(filepath + `/` + file.Name())

			model := &models.Document{
				Title:     data[`title`][0],
				Content:   data[`content`][0],
				Description:data[`description`][0],
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			db.Save(model)


			//bytes, _ := ioutil.ReadFile(filefullpath)
			//rec,_ := regexp.Compile("\\-\\-\\-([^\\-]*)\\-\\-\\-")
			//v := rec.FindStringSubmatch(string(bytes))
			////fmt.Printf("%v",rec.FindStringSubmatch(string(bytes)))
			//fmt.Println(strings.Split(strings.Trim(v[1],"\n"),"\n"))
			//fmt.Println("=================")
			//
			////fmt.Println(filefullpath)
			//output := blackfriday.Run(bytes)
			//markdown := blackfriday.New()
			//node := markdown.Parse(bytes)
			//
			//
			//fmt.Printf("%#v",node.HeadingID)
			//fmt.Printf("%s",output)
			//print
			//opts := html.RendererOptions{
			//	Flags: html.CommonFlags,
			//	//RenderNodeHook: renderHookDropCodeBlock,
			//}
			//renderer := html.NewRenderer(opts)
			//
			//html := markdown.ToHTML(bytes,nil,renderer)

			//
			//<!DOCTYPE html>
			//<html lang="en">
			//<head>
			//<meta charset="UTF-8">
			//<title>Title</title>
			//</head>
			//<body>
			//
			//</body>
			//</html>
			//if err == nil {
			//	//fileinfo,_ := os.Stat(filefullpath)
			//	filename := strings.Replace(file.Name(),".md",".html",1)
			//	ioutil.WriteFile(path.RunRelative("../../../public/html/"+filename), []byte(html), file.Mode())
			//}
			//fmt.Println(file.Name())
		}
	}

	return nil
}

func parseFile(filepath string) map[string][]string {
	// get content
	bytes, _ := ioutil.ReadFile(filepath)

	// 标题，tag处理
	rec, _ := regexp.Compile("\\-\\-\\-([^\\-]*)\\-\\-\\-")

	if !rec.Match(bytes) {
		return nil
	}

	data := make(map[string][]string)

	match := rec.FindStringSubmatch(string(bytes))
	baseInfo := strings.Split(strings.Trim(match[1], "\n"), "\n")
	for _, v := range baseInfo {
		item := strings.Split(v, `:`)
		if item[0] == `tags` {
			data[item[0]] = strings.Split(item[1],`,`)
		} else {
			data[item[0]] = []string{item[1]}
		}
	}

	contents := strings.Split(string(bytes),`---`)

	data[`content`] = []string{contents[2]}

	length := len(contents[2])
	if length > 100 {
		length = 100
	}
	data[`description`] = []string{string([]rune(contents[2])[:length])}

	return data
}
