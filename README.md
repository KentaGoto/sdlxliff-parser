# SDLXLIFF Parser

A simple SDLXLIFF parser written in Go that extracts essential information from an SDLXLIFF file.

## Features
- Extracts the following information from an SDLXLIFF file:
  - Original
  - Datatype
  - Source language
  - Target language
  - Displays the source and target segments of the file.

## Requirements

- [Go](https://golang.org/doc/install) installed on your system.

## Usage

1. Save the provided Go code as `sdlxliff_parser.go` in your desired directory.
2. Open the terminal and navigate to the directory where you saved the `sdlxliff_parser.go` file.
3. Run the command `go build sdlxliff_parser.go` to build the executable.
4. Run the compiled executable using the command `./sdlxliff_parser <path-to-your-sdlxliff-file>` on Unix-based systems or `sdlxliff_parser.exe <path-to-your-sdlxliff-file>` on Windows.

## Example

To parse an SDLXLIFF file named `sample.sdlxliff`, follow these steps:

1. Save the `sdlxliff_parser.go` code in a directory.
2. Open the terminal and navigate to the directory where you saved the `sdlxliff_parser.go` file.
3. Run the command `go build sdlxliff_parser.go`.
4. Run the compiled executable using the command `./sdlxliff_parser sample.sdlxliff` on Unix-based systems or `sdlxliff_parser.exe sample.sdlxliff` on Windows.

After running the parser, the extracted information will be displayed in the terminal.

## License
MIT

## Author
Kenta Goto
