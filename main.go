package main

import (
	"fmt"
	"sync"
)

/*
 *@project: master
 *@author: Administrator
 *@created: 2022/2/23
 *@updated: 2025/3/27
 *@description:
	//////////创建型设计模式
	1.单例模式：确保一个类仅有一个实例
	2.原型模式：复制一个已有对象来生成新对象
	3.建造者模式：通过set来构建一个复杂函数（类似java创建一个类方法一样）
	4.工厂模式：提供一种创建对象的简单方式，由产品到使用(单一类型产品)
	5.抽象工厂模式：创建相关对象的家族（一组相互关联/依赖的对象）（由工厂到产品，产品到使用(一组相关产品)）
	//////////行为设计模式
	1.观察者模式：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖它的对象都会收到通知（观察者模式）
	2.策略模式：定义一系列算法，封装每个算法，并使它们可以相互替换（策略模式）
	3.模板方法模式：定义一个操作中的算法的骨架，而将一些步骤延迟到子类中，使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤
	//////////结构设计模式
	1.适配器模式：将一个类的接口转换成客户希望的另一种接口
	2.组合模式：将多个对象组合在一起进行操作
	3.代理模式：为一个对象提供代理以控制对这个对象的访问
*/

func main() {
	//创建型设计模式

	//单例模式: 确保一个类仅有一个实例，并提供一个访问它的全局访问点
	println(GetInstance("hello"))

	//工厂模式: 提供一种创建对象的简单方式
	/** 理解： 创建怎么样的产品？
	    将对象创建与使用分离，通过专门工厂类控制实例化过程
	由产品到使用(单一类型产品)
	**/
	FactoryTest()

	//抽象工厂模式: 提供一种创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类
	/** 理解：用于创建相关对象的家族（一组相互关联/依赖的对象）
	    工厂的工厂：定义一个抽象工厂接口，由具体工厂子类决定生产哪一组产品
	    产品族：同一工厂生产的多个产品必须属于同一主题/系列（如：同一操作系统风格、同一UI主题）
	    强约束：确保不同产品之间能协同工作（例如：Windows按钮必须搭配Windows风格的文本框）

	由工厂到产品，产品到使用(一组相关产品)
	**/
	AbstractFactoryTest()

	//建造者模式: 将一个复杂对象的构建与它的表示分离，使得同样地构建过程可以创建不同的表示
	//[初始化对象] → [配置部件A] → [配置部件B] → [最终组装] → [交付成品]
	//
	//创造一个产品，一个一个的set值
	BuilderTest()

	//原型模式: 用原型实例指定创建对象的种类，然后通过拷贝这些原型创建新的对象
	PrototypeTest()

}

// Singleton 单例模式
// 确保一个类仅有一个实例，并提供一个访问它的全局访问点
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstance 获取单例
func GetInstance(data string) *Singleton {
	//确保只执行一次
	once.Do(func() {
		instance = &Singleton{data: data}
	})
	return instance
}

/******************************************/

// Factory 工厂接口
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA ConcreteFactoryB 具体工厂
type ConcreteFactoryA struct{}

type ConcreteFactoryB struct{}

// CreateProduct 创建产品
func (a *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteFactoryA{}
}
func (b *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteFactoryB{}
}

// Product 接口定义产品行为
type Product interface {
	Use() string
}

// Use 实现产品行为
func (a *ConcreteFactoryA) Use() string {
	return "ConcreteFactoryA Use"
}
func (b *ConcreteFactoryB) Use() string {
	return "ConcreteFactoryB Use"
}

// FactoryTest 工厂模式
// 工厂模式是一种创建型设计模式，它提供了一种创建对象的简单方式
func FactoryTest() {
	//创建A实例
	var factory Factory
	//由产品到使用
	factory = &ConcreteFactoryA{}
	productA := factory.CreateProduct()
	println(productA.Use())

	factory = &ConcreteFactoryB{}
	productB := factory.CreateProduct()
	println(productB.Use())
}

/******************************************/

// ConcreteProductA1 具体产品
type ConcreteProductA1 struct{}
type ConcreteProductA2 struct{}
type ConcreteProductB1 struct{}
type ConcreteProductB2 struct{}

// AbstractProductA 抽象产品A接口  AbstractProductB 抽象产品B接口
type AbstractProductA interface {
	UseA()
}
type AbstractProductB interface {
	UseB()
}

// UseA 实现产品A和B
func (a *ConcreteProductA1) UseA() {
	fmt.Println("Use ProductA1")
}
func (a *ConcreteProductA2) UseA() {
	fmt.Println("Use ProductA2")
}

func (a *ConcreteProductB1) UseB() {
	fmt.Println("Use ProductB1")
}
func (a *ConcreteProductB2) UseB() {
	fmt.Println("Use ProductB2")
}

// AbstractFactory 抽象工厂接口
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

// ConcreteFactory1 具体工厂1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ConcreteProductA1{}
}
func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ConcreteProductB1{}
}

// ConcreteFactory2 具体工厂2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	return &ConcreteProductB2{}
}

// AbstractFactoryTest 抽象工厂模式
func AbstractFactoryTest() {
	var factory AbstractFactory
	// 使用工厂1创建产品A和B
	factory = &ConcreteFactory1{}
	//从工厂到产品
	productA := factory.CreateProductA()
	productB := factory.CreateProductB()
	productA.UseA()
	productB.UseB()

	// 使用工厂2创建产品A和B
	factory = &ConcreteFactory2{}
	productA = factory.CreateProductA()
	productB = factory.CreateProductB()
	productA.UseA()
	productB.UseB()
}

/******************************************/

// Pizza 复杂产品
type Pizza struct {
	size   string
	crust  string
	cheese bool
}

// PizzaBuilder 建造者接口
type PizzaBuilder interface {
	SetSize(size string) PizzaBuilder
	SetCrust(crust string) PizzaBuilder
	AddCheese() PizzaBuilder
}

// 具体建造者， 还不是最终产品
type concretePizzaBuilder struct {
	pizza Pizza
}

func (c *concretePizzaBuilder) SetSize(size string) PizzaBuilder {
	c.pizza.size = size
	return c
}

func (c *concretePizzaBuilder) SetCrust(crust string) PizzaBuilder {
	c.pizza.crust = crust
	return c
}

func (c *concretePizzaBuilder) AddCheese() PizzaBuilder {
	c.pizza.cheese = true
	return c
}

func NewPizzaBuilder() PizzaBuilder {
	return &concretePizzaBuilder{
		pizza: Pizza{
			size:  "medium",  // 默认值
			crust: "regular", // 默认值
		},
	}
}
func (b *concretePizzaBuilder) Build() Pizza {
	return b.pizza
}

// BuilderTest 建造者模式
func BuilderTest() {
	builder := NewPizzaBuilder()
	myPizza := builder.SetSize("large").
		SetCrust("thin").AddCheese()

	fmt.Printf("My Pizza: %+v\n", myPizza)
}

/******************************************/

type Character struct {
	Name string
}

func (c *Character) Clone() *Character {
	return &Character{
		Name: c.Name,
	}
}
func (c *Character) NewName(name string) {
	c.Name = name
}

// PrototypeTest 原型模式
func PrototypeTest() {
	character := &Character{
		Name: "John",
	}
	clone := character.Clone()
	clone.Name = "Jack"
	fmt.Println("copy Character: " + clone.Name)
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
