package main

import (
	"fmt"
)

type Foo struct {
	content string
}

type FooCollection []*Foo

func (f FooCollection) doSomething() bool {
	return true
}

func (f *FooCollection) doSomethingElse() bool {
	return true
}

func main() {

	// make a few pointer to Foo's
	foo1 := &Foo{content: "hey"}
	foo2 := &Foo{content: "yo"}

	// create a FooCollection using the underlying type, not the alias
	fooCollection := []*Foo{foo1, foo2}

	// call a method that expects a FooCollection, this works
	callDoSomething(fooCollection)

	// call method on FooCollection, gives an error:
	// fooCollection.doSomething undefined
	// (type []*Foo has no field or method doSomething)
	// Question: how would I convert this to a FooCollection directly?
	fooCollection.doSomething()

}

func callDoSomething(fooCollection FooCollection) {

	fooCollection.doSomething()

	fooCollectionPtr := &fooCollection

	result := fooCollectionPtr.doSomethingElse()
	fmt.Printf("%v\n", result)

}
