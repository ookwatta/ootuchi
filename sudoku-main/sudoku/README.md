## 0) File header & imports

```go
package main

import (
	"fmt"
	"os"
)
```

- `package main`: makes this a runnable Go program (`go run . ...`).
- `import "fmt"`: used for printing the board or `Error`.
- `import "os"`: used to read command-line arguments (`os.Args`).

---

## 1) Rule check: `isValid`

```go
func isValid(b [][]byte, r, c int, n byte) bool {
	for i := 0; i < 9; i++ {
		if b[r][i] == n || b[i][c] == n {
			return false
		}
	}
	sr, sc := (r/3)*3, (c/3)*3
	for i := sr; i < sr+3; i++ {
		for j := sc; j < sc+3; j++ {
			if b[i][j] == n {
				return false
			}
		}
	}
	return true
}
```

- `func isValid(...): bool`: answers “Can we place digit `n` at row `r`, col `c`?”
- `for i := 0; i < 9; i++ { ... }`: scan the **row** and **column**.

  - `b[r][i] == n`: if `n` already exists in row `r` → **not valid**.
  - `b[i][c] == n`: if `n` already exists in column `c` → **not valid**.

- `sr, sc := (r/3)*3, (c/3)*3`: top-left of the 3×3 subgrid that contains `(r,c)`.
- Two nested loops `i := sr..sr+2`, `j := sc..sc+2`: scan the **3×3 box**.

  - If any cell in that box equals `n` → **not valid**.

- `return true`: if no conflict in row, column, or box.

---

## 2) Solver (backtracking): `solve`

```go
func solve(b [][]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '.' {
				for n := byte('1'); n <= '9'; n++ {
					if isValid(b, r, c, n) {
						b[r][c] = n
						if solve(b) {
							return true
						}
						b[r][c] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}
```

- Scan the board **left-to-right, top-to-bottom** to find the **first empty** cell (`.`).
- For that empty cell, try digits `'1'`..`'9'`:

  - `isValid(...)`: only try digits that don’t break the rules.
  - Place it: `b[r][c] = n`.
  - Recurse: `if solve(b) { return true }` → if the rest of the board can be solved, we’re done.
  - Otherwise backtrack: `b[r][c] = '.'` and try the next digit.

- If **no** digit works for this empty cell → `return false` (dead end).
- If the loops finish without finding any `.` → the board is full → `return true` (solved).

**Mental model:** “Try a number; if you get stuck, undo and try the next.”

---

## 3) Unique-solution counter: `count`

```go
func count(b [][]byte, lim int) int {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if b[r][c] == '.' {
				cnt := 0
				for n := byte('1'); n <= '9'; n++ {
					if isValid(b, r, c, n) {
						b[r][c] = n
						cnt += count(b, lim)
						b[r][c] = '.'
						if cnt >= lim {
							return cnt
						}
					}
				}
				return cnt
			}
		}
	}
	return 1
}
```

- Same search pattern as `solve`, but instead of stopping at the first solution, it **counts** how many there are.
- `cnt := 0`: local count for this branch.
- For each valid digit:

  - Place it, recurse, add the returned count, then undo.
  - **Early stop**: if `cnt >= lim` (we only care if there are **2 or more**), return immediately.

- If the board has **no empty cells**, that means we found **one** solution → `return 1`.
- Called as `count(copyB(b), 2)` in `main` to check **exactly one** solution.

---

## 4) Board copy: `copyB`

```go
func copyB(s [][]byte) [][]byte {
	d := make([][]byte, 9)
	for i := range s {
		d[i] = append([]byte{}, s[i]...)
	}
	return d
}
```

- Creates a **deep copy** of the 9 rows.
- `append([]byte{}, s[i]...)` copies each row’s bytes so mutations don’t affect the original.
- Used before counting solutions so the counter can mutate freely.

---

## 5) Printer: `printB`

```go
func printB(b [][]byte) {
	for i := range b {
		for j := range b[i] {
			fmt.Printf("%c", b[i][j])
			if j < 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
```

- Prints each row as `digit space digit space ...` (no trailing space).
- Matches the required output format exactly.

---

## 6) The program flow: `main`

```go
func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	b := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		s := os.Args[i+1]
		if len(s) != 9 {
			fmt.Println("Error")
			return
		}
		b[i] = []byte(".........")
		for _, ch := range s {
			if ch != '.' && (ch < '1' || ch > '9') {
				fmt.Println("Error")
				return
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j, ch := range os.Args[i+1] {
			if ch != '.' {
				if !isValid(b, i, j, byte(ch)) {
					fmt.Println("Error")
					return
				}
				b[i][j] = byte(ch)
			}
		}
	}
	if count(copyB(b), 2) != 1 || !solve(b) {
		fmt.Println("Error")
		return
	}
	printB(b)
}
```

Walkthrough:

1. **Argument count check**

   - `if len(os.Args) != 10` → must be program name + **9** row strings.
   - Else print `Error` and quit.

2. **Allocate empty board & validate characters**

   - `b := make([][]byte, 9)`: 9 rows.
   - Loop `i := 0..8`:

     - `s := os.Args[i+1]`: the ith row string.
     - `if len(s) != 9` → `Error` (each row must be 9 chars).
     - `b[i] = []byte(".........")`: start that row as 9 empties.
     - `for _, ch := range s`: check every character is **either** `.` **or** `'1'..'9'`. If not → `Error`.

   _(At this point the board is all dots. We only checked character **types**.)_

3. **Place the given digits (with legality check)**

   - Loop over rows/cols again:

     - `for j, ch := range os.Args[i+1]`: look at each input char.
     - If `ch != '.'` (it’s a digit):

       - `!isValid(b, i, j, byte(ch))` → if placing this given digit breaks Sudoku rules → `Error`.
       - Else actually place it: `b[i][j] = byte(ch)`.

   _(Now the board has all initial clues placed **legally**.)_

4. **Uniqueness + solve**

   - `if count(copyB(b), 2) != 1`: make a copy and count solutions up to 2.

     - If the count isn’t **exactly 1** → `Error`.

   - `|| !solve(b)`: If uniqueness passed but the solver didn’t find a solution (shouldn’t happen given the count) → `Error`.

5. **Print the solved board**

   - `printB(b)` prints in the exact required format.

---

## Tiny example (how a single cell is handled)

Suppose row 0, col 2 is `.`. In `solve`:

- We try `'1'`..`'9'`.
- For each `n`, `isValid(b, 0, 2, n)` checks:

  - Scan row 0 and column 2.
  - Compute box top-left `(sr, sc)` and scan that 3×3 area.

- First `n` that passes:

  - Set `b[0][2] = n`, recurse.
  - If later it fails, we undo (`b[0][2] = '.'`) and try the next `n`.

That’s the **try → recurse → backtrack** rhythm you see in `solve` and the **counting** rhythm you see in `count`.
