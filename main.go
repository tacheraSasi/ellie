package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/command"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

const (
	VERSION = configs.VERSION
)

var ICON any = static.Icon()

// User name from the saved files during initialization
var CurrentUser string = configs.GetEnv("USERNAME")

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		styles.GetErrorStyle().Printf("\nellie Terminated gracefully\n")
		os.Exit(0)
	}()

	// Setup global flags
	showHelp := flag.Bool("help", false, "Show help information")
	showVersion := flag.Bool("version", false, "Show version information")
	flag.Parse()

	// Handle global flags
	if *showVersion {
		styles.GetSuccessStyle().Printf("Ellie CLI v%s\n", VERSION)
		return
	}
	if *showHelp {
		actions.ShowHelp()
		return
	}

	// Initialize configuration
	configs.Init()

	// Get remaining arguments after flags
	args := flag.Args()

	// Interactive mode if no commands
	if len(args) == 0 {
		// Show welcome message before chat
		actions.ShowWelcome()
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	handleCommand(args)
}

func handleCommand(args []string) {
	if len(args) == 0 {
		actions.Chat(configs.GetEnv("OPENAI_API_KEY"))
		return
	}

	cmdName := args[0]

	// Check if the command is an alias
	if actions.ExecuteAlias(cmdName) {
		return
	}

	cmd, exists := command.Registry[cmdName]
	if !exists {
		matches := getClosestMatchingCmd(command.Registry, cmdName)
		// fmt.Println(matches)
		if len(matches) > 0 {
			styles.GetErrorStyle().Printf("Unknown command: %s\n", cmdName)
			styles.GetInfoStyle().Println("Did you mean:")
			for _, m := range matches {
				styles.GetInfoStyle().Printf("  %s\n", m)
			}
		} else {
			styles.GetErrorStyle().Printf("Unknown command: %s\n", cmdName)
			actions.ShowHelp()
		}
		os.Exit(1)
	}

	if cmd.PreHook != nil {
		cmd.PreHook()
	}

	if len(cmd.SubCommands) > 0 && len(args) > 1 {
		handleSubCommand(cmd, args[1:])
		return
	}

	if len(args)-1 < cmd.MinArgs {
		styles.GetErrorStyle().Printf("Invalid usage for %s\n", cmdName)
		styles.GetInfoStyle().Println("Usage:", cmd.Usage)
		os.Exit(1)
	}

	cmd.Handler(args)
}

// Returns a list of command names that closely match the input
func getClosestMatchingCmd(cmdMap map[string]command.Command, cmdArg string) []string {
	var list []string
	for cmd := range cmdMap {
		distance := levenshtein.DistanceForStrings([]rune(cmdArg), []rune(cmd), levenshtein.DefaultOptions)
		maxLen := len(cmdArg)
		if len(cmd) > maxLen {
			maxLen = len(cmd)
		}
		similarity := 1.0 - (float64(distance) / float64(maxLen))
		if similarity > 0.4 {
			list = append(list, cmd)
		}
	}
	return list
}

// Returns a list of subcommand names that closely match the input
func getClosestMatchingSubCmd(subCmdMap map[string]command.Command, subCmdArg string) []string {
	var list []string
	for cmd := range subCmdMap {
		distance := levenshtein.DistanceForStrings([]rune(subCmdArg), []rune(cmd), levenshtein.DefaultOptions)
		maxLen := len(subCmdArg)
		if len(cmd) > maxLen {
			maxLen = len(cmd)
		}
		similarity := 1.0 - (float64(distance) / float64(maxLen))
		if similarity > 0.4 {
			list = append(list, cmd)
		}
	}
	return list
}

func handleSubCommand(parentCmd command.Command, args []string) {
	subCmdName := args[0]
	subCmd, exists := parentCmd.SubCommands[subCmdName]
	if !exists {
		matches := getClosestMatchingSubCmd(parentCmd.SubCommands, subCmdName)
		if len(matches) > 0 {
			styles.GetErrorStyle().Printf("Unknown subcommand: %s\n", subCmdName)
			styles.GetInfoStyle().Println("Did you mean:")
			for _, m := range matches {
				styles.GetInfoStyle().Printf("  %s\n", m)
			}
		} else {
			styles.GetErrorStyle().Printf("Unknown subcommand: %s\n", subCmdName)
			styles.GetInfoStyle().Println("Available subcommands:")
			for name := range parentCmd.SubCommands {
				styles.GetInfoStyle().Printf("  %s\n", name)
			}
		}
		os.Exit(1)
	}

	if subCmd.PreHook != nil {
		subCmd.PreHook()
	}

	// Handle nested subcommands
	if len(subCmd.SubCommands) > 0 && len(args) > 1 {
		handleSubCommand(subCmd, args[1:])
		return
	}

	if len(args)-1 < subCmd.MinArgs {
		styles.GetErrorStyle().Printf("Invalid usage for %s\n", subCmdName)
		styles.GetInfoStyle().Println("Usage:", subCmd.Usage)
		os.Exit(1)
	}

	subCmd.Handler(args)
}

func createServiceCommand(action string) command.Command {
	return command.Command{
		SubCommands: map[string]command.Command{
			"apache":   {Handler: func(args []string) { actions.HandleService(action, "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService(action, "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService(action, "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService(action, "all") }},
		},
	}
}

func greetUser(args []string) {
	hour := time.Now().Hour()
	greeting := styles.GetSuccessStyle().Println
	message := "Good evening!"

	switch {
	case hour < 12:
		message = "Good morning!"
		greeting = styles.GetHighlightStyle().Println
	case hour < 18:
		message = "Good afternoon!"
		greeting = styles.GetInfoStyle().Println
	}

	greeting(message+",", CurrentUser)
}
