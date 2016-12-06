
本目录中的示例代码参考文章[http://www.cnblogs.com/chenssy/archive/2013/05/26/3099495.html](http://www.cnblogs.com/chenssy/archive/2013/05/26/3099495.html)写成。

# 工厂方法模式的特点

* 有一个抽象工厂类，一个抽象产品类；多个实现了抽象工厂的具体工厂类，多个实现了抽象产品的具体产品类。
* 抽象工厂类定义了创建抽象产品类的工厂方法；多个具体工厂类都实现了工厂方法，创建具体产品类作为抽象产品返回。
* 对于相同的创建参数，不同的具体工厂类创建的具体产品是不同的。 

# 示例代码的UML类图

![UML类图](http://ohnogfliw.bkt.clouddn.com/go-design-pattern/02/store.png)