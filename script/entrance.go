package script

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Run(path string) {
	logPrefix := "run script"
	fmt.Printf("%s %s\n", logPrefix, "start ->")

	var (
		prefix string
		args   []string
		suffix string
	)

	switch runtime.GOOS {
	case "windows":
		prefix = "cmd"
		args = []string{"/c"}
		suffix = "bat"
	case "darwin", "linux":
		prefix = "sh"
		suffix = "sh"
	default:
		fmt.Println("Unknown operating system")
		return
	}

	args = append(args, fmt.Sprintf("%s.%s", path, suffix))
	sc := exec.Command(prefix, args...)
	_, err := sc.CombinedOutput()
	if err != nil {
		fmt.Println("run script err:", err)
		return
	}

	fmt.Printf("%s %s\n", logPrefix, "success ->")
}
