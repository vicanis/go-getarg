# go-getarg

Struct tags support to GET parameters

Usage:

```
package main

import (
	"fmt"
	"log"

	getarg "github.com/vicanis/go-getarg"
)

type Object struct {
	FirstName string `getarg:"first_name"`
	LastName  string `getarg:"last_name"`
}

func main() {
	obj := Object{
		FirstName: "John",
		LastName:  "Doe",
	}

	u, err := getarg.Encode(obj)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("struct %#v encoded as url.Values: %s\n", obj, u.Encode())

	decoded := Object{}

	err = getarg.Decode(u, &decoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("url.Values %s decoded as struct: %#v\n", u.Encode(), decoded)
}
```

outputs:

```
struct main.Object{FirstName:"John", LastName:"Doe"} encoded as url.Values: first_name=John&last_name=Doe
url.Values first_name=John&last_name=Doe decoded as struct: main.Object{FirstName:"John", LastName:"Doe"}
```
