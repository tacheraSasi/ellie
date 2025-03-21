package utils

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/tacheraSasi/ellie/styles"
)

// Ads for random promotion messages
var Ads []string = []string{
	"🚀 Boost your productivity with ekilie!",
	"🔥 Check out ekiliSense for smarter school management!",
	"💻 Need a project tracker? Try ekilie!",
}

// GetInput prompts the user for input and returns the trimmed string.
func GetInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	styles.InputPrompt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// RandNum generates a random number between 0 and 100.
func RandNum() int {
	return rand.IntN(100)
}

// RandNumRange generates a random number between min and max.
func RandNumRange(min, max int) int {
	return rand.IntN(max-min+1) + min
}

// IsEven checks if a number is even.
func IsEven(num int) bool {
	return num%2 == 0
}

// IsOdd checks if a number is odd.
func IsOdd(num int) bool {
	return num%2 != 0
}

// RunCommand executes a shell command and prints the output or error.
func RunCommand(cmdArgs []string, errMsg string) {
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%s %s\n", errMsg, err)
		return
	}
	if len(output) > 0 {
		fmt.Printf("Output:\n%s\n", output)
	}
}

// IsLinux returns true if the OS is Linux.
func IsLinux() bool {
	return strings.Contains(runtime.GOOS, "linux")
}

// IsMac returns true if the OS is macOS.
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsWindows returns true if the OS is Windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// RenderMarkdown renders Markdown input using Glamour.
func RenderMarkdown(input string) (string, error) {
	rendered, err := glamour.Render(input, "dark")
	if err != nil {
		return "", err
	}
	return rendered, nil
}

// Exists checks if a file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ReadFile reads a file and returns its content as a string.
func ReadFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes a string to a file.
func WriteFile(filePath, content string) error {
	return os.WriteFile(filePath, []byte(content), 0644)
}

// AppendToFile appends content to a file.
func AppendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content + "\n")
	return err
}

// CurrentTimestamp returns the current timestamp as a formatted string.
func CurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Sleep pauses execution for a given number of seconds.
func Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// ClearScreen clears the console screen based on the OS.
func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// GetRandomAd returns a random promotional message from Ads.
func GetRandomAd() string {
	return Ads[rand.IntN(len(Ads))]
}

func IsErr(err error,msg string){
	if err != nil{
		ErrorStyle.Println(msg,err)
		return
	}
}
