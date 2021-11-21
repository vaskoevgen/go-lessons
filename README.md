
https://golangify.com/

## Install GO lang:

## 1. Download the Go installer.

```
wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
```

## 2. Extract the archive you downloaded into `/usr/local`, creating a Go tree in `/usr/local/go`: 


```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz
```

## 3. Add `/usr/local/go/bin` to the PATH environment variable:

```
export PATH=$PATH:/usr/local/go/bin
```

## 4. Verify that you've installed Go by opening a command prompt and typing the following command:

```
go version
```

> # Greetings GO module Greetings:

## 1. Create directory:

```
mkdir greetings
cd greetings
```

## 2. Start your module:

```
go mod init example.com/greetings
```

## 3. Write code in file `greetings.go`:

```
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```


> # Create GO module `Hello`:

## 1. Create directory `hello`

```
<home>/
 |-- greetings/
 |-- hello/
```

```
cd ..
mkdir hello
cd hello
```

## 2. Start your module:

`go mod init example.com/hello`

## 3. Write code in file `hello.go`:

```
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```

## 4. The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency.

```
go mod edit -replace example.com/greetings=../greetings
go mod tidy
```