// Global variables
let currentInstallMethod = 'brew';
let currentDemo = 'git';

// DOM elements
const mobileToggle = document.getElementById('mobileToggle');
const installBtn = document.getElementById('installBtn');
const demoBtn = document.getElementById('demoBtn');
const copyBtn = document.getElementById('copyBtn');
const typewriter = document.getElementById('typewriter');
const terminalOutput = document.getElementById('terminalOutput');
const codeTitle = document.getElementById('codeTitle');
const codeContent = document.getElementById('codeContent');
const demoOutput = document.getElementById('demoOutput');

// Typewriter effect
const commands = [
  'ellie git commit',
  'ellie todo add "Fix login bug"',
  'ellie switch api',
  'ellie start-day',
  'ellie network-status'
];

let currentCommandIndex = 0;
let currentCharIndex = 0;
let isDeleting = false;

function typewriterEffect() {
  const currentCommand = commands[currentCommandIndex];
  
  if (isDeleting) {
    typewriter.textContent = currentCommand.substring(0, currentCharIndex - 1);
    currentCharIndex--;
  } else {
    typewriter.textContent = currentCommand.substring(0, currentCharIndex + 1);
    currentCharIndex++;
  }
  
  if (!isDeleting && currentCharIndex === currentCommand.length) {
    setTimeout(() => {
      isDeleting = true;
    }, 2000);
  } else if (isDeleting && currentCharIndex === 0) {
    isDeleting = false;
    currentCommandIndex = (currentCommandIndex + 1) % commands.length;
  }
  
  const typingSpeed = isDeleting ? 50 : 100;
  setTimeout(typewriterEffect, typingSpeed);
}

// Terminal output examples
const terminalOutputs = {
  'ellie git commit': `📝 Conventional Commit Builder
─────────────────────────────
🔧 Type: feat
🎯 Scope: auth
📌 Description: Add OAuth2 support
✅ Successfully committed and pushed!`,
  
  'ellie todo add "Fix login bug"': `✅ Added todo #1: Fix login bug`,
  
  'ellie switch api': `✅ Switched to project 'api'
📂 /Users/dev/projects/api
🔧 Starting development server...`,
  
  'ellie start-day': `🌅 Starting your development day...
✅ Opening applications...  
✅ Starting services...
✅ Checking Git repositories...
🚀 Your development environment is ready!`,
  
  'ellie network-status': `🌐 Network Status
─────────────────
📡 WiFi: Connected
🌍 Internet: Active
⚡ Speed: 150 Mbps
📍 IP: 192.168.1.100`
};

function updateTerminalOutput() {
  const currentCommand = commands[currentCommandIndex];
  if (terminalOutputs[currentCommand]) {
    terminalOutput.textContent = terminalOutputs[currentCommand];
  }
}

// Installation methods
const installMethods = {
  brew: {
    title: 'Homebrew Installation',
    code: `# One-time setup
brew tap beyondEllie/ellie

# Install CLI
brew install ellie`
  },
  intel: {
    title: 'macOS Intel Installation',
    code: `curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_amd64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_amd64.tar.gz`
  },
  silicon: {
    title: 'macOS Apple Silicon Installation',
    code: `curl -O -L https://github.com/tacheraSasi/ellie/releases/download/0.0.91/ellie_mac_arm64.tar.gz
sudo tar -C /usr/local/bin -xzvf ellie_mac_arm64.tar.gz`
  },
  windows: {
    title: 'Windows Installation',
    code: `<p>1. Download the executable:</p>\
<a href="./downloads/ellie.exe" class="btn-primary mb-4" download>🚀 Download ellie.exe</a>\
<p class="mt-4">2. Open a terminal (PowerShell or Command Prompt) and navigate to the directory where you saved the file.</p>\
<p>3. Run the following command to get started:</p>\
<code>.\\ellie.exe --help</code>`
  },
  linux: {
    title: 'Linux Installation',
    code: `<p>1. Download the executable:</p>\
<a href="./downloads/ellie_linux" class="btn-primary mb-4" download>🚀 Download ellie_linux</a>\
<p class="mt-4">2. Open a terminal and navigate to the directory where you saved the file.</p>\
<p>3. Make the file executable:</p>\
<code>chmod +x ./ellie_linux</code>\
<p class="mt-4">4. Run the following command to get started:</p>\
<code>./ellie_linux --help</code>`
  }
};

// Demo content
const demoContent = {
  git: `$ ellie git commit
📝 Conventional Commit Builder
─────────────────────────────
🔧 Type (feat, fix, docs, style, refactor, perf, test, chore, revert) ➜ feat
🎯 Scope (optional) ➜ auth
📌 Description ➜ Add OAuth2 support
💬 Body (optional):
◎ Press Enter twice to finish:
Implemented Google and GitHub providers
Updated configuration schema

💥 Breaking change? (Y/n) ➜ y
📣 Breaking change details ➜ Changed config format
🔗 Issue number (optional) ➜ 42

✨ Commit Preview:
──────────────────
feat(auth): Add OAuth2 support

Implemented Google and GitHub providers
Updated configuration schema

BREAKING CHANGE: Changed config format

Refs #42
──────────────────
✅ Successfully committed and pushed!`,

  todo: `$ ellie todo add "Fix login bug" api high
✅ Added todo #1: Fix login bug [api] 🔴 High

$ ellie todo list
Your todos:
📁 api:
  ❌ #1: Fix login bug 🔴 High
  ✅ #2: Update docs 🟡 Medium
📁 frontend:
  ❌ #3: Responsive design 🟢 Low

$ ellie todo complete 1
✅ Completed todo #1: Fix login bug`,

  projects: `$ ellie project add api ~/projects/api "API Service" backend,nodejs
✅ Added project 'api'

$ ellie project list
📁 api
   📝 API Service
   📂 /Users/me/projects/api
   🏷️  backend, nodejs
   ⏰ Last used: 2 hours ago

📁 frontend
   📝 React Dashboard
   📂 /Users/me/projects/dashboard
   🏷️  frontend, react
   ⏰ Last used: 1 day ago

$ ellie switch api
✅ Switched to project 'api'
📂 /Users/me/projects/api
🔧 Starting development server...`
};

// Event listeners
document.addEventListener('DOMContentLoaded', function() {
  // Start typewriter effect
  setTimeout(typewriterEffect, 1000);
  
  // Update terminal output periodically
  setInterval(updateTerminalOutput, 4000);
  
  // Mobile menu toggle
  if (mobileToggle) {
    mobileToggle.addEventListener('click', function() {
      const navLinks = document.querySelector('.nav-links');
      navLinks.classList.toggle('active');
    });
  }
  
  // Install button scroll
  if (installBtn) {
    installBtn.addEventListener('click', function() {
      document.getElementById('installation').scrollIntoView({
        behavior: 'smooth'
      });
    });
  }
  
  // Demo button scroll
  if (demoBtn) {
    demoBtn.addEventListener('click', function() {
      document.querySelector('.demo').scrollIntoView({
        behavior: 'smooth'
      });
    });
  }
  
  // Copy button functionality
  if (copyBtn) {
    copyBtn.addEventListener('click', function() {
      const code = codeContent.textContent;
      navigator.clipboard.writeText(code).then(function() {
        const originalText = copyBtn.querySelector('.copy-text').textContent;
        copyBtn.querySelector('.copy-text').textContent = 'Copied!';
        copyBtn.style.background = 'var(--success)';
        
        setTimeout(() => {
          copyBtn.querySelector('.copy-text').textContent = originalText;
          copyBtn.style.background = 'var(--primary)';
        }, 2000);
      });
    });
  }
  
  // Installation method switching
  const installMethodElements = document.querySelectorAll('.install-method');
  installMethodElements.forEach(method => {
    method.addEventListener('click', function() {
      const methodType = this.getAttribute('data-method');
      switchInstallMethod(methodType);
    });
  });
  
  // Demo tab switching
  const demoTabs = document.querySelectorAll('.demo-tab');
  demoTabs.forEach(tab => {
    tab.addEventListener('click', function() {
      const demoType = this.getAttribute('data-demo');
      switchDemo(demoType);
    });
  });
  
  // Smooth scrolling for navigation links
  const navLinks = document.querySelectorAll('.nav-links a[href^="#"]');
  navLinks.forEach(link => {
    link.addEventListener('click', function(e) {
      e.preventDefault();
      const targetId = this.getAttribute('href').substring(1);
      const targetElement = document.getElementById(targetId);
      if (targetElement) {
        targetElement.scrollIntoView({
          behavior: 'smooth',
          block: 'start'
        });
      }
    });
  });
  
  // Navbar background on scroll
  window.addEventListener('scroll', function() {
    const navbar = document.querySelector('.navbar');
    if (window.scrollY > 100) {
      navbar.style.background = 'rgba(255, 255, 255, 0.98)';
    } else {
      navbar.style.background = 'rgba(255, 255, 255, 0.95)';
    }
  });
  
  // Intersection Observer for animations
  const observerOptions = {
    threshold: 0.1,
    rootMargin: '0px 0px -50px 0px'
  };
  
  const observer = new IntersectionObserver(function(entries) {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.style.opacity = '1';
        entry.target.style.transform = 'translateY(0)';
      }
    });
  }, observerOptions);
  
  // Observe feature cards
  const featureCards = document.querySelectorAll('.feature-card');
  featureCards.forEach(card => {
    card.style.opacity = '0';
    card.style.transform = 'translateY(30px)';
    card.style.transition = 'all 0.6s ease';
    observer.observe(card);
  });
  
  // Observe steps
  const steps = document.querySelectorAll('.step');
  steps.forEach((step, index) => {
    step.style.opacity = '0';
    step.style.transform = 'translateY(30px)';
    step.style.transition = `all 0.6s ease ${index * 0.2}s`;
    observer.observe(step);
  });
});

function switchInstallMethod(method) {
  currentInstallMethod = method;

  // Update active state
  document.querySelectorAll('.install-method').forEach(m => {
    m.classList.remove('active');
  });
  document.querySelector(`[data-method="${method}"]`).classList.add('active');

  // Update code content
  const methodData = installMethods[method];
  if (methodData) {
    codeTitle.textContent = methodData.title;
    if (method === 'windows' || method === 'linux') {
      copyBtn.style.display = 'none';
      codeContent.innerHTML = methodData.code;
    } else {
      copyBtn.style.display = 'flex';
      codeContent.innerHTML = `<code>${methodData.code}</code>`;
    }
  }
}

function switchDemo(demo) {
  currentDemo = demo;
  
  // Update active state
  document.querySelectorAll('.demo-tab').forEach(tab => {
    tab.classList.remove('active');
  });
  document.querySelector(`[data-demo="${demo}"]`).classList.add('active');
  
  // Update demo content with typing effect
  if (demoOutput) {
    demoOutput.textContent = '';
    typeText(demoContent[demo], demoOutput, 20);
  }
}

function typeText(text, element, speed = 50) {
  let index = 0;
  element.textContent = '';
  
  function type() {
    if (index < text.length) {
      element.textContent += text.charAt(index);
      index++;
      setTimeout(type, speed);
    }
  }
  
  type();
}

// Initialize demo content
document.addEventListener('DOMContentLoaded', function() {
  setTimeout(() => {
    if (demoOutput) {
      typeText(demoContent.git, demoOutput, 15);
    }
  }, 500);
});