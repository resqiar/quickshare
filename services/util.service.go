package services

import (
	"bytes"
	"log"
	"sync"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	html "github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/anchor"
)

type UtilService interface {
	ParseMD(s []byte) []byte
}

type UtilServiceImpl struct{}

var (
	engine = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("paraiso-dark"),
				highlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
			&anchor.Extender{
				Texter:   anchor.Text("#"),
				Position: anchor.After,
			},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	sanitizePolicy = bluemonday.UGCPolicy().AllowAttrs("style").OnElements("p", "span", "pre")
	bufferPool     = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)

func (service *UtilServiceImpl) ParseMD(s []byte) []byte {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	// since we are reusing resources, better reset it first
	buf.Reset()

	if err := engine.Convert([]byte(s), buf); err != nil {
		log.Println("Error parsing MD:", err)
		return nil
	}

	sanitized := sanitizePolicy.SanitizeBytes(buf.Bytes())
	return sanitized
}
