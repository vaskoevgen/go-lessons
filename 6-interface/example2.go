// We start by defining our Animal interface:

type Animal interface {
    Speak() string
}
pretty simple: we define an Animal as being any type that has a method named Speak. The Speak method takes no arguments and returns a string. Any type that defines this method is said to satisfy the Animal interface. There is no implements keyword in Go; whether or not a type satisfies an interface is determined automatically. Let’s create a couple of types that satisfy this interface:

type Dog struct {
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
    return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}

// We now have four different types of animals: A dog, a cat, a llama, and a Java programmer. In our main() function, we can create a slice of Animals, and put one of each type into that slice, and see what each animal says. Let’s do that now:

func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}