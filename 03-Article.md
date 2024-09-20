# Product Development Cradle to Grave - Part 3

This series of articles is designed as a tutorial for how to write a portable full stack service using modern techniques. I will take you on a journey from concept through deployment. The series will be detailed and therefore is closer to a book than a bunch of blog posts. Hopefully we'll all learn something along the way. I'll strive to make the series useful for software engineers at any point in their career. I will incorporate how I approach doing product development which has worked well when working with teams over my career.

Let's continue with the tasks listed in the estimate.

# Revisit Assessment

Let's see how we are doing asking the original task list. I'll add a column:

| Task                                                            | Effort | Actual |
| --------------------------------------------------------------- | ------ | ------ |
| Setup dev environment for Go, C++ and VS Code                   | 1      | 3      |
| Write slot machine "business logic" using TDD                   | 3      | 1      |
| Create API to use slot machine business logic                   | 2      | 1+?    |
| Create a simple web page<sup>[1]</sup>                          | 3      |        |
| Write design the API to invoke the business logic<sup>[2]</sup> | 2      |        |
| Write the code to format the web page                           | 2      |        |
| Write code to display msg from server<sup>[3]</sup>             | 3      |        |
| Total Effort                                                    | 16     |        |

It's difficult to assess actual assessment because I'm writing these articles at the same time. However, I think I'm on track. Notice how the actual effort to complete setup and writing logic are reversed. Also notice the `?` for creating the API. This is typical.

For the setup case, I spent less than 3 hours, however my computer was already setup so I've adjusted it to account for what I forgot about initial conditions for a typical developer starting a new project with a new stack.

It's also common for work to be completed in "chunks". A chunk of 4-6 hours is typical. Most people lose focus after a few hours. Everyone is different, but the rule of thumb I use is a function of how frequently I make mistakes. When the frequency becomes noticeable, time to take a break.

Quality is all about "yield". It's all about a personal commitment to building in quality from the ground up. TDD is a tool. **How** you do your work is a continuos activity that either creates and asset or a liability. Striving for 100% yield means work is only complete when all defects have been removed and the asset is ready for "validation". Verification is a continuous process. Escaped defects are caught by QA, which may not be sufficient in some Agile development teams. This results in the dreaded fix-on-fix cycle. I could point to the "Personal Software Process (PSP)" training I took, but everyone is different, so I'll refrain and stop the pontification. Let's get back to work.

Do I work on enhancing the API or creating the simple web page at this point? I could even work on the API. It's not necessary to follow the sequence outlined in the task list. What should I choose? It depends. Mostly on what you are capable of doing at the time. If you are focused, do harder things. If you don't have a lot of time (maybe a meeting is coming up, lunch?, beers and/or dinner soon?), these things will force you to rush.

I have 2 hours before lunch so I think I can do the web page and some of the API. This will allow me to use the "define the output first" approach. We can also use our "pair programmer" which adds some interest and otherwise less than exciting task. Doing the webpage will help me think about the API. So, let's do it!

# Your Personal Pair Programmer

It's 2024. ChatGPT is a relatively new and constantly improving tool we can use. There's no reason you can't use it creatively. I use it as my personal pair programmer. Let's face it, having a dedicated pair programmer as is central to Agile methodologies is a luxury few of us have. It's also not emphasized much in some practices.

I actually posted about using [AI as a pair programmer](https://www.linkedin.com/posts/joebologna_ai-artificialintelligence-ai-activity-7240476335037366272-O5gg?utm_source=share&utm_medium=member_desktop) on LinkedIn and got all positive reactions. (Yeah!)

## Methodology

Agile is methodology. It can be used for many things. It has "camps", SCRUM is a common camp for instance. In reality, Agile started out at "Extreme Programming" when I learned it. Over time, it was produced and adopted widely. However, if you actually study the genesis of SCRUM you will discover it was not codified by a software engineer. I suppose this is OK, however, as a practicing software engineer trained in CMMI principles from the SEI, I can spot the flaws.

Our objective is to get this product built, so I'll refrain from a methodology discussion. Suffice to say, using a pair programmer is essential to agility and yield. TDD is essential to yield. I've seen many arguments claiming that using TDD is slower than using Agile methods. Ok, fine. If yield is ignored yield, mistakes are factored out of the yield equation. This pawns off unfinished products downstream. You may actually be the guy downstream.

Suffice to say, using a AI as a pair programmer is better than nothing. Since it's just me doing this project soup to nuts, that's what I'll do.

## Prompt for a Simple Web Page

```
Prompt: write me code to display a web page like this:

[number][number][number] [spin button][reset button][add 10 credits]
balance: {number of credits}
```

Response:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Spin Game</title>
    <style>
      body {
        font-family: Arial, sans-serif;
        text-align: center;
        margin-top: 50px;
      }
      .numbers {
        font-size: 24px;
        margin-bottom: 20px;
      }
      button {
        margin: 5px;
        padding: 10px 20px;
        font-size: 16px;
      }
      .balance {
        margin-top: 20px;
        font-size: 20px;
      }
    </style>
  </head>
  <body>
    <div class="numbers">
      <span id="number1">0</span>
      <span id="number2">0</span>
      <span id="number3">0</span>
    </div>

    <button id="spinButton">Spin</button>
    <button id="resetButton">Reset</button>
    <button id="addCreditsButton">Add 10 Credits</button>

    <div class="balance">Balance: <span id="balance">0</span> credits</div>

    <script>
      let balance = 0;

      document.getElementById('spinButton').addEventListener('click', function () {
        // Generate three random numbers between 0 and 9
        document.getElementById('number1').textContent = Math.floor(Math.random() * 10);
        document.getElementById('number2').textContent = Math.floor(Math.random() * 10);
        document.getElementById('number3').textContent = Math.floor(Math.random() * 10);
      });

      document.getElementById('resetButton').addEventListener('click', function () {
        // Reset the numbers to 0
        document.getElementById('number1').textContent = 0;
        document.getElementById('number2').textContent = 0;
        document.getElementById('number3').textContent = 0;
      });

      document.getElementById('addCreditsButton').addEventListener('click', function () {
        // Add 10 credits to the balance
        balance += 10;
        document.getElementById('balance').textContent = balance;
      });
    </script>
  </body>
</html>
```

What does this look like?

![Rock Page](./rock-page.png)

Good enough to start with - that took all of 10 seconds. Nice.

## Extending the API

Now that we have some idea of the output required, we know the API needs provisions to:

- Provide the wheel values (almost done)
- Spin (done)
- Reset the balance
- Add 10 to the balance
- Provide the balance

### Go Interfaces

Let's start with the Go package. This is a good time to introduce Go interfaces. A Go interface can actually define the API for us. We will be using this interface later. Now is a good time to define it.

As an aside, you can write your own code generators in Go easily. This is because the language is well defined using something call `AST`. Feel free to look into this. It's bad-ass.

Let's pawn this off on our pair programmer :)

#### ChatGPT Interaction to Generate the Go Interface

Prompts:

1. write me a go interface that has the following capabilities:

- a spin function that accepts a seed and returns a slice of ints
- a reset function
- a update balance function that accepts an int
- a getbalance function that returns an int

2. add a construct to force implementation of the interface

Output after the 2nd prompt:

In Go, you can enforce the implementation of an interface by using a compile-time assertion. This is done by creating a variable of the interface type and assigning an instance of the struct that implements the interface to it. If the struct does not implement the interface, the code will not compile.

Here's how you can modify the previous example to include a construct that forces the implementation of the `Spinner` interface:

```go
package main

import "fmt"

// Spinner interface defines the required methods for a spinner.
type Spinner interface {
	Spin(seed int) []int          // Spin accepts a seed and returns a slice of ints
	Reset()                       // Reset resets the spinner's state
	UpdateBalance(amount int)     // UpdateBalance updates the balance by the given amount
	GetBalance() int              // GetBalance returns the current balance
}

// MySpinner struct implements the Spinner interface.
type MySpinner struct {
	balance int
}

// Ensure MySpinner implements the Spinner interface.
var _ Spinner = (*MySpinner)(nil) // This line enforces the implementation at compile time.

// Spin generates a slice of random integers based on the seed.
func (s *MySpinner) Spin(seed int) []int {
	// Example implementation: generate a slice of integers based on the seed
	return []int{seed, seed + 1, seed + 2} // Just a placeholder
}

// Reset resets the spinner's state.
func (s *MySpinner) Reset() {
	s.balance = 0 // Reset balance to 0
}

// UpdateBalance updates the balance by the given amount.
func (s *MySpinner) UpdateBalance(amount int) {
	s.balance += amount
}

// GetBalance returns the current balance.
func (s *MySpinner) GetBalance() int {
	return s.balance
}

func main() {
	// Example usage
	var spinner Spinner = &MySpinner{}
	spinner.UpdateBalance(100)
	spinResult := spinner.Spin(42)
	currentBalance := spinner.GetBalance()

	// Output the results
	fmt.Println("Spin Result:", spinResult)
	fmt.Println("Current Balance:", currentBalance)

	// Reset the spinner
	spinner.Reset()
	fmt.Println("Balance after reset:", spinner.GetBalance())
}
```

Explanation of the Enforcement:

- The line `var _ Spinner = (*MySpinner)(nil)` is a compile-time assertion that checks whether `MySpinner` implements the `Spinner` interface. If `MySpinner` does not implement all the methods defined in the `Spinner` interface, the code will not compile, and you will receive an error.
- This pattern is commonly used in Go to ensure that a struct adheres to an interface, providing a safeguard against accidental changes that might break the contract defined by the interface.

#### Evaluating the ChatGPT Response

The response has more than we bargained for. It includes a struct to hold the balance and getter and update methods to mutate it. The main() even does some unit testing, smart.

### Updating the Go API

Using this data we can now update `vslot/vslot.go` and `vslot_test.go`.

There is no compelling reason to put the interface in it's own file just yet. We'll make this refactor later when using the code generator. For now, lets just update `vslot.go`.

### TDD to the rescue

Ok, so writing the unit tests for this wasn't quite as simple as I expected. So, I'm into some refactoring. The issue is regarding the behavior of `rand`. This code fails to set is_win to true:

```go
func TestBalance(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)

	v := MyVSlot{}
	a.Equal(v.GetBalance(), 0)

	spin_result := v.Spin(0)
	is_win := spin_result == spin0 // spin0 is [2,8,8], spin_result is [5,5,4]
	a.True(is_win)

	if is_win {
		v.UpdateBalance(10)
		a.Equal(10, v.GetBalance())
	} else {
		a.FailNow("should have won")
	}
}
```

I figured this out using the Go debugger. I'm not sure why, I suspect it is because Go tests are run in parallel, but since I used the debugger to perform the test, I suspect it's something different.

Noodle... oh, crap. Globals are fine in Go tests, but `rand.Intn(10)` is called when the globals are initialized. This is wrong because the seed is not set yet. rand does not report an error, I suppose it just uses a random seed. It really doesn't matter. The globals need to be set in a setup function.

Thinking about this further, it seems initializing the random number generator should be more controlled to avoid this from creeping in at a later date. It's likely that at some time in the future the code will be modified by someone who is not using TDD, not running the unit tests or future units tests are flawed. So we should consider adding some protections. For now, let's just put it on the list. We can do this using `// REVIEW` in the code, my handy Todo+ extension will flag it.

We'll also likely need to use globals, in the tests. This should be done with care. We will use a `TestMain()` function to accomplish this.

### Refactored Go Code

At this point, it was time to reconsider what I'd done thus far. I concluded that a refactor was needed. After trying a few things to get the tests to work, it became obvious that I'd neglected to create a "New" function to create a MyVSlot object. So, I refactored by hand (instead of using AI).

This is the result:

**vslot/vslot.go**

```go
package vslot

import "math/rand"

// VSlot interface defines the required methods for implementing a virtual slot machine.
type VSlot interface {
	Spin() [3]int             // Spin accepts a seed and returns a slice of ints
	Reset()                   // Reset resets the spinner's state
	UpdateBalance(amount int) // UpdateBalance updates the balance by the given amount
	GetBalance() int          // GetBalance returns the current balance
}

// MyVSlot struct implements the Spinner interface.
type MyVSlot struct {
	balance int
}

// Ensure MySpinner implements the Spinner interface.
var _ VSlot = (*MyVSlot)(nil) // This line enforces the implementation at compile time.

func NewMyVSlot(seed int64) *MyVSlot {
	m := MyVSlot{}
	m.balance = 0
	rand.Seed(seed)
	return &m
}

// Spin generates a slice of random integers based on the seed.
func (s *MyVSlot) Spin() [3]int {
	return [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}

// Reset resets the spinner's state.
func (s *MyVSlot) Reset() {
	s.balance = 0 // Reset balance to 0
}

// UpdateBalance updates the balance by the given amount.
func (s *MyVSlot) UpdateBalance(amount int) {
	s.balance += amount
}

// GetBalance returns the current balance.
func (s *MyVSlot) GetBalance() int {
	return s.balance
}
```

**vslot/vslot_test.go**

```go
package vslot

import (
	"flag"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	spin0, spin1 [3]int
	myVSlot      *MyVSlot
)

// might need to run with -parallel=1
func TestMain(m *testing.M) {
	flag.Parse()
	setup()

	code := m.Run()

	// teardown()

	os.Exit(code)
}

func setup() {
	myVSlot = NewMyVSlot(0)
	spin0 = [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
	spin1 = [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}

func TestSpin(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)
}

func TestTwoSpins(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	spin_result = myVSlot.Spin()
	a.Equal(spin1, spin_result)
}

func TestBalance(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)
	a.Equal(myVSlot.GetBalance(), 0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	myVSlot.UpdateBalance(10)
	a.Equal(10, myVSlot.GetBalance())
}

func TestReset(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)
	a.Equal(myVSlot.GetBalance(), 0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	myVSlot.UpdateBalance(10)
	a.Equal(10, myVSlot.GetBalance())

	myVSlot.Reset()
	a.Equal(0, myVSlot.GetBalance())
}
```

Let's run the tests and check code coverage:

```console
go test -cover
```

Output:

```console
PASS
coverage: 100.0% of statements
ok  	slots/vslot	0.447s
```

Yeah! We now have a functional vslot package.

#### ChatGPT Interaction to Generate the C++ Code

Let's try to get ChatGPT to do this for us. Maybe it will just rewrite the Go code.

Prompt: convert this code to C++:

_paste vslot/vslot.go here_

Response: Here's the equivalent C++ code for your Go implementation:

```c++
#include <iostream>
#include <cstdlib>
#include <ctime>
#include <array>

// VSlot interface defines the required methods for implementing a virtual slot machine.
class VSlot {
public:
    virtual std::array<int, 3> Spin() = 0;             // Spin generates a slice of random integers
    virtual void Reset() = 0;                          // Reset resets the spinner's state
    virtual void UpdateBalance(int amount) = 0;        // UpdateBalance updates the balance by the given amount
    virtual int GetBalance() const = 0;                // GetBalance returns the current balance
    virtual ~VSlot() = default;                        // Virtual destructor for proper cleanup
};

// MyVSlot class implements the VSlot interface.
class MyVSlot : public VSlot {
private:
    int balance;

public:
    MyVSlot(int64_t seed) : balance(0) {
        std::srand(seed);
    }

    std::array<int, 3> Spin() override {
        return {std::rand() % 10 + 1, std::rand() % 10 + 1, std::rand() % 10 + 1};
    }

    void Reset() override {
        balance = 0; // Reset balance to 0
    }

    void UpdateBalance(int amount) override {
        balance += amount;
    }

    int GetBalance() const override {
        return balance;
    }
};

int main() {
    MyVSlot slotMachine(time(0)); // Initialize with current time as seed

    // Example usage
    auto result = slotMachine.Spin();
    std::cout << "Spin result: " << result[0] << " " << result[1] << " " << result[2] << std::endl;

    slotMachine.UpdateBalance(100);
    std::cout << "Balance after update: " << slotMachine.GetBalance() << std::endl;

    slotMachine.Reset();
    std::cout << "Balance after reset: " << slotMachine.GetBalance() << std::endl;

    return 0;
}
```

Not bad, let's see if it will add the unit tests.

Prompt: replace the example using main to unit tests using doctest.h

Response:

```c++
#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"
#include <array>
#include <cstdlib>
#include <ctime>

// VSlot interface defines the required methods for implementing a virtual slot machine.
class VSlot {
public:
    virtual std::array<int, 3> Spin() = 0;
    virtual void Reset() = 0;
    virtual void UpdateBalance(int amount) = 0;
    virtual int GetBalance() const = 0;
    virtual ~VSlot() = default;
};

// MyVSlot class implements the VSlot interface.
class MyVSlot : public VSlot {
private:
    int balance;

public:
    MyVSlot(int64_t seed) : balance(0) {
        std::srand(seed);
    }

    std::array<int, 3> Spin() override {
        return {std::rand() % 10 + 1, std::rand() % 10 + 1, std::rand() % 10 + 1};
    }

    void Reset() override {
        balance = 0;
    }

    void UpdateBalance(int amount) override {
        balance += amount;
    }

    int GetBalance() const override {
        return balance;
    }
};

TEST_CASE("Testing MyVSlot class") {
    MyVSlot slotMachine(12345); // Initialize with a fixed seed for reproducibility

    SUBCASE("Spin generates three numbers between 1 and 10") {
        auto result = slotMachine.Spin();
        CHECK(result[0] >= 1);
        CHECK(result[0] <= 10);
        CHECK(result[1] >= 1);
        CHECK(result[1] <= 10);
        CHECK(result[2] >= 1);
        CHECK(result[2] <= 10);
    }

    SUBCASE("Balance updates correctly") {
        slotMachine.UpdateBalance(100);
        CHECK(slotMachine.GetBalance() == 100);

        slotMachine.UpdateBalance(-50);
        CHECK(slotMachine.GetBalance() == 50);
    }

    SUBCASE("Reset sets balance to zero") {
        slotMachine.UpdateBalance(100);
        slotMachine.Reset();
        CHECK(slotMachine.GetBalance() == 0);
    }
}
```

Bad ass!

We need to make a few tweaks to test.

1. Change the test name to MyVSlotTest
2. Update tests/CMakeLists.txt, setting the test name to MyVSlotTest

Let's test it:

```console
cd build; cmake ..; make; ctest; cd ..
```

Output:

```console
-- File already exists: /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/doctest.h
-- Configuring done (0.0s)
-- Generating done (0.0s)
-- Build files have been written to: /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
[ 50%] Built target HelloWorld
[100%] Built target test_vslot
Test project /Users/joeb/Projects/2024/rockyourslotsoff/src/Cpp/build
    Start 1: MyVSlotTest
1/1 Test #1: MyVSlotTest ......................   Passed    0.00 sec

100% tests passed, 0 tests failed out of 1

Total Test time (real) =   0.01 sec
```

Success (happy dance).

### Evaluate C++ Tests

It's obvious that ChatGPT tested the MyVSlot class differently than I did. It did use a fixed seed, which I expected - so it passes the hallucination test. However, I have no issue with the actual implementation because it is functionally equivalent. I can't think of a reason to make changes or enhancements. So we're done!

# Summary

At this point we've finished the "business logic". But we've not created the API. So, it's pretty obvious that getting to this point required more effort than expected. Not to worry, using TDD and interfaces things will go a lot faster as we progress. In theory, the API should be a piece of cake now.

I'm familiar with the Go tools for making APIs. I'm not familiar with C++ yet, so this will take a little more work.

However, initially, we will generate the webpage by just spitting out a static representation of the slot machine, then refreshing the entire page. When this works, we'll make the page dynamic by using JSON/RPC calls from the web page to the server over HTTP, probably a websocket.
