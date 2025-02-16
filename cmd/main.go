package main

import (
	"github.com/spf13/cobra"
	openai "heidi/m/pkg"
)

func main() {

	var instruct string
	var d, c, e bool
	var cmpl = &cobra.Command{
		Use:   "ask [ask a word, phrase or sentence]",
		Short: "check the meaning of it",
		Long:  `print the meaning of the word, phrase or sentence`,

		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case d == true:
				instruct = "请翻译成德语"
			case c == true:
				instruct = "请翻译成中文"
			case e == true:
				instruct = "请翻译成英语"
			default:
				instruct = "请提出问题"
			}
			openai.Cmpl(instruct)
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
	cmpl.Flags().BoolVarP(&d, "deutsch", "d", false, "transfer to de")
	cmpl.Flags().BoolVarP(&c, "chinese", "c", false, "transfer to cn")
	cmpl.Flags().BoolVarP(&e, "english", "e", false, "transfer to en")

	rootCmd.AddCommand(conv)
	rootCmd.Execute()
}
