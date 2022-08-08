# go-colorcolorful print for go

## install

```shell
go get -u github.com/merryituxz/go-color
```



## usage

```go
func main() {
	color.Red().Println("hello").
		Green().Println("hello").
		Yellow().Println("hello").
		Blue().Println("hello").
		Magenta().Println("hello").
		Cyan().Println("hello")
}
```

**output:**

<font style="color:red">red</font>

<font style="color:green">green</font>

<font style="color:yellow">yellow</font>

<font style="color:blue">blue</font>

<font style="color:magenta">magenta</font>

<font style="color:cyan">cyan</font>

<br/>

```go
func main() {
    color.Red().Println("hello").
    	Println("hello").
    	Println("hello")
}
```

**output:**

<font style="color:red">red</font>

<font style="color:red">red</font>

<font style="color:red">red</font>

<br/>

```go
func main() {
    color.Println("black")
    color.Red().Println("red")
    color.Println("black")
}
```

**output:**

<font style="color:black">black</font>

<font style="color:red">red</font>

<font style="color:black">black</font>