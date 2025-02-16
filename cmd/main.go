package main

import (
	"github.com/spf13/cobra"
	openai "heidi/m/pkg"
)

func main() {

	var cmpl = &cobra.Command{
		Use:   "ask [ask a word, phrase or sentence]",
		Short: "check the meaning of it",
		Long:  `print the meaning of the word, phrase or sentence`,

		Run: func(cmd *cobra.Command, args []string) {
			openai.Cmpl()
		},
	}

	var conv = &cobra.Command{
		Use:   "conv [start a conversation]",
		Short: "start a conversation",
		Long:  `start a conversation`,

		Run: func(cmd *cobra.Command, args []string) {
			openai.Conv()
		},
	}

	var rootCmd = &cobra.Command{Use: "ai_asker"}

	rootCmd.AddCommand(cmpl)
	rootCmd.AddCommand(conv)
	rootCmd.Execute()
}
