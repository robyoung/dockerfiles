package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strings"
	"syscall"
	"time"
)


type Config struct {
	Home string
	Dev  string
	Pwd  string
}

func getConfig() Config {
	return Config{
		os.Getenv("HOME"),
		os.Getenv("DEV_DIR"),
		os.Getenv("PWD"),
	}
}

func runAsParent(args ...string) {
	fmt.Printf("Running: %s\n", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	start := time.Now()
	if err := cmd.Run(); err != nil {
		log.Fatal("Failed to run command: ", err)
	}
	elapsed := time.Now().Sub(start)
	fmt.Println("Elapsed", elapsed.String())
}

func main() {
	config := getConfig()
	if err := syscall.Setgroups([]int{999}); err != nil {
		log.Fatal("Failed to set groups: ", err)
	}
	args := []string{
		"docker", "run", "-ti",
			"-u", "robyoung",
			"-e", "VIM_EXTRA_PLUGINS=1",
			"-e", fmt.Sprintf("VIMHOME=%s/.vim", config.Home),
			"-e", "FZF_DEFAULT_COMMAND=rg --files",
			"-w", config.Pwd,
			"-v", fmt.Sprintf("%[1]s/%[2]s:%[1]s/%[2]s", config.Home, config.Dev),
			"-v", fmt.Sprintf("%[1]s/dev-vim:%[1]s/.vim", config.Home),
			"-v", fmt.Sprintf("%[1]s/%[2]s/personal/dotfiles/vimrc:%[1]s/.vimrc", config.Home, config.Dev),
			"-v", fmt.Sprintf("%[1]s/.viminfo:%[1]s/.viminfo", config.Home),
			"-v", fmt.Sprintf("%[1]s/%[2]s/github/junegunn/fzf:%[1]s/.fzf", config.Home, config.Dev),
			"vim",
				"vim",
				"-i", fmt.Sprintf("%[1]s/.vim/viminfo", config.Home),
	}
	args = append(args, os.Args[1:]...)
	runAsParent(args...)
}
