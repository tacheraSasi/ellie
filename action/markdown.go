package actions

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/tacheraSasi/ellie/styles"
	"github.com/tacheraSasi/ellie/utils"
)

// MarkdownRender renders a markdown file in the terminal
func MarkdownRender(args []string) {
	if len(args) < 2 {
		styles.GetErrorStyle().Println("❌ Please provide a markdown file path")
		styles.GetInfoStyle().Println("Usage: ellie md <filename>")
		return
	}

	filePath := args[1]
	
	// Validate file extension
	if !strings.HasSuffix(strings.ToLower(filePath), ".md") && !strings.HasSuffix(strings.ToLower(filePath), ".markdown") {
		styles.GetErrorStyle().Printf("❌ File must have a .md or .markdown extension: %s\n", filePath)
		return
	}

	// Check if file exists
	if !utils.Exists(filePath) {
		styles.GetErrorStyle().Printf("❌ File not found: %s\n", filePath)
		return
	}

	// Get absolute path for better error messages
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		styles.GetErrorStyle().Printf("❌ Error resolving file path: %v\n", err)
		return
	}

	// Read the markdown file
	content, err := utils.ReadFile(absPath)
	if err != nil {
		styles.GetErrorStyle().Printf("❌ Error reading file %s: %v\n", absPath, err)
		return
	}

	if strings.TrimSpace(content) == "" {
		styles.GetWarningStyle().Printf("⚠️ File %s is empty\n", absPath)
		return
	}

	// Get appropriate theme for rendering
	theme := getThemeForRendering()
	
	// Render the markdown with the selected theme
	rendered, err := renderMarkdownWithTheme(content, theme)
	if err != nil {
		styles.GetErrorStyle().Printf("❌ Error rendering markdown: %v\n", err)
		return
	}

	// Display the rendered output
	fmt.Printf("\n📝 %s\n\n", absPath)
	fmt.Print(rendered)
}

// getThemeForRendering determines the best theme to use based on current theme settings
func getThemeForRendering() string {
	currentTheme := styles.GetTheme()
	
	switch currentTheme {
	case "light":
		return "light"
	case "dark":
		return "dark"
	case "auto":
		// For auto mode, we'll default to dark for now
		// In a more advanced implementation, this could detect terminal background
		return "dark"
	default:
		return "dark"
	}
}

// renderMarkdownWithTheme renders markdown content with the specified theme
func renderMarkdownWithTheme(content, theme string) (string, error) {
	// Available themes that work well
	validThemes := map[string]bool{
		"dark":            true,
		"light":           true,
		"notty":           true,
		"pink":            true,
		"solarized-dark":  true,
		"solarized-light": true,
		"dracula":         true,
		"no-color":        true,
		"auto":            true,
	}

	// Use the theme if it's valid, otherwise fall back to dark
	if !validThemes[theme] {
		theme = "dark"
	}

	// Render using glamour with the selected theme
	rendered, err := glamour.Render(content, theme)
	if err != nil {
		// If the theme fails, try with a basic fallback
		rendered, err = glamour.Render(content, "dark")
		if err != nil {
			return "", fmt.Errorf("failed to render markdown: %v", err)
		}
	}

	return rendered, nil
}