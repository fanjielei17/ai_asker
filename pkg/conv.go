package pkg

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// openai conversation
func Conv() {
	// 读取 OpenAI API Key
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	// 初始化对话历史
	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: "你是一个 AI 助手"},
	}
	for {
		// 获取用户输入
		reader := bufio.NewReader(os.Stdin) // 读取标准输入
		fmt.Println("请输入问题（按 Ctrl+D 结束）：")

		var question strings.Builder
		for {
			line, err := reader.ReadString('\n') // 读取一行，保留 `\n`
			if err != nil {
				fmt.Println("\n输入结束（EOF）")
				break
			}
			question.WriteString(line)
		}

		// 添加用户输入到对话历史
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

		// 添加 AI 回复到对话历史
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: fullResponse,
		})
	}
}
