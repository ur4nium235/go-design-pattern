ABSTRACT FACTORY
===

> Cho phép nhóm các object liên quan đến nhau.  

Giả sử ta có 2 nhà sản xuất là: 

* `nike`
* `adidas`

Hãy tưởng chúng ta cần mua bộ đồ thể thao gồm có `shoe - giầy` và `short - quần ngắn`  
```go
type shoe struct {
	logo string
	size int
}
type short struct {
	logo string
	size int
}
```

Chúng ta cũng có 2 interface sản phẩm: 
* `iShoe` - interface này được implement bởi `nikeShoe` và `adidasShoe` 
```go
type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
```
* `iShort` - interface này được implement bởi `nikeShort` và `adidasShort` 
```go
type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}
```

Chúng ta muốn mua một bộ đồ đầy đủ (gồm 2 món đồ trên) của một nhà sản xuất, tức là `nike`

```go
type nikeShoe struct {
	shoe
}

type nikeShort struct {
	short
}
```
hoặc `adidas`

```go
type adidasShoe struct {
	shoe
}

type adidasShort struct {
	short
}
```  
Cả 2 nhà sản xuất đều implement interface `iSportsFactory` 
```go
type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

type adidas struct {
}

type nike struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short{
			logo: "adidas",
			size: 14,
		},
	}
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short{
			logo: "nike",
			size: 14,
		},
	}
}

```


