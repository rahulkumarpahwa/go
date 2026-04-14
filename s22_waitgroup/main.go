package main

import (
	"fmt"
	"sync"
)

// wait group : mechanism, in the previous part of the goroutines we see that until all of our go routines gets executed we need to pass the sleep in the main function.
// it is okay for the example but not good for the real world as we don't know how much time it would take to get completed.
// to handle that we use the waitgroup.

// we want to synchronous the way that when all the go routines get completed then only the main method should completed. for that we use the wait group.

func task(id int, w *sync.WaitGroup) { // we have to pass the waitgroup as the pointer here from the sync package.

	// after the task has been done we have to remove the waitgroup which is done with 'Done()' method with the use of defer keyword. this 'defer' makes sure that the line runs at the end with the defer keyword.
	defer w.Done()
	// this will minus the 1 we have added while making the call in the main.
	fmt.Println("Doing task with id : ", id)
}

func main() {
	var wg sync.WaitGroup
	// A WaitGroup is a counting semaphore typically used to wait for a group of goroutines or tasks to finish.

	// here we don't know how much time does the go routine will take to finish, so we can't tell how long we need to sleep the main goroutine.
	// so, to snchronize them, we will need something which will make sure that all the goroutines which are running in background other than main will finished and then after the main will be finished. so, for that we will use the waitgroup!

	for i := 0; i <= 10; i++ {
		wg.Add(1) // this waitgroup wg acts as a counter so we will add the one in it.
		// we added one because we start one goroutines. as the next go routines starts one will be added. so, each time one go rotines add one.
		// the above defer keyword Done() method will remove that one.
		go task(i, &wg)
		// we pass here the pointer of the waitgroup so as to get call the Done() with the defer.
	}

	wg.Wait() //here we have to do.
	// so, program will wait until our waitgroup not gets to zero by-default.

	// so in waitgroup, there are three things : 1. Done, 2. Wait 3. Add

	//---- when used with anonmous method then we don't need to pass the wg in the that as we call directly call as:
	var wg2 sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg2.Add(1)
		go func(i int) {
			fmt.Println("doing task by anonmous fxn: ", i)
			wg2.Done()
		}(i)

	}
	wg2.Wait()

}

/*
NOTE :

Timeline of execution with waitgroup
main:
  Add(1) Add(1) Add(1)
  start goroutines
  Wait()  ← BLOCKED

goroutines:
  run → Done()
  run → Done()
  run → Done()

counter → 3 → 2 → 1 → 0

main:
  resumes → exits
*/

/*
OTHER NOTE :

Nice — now you’re asking the **right low-level question** 🔥
Let’s connect `WaitGroup` with the scheduler + goroutines.

---

# 🧠 What problem does `WaitGroup` solve?

We already saw:

* Goroutines run **independently**
* `main()` exits → everything dies

So we need:

> “Don’t let `main` exit until all goroutines finish”

👉 That’s exactly what `sync.WaitGroup` does

---

# ⚙️ Internally: What is a WaitGroup?

A `WaitGroup` is basically:

> ✅ A **counter + waiting mechanism**

---

## 🧩 Core idea

It tracks:

```text
Number of active goroutines
```

---

# 🔄 How it works step-by-step

### 1. `wg.Add(1)`

```go
wg.Add(1)
```

👉 Means:

> “One more goroutine is going to run”

Internally:

```text
counter = counter + 1
```

---

### 2. Start goroutine

```go
go task(i, &wg)
```

Now scheduler runs it whenever it wants.

---

### 3. `wg.Done()`

Inside your goroutine:

```go
defer wg.Done()
```

👉 Means:

> “This goroutine has finished”

Internally:

```text
counter = counter - 1
```

---

### 4. `wg.Wait()`

```go
wg.Wait()
```

👉 This is the key part:

> “Block until counter becomes 0”

---

# 🔥 What actually happens during `Wait()`

This is where it connects to scheduler 👇

When `main` calls:

```go
wg.Wait()
```

👉 The **main goroutine gets blocked**

---

## 🚨 Important:

Blocked ≠ stopped program

It means:

* Main goroutine is paused
* Scheduler runs other goroutines

---

# 🧠 Scheduler + WaitGroup interaction

```text
Main goroutine → calls Wait() → goes to sleep

Scheduler:
    runs G1
    runs G2
    runs G3
```

Each goroutine finishes:

```text
Done() → counter decreases
```

---

## 🎯 When counter becomes 0

* WaitGroup “wakes up” the main goroutine
* `Wait()` returns
* `main()` exits safely ✅

---

# 📊 Timeline of execution

```text
main:
  Add(1) Add(1) Add(1)
  start goroutines
  Wait()  ← BLOCKED

goroutines:
  run → Done()
  run → Done()
  run → Done()

counter → 3 → 2 → 1 → 0

main:
  resumes → exits
```

---

# ⚡ Key Insight

> `WaitGroup` does NOT run goroutines
> It only **controls when main is allowed to exit**

---

# 🧩 Why this works (deep understanding)

Without `Wait()`:

```text
main exits → scheduler stops → goroutines die
```

With `Wait()`:

```text
main blocked → scheduler keeps running goroutines → all finish → main exits
```

---

# 🚨 Common mistakes

### ❌ Forgetting `Add()` before goroutine

```go
go task(i)
wg.Add(1) // WRONG order
```

👉 Race condition (goroutine may finish before Add)

---

### ❌ Forgetting `Done()`

👉 `Wait()` will block forever (deadlock)

---

### ❌ Passing WaitGroup by value

```go
func task(wg sync.WaitGroup) // ❌
```

👉 Always use pointer:

```go
func task(wg *sync.WaitGroup) // ✅
```

---

# 🧠 Mental Model (Best one)

> WaitGroup is like a **countdown latch**

* Add → “tasks started”
* Done → “task finished”
* Wait → “wait until all done”

---

# 🔥 Final one-liner

> “WaitGroup blocks the main goroutine, giving the scheduler time to run all other goroutines until they signal completion.”
*/
