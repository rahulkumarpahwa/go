package main

import (
	"fmt"
	"time"
)

func task(id int) {
	// printing id so as to get the order of tasks.
	fmt.Println("doing task", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		// this is blocking and we want all of them to execute together so as to get the fast speed.
		// task(i)

		// writing go will make this 11 task func's run concurrently (by interleaving or parallelism, based upon the CPU cores) each in new lightweight thread called as goroutine, which is not a complete thread but the part of OS thread.
		// If the CPU has one core, then it will juggle between the tasks, which looks they all run at once to have the concurrency.
		// if multiple cores then it will run few on one and few on another where it will interleave (juggle) based upon the count of the cores.

		// when go is written here in front of the method, then golang will not directly run it, but will put in the 'scheduler' (called as run queue) and it manages and then scheduler, schedules the goroutines. Means, it will not run right away but goes to scheduler and then runs.

		go task(i)

		// as the above is non blocking, these 11 task will scheduled and as the main progranm does not have anything to do further, the main goroutine will end and scheduled task will never executed.
	}

	// simulate the block main so as to make the tasks run which are scheduled.
	time.Sleep(time.Second)

	// now when printed, the task will not run in order series wise, as the tasks are not blocking now, but they are scheduled and run concurrently by the 'run queue' which is scheduler and a 'ready to go list'. (refer to question IMPORTANT 2 BELOW) and now are concurrent and very fast.

	// we can do the anonymous function as well
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("doing task by anonmous method", i)
		}(i)
	}

}

/*
IMPORTANT QUESTION:

why do we need to have the sleep in the main as we have the tasks started running immedailtely together all the tasks at once and they just print a statement and then why can't they be completed before the main is completed?

Answer :
Good question — this is one of the most important (and confusing) parts of understanding goroutines.

Let’s break it down **very clearly** 👇

---

## 🚨 Core Issue: `main` exits too fast

When your program starts, **`main()` itself runs as a goroutine** (called the *main goroutine*).

Now look at your loop:

```go
for i := 0; i <= 10; i++ {
    go task(i)
}
```

What actually happens:

1. `go task(i)` does **NOT run the function immediately**
2. It just **schedules** the goroutine to run
3. Control returns instantly to the loop
4. Loop finishes **very fast (microseconds)**
5. `main()` reaches the end → **program exits**

👉 When `main()` exits → **ALL goroutines are killed immediately**

---

## ⚡ Why tasks don’t finish (even though they are “started”)

Even though you wrote:

```go
go task(i)
```

This only means:

> “Hey scheduler, please run this sometime”

It does **NOT guarantee execution before `main` ends**

---

## 🧠 Important Concept

Goroutines are:

* **asynchronous**
* **non-blocking**
* **scheduled, not executed instantly**

So this is what actually happens in timeline:

```
main:   start loop → schedule goroutines → finish loop → exit
tasks:  maybe start… maybe not… 💀 killed if main exits
```

---

## 🤔 Then why `time.Sleep()` works?

```go
time.Sleep(time.Second)
```

This forces `main()` to:

> “wait for 1 second before exiting”

During that time:

* Scheduler gets a chance to run your goroutines
* They execute and print output

---

## ❗ But Sleep is NOT the correct solution

`time.Sleep()` is just a **hack**, not proper synchronization.

Problems:

* What if tasks take longer than 1 second?
* What if they finish in 10ms? → waste time

---

## ✅ Correct Way: Use `sync.WaitGroup`

This is the proper way to wait for goroutines:

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // waits until all goroutines call Done()
}
```

---

## 🧩 Simple Analogy

Think of `main()` like a manager:

* You assign 11 workers (goroutines)
* Then immediately leave the office

💥 Workers get kicked out because office is closed

👉 `time.Sleep()` = manager waits randomly
👉 `WaitGroup` = manager waits until all workers report completion ✅

---

## 💡 Final Takeaway

* `go task()` ≠ immediate execution
* Goroutines need **time + scheduler opportunity**
* Program exits when `main()` ends
* Use **WaitGroup**, not `Sleep`, in real applications

---

If you want, I can next explain **how Go scheduler actually decides which goroutine runs (M:N model)** — that will connect perfectly with your earlier question.

*/

/*
IMPORTANT QUESTION 2:

but it they are scheduled (in some list , i don't know where and how) then how they are executed then concurrent?

Great — this is *exactly* the right doubt.
You’re thinking:

> “If goroutines are just sitting in some queue, how are they actually running concurrently?”

Let’s clear this properly 👇

---

# 🧠 First: Yes, they are in a queue (but not stuck there)

Goroutines are placed in a **run queue** (inside `P` — processor).

But important:

> That queue is not a waiting room forever — it’s a **ready-to-run list**

---

# ⚙️ What actually happens

Each **P (processor)** is constantly doing this loop:

```text
while true:
    pick a goroutine from queue
    run it for a short time
    if not finished → put it back
```

👉 This loop runs continuously on OS threads.

---

# 🔄 How concurrency happens (Single Core)

Let’s say:

* 1 CPU core
* 1 P
* 1 thread (M)
* 3 goroutines: G1, G2, G3

Execution looks like:

```text
G1 runs → paused
G2 runs → paused
G3 runs → paused
G1 runs again → paused
...
```

⚡ This switching is **very fast (microseconds)**

👉 So it *looks like* everything is running together
👉 This is called **concurrency (interleaving)**

---

# ⚡ How parallelism happens (Multi-Core)

Now suppose:

* 4 CPU cores
* 4 P
* multiple threads

```text
Core1 → G1 running
Core2 → G2 running
Core3 → G3 running
Core4 → G4 running
```

👉 Now goroutines are literally running **at the same time**

---

# 🔥 So where does “concurrent execution” come from?

It comes from **two mechanisms combined**:

### 1. Fast switching (time slicing)

* Goroutine runs for a short time
* Scheduler switches to another

### 2. Multiple cores

* Different goroutines run simultaneously

---

# ⛔ Important: Goroutines don’t block the scheduler

If one goroutine does:

```go
time.Sleep(...)
```

or network call:

👉 Scheduler immediately runs another goroutine

So queue keeps moving.

---

# 🧩 Think of it like CPU sharing

Imagine 1 chef (CPU) and 3 dishes (goroutines):

```text
Cook dish A → pause
Cook dish B → pause
Cook dish C → pause
Repeat...
```

👉 All dishes progress together
👉 That’s concurrency

Now 3 chefs:

```text
Chef1 → A
Chef2 → B
Chef3 → C
```

👉 That’s parallelism

---

# 🔁 Back to your confusion (very important)

> “If they are in a list, how are they running?”

Because:

* They are not *stuck* in the list
* The scheduler is **continuously pulling and executing them**
* The list is just a **source of ready work**

---

# 💡 Final Mental Model

> “Goroutines are like tasks in a queue, and the scheduler is constantly picking, running, pausing, and switching between them extremely fast.”

---

# 🚨 Why your earlier code behaved like that

Because:

* Goroutines were **added to queue**
* But `main` exited before scheduler got enough time to run them

👉 So they *never got enough CPU time*

---

If you want, next I can show you:

👉 **exact timeline (nanosecond level) of your code execution**
👉 OR
👉 **why order of prints is random (even on same CPU)**


*/
