package utils

import (
	"bytes"
	"context"
	"io"
	"strings"

	"github.com/a-h/templ"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yuin/goldmark"
)

func RenderMarkdown(data []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		var buf bytes.Buffer
		err := goldmark.Convert([]byte(strings.Join(data, "\n")), &buf)
		if err != nil {
			return err
		}
		html := bluemonday.UGCPolicy().SanitizeBytes(buf.Bytes())
		_, err = io.WriteString(w, string(html))
		return err
	})
}
