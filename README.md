![logo](https://user-images.githubusercontent.com/6661165/78040587-9cc4d000-73aa-11ea-9710-567e714bdf59.png)

# lazyhub

[![GoDoc](https://godoc.org/github.com/jroimartin/gocui?status.svg)](https://godoc.org/github.com/jroimartin/gocui)
---
:octocat: lazyhub - Terminal UI Client for GitHub using gocui.

# Demo

![demo](https://user-images.githubusercontent.com/6661165/77839109-f5cb1300-71b4-11ea-886d-924e6efe1b71.gif)

# Features

* 🚀Check the trending repositories on GitHub today
* 🔍Search repositories
* 📘Read the README
* 📄Copy the clone command to clipboard
* 💻Open the repository page on your browser


# Install

Using brew

```
brew tap ryo-ma/lazyhub
brew install lazyhub
```

Using go get

```
go get -u github.com/ryo-ma/lazyhub
```

# Binary Download

[Binary releases are available](https://github.com/ryo-ma/lazyhub/releases/tag/v0.0.1)

# Usage

Run the following command.

```
lazyhub
```

# Keys

* <kbd>j</kbd> / <kbd>DownArrow(↓)</kbd>
Move down a line
* <kbd>k</kbd> / <kbd>UpArrow(↑)</kbd>
Move up a line
* <kbd>q</kbd> / <kbd>CTRL+C</kbd>
Quit
* <kbd>CTRL+D</kbd>
Move down 5 lines
* <kbd>CTRL+U</kbd>
Move up 5 lines
* <kbd>x</kbd>
Back panel
* <kbd>Enter</kbd> / <kbd>r</kbd>
Open the README
* <kbd>c</kbd>
Copy the clone command to clipboard
* <kbd>o</kbd>
Open the repository page on your browser

# Using API

* https://github.com/huchenme/github-trending-api

# LICENSE

Apache LICENSE 2.0

[LICENSE](./LICENSE)
