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
```

Let's start VSCode from the command line, more about why to do this later.

```console
code .
```

In VSCode use the `File/Save Workspace As...` function to create the workspace. Then run the program using `Run/Start Debugging` or press F5. You may be prompted to install some tools and a popup may occur from MacOS asking you to let VSCode run the developer tools, authenticate to allow it. If successful, the debug output window will display:

```
Starting: /Users/joeb/go/bin/dlv dap --listen=127.0.0.1:61338 --log-dest=3 from /Users/joeb/Projects/2024/rockyourslotsoff/src/Go
DAP server listening at: 127.0.0.1:61338
Type 'dlv help' for list of commands.
Hello World!
Process 48928 has exited with status 0
Detaching
```

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

Let's start VSCode from the command line, more about why to do this later.

```console
cd src/Cpp
code .
```

In VSCode use the `File/Save Workspace As...` function to create the workspace. This will be a second workspace, we'll be using a lot of them.

C++ is a little more complicated to compile and build than Go. Install cmake using homebrew, we'll need to use it now.

Let's create the file cmake needs:

```console
tee CMakeLists.txt <<\EOF
cmake_minimum_required(VERSION 3.10)
project(HelloWorld)

set(CMAKE_CXX_STANDARD 11)

add_executable(HelloWorld main.cpp)
EOF
```

Create the build directory and let cmake do it's magic:

```console
mkdir build
cd build
cmake ..
```

Output:

```console
-- The C compiler identification is AppleClang 15.0.0.15000309
-- The CXX compiler identification is AppleClang 15.0.0.15000309
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
-- Check for working C compiler: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/cc - skipped
-- Detecting C compile features
-- Detecting C compile features - done
-- Detecting CXX compiler ABI info
-- Detecting CXX compiler ABI info - done
-- Check for working CXX compiler: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/c++ - skipped
-- Detecting CXX compile features
-- Detecting CXX compile features - done
-- Configuring done (2.9s)
-- Generating done (0.0s)
-- Build files have been written to: /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
```

We can build the binary now:

```console
make
[ 50%] Building CXX object CMakeFiles/HelloWorld.dir/main.cpp.o
[100%] Linking CXX executable HelloWorld
[100%] Built target HelloWorld
```

Let's run it:

```console
./HelloWorld
Hello World!
```

We're all setup to start writing some code or are we?

### Using TDD with Go

As I'm mentioned we'll be using Test Driven Development (TDD). If you want to get formally trained in the process you can find a lot of material. For now, we'll just be pragmatic about it.

TDD is all about building stuff in small increments. If you have studied `Software Craftsmanship`, as taught by **Uncle Bob** for decades, you'll be well prepared. He's easy to find on YouTube, for example [Uncle Bob on TDD](https://youtu.be/GvAzrC6-spQ). James Grenning has a very comprehensive training program and some [YouTube videos of workshops he's given](https://www.youtube.com/results?search_query=james+grenning+tdd).

Go is particularly easy to unit test. Let's create a skeleton for the "virtual slot machine" and test it:

```console
mkdir vslot
cd vslot
tee -a vslot.go <<\EOF
package vslot

func Spin() [3]int {
	return [3]int{1, 2, 3}
}
EOF
```

```console
tee -a vslot_test.go <<\EOF
package vslot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	a := assert.New(t)
	a.Equal([3]int{1, 2, 3}, Spin())
}
EOF
```

We need to get the modules before this will run:

```console
cd ..
go mod tidy
cd vslot
```

Output:

```console
go: finding module for package github.com/stretchr/testify/assert
go: found github.com/stretchr/testify/assert in github.com/stretchr/testify v1.9.0
```

Now test it:

```console
go test ./...
```

Output:

```console
ok      slots/vslot     0.329s
```

We now have a basic framework for writing the spin function.

### Using TDD with C++

Unit testing C++ is more involved than Go, but let's use a new method that only requires a header file.

The GitHub project is [doctest](https://github.com/doctest/doctest).

Let's just overwrite the CMakeLists.txt file with the unit test included:

```console
tee CMakeLists.txt <<\EOF
cmake_minimum_required(VERSION 3.10)
project(HelloWorld)

set(CMAKE_CXX_STANDARD 11)

set(FILE_URL "https://raw.githubusercontent.com/doctest/doctest/master/doctest/doctest.h")
set(LOCAL_FILE "${CMAKE_CURRENT_SOURCE_DIR}/doctest.h")

if(NOT EXISTS ${LOCAL_FILE})
    message(STATUS "Downloading file from ${FILE_URL}...")
    execute_process(
        COMMAND curl -s -L -o ${LOCAL_FILE} ${FILE_URL}
        RESULT_VARIABLE result
    )

    if(result)
        message(FATAL_ERROR "Failed to download file from ${FILE_URL}")
    else()
        message(STATUS "File downloaded successfully.")
    endif()
else()
    message(STATUS "File already exists: ${LOCAL_FILE}")
endif()

enable_testing()
add_subdirectory(tests)

add_executable(HelloWorld main.cpp)
EOF
```

Now add the test:

```console
mkdir tests
tee tests/CMakeLists.txt <<\EOF
add_executable(test_vslot ../vslot.cpp)
add_test(NAME TestSpin COMMAND test_vslot)
EOF
```

Update the Makefile and run tests:

```console
cmake .. && make && ctest
```

Output:

```console
-- The C compiler identification is AppleClang 15.0.0.15000309
-- The CXX compiler identification is AppleClang 15.0.0.15000309
-- Detecting C compiler ABI info
-- Detecting C compiler ABI info - done
-- Check for working C compiler: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/cc - skipped
-- Detecting C compile features
-- Detecting C compile features - done
-- Detecting CXX compiler ABI info
-- Detecting CXX compiler ABI info - done
-- Check for working CXX compiler: /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/c++ - skipped
-- Detecting CXX compile features
-- Detecting CXX compile features - done
-- File already exists: /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/doctest.h
-- Configuring done (0.9s)
-- Generating done (0.0s)
-- Build files have been written to: /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
[ 25%] Building CXX object CMakeFiles/HelloWorld.dir/main.cpp.o
[ 50%] Linking CXX executable HelloWorld
[ 50%] Built target HelloWorld
[ 75%] Building CXX object tests/CMakeFiles/test_vslot.dir/__/vslot.cpp.o
[100%] Linking CXX executable test_vslot
[100%] Built target test_vslot
Test project /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
    Start 1: TestSpin
1/1 Test #1: TestSpin .........................   Passed    0.33 sec

100% tests passed, 0 tests failed out of 1

Total Test time (real) =   0.33 sec
```

### Adding Behavior Using TDD with Go

Now that we have a basic TDD process, we can start creating more useful behavior. Let's have Spin return a random number in position 0 of the slice. This is a little tricky to test. We'll add a parameter to the Spin function to accept a seed. The same seed with return the same sequence of numbers so we can do the asserts.

Let's overwrite vslot.go:

```console
tee vslot/vslot.go <<\EOF
package vslot

import "math/rand"

func Spin(seed int64) [3]int {
	rand.Seed(seed)
	// Generate a random integer between 1 and 10
	value1 := rand.Intn(10) + 1
	return [3]int{value1, 2, 3}
}
EOF
```

Let's overwrite vslot_test.go:

```console
tee vslot/vslot_test.go <<\EOF
package vslot

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)
	// Generate a random integer between 1 and 10
	value1 := rand.Intn(10) + 1

	a.Equal([3]int{value1, 2, 3}, Spin(0))
}
EOF
```

Force running the test:

```console
cd vslot; go test -count=1 ./...; cd ..
```

Output:

```console
ok      slots/vslot     0.354s
```

We now have a functioning Spin function. It must be supplied with a random number in production. This is typically accomplished by supplying this parameter: `time.Now().UnixNano()`.

### Adding Behavior Using TDD with C++

Overwrite the vslot.cpp file:

```console
tee vslot.cpp <<\EOF
#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"
#include <array>
#include <cstdlib>

std::array<int, 3> spin(unsigned int seed) {
  std::srand(seed);
  unsigned int value1 = (std::rand() % 10) + 1;
  return {int(value1), 2, 3};
}

TEST_CASE("check if spin worked") {
  std::srand(0);
  int value1 = int((std::rand() % 10) + 1);
  std::array<int, 3> expected = {value1, 2, 3};
  std::array<int, 3> result = spin(0);
  CHECK(result == expected);
}
EOF
```

Re-run the test:

```console
cd build; make; ctest; cd ..
```

Output:

```console
Test project /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
    Start 1: TestSpin
1/1 Test #1: TestSpin .........................   Passed    0.00 sec

100% tests passed, 0 tests failed out of 1

Total Test time (real) =   0.01 sec
```

We now have a functioning Spin function. It must be supplied with a random number in production. This is typically accomplished by supplying this parameter: `std::time(nullptr)`.

### Flesh Out Full Functionality using TDD with Go

Iterate on the method above where all 3 values returned by the spin function in Go and C++ are random. The result is surprisingly simple.

Set the contents of `vslot.go` to:

```go
package vslot

import "math/rand"

func Spin(seed int64) [3]int {
	rand.Seed(seed)
	return [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}
```

Set the contents of `vslot_test.go` to:

```go
package vslot

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)

	a.Equal([3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}, Spin(0))
}
```

Re-run the tests:

```console
cd vslot; go test -count=1 ./...; cd ..
```

Output:

```console
ok      slots/vslot     0.340s
```

### Flesh Out Full Functionality using TDD with C++

Let's use a similar strategy for C++.

```console
cd build; make; ctest; cd ..
```

Output:

```console
Test project /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
    Start 1: TestSpin
1/1 Test #1: TestSpin .........................   Passed    0.00 sec

100% tests passed, 0 tests failed out of 1

Total Test time (real) =   0.01 sec
```

# Summary

The spin logic is surprisingly simple and easy to test using the right tools. We're ready to move on to the next task in Part 3.
