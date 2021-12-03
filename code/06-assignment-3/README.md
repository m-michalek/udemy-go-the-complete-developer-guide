# Here are the directions:

1. Create a program that reads the contents of a text file then prints its contents to the terminal.
2. The file to open should be provided as an argument to the program when it is executed at the terminal.  For example, 'go run main.go myfile.txt' should open up the myfile.txt file
3. To read in the arguments provided to a program, you can reference the variable 'os.Args', which is a slice of type string
4. To open a file, check out the documentation for the 'Open' function in the 'os' package - https://golang.org/pkg/os/#Open
5. What interfaces does the 'File' type implement? 
6. If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy function!