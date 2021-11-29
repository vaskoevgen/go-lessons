
https://gobyexample.com/

https://www.tutorialspoint.com/go/index.htm

https://metanit.com/go/tutorial/1.1.php

https://golangify.com/

https://github.com/fsufitch/sample-go-app-v2

## Install GO lang:

## 1. Download the Go installer.

```
wget https://golang.org/dl/go1.17.3.linux-amd64.tar.gz
```

## 2. Extract the archive you downloaded into `/usr/local`, creating a Go tree in `/usr/local/go`: 


```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
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
mkdir firstmodule
cd first module
```

## 2. Start your module:

```
go mod init github.com/vaskoevgen/go-lessons
```

## 3. Write code in file `helloworld/helloworld.go`:

```
package helloworld

import (
	"fmt"
)

func Hello() {
    fmt.Println("Hello, World!")
}
```