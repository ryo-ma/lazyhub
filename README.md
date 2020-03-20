# lazyhub

[![GoDoc](https://godoc.org/github.com/jroimartin/gocui?status.svg)](https://godoc.org/github.com/jroimartin/gocui)

:octocat: lazyhub - Terminal UI Client for GitHub using gocui.

<img src="https://user-images.githubusercontent.com/6661165/77137331-c6c7e980-6ab0-11ea-9b33-f663bee6ad81.gif" width="1000">

# Features

* Check the trending repositories on GitHub today
* Search repositories
* Read a README
* Copy the clone command to clipboard
* Open the repository page on your browser


# Install

```
go get -u github.com/ryo-ma/lazyhub
```

# Usage

Run the following command.

```
lazyhub
```

# Keys

* <kbd>j</kbd> / <kbd>DownArrow(↓)</kbd>
Move down a line
* <kbd>k</kbd> / <kbd>DownUp(↑)</kbd>
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

# LICENSE

Apache LICENSE 2.0

[LICENSE](./LICENSE)
