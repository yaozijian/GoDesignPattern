   七八年前买过GoF编著的著名的《设计模式——可复用面向对象软件的基础》一书，看了一点就看不下去了，实在是讲得比较高深，不好懂。两三年前看过电子版的《设计模式之禅》一书，觉得很好懂，仔细读了约10章，还在为知笔记里面做了比较详细的笔记。只是完整版的电子书需要花钱买，所以没有看完。近期在博客园看到一个[讲设计模式的系列文章](http://www.cnblogs.com/chenssy/p/3357683.html)，觉得非常好，就全部收集到为知笔记了。想想过去的经历，真的是该仔细学习下设计模式，同时稍微了解下UML图的画法了。为此，决定仔细学习下设计模式。这次学习主要参考博客园看到的将设计模式的系列文章，示例代码参考文章中的代码，但有以下目标：

* 使用Go语言编码：系列文章使用Java编码，而我将使用Go语言编码。在学习的过程中比较Go与Java的差别
* 给出完整代码：系列中的部分文章仅给出了核心代码，而没有完整代码，我学习过程中当然需要完整可用的代码。
* 代码上传到个人Github仓库中：地址为[https://github.com/yaozijian/GoDesignPattern](https://github.com/yaozijian/GoDesignPattern)

# Github中代码的组织和编译方法

* 每个目录存放一个设计模式的示例代码，目录名称指示设计模式名称。
* 为明确区分设计模式代码和客户端代码，将客户端代码写在名字全部为数字的.go主文件（如1.go,2.go等）中，而其他文件则包含实现设计模式的代码。主文件的包名为main,将编译成可执行程序;其他文件的包名为m,将编译成包m。
* 目录中design.simp为使用Software Ideas Modeler绘制的UML图；png文件为UML图的截图。
* 在每个设计模式的代码目录中执行../build.sh可完成编译，生成名字全部为数字的可执行程序。
* 在顶层目录中执行clear.sh可以删除每个目录中的可执行程序，用于清理。
* build.sh和clear.sh遵循可在Linux中运行。这两个脚本文件比较简单，可很容易地写出可在windows中运行的相应脚本。

学习过程中发现，为实现设计模式示例代码，要关注Go与Java的一些差别，其中主要的几点是：

# 构造方法
Go没有构造方法的概念，构造一个结构体可以使用new或者结构体字面量。new能创建出一个零值的结构体，返回其指针，但没法传入初始化参数；使用结构体字面量时可以使用一些参数初始化各个字段，有点类似构造方法了。但最好的模拟构造方法的方法可能是，使用一个NewXXX函数，这个函数接受构造结构体的各种参数，返回结构体指针。
Java中类可以有私有构造方法，使得用户不能直接构造类的实例。Go中模拟这种私有构造方法的方法是，将结构体声明为非导出的，即名称的首字符不是大写字母，这样用户就不能使用new或者字面量构造结构体了；然后提供一个NewXXX函数用于构造结构体，用户只能通过这个NewXXX函数来构造结构体。

# 抽象方法
Java中的抽象方法只定义方法的签名，而将方法的实现交由子类处理，含有抽象方法的类是抽象类，无法直接实例化。这在Go中是无法实现的，Go没有抽象方法的概念，无法实现Java中那种调用抽象方法时，到底调用的是哪个子类的方法只能在运行时确定的行为。Go中不经过接口的方法调用都是静态绑定的，即具体调用哪个类型的方法在编译时就可以确定。以系列文章中第23种设计模式的示例代码为例，Java代码为：

```
public abstract class CaffeineBeverage {
    final void prepareRecipe(){
        boilWater();
        brew();
        pourInCup();
        addCondiments();
    }
    abstract void brew();
    abstract void addCondiments();
    void boilWater(){ System.out.println("Boiling water...");}
    void pourInCup(){ System.out.println("Pouring into Cup...");}
}
```

这里的抽象类CaffeineBeverage的方法prepareRecipe在运行时调用抽象方法brew和addCondiments时，到底调用的是什么类的这两个方法，是由当前调用方法prepareRecipe的是哪个具体的非抽象子类的实例确定的，也就是无法在编译时确定。Go中要实现这种动态绑定行为，只能通过接口调用方法。如果在Go中这样写：

```
func (d *drink) Prepare(){
  d.boilWater()
  d.brew()
  d.pourInCup()
  d.addCondiments()
}
```

则Prepare方法调用的四个方法，肯定都是drink结构体的方法，不可能是其子类(Go没有这个概念，可以近似地认为其他嵌套了drink的结构体是drink的子类)的方法。为实现类似Java中抽象方法的行为，可使用函数成员：

1.这样定义drink结构体：

```
drink struct{
   _brew func()
   _addCondiments func()
}
```

2.这样定义方法brew和addCondiments：

```
func (d *drink) brew(){ d._brew() }
func (d *drink) addCondiments(){ d._addCondiments() }
```

3.在构造子类对象的时候设置_brew 和_addCondiments这两个函数成员的值，相关代码为：

```
type(
  coffie struct{ drink }
)

func NewCoffie() Drink{
  t := &coffie{}
  t._brew = t.brew
  t._addCondiments = t.addCondiments
  return t
}

func (c *coffie) brew(){
  fmt.Println("用沸水冲泡咖啡粉")
}

func (c *coffie) addCondiments(){
  fmt.Println("加入牛奶和糖")
}
```

这样，父类drink调用方法brew和addCondiments时，虽然直接调用的是父类的这两个方法，但最终调用的是子类对象的方法。无法在编译时确定父类最终调用哪个子类的方法，只能在运行时动态确定。这就近似地实现了抽象方法的行为。
