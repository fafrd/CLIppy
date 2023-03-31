package brain

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
)

const tokens = 200

type Brain struct {
	apiKey string
	model  string
}

func NewBrain(model string) (*Brain, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("undefined env var OPENAI_API_KEY")
	}

	return &Brain{
		apiKey: apiKey,
		model:  model,
	}, nil
}

func (b *Brain) InterpretGoal(goal string) (string, error) {
	prompt := fmt.Sprintf(`Interpolate the goal statement into the following statement, as if you are Clippy, the helpful assistant. Don't be too verbose, though. Just reiterate the goal.

Statement: It looks like you are trying to [goal].
	
Goal: %s`, goal)

	return b.genDialogue(prompt)
}

/*
Assume we are using Ubuntu Linux. Speak as if you are Clippy, the helpful assistant. Our goal is to Run a minecraft server.

So far:
The user installed openjdk-8-jre.
The user tried to run wget to download the minecraft jar, but wget was not installed.
The user successfully installed wget.
The user successfully downloaded https://launcher.mojang.com/v1/objects/bb2b6b1aefcd70dfd1892149ac3a215f6c636b07/server.jar.

What is our current goal, and what command should the user run?

Speak as if you are Clippy, the helpful assistant. Speak in the format of: "Our current goal is _. You should _."

Our current goal is to run the minecraft server. You should run the command 'java -Xmx1024M -Xms1024M -jar server.jar nogui'.
*/

func (b *Brain) GetNextStatement(lastCommandRun string, lastCommandOutput string) (string, error) {
	// 1. Get outcome from prev command in form of "The user ___."
	// 2. Get next command in form of "Our current goal is ___. You should run ___."
	return "", nil // TODO
}

func (b *Brain) genDialogue(prompt string) (string, error) {

	ctx := context.Background()
	client := gpt3.NewClient(b.apiKey)

	fmt.Printf("### Sending request to OpenAI:\n%s\n\n", prompt)

	messages := []gpt3.ChatCompletionRequestMessage{
		{
			Role:    "user",
			Content: prompt,
		},
	}
	request := gpt3.ChatCompletionRequest{
		Model:       b.model,
		Messages:    messages,
		MaxTokens:   tokens,
		Temperature: gpt3.Float32Ptr(0.0),
	}
	resp, err := client.ChatCompletion(ctx, request)

	if err != nil {
		fmt.Printf("### ERROR from OpenAI:\n%s\n\n", err)
		return "", err
	}

	trimmedResponse := strings.TrimSpace(resp.Choices[0].Message.Content)
	fmt.Printf("### Received response from OpenAI:\n%s\n\n\n", trimmedResponse)
	return trimmedResponse, nil
}
