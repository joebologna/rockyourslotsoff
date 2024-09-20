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

#### ChatGPT Interaction to Generate the Interface

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

### Using Globals in Go Tests

This should be done with care. We will use a `TestMain()` function to accomplish this.
