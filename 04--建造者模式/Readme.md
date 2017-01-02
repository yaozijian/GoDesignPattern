
本目录中的示例代码参考文章[http://www.cnblogs.com/chenssy/p/3307787.html](http://www.cnblogs.com/chenssy/p/3307787.html)写成。

# 建造者模式的特点

* 将复杂对象的构建与表示分离，使得同样的构建过程可以用于创建不同的表示。示例代码中，将不同的MealBuilder传递给waiter::Construct()方法中相同的构建代码，可以获得不同的表示，即不同的Meal。
* 建造者返回给用户一个完整的产品对象，让用户无须关心对象所包含的额外属性和组建方式。
* 示例代码中包m导出给用户使用的全部是interface，这就是面向接口编程的体现。

# 示例代码的UML类图

![UML类图](http://ohnogfliw.bkt.clouddn.com/go-design-pattern/04/meal.png)