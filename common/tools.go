package common

import (
	"time"

	"github.com/tacheraSasi/ellie/types"
)

// ServerInitSession represents a server environment setup session
type ServerInitSession struct {
	OS             string
	ServerName     string
	SuccessCount   int
	SkippedCount   int
	FailedCount    int
	StartTime      time.Time
	EndTime        time.Time
	InstalledTools []string
	FailedTools    []string
	Framework      string
}

// Framework represents a server framework with required tools and setup commands
type Framework struct {
	Name          string
	Description   string
	RequiredTools []string
	SetupCommands []string
}

// frameworks is a map of available server frameworks
var Frameworks = map[string]Framework{
	"general": {
		Name:          "General",
		Description:   "General server setup with common tools",
		RequiredTools: []string{"Git", "Node.js", "Python", "Docker", "NGINX"},
		SetupCommands: []string{},
	},
	"laravel": {
		Name:          "Laravel",
		Description:   "PHP-based Laravel framework server",
		RequiredTools: []string{"PHP", "Composer", "Node.js", "MySQL", "NGINX"},
		SetupCommands: []string{
			"composer global require laravel/installer",
			"laravel new project --no-interaction",
		},
	},
	"nodejs": {
		Name:          "Node.js",
		Description:   "Node.js server with Express",
		RequiredTools: []string{"Node.js", "Yarn"},
		SetupCommands: []string{
			"mkdir project && cd project",
			"yarn init -y",
			"yarn add express",
		},
	},
	"django": {
		Name:          "Django",
		Description:   "Python-based Django framework server",
		RequiredTools: []string{"Python", "PostgreSQL"},
		SetupCommands: []string{
			"pip install django psycopg2-binary",
			"django-admin startproject project",
		},
	},
	"rails": {
		Name:          "Ruby on Rails",
		Description:   "Ruby-based web framework",
		RequiredTools: []string{"Ruby", "Node.js", "Yarn", "PostgreSQL"},
		SetupCommands: []string{
			"gem install rails",
			"rails new project --database=postgresql",
		},
	},
	"flask": {
		Name:          "Flask",
		Description:   "Lightweight Python framework",
		RequiredTools: []string{"Python", "PostgreSQL"},
		SetupCommands: []string{
			"pip install Flask gunicorn psycopg2-binary",
			"mkdir project && cd project",
			"echo 'from flask import Flask\napp = Flask(__name__)\n@app.route(\"/\")\ndef hello():\n    return \"Hello, World!\"' > app.py",
		},
	},
	"spring_boot": {
		Name:          "Spring Boot",
		Description:   "Java-based enterprise framework",
		RequiredTools: []string{"Java", "Maven"},
		SetupCommands: []string{
			"curl https://start.spring.io/starter.tgz -d dependencies=web -d type=maven-project -d baseDir=project | tar -xzvf -",
		},
	},
	"nextjs": {
		Name:          "Next.js",
		Description:   "React framework for server-rendered apps",
		RequiredTools: []string{"Node.js", "Yarn"},
		SetupCommands: []string{
			"yarn create next-app project",
		},
	},
	"gin": {
		Name:          "Gin",
		Description:   "Go web framework",
		RequiredTools: []string{"Go"},
		SetupCommands: []string{
			"mkdir project && cd project",
			"go mod init project",
			"go get github.com/gin-gonic/gin",
		},
	},
	"nuxtjs": {
		Name:          "Nuxt.js",
		Description:   "Vue.js framework for server-side rendering",
		RequiredTools: []string{"Node.js", "Yarn"},
		SetupCommands: []string{
			"yarn create nuxt-app project",
		},
	},
	"fastapi": {
		Name:          "FastAPI",
		Description:   "Modern Python API framework",
		RequiredTools: []string{"Python", "PostgreSQL"},
		SetupCommands: []string{
			"pip install fastapi uvicorn psycopg2-binary",
			"mkdir project && cd project",
			"echo 'from fastapi import FastAPI\napp = FastAPI()\n\n@app.get(\"/\")\ndef read_root():\n    return {\"Hello\": \"World\"}' > main.py",
		},
	},
}

var Tools []types.DevTool = []types.DevTool{
	{
		Name:           "Git",
		Description:    "Version control system",
		CheckCmd:       "git --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":     "brew install git",
			"linux":   "sudo apt-get install git -y",
			"windows": "choco install git -y",
		},
		Configure: map[string]string{
			"common": "git config --global core.autocrlf input && git config --global init.defaultBranch main",
		},
	},
	{
		Name:           "Node.js",
		Description:    "JavaScript runtime",
		CheckCmd:       "node --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":     "brew install node",
			"linux":   "curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash - && sudo apt-get install -y nodejs",
			"windows": "choco install nodejs-lts",
		},
	},
	{
		Name:           "Docker",
		Description:    "Containerization platform",
		CheckCmd:       "docker --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install --cask docker",
			"linux":   "curl -fsSL https://get.docker.com | sh",
			"windows": "choco install docker-desktop",
		},
	},
	{
		Name:           "Go",
		Description:    "Go programming language",
		CheckCmd:       "go version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install go",
			"linux":   "sudo apt install golang -y",
			"windows": "choco install golang",
		},
	},
	{
		Name:           "Python",
		Description:    "Python programming language",
		CheckCmd:       "python3 --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":     "brew install python",
			"linux":   "sudo apt install python3 -y",
			"windows": "choco install python",
		},
	},
	{
		Name:           "VS Code",
		Description:    "Code editor",
		CheckCmd:       "code --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install --cask visual-studio-code",
			"linux":   "sudo snap install --classic code",
			"windows": "choco install vscode",
		},
	},
	{
		Name:           "Yarn",
		Description:    "Modern package manager",
		CheckCmd:       "yarn --version",
		DefaultInstall: true,
		Install: map[string]string{
			"mac":     "brew install yarn",
			"linux":   "npm install -g yarn",
			"windows": "choco install yarn",
		},
	},
	{
		Name:           "PostgreSQL",
		Description:    "Relational database",
		CheckCmd:       "psql --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install postgresql",
			"linux":   "sudo apt install postgresql postgresql-contrib -y",
			"windows": "choco install postgresql",
		},
		Configure: map[string]string{
			"linux": "sudo systemctl enable postgresql && sudo systemctl start postgresql",
			"mac":   "brew services start postgresql",
		},
	},
	{
		Name:           "Redis",
		Description:    "In-memory data store",
		CheckCmd:       "redis-cli --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install redis",
			"linux":   "sudo apt install redis -y",
			"windows": "choco install redis",
		},
		Configure: map[string]string{
			"linux": "sudo systemctl enable redis && sudo systemctl start redis",
			"mac":   "brew services start redis",
		},
	},
	{
		Name:           "AWS CLI",
		Description:    "Amazon Web Services CLI",
		CheckCmd:       "aws --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install awscli",
			"linux":   "curl 'https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip' -o awscliv2.zip && unzip awscliv2.zip && sudo ./aws/install",
			"windows": "choco install awscli",
		},
	},
	{
		Name:           "Terraform",
		Description:    "Infrastructure as code tool",
		CheckCmd:       "terraform --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install terraform",
			"linux":   "sudo apt install terraform -y",
			"windows": "choco install terraform",
		},
	},
	{
		Name:           "kubectl",
		Description:    "Kubernetes cluster manager",
		CheckCmd:       "kubectl version --client",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install kubectl",
			"linux":   "sudo apt install kubectl -y",
			"windows": "choco install kubernetes-cli",
		},
	},
	{
		Name:           "Helm",
		Description:    "Kubernetes package manager",
		CheckCmd:       "helm version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install helm",
			"linux":   "curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash",
			"windows": "choco install kubernetes-helm",
		},
	},
	{
		Name:           "NGINX",
		Description:    "Web server and reverse proxy",
		CheckCmd:       "nginx -v",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install nginx",
			"linux":   "sudo apt install nginx -y",
			"windows": "choco install nginx",
		},
	},
	{
		Name:           "GitHub CLI",
		Description:    "GitHub command-line tool",
		CheckCmd:       "gh --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install gh",
			"linux":   "sudo apt install gh -y",
			"windows": "choco install gh",
		},
	},
	{
		Name:           ".NET SDK",
		Description:    ".NET development platform",
		CheckCmd:       "dotnet --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install dotnet",
			"linux":   "wget https://dot.net/v1/dotnet-install.sh -O dotnet-install.sh && chmod +x ./dotnet-install.sh && ./dotnet-install.sh",
			"windows": "choco install dotnet-sdk",
		},
	},
	{
		Name:           "Java",
		Description:    "OpenJDK development kit",
		CheckCmd:       "javac --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install openjdk",
			"linux":   "sudo apt install openjdk-17-jdk -y",
			"windows": "choco install openjdk",
		},
	},
	{
		Name:           "PHP",
		Description:    "PHP runtime",
		CheckCmd:       "php --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install php",
			"linux":   "sudo apt install php -y",
			"windows": "choco install php",
		},
	},
	{
		Name:           "Ansible",
		Description:    "Configuration management",
		CheckCmd:       "ansible --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install ansible",
			"linux":   "sudo apt install ansible -y",
			"windows": "choco install ansible",
		},
	},
	{
		Name:           "Vagrant",
		Description:    "Virtual machine manager",
		CheckCmd:       "vagrant --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install vagrant",
			"linux":   "sudo apt install vagrant -y",
			"windows": "choco install vagrant",
		},
	},
	{
		Name:           "Prettier",
		Description:    "Code formatter",
		CheckCmd:       "prettier --version",
		DefaultInstall: true,
		Install: map[string]string{
			"common": "npm install -g prettier",
		},
	},
	{
		Name:           "ESLint",
		Description:    "JavaScript linter",
		CheckCmd:       "eslint --version",
		DefaultInstall: true,
		Install: map[string]string{
			"common": "npm install -g eslint",
		},
	},
	{
		Name:           "Rust",
		Description:    "Rust programming language",
		CheckCmd:       "rustc --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh",
			"linux":   "curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh",
			"windows": "choco install rust",
		},
	},
	{
		Name:           "TypeScript",
		Description:    "TypeScript compiler",
		CheckCmd:       "tsc --version",
		DefaultInstall: true,
		Install: map[string]string{
			"common": "npm install -g typescript",
		},
	},
	{
		Name:           "MongoDB",
		Description:    "NoSQL database",
		CheckCmd:       "mongod --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew tap mongodb/brew && brew install mongodb-community",
			"linux":   "sudo apt install mongodb -y",
			"windows": "choco install mongodb",
		},
		Configure: map[string]string{
			"mac":   "brew services start mongodb/brew/mongodb-community",
			"linux": "sudo systemctl enable mongod && sudo systemctl start mongod",
		},
	},
	{
		Name:           "Composer",
		Description:    "PHP dependency manager",
		CheckCmd:       "composer --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install composer",
			"linux":   "sudo apt install composer -y",
			"windows": "choco install composer",
		},
	},
	{
		Name:           "MySQL",
		Description:    "MySQL database server",
		CheckCmd:       "mysql --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install mysql",
			"linux":   "sudo apt install mysql-server -y",
			"windows": "choco install mysql",
		},
	},
	{
		Name:           "Ruby",
		Description:    "Ruby programming language",
		CheckCmd:       "ruby --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install ruby",
			"linux":   "sudo apt install ruby-full -y",
			"windows": "choco install ruby",
		},
	},
	{
		Name:           "Maven",
		Description:    "Java project management and build tool",
		CheckCmd:       "mvn --version",
		DefaultInstall: false,
		Install: map[string]string{
			"mac":     "brew install maven",
			"linux":   "sudo apt install maven -y",
			"windows": "choco install maven",
		},
	},
}
