/*
● THE io.Writer INTERFACE — one of Go's most important interfaces
  ○ Defined as: type Writer interface { Write(p []byte) (n int, err error) }
  ○ Any type with a Write([]byte) (int, error) method satisfies io.Writer
● os.File IMPLEMENTS io.Writer — os.Create returns an *os.File, which has a Write method
● WRITING BYTES TO A FILE — Write takes a []byte, so strings must be converted
● BUILDING ON LESSON 139-141: interfaces, polymorphism, and the power of implicit satisfaction
*/
package main

import (
	"log"
	"os"
)

// Commented-out code shows the NEXT step: making a custom type that accepts
// an io.Writer, so you can write to ANY destination (file, network, buffer).
// This is the POLYMORPHIC power of interfaces — see Lesson 141.

// type person struct {
// 	first string
// }

// // func (r receiver) identifier(p parameter(s)) (return(s)) { code }
// func (p person) writeOut(w io.Writer) error {
// 	_, err := w.Write([]byte(p.first))
// 	return err
// }

func main() {
	// os.Create creates (or truncates) a file and returns *os.File and an error.
	// *os.File implements io.Writer, io.Reader, io.Closer, and more.
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err) // log.Fatalf logs the error and calls os.Exit(1)
	}
	defer f.Close() // defer ensures the file is closed when main() returns (see Lesson 137)

	// Convert the string to a []byte — Write requires a byte slice, not a string.
	// This is the same Write method that satisfies the io.Writer interface.
	s := []byte("Hello gophers!")

	// f.Write writes the byte slice to the file.
	// The first return value (int) is the number of bytes written — discarded here with _.
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

/*
REMARKS:
- Key takeaway 1: os.File satisfies io.Writer because it has a Write([]byte) (int, error) method
- Key takeaway 2: Write takes []byte, so use []byte("string") to convert — Go does NOT auto-convert
- Key takeaway 3: The commented-out code previews how you'd use io.Writer as a PARAMETER TYPE,
  letting the same function write to files, buffers, network connections, etc.
- BUILDING ON LESSON 137: defer f.Close() ensures cleanup even if an error occurs later
- BUILDING ON LESSON 139-141: io.Writer is a real-world example of the interface pattern —
  implicit satisfaction means *os.File doesn't need to declare "implements io.Writer"
- Python comparison: In Python, you'd use duck typing with file-like objects.
  Go's interfaces give you the same flexibility but with COMPILE-TIME type checking.
- Common pitfall: always check the error from os.Create — file operations can fail
  (permissions, disk full, invalid path)
*/
