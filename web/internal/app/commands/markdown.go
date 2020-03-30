package commands

import (
	"fmt"
	"github.com/firmeve/firmeve/kernel/contract"
	"github.com/firmeve/firmeve/support/path"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
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

	if err := getAllFile(path); err != nil {
		logger.Error("%#v", err)
	}

	// 先解析一个单文件
	md := []byte("## markdown document")
	output := markdown.ToHTML(md, nil, nil)
	fmt.Println(cmd.Flag(`path`).Value.String())
	fmt.Printf("%s", output)
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

func getAllFile(filepath string) error {
	// 遍历所有目录和文件
	resources, err := ioutil.ReadDir(filepath)
	if err != nil {
		return err
	}

	for _, file := range resources {
		if file.IsDir() {
			//getAllFile()
		} else {
			filefullpath := filepath + `/` + file.Name()
			bytes, err := ioutil.ReadFile(filefullpath)

			opts := html.RendererOptions{
				Flags: html.CommonFlags,
				//RenderNodeHook: renderHookDropCodeBlock,
			}
			renderer := html.NewRenderer(opts)

			html := markdown.ToHTML(bytes,nil,renderer)


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
			if err == nil {
				//fileinfo,_ := os.Stat(filefullpath)
				filename := strings.Replace(file.Name(),".md",".html",1)
				ioutil.WriteFile(path.RunRelative("../../../public/html/"+filename), []byte(html), file.Mode())
			}
			fmt.Println(file.Name())
		}
	}

	return nil
}
