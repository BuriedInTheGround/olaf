# olaf

A simple multi-clipboard system.

It allows you to store the (textual) contents of the system clipboard in up to
four virtual clipboards, so you can switch between several copied contents.

## Installation

On Windows and Linux, you can use [the pre-built binaries].

If your system has [Go 1.18+], you can build from source.

```
go install interrato.dev/olaf@latest
```

Help from new packagers is very welcome.

## Usage

The usage is very simple. First, you need to start `olaf`: you can do it from
the CLI or, if you prefer, by double-clicking the executable.

> [!NOTE]
> When using the CLI, you may want to run `olaf` as a background task.
> On bash-like shells, this is usually accomplished by appending the
> ampersand symbol, as in the following example.
>
> ```
> $ olaf &
> ```

Once `olaf` is running, follow these steps:

1. Copy any text using `Ctrl+c` or your preferred method.
2. Press `Ctrl+u` to save what you've copied to the virtual clipboard `u`.
3. Do your stuff.
4. Press `Alt+u` to copy the contents of the virtual clipboard `u` back to the
   system clipboard.
5. Paste with `Ctrl+v` or your preferred method.

Four virtual clipboards are currently available: `u`, `i`, `o`, and `p`.

<!-- References -->

[the pre-built binaries]: https://github.com/BuriedInTheGround/olaf/releases "GitHub releases page for olaf"
[Go 1.18+]: https://go.dev/dl "The Go programming language downloads page"
