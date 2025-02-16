package pkg

import (
	"bufio"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

// openapi completion
func Cmpl(instruct string) {

	if strings.TrimSpace(instruct) == "" {
		instruct = "请输入问题"
	}
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: instruct},
	}

	// 获取用户输入
	reader := bufio.NewReader(os.Stdin) // 读取标准输入
	fmt.Println(instruct + "\npress Ctrl+D to finish")

	var question strings.Builder
	for {
		line, err := reader.ReadString('\n') // 读取一行，保留 `\n`
		if err != nil {
			fmt.Println("\n")
			break
		}
		question.WriteString(line)
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question.String(),
	})

	// make a request
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4oMini, // 或 GPT3Dot5Turbo
		Messages: messages,
		Stream:   true, // 启用流式响应
	}

	// 发送请求并流式读取回复
	stream, err := client.CreateChatCompletionStream(context.Background(), req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer stream.Close()

	fmt.Print("AI：")
	var fullResponse string
	for {
		response, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Print(response.Choices[0].Delta.Content) // 实时输出流数据
		fullResponse += response.Choices[0].Delta.Content
	}
	fmt.Println() // 输出换行

}
