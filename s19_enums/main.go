package main

import "fmt"

//enums : enumerated types.

func changeOrderStatus(status string) {
	fmt.Println("Status chnaged to : ", status)
}

// for example we have the multiple status which we need to pass and such, need to have a fixed value of the status so then we will use the enums. also, it is easy to create the group of status with that.
// but golang does not have the direct datatype enums.
// here we will do with keyword const and custom types using the "type" keyword.

type OrderStatus int

const (
	Recieved OrderStatus = iota // read about it below.
	Confirmed
	Prepared
	Delivered
)

// now we will create the function which will recieve the parameter of type OrderStatus as:
func changeOrderStatus2(status OrderStatus) {
	fmt.Println("Status changed to : ", status)
}

type OrderStatus2 string

//-------when status values in created enum is string------
const (
	Recieved2  OrderStatus2 = "recieved"
	Confirmed2 OrderStatus2 = "confirmed"
	Prepared2  OrderStatus2 = "prepared"
	Delivered2 OrderStatus2 = "delivered"
)

func changeOrderStatus3(status OrderStatus2) {
	fmt.Println("Status changed to : ", status)
}

func main() {
	changeOrderStatus("applied ")
	changeOrderStatus("confirmed to espanol")
	//now in the changeOrderStatus2 we can pass the string but the enum created using the const and iota as :
	changeOrderStatus2(Delivered) //O/p : 3

	// it also helps in the typo problem with string values that are passed.
	// when the same funtion which change the status need to be passed as the string then we have to remember what string values can be passed but now the values are with type OrderStatus which makes it easy for us to have the values which needs to be passed. we can check by hovering over the method.
	// these statuses will be integer when using the iota

	//enum with string
	changeOrderStatus3(Recieved2)
}

/*
iota is a predeclared identifier in Go that represents the zero-indexed ordinal number of the current constant specification within a const block, automatically incrementing with each line.

Definition and Behavior :
In Go, iota is a predeclared identifier used exclusively within constant declarations (const blocks). It represents the untyped integer ordinal number of the current constant specification, starting at 0 for the first constant in the block and incrementing by 1 for each subsequent constant in the same block (). Each new const block resets iota to 0, making its scope limited to the block in which it appears.

iota is not a keyword but a built-in identifier, and it cannot be used outside const declarations. Its primary purpose is to simplify the creation of sequential constants, such as enumerations or bitwise flags, without manually specifying each value.

Basic Example :
const (
a = iota // 0
b        // 1
c        // 2
)
Here, a is assigned 0, b automatically becomes 1, and c becomes 2. This demonstrates iota's automatic incrementing behavior.

example :

Think of `iota` as a **"counter"** that Go uses specifically when you are making a list of constants. Instead of you typing out 0, 1, 2, 3 manually, `iota` does the counting for you.

Here is the breakdown in plain English:

### 1. It’s an Automatic Counter

When you start a `const` block, `iota` starts at **0**. Every time you move to a new line in that block, the counter goes up by **1**.

### 2. It "Remembers" the Pattern

You don't have to keep typing `iota` on every line. If you define the first constant with `iota`, the lines below it will automatically follow that pattern until the block ends.

### 3. It Resets Every Time

Every time you start a **brand new** `const` block, the counter resets to **0**. It doesn't care what the number was in the previous block.

---

### Why use it?

It’s mostly used to create **Enums** (lists of related items). Here are two common ways it's used:

**The Basic List**

```go
const (
    Red    = iota // 0
    Blue          // 1
    Green         // 2
)

```

**Skipping Values**
If you want to skip a number, you can use the "blank identifier" (`_`):

```go
const (
    _ = iota // 0 is skipped
    A        // 1
    B        // 2
)

```

**Powers of 2 (Bitshifting)**
`iota` is great for math too. You can use it to create values like 1, 2, 4, 8 easily:

```go
const (
    Read   = 1 << iota // 1 (1 << 0)
    Write  = 1 << iota // 2 (1 << 1)
    Execute = 1 << iota // 4 (1 << 2)
)

```

### Key Rules to Remember

* **Only for Constants:** You can't use `iota` inside a regular function with `var`.
* **Starts at Zero:** It always begins at 0 unless you do math to it (like `iota + 1`).

*/
