# design-pattern
设计模式 GoLang实现

---
## 模型分类
* **创建型模式**：五种。*工厂方法模式*、*抽象工厂模式*、*单例模式*、*建造者模式*、*原型模式*。
* **结构性模式**：七种。*适配器模式*、装饰者模式、代理模式、外观模式、桥接模式、组合模式、享元模式。
* **行为型模式**：十一种。策略模式、模板方法模式、观察者模式、迭代子模式、责任链模式、命令模式、备忘录模式、状态模式、访问者模式、中介者模式、解释器模式。
* **并发型模式**、**线程池模式**。

## 项目目录
```
.
├── LICENSE
├── README.md
├── creation 代码案例
│   ├── abstract 抽象工厂模式
│   │   ├── abstract_test.go
│   │   ├── doc.go
│   │   ├── models.go
│   │   └── types.go
│   ├── adapter 适配器模式
│   │   ├── adapter_test.go
│   │   ├── doc.go
│   │   ├── models.go
│   │   └── types.go
│   ├── builder 建造者模式
│   │   ├── builder_test.go
│   │   ├── doc.go
│   │   ├── models.go
│   │   └── types.go
│   ├── factory 工厂方法模式
│   │   ├── doc.go
│   │   ├── factory_test.go
│   │   ├── models.go
│   │   └── types.go
│   ├── prototype 原型模式
│   │   ├── doc.go
│   │   ├── models.go
│   │   └── prototype_test.go
│   └── singleton 单例模式
│       ├── doc.go
│       ├── models.go
│       ├── singleton_test.go
│       └── types.go
└── go.mod
```