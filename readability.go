package readability

import (
	"context"

	"gitee.com/baixudong/cmd"
	"github.com/tidwall/gjson"

	_ "embed"
)

//go:embed readability.js
var readbilityJs string

type Client struct {
	cmd *cmd.JyClient
}
type ParseOption struct {
	MaxElemsToParse   int    `json:"maxElemsToParse,omitempty"`
	NbTopCandidates   int    `json:"nbTopCandidates,omitempty"`
	CharThreshold     int    `json:"charThreshold,omitempty"`
	ClassesToPreserve string `json:"classesToPreserve,omitempty"`
	KeepClasses       bool   `json:"keepClasses,omitempty"`
	Serializer        string `json:"serializer,omitempty"`
	DisableJsonLd     bool   `json:"disableJSONLD,omitempty"`
}

func NewClient(ctx context.Context) (*Client, error) {
	cli, err := cmd.NewJsClient(ctx, cmd.JsClientOption{
		Script: readbilityJs,
		Names:  []string{"clear"},
	})
	return &Client{cmd: cli}, err
}

func (obj *Client) Parse(url, content string, options ...ParseOption) (gjson.Result, error) {
	var option ParseOption
	if len(options) > 0 {
		option = options[0]
	}
	return obj.cmd.Call("clear", url, content, option)
}
func (obj *Client) Close() {
	obj.cmd.Close()
}
