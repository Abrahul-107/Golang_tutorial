# Golang_tutorial
Go language basics to advance 



## _<span style="color:green">1._Difference Between var and "=_</span>_
| Feature                            | `var`                                     | `:=`                                    |
|------------------------------------|-------------------------------------------|-----------------------------------------|
| Scope                              | Can be used inside and outside of functions | Can only be used inside functions   |
| Declaration and assignment        | Declaration and value assignment can be done separately | Declaration and value assignment must be done in the same line |



## _<span style="color:green">2._Numeric types in golang_</span>_

### Unsigned Integers

- **uint8**: Unsigned 8-bit integers (0 to 255)
- **uint16**: Unsigned 16-bit integers (0 to 65535)
- **uint32**: Unsigned 32-bit integers (0 to 4294967295)
- **uint64**: Unsigned 64-bit integers (0 to 18446744073709551615)

### Signed Integers

- **int8**: Signed 8-bit integers (-128 to 127)
- **int16**: Signed 16-bit integers (-32768 to 32767)
- **int32**: Signed 32-bit integers (-2147483648 to 2147483647)
- **int64**: Signed 64-bit integers (-9223372036854775808 to 9223372036854775807)



## _<span style="color:green">3._Floating Point Data Types_</span>_

1. **float32**: IEEE-754 32-bit floating-point numbers
2. **float64**: IEEE-754 64-bit floating-point numbers

### Complex Data Types

1. **complex64**: Complex numbers with float32 real and imaginary parts
2. **complex128**: Complex numbers with float64 real and imaginary parts


## _<span style="color:green">4.User Input Go Program</span>_
This Go program demonstrates how to read user input from the command line.

## Overview

This program is designed to illustrate the process of reading user input in a Go program. It uses the `bufio` package to create a reader that reads input from the standard input (keyboard).

## How it Works

1. The program prompts the user to enter their roll number.

2. It reads the input from the user until a newline character (`'\n'`) is encountered.

3. It prints the entered input along with its type.

## _<span style="color:green">Defer in golang</span>_

**--> A **defer** statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns. Defer is commonly used to simplify functions that perform various clean-up actions.**

# CopyFile Function Documentation

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)  // Open the source file
    if err != nil {
        return  // If opening the source file fails, return immediately
    }

    dst, err := os.Create(dstName)  // Create the destination file
    if err != nil {
        return  // If creating the destination file fails, return immediately
    }

    written, err = io.Copy(dst, src)  // Copy content from source to destination
    dst.Close()  // Close the destination file
    src.Close()  // Close the source file
    return  // Return the number of bytes written and any error encountered
}
```
## Overview

The `CopyFile` function copies the contents of a source file to a destination file. This function handles the opening, reading, writing, and closing of files. By using `defer`, it ensures that files are properly closed even if an error occurs.

## Function Signature

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)  // Open the source file
    if err != nil {
        return  // If opening the source file fails, return immediately
    }
    defer src.Close()  // Ensure the source file is closed when the function completes

    dst, err := os.Create(dstName)  // Create the destination file
    if err != nil {
        return  // If creating the destination file fails, return immediately
    }
    defer dst.Close()  // Ensure the destination file is closed when the function completes

    written, err = io.Copy(dst, src)  // Copy content from source to destination
    return  // Return the number of bytes written and any error encountered
}

```

## Uses of defer in Go

**Resource Management:** defer is often used to release resources such as files, network connections, or database connections.

**Simplifying Cleanup Code:** defer helps in writing cleaner code by placing cleanup code next to the resource allocation code, which makes the code easier to read and maintain.

**Handling Panics:** defer can be used to handle panics in a controlled manner, allowing for necessary cleanup before the program exits or recovers.




