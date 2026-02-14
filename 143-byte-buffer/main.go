/*
● bytes.Buffer — an IN-MEMORY buffer that implements io.Writer (and io.Reader)
  ○ No file I/O needed — great for building strings, testing, and piping data
● ANOTHER io.Writer IMPLEMENTATION — just like *os.File in Lesson 142, bytes.Buffer has Write()
● Buffer METHODS — NewBufferString, WriteString, Write, Reset, String
● BUILDING ON LESSON 142: both os.File and bytes.Buffer satisfy io.Writer,
  demonstrating POLYMORPHISM — same interface, different underlying types
*/
package main

import (
	"bytes"
	"fmt"
)

// Commented-out code from Lesson 142 shows the progression:
// 1. Lesson 142: write to a FILE using os.File (io.Writer)
// 2. Lesson 143: write to a BUFFER using bytes.Buffer (io.Writer)
// 3. Next step: use io.Writer as a parameter so one function works with BOTH

// type person struct {
// 	first string
// }

// // func (r receiver) identifier(p parameter(s)) (return(s)) { code }
// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// The os.File code from Lesson 142 is commented out — we're swapping
	// the file-based writer for an in-memory buffer.

	// f, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }
	// defer f.Close()

	// s := []byte("Hello gophers!")

	// _, err = f.Write(s)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// bytes.NewBufferString creates a *bytes.Buffer initialized with the given string.
	// The buffer holds the bytes in memory — no file needed.
	// review https://pkg.go.dev/bytes#Buffer.Write
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String()) // String() returns the buffer contents as a string → "Hello "

	// WriteString appends a string directly to the buffer (no []byte conversion needed).
	b.WriteString("Gophers!")
	fmt.Println(b.String()) // → "Hello Gophers!"

	// Reset empties the buffer — length goes to 0 but capacity is retained.
	// Useful for reusing a buffer without allocating new memory.
	b.Reset()

	b.WriteString("Zaphod Beeblebrox")
	fmt.Println(b.String()) // → "Zaphod Beeblebrox"

	// Write is the io.Writer method — takes []byte, same signature as os.File.Write.
	// This is what makes bytes.Buffer satisfy the io.Writer interface.
	b.Write([]byte(": President of the Universe"))
	fmt.Println(b.String()) // → "Zaphod Beeblebrox: President of the Universe"
}

/*
REMARKS:
- Key takeaway 1: bytes.Buffer is an IN-MEMORY io.Writer — no disk I/O, no error handling for file ops
- Key takeaway 2: WriteString is a convenience method; Write([]byte(...)) does the same thing
  but WriteString avoids the manual string-to-byte conversion
- Key takeaway 3: Reset() clears the buffer for reuse — efficient because it keeps the
  underlying allocated memory
- BUILDING ON LESSON 142: Lesson 142 used os.File.Write, this lesson uses bytes.Buffer.Write —
  SAME interface method, DIFFERENT underlying storage (disk vs memory)
- BUILDING ON LESSON 139-141: This is POLYMORPHISM in action. The commented-out writeOut()
  function accepts io.Writer, so it would work with BOTH os.File AND bytes.Buffer unchanged
- Python comparison: bytes.Buffer is like Python's io.StringIO or io.BytesIO —
  an in-memory stream that behaves like a file
- Common pattern: bytes.Buffer is heavily used in Go for building strings efficiently
  (like strings.Builder), writing test output, and HTTP response bodies
*/
