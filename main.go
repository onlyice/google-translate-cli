package main

import (
	_ "embed"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/urfave/cli/v2"
	"html/template"
	"log"
	"net/url"
	"os"
)

//go:embed result.html
var htmlTepl string
var tmpl = template.Must(template.New("result").Parse(htmlTepl))

func main() {
	app := &cli.App{
		Usage: "tranlate text using Google Translate API",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "from",
				Aliases: []string{"f"},
				Value:   "auto",
				Usage:   "the text language (en, de, ...)",
			},
			&cli.StringFlag{
				Name:    "to",
				Aliases: []string{"t"},
				Value:   "zh_CN",
				Usage:   "the language in which the text should be translated",
			},
			&cli.StringFlag{
				Name:  "host",
				Value: "google.cn",
				Usage: "the host of Google translate API",
			},
			&cli.StringFlag{
				Name:     "text",
				Usage:    "the text to be translated",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "gd",
				Usage: "output html content for GoldenDict",
			},
		},
		Action: func(c *cli.Context) error {
			text := c.String("text")

			gd := c.Bool("gd")
			if gd {
				// GoldenDict 会把文本用 percentage encoding 转换之后传进来
				var err error
				text, err = url.QueryUnescape(text)
				if err != nil {
					return err
				}
			}

			translated, err := gtranslate.TranslateWithParams(
				text,
				gtranslate.TranslationParams{
					From:       c.String("from"),
					To:         c.String("to"),
					GoogleHost: c.String("host"),
				},
			)
			if err != nil {
				panic(err)
			}

			if gd {
				tmpl.Execute(os.Stdout, &struct {
					Text, Translated string
				}{
					Text:       c.String("text"),
					Translated: translated,
				})
			} else {
				fmt.Println(translated)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
