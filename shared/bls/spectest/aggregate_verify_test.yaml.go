// Code generated by yaml_to_go. DO NOT EDIT.
// source: aggregate_verify.yaml

package spectest

type AggregateVerifyTest struct {
	Input struct {
		Pubkeys   []string `json:"pubkeys"`
		Messages  []string `json:"messages"`
		Signature string   `json:"signature"`
	} `json:"input"`
	Output bool `json:"output"`
}