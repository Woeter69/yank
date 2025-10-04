package yank

import (
	"fmt"
	"os"
	"errors"
	"github.com/spf13/cobra"
	"github.com/Woeter69/yank/internal/install"
	"github.com/Woeter69/yank/internal/utils"
	"github.com/Woeter69/yank/internal/remove"
)

func main() {
	rootCmd := &cobra.Command {
		Use: "yank",
		Short: "Yank package manager",
	}

	installCmd := &cobra.Command {
		Use: "install [package]",
		Short: "Install a package",
	} 

	removeCmd := &cobra.Command {
		Use: "remove [package]",
		Short: "Remove a package",
	}

}
