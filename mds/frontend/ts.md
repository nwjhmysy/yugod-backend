js中的Symbols类型

ChatGPT

在 JavaScript 中，`Symbol` 是一种基本数据类型，它在 ECMAScript 6 (ES6) 中被引入。它是一种唯一且不可变的数据类型，通常用于对象属性的键值，以确保属性名的唯一性。

以下是关于 `Symbol` 的一些重要特点：

1. **唯一性**: 每个通过 `Symbol` 创建的值都是唯一的，即使它们有相同的描述。

   ```
   const symbol1 = Symbol('foo');
   const symbol2 = Symbol('foo');
   
   console.log(symbol1 === symbol2); // false
   ```
   
2. **不可变性**: `Symbol` 是不可变的，一旦创建，它的值不可更改。

   ```
   const symbol = Symbol('foo');
   symbol = Symbol('bar'); // Error: Cannot assign to read only property
   ```
   
3. **作为属性名**: `Symbol` 可以作为对象的属性名，保证了属性的唯一性。

   ```
   const uniqueSymbol = Symbol('unique');
   
   const obj = {
     [uniqueSymbol]: 'Hello'
   };
   
   console.log(obj[uniqueSymbol]); // Hello
   ```
   
4. **不能使用 `new` 关键字**: 不能使用 `new` 关键字来创建 `Symbol`。

   ```
   const symbol = new Symbol(); // Error: Symbol is not a constructor
   ```
   
5. **全局注册表 (Global Symbol Registry)**: 通过 `Symbol.for(key)` 可以在全局注册表中创建一个可重用的 `Symbol`。

   ```
   const globalSymbol = Symbol.for('foo');
   const localSymbol = Symbol('foo');
   
   console.log(globalSymbol === localSymbol); // false
   ```
   
6. **获取描述信息**: 可以通过 `Symbol.description` 获取 `Symbol` 的描述信息。

   ```
   const symbol = Symbol('foo');
   console.log(symbol.description); // foo
   ```
   
7. **Well-known Symbols**: ES6 引入了一组预定义的 `Symbol`，它们被用于对象的内部行为，如 `Symbol.iterator` 用于指定一个对象的默认迭代器。

   ```
   const arr = [1, 2, 3];
   const iterator = arr[Symbol.iterator]();
   console.log(iterator.next()); // { value: 1, done: false }
   ```

`Symbol` 类型为 JavaScript 提供了一种新的方式来确保属性名的唯一性，特别是在涉及到对象的属性名需要在多个地方使用的情况下，这种特性非常有用。



### 二.项目中的`xx.d.ts`文件

declare —— 声明

`xx.d.ts` 文件通常是 TypeScript 项目中的类型声明文件。`.d.ts` 扩展名代表 "TypeScript Declaration"，它们用于描述 JavaScript 代码中的类型信息。在 TypeScript 中，类型声明文件用于描述 JavaScript 库或模块的类型信息，以便在 TypeScript 项目中进行类型检查和代码补全。

在 TypeScript 中，你可以使用 `declare` 关键字来声明变量、函数、类等的类型而不进行实际的定义。`declare` 关键字告诉编译器这些实体已经存在，通常是由外部代码提供。这样，TypeScript 编译器就不会对这些声明进行检查或编译。`declare` 主要用于描述全局变量、函数、类以及模块等的类型，以便让 TypeScript 编译器知道它们的存在，但不对其进行实际的实现或定义。

`declare`**关键字**

例如，在一个 `.d.ts` 文件中，你可以这样使用 `declare` 关键字来**声明一个全局变量**：

```
declare var myGlobal: number;
```

这个声明告诉 TypeScript 编译器在编译时会有一个名为 `myGlobal` 的全局变量，它的类型为 `number`，但编译器并不会去检查它的实际定义。这就使得你可以在 TypeScript 中使用这个变量而不会报错，但是需要确保在实际运行环境中有该变量的定义。