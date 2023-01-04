package test

import (
	"context"
	"ember/api/openai/common"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"os"
	"strings"
	"testing"
)

func TestChatGptModel(t *testing.T) {
	f, _ := os.Getwd()
	p := f[:strings.LastIndex(f, "/")] + "/env/keys.yml"
	ak := common.GetApiKey(p)
	c := gogpt.NewClient(ak)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "ada",
		MaxTokens: 5,
		Prompt:    "Lorem ipsum",
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
