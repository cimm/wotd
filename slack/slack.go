package slack

import (
	"bytes"
	"encoding/json"
	"github.com/cimm/wotd/wikitionary"
	"log"
	"net/http"
	"text/template"
)

const (
	contentType = "application/json"
	userAgent   = "WotdBot/0.1 (+https://github.com/cimm/wotd)"
	wotdTmpl    = `*<{{.Url}}|{{.Word}}>*
{{- range .Definitions}}
> {{.}}
{{- end}}`
)

func FormatWotd(wotd wikitionary.Wotd) string {
	message := new(bytes.Buffer)
	t := template.Must(template.New("wotd").Parse(wotdTmpl))
	t.Execute(message, wotd)
	return message.String()
}

func Post(webhook, message string) {
	payload := map[string]string{
		"text": message,
	}
	jsonPayload, err := json.MarshalIndent(payload, "", "")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	c := http.Client{}
	req, err := http.NewRequest(http.MethodPost, webhook, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", userAgent)

	_, err = c.Do(req)
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
