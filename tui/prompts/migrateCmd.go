package prompts

import "github.com/sveltinio/prompti/confirm"

// ConfirmMigration renders a confirmation dialog box.
func ConfirmMigration(question string) (bool, error) {
	return confirm.Run(&confirm.Config{Question: question})
}
