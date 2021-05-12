package main

import (
	_ "embed"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/urfave/cli/v2"
	"html/template"
	"log"
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
				Name:     "html",
				Usage: "output html content",
			},
		},
		Action: func(c *cli.Context) error {
			translated, err := gtranslate.TranslateWithParams(
				c.String("text"),
				gtranslate.TranslationParams{
					From:       c.String("from"),
					To:         c.String("to"),
					GoogleHost: c.String("host"),
				},
			)
			if err != nil {
				panic(err)
			}

			if c.Bool("html") {
				tmpl.Execute(os.Stdout, &struct {
					Text, Translated string
				}{
					Text: c.String("text"),
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
