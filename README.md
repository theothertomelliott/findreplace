# findreplace

A quick command line tool for find/replace operations in file trees. Saves messing about with sed for simple operations.

## Installation

   go get github.com/theothertomelliott/findreplace

## Usage

	findreplace findString replaceString

The above will recursively replace all instances of `findString` with `replaceString` in the current working directory's heirarchy.

The input strings are treated as literals, wildcards and regular expressions are not supported.

## Limitations

findreplace will skip over all non-regular files (directories, symlinks, etc), but it does *not* skip hidden or binary files.
