// LESSON: Exercise - SHA-256 checksum verification
package main

/*
● When I ran a SHA-256 checksum on this text file, I got this hash:
	○ 7c6c8937b2a120af15849db05c9f46326761e0eec852a2e973b1e0b6acd59a01
● Download the text file associated with this hands-on exercise. Run a SHA-256 checksum
on it. Do you get the same hash?
	○ shasum -a 256 /path/to/file
		■ . is the current directory
		■ if you're in the directory with the file
			● shasum -a 256 ./file
● change the file by one character, then run SHA-256 again
	○ 9be13f9173f28ce3dd89c72aad7f5b0549a0641feb869509c7f96e8dc8b6ea8e
*/