# Product Development Cradle to Grave - Part 2

This series of articles is designed as a tutorial for how to write a portable full stack service using modern techniques. I will take you on a journey from concept through deployment. The series will be detailed and therefore is closer to a book than a bunch of blog posts. Hopefully we'll all learn something along the way. I'll strive to make the series useful for software engineers at any point in their career. I will incorporate how I approach doing product development which has worked well when working with teams over my career.

Let's get started on the tasks listed in the estimate.

# Setup Development Environment

### Clone/Create Repository

First, we need a directory and initialize it as a Git repository. If you intend to use GitHub or Bitbucket, it's easiest to create a repository in the cloud and clone it locally. I'll skip the Git tutorial and assume you've created a directory and that has been initialized. If you wish, you can just clone this repository and use it:

> git clone https://github.com/joebologna/rockyourslotsoff.git

## Install Compilers and VS Code

We'll be using Go, C++ and VS Code on MacOS. The procedure for other OSes will be different, but I won't cover those details here for now.

Asking ChatGPT for the procedure provides a good overview, let's use it:

Prompt: How do i install compilers on mac os?

Response:

To install compilers on macOS, you typically have a few options depending on the programming languages you want to use. Here are the steps for some common compilers:

### 1. **Xcode Command Line Tools (C/C++/Objective-C)**

The easiest way to get a C/C++ compiler on macOS is to install the Xcode Command Line Tools.

1. Open the Terminal application.
2. Run the following command:
   ```bash
   xcode-select --install
   ```
3. A dialog will appear prompting you to install the tools. Click "Install" and follow the instructions.

### 2. **Homebrew (Various Languages)**

Homebrew is a package manager for macOS that makes it easy to install various software, including compilers.

1. If you don't have Homebrew installed, open Terminal and run:
   ```bash
   /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
   ```
2. Once Homebrew is installed, you can install various compilers. For example:
   - For GCC (GNU Compiler Collection):
     ```bash
     brew install gcc
     ```
   - For other languages, you can search for the package:
     ```bash
     brew search <language>
     ```

### 3. **Installing Python (Python Compiler)**

If you need Python, you can install it via Homebrew:

```bash
brew install python
```

### 4. **Installing Go (Go Compiler)**

To install Go:

```bash
brew install go
```

### 5. **Installing VS Code**

You can install VSCode using Homebrew:

```bash
brew install --cask visual-studio-code
```

After everything is setup, you will want to install some extensions. When VSCode detects you are using Go or C++, it will prompt to install extensions. You should install them.

## Write Slot Machine "Business Logic" using TDD

At this point we're ready to start writing some code. We're writing code for Go and C++, so we'll just use a subdirectory for each.

```
mkdir -p src/Go src/Cpp
```

### Setup Go Workspace

VSCode really likes to use workspaces, especially for Go. The simplest way to setup a workspace is to just create a hello.go program, open the directory in VSCode and use the menu options `File/Save Workspace As...` function.

```go
cat <<\EOF >src/Go/main.go
package main

import "fmt"

func main() {
        fmt.Println("Hello World!")
}
EOF
```

```console
cd src/Go
go mod init slots
go mod tidy
go run .
```

> Hello World!

```console
code .
```

In VSCode use the `File/Save Workspace As...` function to create the workspace.

### Setup C++ Workspace

VSCode really likes to use workspaces, especially for Go. The simplest way to setup a workspace is to just create a hello.go program, open the directory in VSCode and use the menu options `File/Save Workspace As...` function.

```go
cd -
cat <<\EOF >src/Cpp/main.cpp
#include <iostream>

int main() {
  std::cout << "Hello World!" << std::endl;
  return 0;
}
EOF
```

```console
cd src/Cpp
c++ main.cpp
./a.out
rm a.out
```

> Hello World!

```console
code .
```

In VSCode use the `File/Save Workspace As...` function to create the workspace.
