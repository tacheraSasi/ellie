package command

import (
	"flag"
	"fmt"

	actions "github.com/tacheraSasi/ellie/action"
	"github.com/tacheraSasi/ellie/configs"
	"github.com/tacheraSasi/ellie/static"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/types"
)

var Registry = map[string]Command{
	"run": {
		Handler: actions.Run,
	},
	"::": {
		Usage:   "ellie :: run docker container for me please",
		MinArgs: 1,
		Handler: actions.SmartRun,
	},
	"code":{
		Usage: "ellie code",
		Handler: func(_ []string) {
			actions.StartEllieCode()
		},
	},
	"user-env": {
		Handler: func(s []string) {
			// Create user context
			userCtx := types.NewUserContext()

			// Add system message with instructions and context
			instructions := fmt.Sprintf(`!!!!!!!!!!!!!!!!!!!!!IMPORTANT YOU ARE ELLIE note: %s `, static.Instructions(*userCtx))
			fmt.Println(instructions)
		},
	},
	"focus": {
		PreHook: func() { styles.InfoStyle.Println("Activating focus mode...") },
		Handler: actions.Focus,
	},
	"pwd": {
		Handler: func(_ []string) { actions.Pwd() },
	},
	"size": {
		MinArgs: 1,
		Handler: func(s []string) { actions.Size() },
	},
	"open-explorer": {
		Handler: func(_ []string) { actions.OpenExplorer() },
	},
	"open": {
		Usage:   "open <path>",
		MinArgs: 1,
		Handler: func(args []string) {
			actions.OpenExplorer(args[1])
		},
	},
	"play": {
		MinArgs: 1,
		Usage:   "play <media>",
		PreHook: func() { styles.InfoStyle.Println("Initializing media player...") },
		Handler: actions.Play,
	},
	"setup-git": {
		Handler: func(args []string) {
			actions.GitSetup(configs.GetEnv("PAT"), configs.GetEnv("USERNAME"))
		},
	},
	"sysinfo": {
		Handler: func(_ []string) { actions.SysInfo() },
	},
	"dev-init": {
		Handler: func(args []string) {
			fs := flag.NewFlagSet("dev-init", flag.ExitOnError)
			allFlag := fs.Bool("all", false, "Install all recommended tools")
			fs.Parse(args[1:])
			actions.DevInit(*allFlag)
		},
	},
	"server-init": {
		Handler: func(_ []string) { actions.ServerInit() },
	},
	"install": {
		MinArgs: 1,
		Usage:   "install <package>",
		Handler: func(args []string) { actions.InstallPackage(args[1]) },
	},
	"update": {
		Handler: func(_ []string) { actions.UpdatePackages() },
	},
	"list": {
		MinArgs: 1,
		Usage:   "list <directory>",
		Handler: func(args []string) { actions.ListFiles(args[1]) },
	},
	"create-file": {
		MinArgs: 1,
		Usage:   "create-file <path>",
		Handler: func(args []string) { actions.CreateFile(args[1]) },
	},
	"network-status": {
		Handler: func(_ []string) { actions.NetworkStatus() },
	},
	"connect-wifi": {
		MinArgs: 2,
		Usage:   "connect-wifi <SSID> <password>",
		Handler: func(args []string) { actions.ConnectWiFi(args[1], args[2]) },
	},
	"greet": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", configs.GetEnv("USERNAME"))
		},
	},
	"send-mail": {
		Handler: func(_ []string) { actions.Mailer() },
	},
	"chat": {
		Handler: func(_ []string) { actions.Chat(configs.GetEnv("OPENAI_API_KEY")) },
	},
	"review": {
		Usage:   "review <filename/filepath>",
		MinArgs: 1,
		Handler: func(args []string) { actions.Review(args[1]) },
		// PreHook: ,
	},
	"security-check": {
		Usage:   "security-check <path>",
		MinArgs: 1,
		Handler: func(args []string) { actions.SecurityCheck(args[1]) },
		// PreHook: ,
	},
	"git": {
		SubCommands: map[string]Command{
			"status":        {Handler: func(_ []string) { actions.GitStatus() }},
			"push":          {Handler: func(_ []string) { actions.GitPush() }},
			"commit":        {Handler: func(args []string) { actions.GitConventionalCommit() }},
			"pull":          {Handler: func(_ []string) { actions.GitPull() }},
			"branch-create": {Handler: func(_ []string) { actions.GitBranchCreate() }},
			"branch-switch": {Handler: func(_ []string) { actions.GitBranchSwitch() }},
			"branch-delete": {Handler: func(_ []string) { actions.GitBranchDelete() }},
			"stash-save":    {Handler: func(_ []string) { actions.GitStashSave() }},
			"stash-pop":     {Handler: func(_ []string) { actions.GitStashPop() }},
			"stash-list":    {Handler: func(_ []string) { actions.GitStashList() }},
			"tag-create":    {Handler: func(_ []string) { actions.GitTagCreate() }},
			"tag-list":      {Handler: func(_ []string) { actions.GitTagList() }},
			"tag-delete":    {Handler: func(_ []string) { actions.GitTagDelete() }},
			"log":           {Handler: func(_ []string) { actions.GitLogPretty() }},
			"diff":          {Handler: func(_ []string) { actions.GitDiff() }},
			"merge":         {Handler: func(_ []string) { actions.GitMerge() }},
			"rebase":        {Handler: func(_ []string) { actions.GitRebase() }},
			"cherry-pick":   {Handler: func(_ []string) { actions.GitCherryPick() }},
			"reset":         {Handler: func(_ []string) { actions.GitReset() }},
			"bisect":        {Handler: func(_ []string) { actions.GitBisect() }},
			"bisect-good":   {Handler: func(_ []string) { actions.GitBisectGood() }},
			"bisect-bad":    {Handler: func(_ []string) { actions.GitBisectBad() }},
			"bisect-reset":  {Handler: func(_ []string) { actions.GitBisectReset() }},
		},
	},
	"start": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("start", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("start", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("start", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("start", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"stop": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("stop", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("stop", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("stop", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("stop", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"restart": {
		SubCommands: map[string]Command{
			"apache":   {Handler: func(args []string) { actions.HandleService("restart", "apache") }},
			"mysql":    {Handler: func(args []string) { actions.HandleService("restart", "mysql") }},
			"postgres": {Handler: func(args []string) { actions.HandleService("restart", "postgres") }},
			"all":      {Handler: func(args []string) { actions.HandleService("restart", "all") }},
			"list":     {Handler: func(args []string) { actions.ListServices() }},
		},
	},
	"config": {
		Handler: func(_ []string) { configs.Init() },
	},
	"reset-config": {
		Handler: func(_ []string) { configs.ResetConfig() },
	},
	"whoami": {
		Handler: func(_ []string) {
			styles.Highlight.Println("Your majesty,", configs.GetEnv("USERNAME"))
		},
	},
	"alias": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 1,
				Usage:   "alias add <name>=\"<command>\"",
				Handler: actions.AliasAdd,
			},
			"list": {
				Handler: actions.AliasList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "alias delete <name>",
				Handler: actions.AliasDelete,
			},
		},
	},
	"todo": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 1,
				Usage:   "todo add \"<task>\" [category] [priority]",
				Handler: actions.TodoAdd,
			},
			"list": {
				Handler: actions.TodoList,
			},
			"complete": {
				MinArgs: 1,
				Usage:   "todo complete <id>",
				Handler: actions.TodoComplete,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "todo delete <id>",
				Handler: actions.TodoDelete,
			},
			"edit": {
				MinArgs: 3,
				Usage:   "todo edit <id> <field> <value>",
				Handler: actions.TodoEdit,
			},
		},
	},
	"project": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 2,
				Usage:   "project add <name> <path> [description] [tags...]",
				Handler: actions.ProjectAdd,
			},
			"list": {
				Handler: actions.ProjectList,
			},
			"delete": {
				MinArgs: 1,
				Usage:   "project delete <name>",
				Handler: actions.ProjectDelete,
			},
			"search": {
				MinArgs: 1,
				Usage:   "project search <query>",
				Handler: actions.ProjectSearch,
			},
		},
	},
	"switch": {
		MinArgs: 1,
		Usage:   "switch <project-name>",
		Handler: actions.ProjectSwitch,
	},
	"history": {
		Handler: actions.History,
	},
	"start-day": {
		Handler: actions.StartDay,
	},
	"day-start": {
		SubCommands: map[string]Command{
			"add": {
				MinArgs: 2,
				Usage:   "day-start add <type> <value>",
				Handler: actions.DayStartConfigAdd,
			},
			"list": {
				Handler: actions.DayStartConfigList,
			},
		},
	},

	//Pending commands
	"weather": {
		Handler: func(args []string) { actions.Weather() },
	},
	"joke": {
		Handler: func(args []string) { actions.Joke() },
	},
	"remind": {
		Handler: func(_ []string) { actions.Remind() },
	},
	"about": {
		Handler: actions.ShowAbout,
	},
	"theme": {
		SubCommands: map[string]Command{
			"set": {
				MinArgs: 1,
				Usage:   "theme set <light|dark|auto>",
				Handler: func(args []string) {
					mode := args[1]
					if mode != "light" && mode != "dark" && mode != "auto" {
						styles.GetErrorStyle().Println("Invalid theme. Use 'light', 'dark', or 'auto'.")
						return
					}
					styles.SetTheme(mode)
					styles.GetSuccessStyle().Printf("Theme set to %s.\n", styles.GetTheme())
				},
			},
			"show": {
				Handler: func(_ []string) {
					styles.GetInfoStyle().Printf("Current theme: %s\n", styles.GetTheme())
				},
			},
		},
	},
	"md": {
		Usage:   "md <filename>",
		MinArgs: 1,
		Handler: actions.MarkdownRender,
	},
}
