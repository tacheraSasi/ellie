# 📝 Ellie CLI - Markdown Rendering Feature

This document demonstrates the new `ellie md` command functionality.

## ✨ Features

- **Theme-aware rendering** that adapts to your current Ellie theme
- **Beautiful syntax highlighting** for code blocks  
- **Full markdown support** including tables, lists, and blockquotes
- **File validation** with support for `.md` and `.markdown` extensions
- **Error handling** for missing or invalid files

## 🚀 Usage

```bash
# Render a markdown file
ellie md README.md

# Set theme and render
ellie theme set dark
ellie md documentation.md

# Works with both extensions
ellie md notes.markdown
```

## 🎨 Theme Support

| Theme Mode | Description |
|------------|-------------|
| `dark` | Optimized for dark terminals |
| `light` | Optimized for light terminals |
| `auto` | Automatically detects terminal |

## 💻 Code Examples

### Go Code
```go
func renderMarkdown(content, theme string) (string, error) {
    return glamour.Render(content, theme)
}
```

### Python Code  
```python
def hello_ellie():
    print("Hello from Ellie CLI!")
    return "markdown rendered successfully"
```

## ✅ Benefits

- Seamless integration with existing Ellie CLI
- No additional dependencies (uses existing Glamour)
- Consistent with Ellie's theme system
- Beautiful output for documentation review

---

**Powered by Glamour & Ellie CLI** 🎯