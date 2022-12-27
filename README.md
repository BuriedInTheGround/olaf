# olaf

A simple multi-clipboard system.

It allows saving the (textual) content of the system clipboard to up-to-four
virtual clipboards.

## Installation

<!---
On Windows, Linux, and macOS, you can use [the pre-built binaries].
-->

If your system has [Go 1.18+], you can build from source.
```
go install interrato.dev/olaf@latest
```

## Usage

The usage is very simple. First, you need to start `olaf`: you can do it
from the CLI or by double-clicking the executable, if you're more comfortable
this way.

> **Note**
> When running it using the CLI you may want to run `olaf` as a background
> task. On bash-like shells, this is usually accomplished by appending the
> ampersand symbol, like shown in the following example.
> ```
> $ olaf &
> ```

Once `olaf` is running, follow these steps:
1. Copy any text with `Ctrl+c` or with your preferred method
2. Press `Ctrl+u` to save what you've copied to virtual clipboard `u`
3. Do your stuff
4. Press `Alt+u` to copy back the content from virtual clipboard `u` to system clipboard
5. Paste with `Ctrl+v` or with your preferred method

Four virtual clipboards are currently available: `u`, `i`, `o`, and `p`.

<!-- References -->

[the pre-built binaries]: https://github.com/BuriedInTheGround/olaf/releases "GitHub releases page for olaf"
[Go 1.18+]: https://go.dev/dl "The Go programming language downloads page"
