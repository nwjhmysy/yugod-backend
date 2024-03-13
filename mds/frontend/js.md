###   1.直接取对象属性`{}[]`

```javascript
let b = "a";
const obj = {
  a: 1,
  b: 2,
  c: 3
}[b];
console.log(obj);//1
```

### 2.可选链`(?.)` 运算符

**可选** 的链接运算符 ( **`?.`**) 使您能够读取位于连接对象链深处的属性值，而无需检查链中的每个引用是否有效。

`?.`运算符类似于 链接运算符，除了 如果`.`引用为空（[`null`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/null)或 [`undefined`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/undefined)）不会导致错误，表达式短路并返回值为 `undefined`。与函数调用一起使用时， `undefined`如果给定函数不存在，则返回。

当存在引用可能丢失的可能性时，当访问链接属性时，这会导致更短和更简单的表达式。当无法保证需要哪些属性时，它在探索对象的内容时也很有帮助。

可选链接不能用于未声明的根对象，但可以用于未定义的根对象。

```javascript
const adventurer = {
  name: 'Alice',
  cat: {
    name: 'Dinah'
  }
};

const dogName = adventurer.dog?.name;
console.log(dogName);
// expected output: undefined

console.log(adventurer.someNonExistentMethod?.());
// expected output: undefined
```

### 3.ES2020的条件加载—`import(//path)`

```
improt是静态加载，同步加载，先于模块内的其他语句执行，所以只能放在模块首部，所以不支持条件语句等
import a from './'
```

```javascript
es2020的import()方法是异步加载，支持条件语句，返回Promise对象

if(true){
   import('./')
     .then(module =>{
	console.log(module)
 })
} else {
  import('./a')
     .then(module =>{
	console.log(module)
 })
}
```

Import( )并不会在一进入就加载对应模块，可以做到按需加载。

`import()`函数可以用在任何地方，不仅仅是模块，非模块的脚本也可以使用。它是运行时执行，也就是说，什么时候运行到这一句，就会加载指定的模块。另外，`import()`函数与所加载的模块没有静态连接关系，这点也是与`import`语句不相同。`import()`类似于 Node 的`require`方法，区别主要是前者是异步加载，后者是同步加载。

```javascript
Promise.all([
  import('./module1.js'),
  import('./module2.js'),
  import('./module3.js'),
])
.then(([module1, module2, module3]) => {
   ···
});
```

### 4.多行省略号

```css
{
  overflow: hidden;
  display: -webkit-box;
  text-overflow: ellipsis;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: (n行省略);
}
```

### 5.`ES2016-ES2021`新特性纵览

#### `ES2016`

```
node 支持：node v7.5.0 以上支持 100%。

新增规定只要函数参数使用了默认值、解构赋值、或者扩展运算符，那么函数内部就不能显式设定为严格模式，否则会报错。
新增了数组实例的 includes 方法和指数运算符。
```

#### `ES2017`

```
node 支持：node v9.11.2 以上支持 100%。

1.Object.getOwnPropertyDescriptors() 返回指定对象所有自身属性（非继承属性）的描述对象。
2.Object.values 和 Object.entries 供 for...of 循环使用。
3.padStart() 用于头部补全，padEnd() 用于尾部补全。
4.引入了 async 函数。
5.允许函数的最后一个参数有尾逗号。
6.引入 SharedArrayBuffer，允许 Worker 线程与主线程共享同一块内存。
```

#### `ES2018`

```
node 支持：node v12.4.0 以上支持 100%。

1.正则：引入s修饰符，使得.可以匹配任意单个字符，后行断言。
2.引入了“异步遍历器”（Async Iterator），为异步操作提供原生的遍历器接口。
3.模板字符串：遇到不合法的字符串转义，就返回 undefined，而不是报错。
4.对象引入扩展运算符。
5.Promise.finally() 不管 Promise 对象最后状态如何，都会执行的操作。
```

#### `ES2019`

```
node 支持：node v12.4.0 以上支持 100%。

1.新增 trimStart() 和 trimEnd()。
2.toString() 返回函数代码本身，以前会省略注释和空格，现在明确要求返回一模一样的原始代码。
3.允许 catch 语句省略参数。
4.Array.prototype.sort() 的默认排序算法必须稳定。
5.Symbol 提供了一个实例属性 description，直接返回 Symbol 的描述。
```

#### `ES2020`

```javascript
node 支持：node v14.5.0 以上支持 100%。

1.import() 动态异步加载。
2.globalThis 打通所有环境。
3.String.prototype.matchAll() 可以一次性取出所有匹配。不过，它返回的是一个遍历器（Iterator），而不是数组。
4.BigInt 大整数计算保持精度，无位数限制 BigInt('123') 123n typeof 123n === 'bigint'。
5.Promise.allSettled() 只有等到所有这些参数实例都返回结果，不管是 fulfilled 还是 rejected，包装实例才会结束。
6.链判断运算符 ?.
(直接在链式调用的时候判断，左侧的对象是否为 null 或 undefined，如果是，就不再往下运算，而是返回 undefined
或者当左侧对象不确定是否有某一属性时)
7.null判断运算符
注：只有运算符左侧的值为 null 或 undefined 时，才会返回右侧的值
    let num = person.n ?? 100
    ?? 和 && || 同时用必须用括号表明优先级否则会报错
    && 和 || 的优先级孰高孰低，如果多个逻辑运算符一起使用，必须用括号表明优先级，否则会报错

```

#### `ES2021`

```
Promise.any() 只要参数实例有一个变成 fulfilled 状态，包装实例就会变成 fulfilled 状态；如果所有参数实例都变成 rejected 状态，包装实例就会变成 rejected 状态。
replaceAll() 一次性替换所有匹配 'aabbcc'.replaceAll('b', '_') // 'aa__cc'。
```

### 6.迭代器(`Iterator`)

ES5语法模拟迭代器

```javascript
function createIterator(items) {
    var i = 0;
    
    return { // 返回一个迭代器对象
        next: function() { // 迭代器对象一定有个next()方法
            var done = (i >= items.length);
            var value = !done ? items[i++] : undefined;
            
            return { // next()方法返回结果对象
                value: value,
                done: done
            };
        }
    };
}

var iterator = createIterator([1, 2, 3]);

console.log(iterator.next());  // "{ value: 1, done: false}"
console.log(iterator.next());  // "{ value: 2, done: false}"
console.log(iterator.next());  // "{ value: 3, done: false}"
console.log(iterator.next());  // "{ value: undefiend, done: true}"
// 之后所有的调用都会返回相同内容
console.log(iterator.next());  // "{ value: undefiend, done: true}"

//如果调用next()执行最后一个元素，那么返回对象中的done是true
```

遍历迭代器

```javascript
//错误,因为在while条件语句里也会迭代
while(!iterator.next().done){
    console.log(iterator.next().value);
}

//正确
let a;
while (a = iterator.next(), !a.done) {
    console.log(a);
};
```

### 7.生成器(`function*`)

生成器是一种返回迭代器的函数，通过`function`关键字后的星号（*）来表示，函数中会用到新的关键字`yield`

```javascript
function *createIterator(items) {
    for(let i=0; i<items.length; i++) {
        yield items[i];
    }
}

let iterator = createIterator([1, 2, 3]);

// 既然生成器返回的是迭代器，自然就可以调用迭代器的next()方法
console.log(iterator.next());  // "{ value: 1, done: false}"
console.log(iterator.next());  // "{ value: 2, done: false}"
console.log(iterator.next());  // "{ value: 3, done: false}"
console.log(iterator.next());  // "{ value: undefiend, done: true}"
// 之后所有的调用都会返回相同内容
console.log(iterator.next());  // "{ value: undefiend, done: true}"

上面，我们用ES6的生成器，大大简化了迭代器的创建过程。我们给生成器函数createIterator()传入一个items数组，函数内部，for循环不断从数组中生成新的元素放入迭代器中，每遇到一个yield语句循环都会停止；每次调用迭代器的next()方法，循环便继续运行并停止在下一条yield语句处。
```

#### 生成器的创建方式

```javascript
//生成器是个函数：
function *createIterator(items) { ... }

//可以用函数表达式方式书写：
let createIterator = function *(item) { ... }

//也可以添加到对象中，ES5风格对象字面量：
let o = {
    createIterator: function *(items) { ... }
};
let iterator = o.createIterator([1, 2, 3]);
                                       
//ES6风格的对象方法简写方式：
let o = {
    *createIterator(items) { ... }
};
let iterator = o.createIterator([1, 2, 3]);
```

#### 可迭代(`iterable`)对象

所有的集合对象（**数组、Set集合及Map集合**）和**字符串**都是可迭代对象

#### 访问默认迭代器

```javascript
可迭代对象，都有一个Symbol.iterator方法，for-of循环时，通过调用colors数组的Symbol.iterator方法来获取默认迭代器的，这一过程是在JavaScript引擎背后完成的。

let values = [1, 2, 3];
let iterator = values[Symbol.iterator]();

console.log(iterator.next());  // "{ value: 1, done: false}"
console.log(iterator.next());  // "{ value: 2, done: false}"
console.log(iterator.next());  // "{ value: 3, done: false}"
console.log(iterator.next());  // "{ value: undefined, done: true}"

在这段代码中，通过Symbol.iterator获取了数组values的默认迭代器，并用它遍历数组中的元素。在JavaScript引擎中执行for-of循环语句也是类似的处理过程。
```

用`Symbol.iterator`属性来检测对象是否为可迭代对象：

```javascript
function isIterator(object) {
    return typeof object[Symbol.iterator] === "function";
}

console.log(isIterable([1, 2, 3]));  // true
console.log(isIterable(new Set()));  // true
console.log(isIterable(new Map()));  // true
console.log(isIterable("Hello"));  // true
```

### 8.`async..await`生效条件

```
await后面的方法调用必须返回一个Promise对象
```

### 9.关于`Promise.resolve()`方法

```javascript
Promise.resolve('foo')
//等价于
new Promise(resolve=>resolve('foo'))
```

```
Promise.resolve方法允许调用时不带参数，直接返回一个resolved状态的 Promise 对象。

所以，如果希望得到一个 Promise 对象，比较方便的方法就是直接调用Promise.resolve方法。
```

```javascript
//需要注意的是，立即resolve的 Promise 对象，是在本轮“事件循环”（event loop）的结束时，而不是在下一轮“事件循环”的开始时。
setTimeout(function () {
  console.log(3);
}, 0);
Promise.resolve().then(function () {
  console.log(2);
});
console.log(1);
//执行结果
1
2
3
```

### 10.箭头函数中`this`的指向问题

概括：

箭头函数体内的`this`对象，就是定义**该函数所在的作用域指向的对象**，而不是使用时所在的作用域指向的对象。（作用域是指函数内）

```js
var name = 'window'; // 其实是window.name = 'window'

var A = {
   name: 'A',
   sayHello: function(){
      console.log(this.name)
   }
}

A.sayHello();// 输出A

var B = {
  name: 'B'
}

A.sayHello.call(B);//输出B

A.sayHello.call();//不传参数指向全局window对象，输出window.name也就是window
```

从上面可以看到，sayHello这个方法是定义在A对象中的，当我们使用call方法，把其指向B对象，最后输出了B；可以得出，sayHello的this只跟使用时的对象有关。

改造一下：

```js
var name = 'window'; 

var A = {
   name: 'A',
   sayHello: () => {
      console.log(this.name)
   }
}

A.sayHello();// 还是以为输出A ? 错啦，其实输出的是window
```

“**该函数所在的作用域指向的对象**”，作用域是指函数内部，这里的箭头函数，也就是sayHello，所在的作用域其实是最外层的js环境，因为没有其他函数包裹；然后最外层的js环境指向的对象是winodw对象，所以这里的this指向的是window对象。

改造永远绑定`this`指向的箭头函数

```js
var name = 'window'; 

var A = {
   name: 'A',
   sayHello: function(){
      var s = () => console.log(this.name)
      return s//返回箭头函数s
   }
}

var sayHello = A.sayHello();//这里的sayHello是返回的箭头函数，并且指向A
sayHello();// 输出A 

var B = {
   name: 'B';
}

sayHello.call(B); //还是A
sayHello.call(); //还是A
```

### 11.全局对象`window`或`globe`

需要开启非严格模式

### 12.空值合并操作符`??`

在编写代码时，如果某个属性不为 null 和 undefined，那么就获取该属性，如果该属性为 null 或 undefined，则取一个默认值：

```javascript
const name = dogName ? dogName : 'default'; 
//可以通过 || 来简化：
const name =  dogName || 'default'; 
```

但是 || 的写法存在一定的缺陷，**当 dogName 为 0 或 false 的时候也会走到 default 的逻辑**。

所以 ES2020 引入了 ?? 运算符。**只有 ?? 左边为 null 或 undefined时才返回右边的值：**

```javascript
const dogName = 0; 
const name1 =  dogName ?? 'default';  // name1 = 0;
const name2 = dogName || 'default';		//name = 'defalut';
```

### 13.规范（待确认）

```
1.当嵌套括号左侧在一行时，右侧括号也应在一行
{({
	...
})}
2.当嵌套括号左侧不在一行时，右侧括号也应不在一行
{
	(
		{...}
	)
}
```

### 14.项目中api的封装（类式封装）

例如：user系列的请求

1.将每个user相关的请求封装为类中的方法（UserApi类）

```javascript
export const getUsersApi = async (token?: string) =>
  new UsersApi(await AuthManager.getInstance().getConfiguration(token))

//AuthManager类的作用：生成一个关于配置的实例化对象
async getConfiguration(token?: string): Promise<Configuration> {
    const conf = new Configuration()
    // conf.accessToken = await this.getAccessToken()
    conf.accessToken = token
    conf.basePath = API_BASE_PATH
    conf.baseOptions = {
      headers: {
        'X-Client-Platform': 'web',
        'X-Client-Version': APP_VERSION,
      },
    }
    return conf
  }

```

2.UserApi

```javascript
export class UsersApi extends BaseAPI {
	public getUsers(options?: AxiosRequestConfig) {
  	return UsersApiFp(this.configuration).getUsers(options).then((request) => request(this.axios, this.basePath));
  }
  //...各个user请求相关的方法
}
```

3.真正调用的其实是`UsersApiFp`方法返回的对象中的`getUsers`方法

```javascript
export const UsersApiFp = function(configuration?: Configuration) {
    const localVarAxiosParamCreator = UsersApiAxiosParamCreator(configuration)
    return {
    	async getUsers(options?: AxiosRequestConfig): Promise<(axios?: AxiosInstance, basePath?: string) => AxiosPromise<GetUsersResponse>> {
      const localVarAxiosArgs = await localVarAxiosParamCreator.getUsers(options);
      return createRequestFunction(localVarAxiosArgs, globalAxios, BASE_PATH, configuration);
  		//...
      },
    }
}
```

4.在`createRequestFunction`中调用`anxios`方法

```javascript
export const createRequestFunction = function (axiosArgs: RequestArgs, globalAxios: AxiosInstance, BASE_PATH: string, configuration?: Configuration) {
    return <T = unknown, R = AxiosResponse<T>>(axios: AxiosInstance = globalAxios, basePath: string = BASE_PATH) => {
        const axiosRequestArgs = {...axiosArgs.options, url: (configuration?.basePath || basePath) + axiosArgs.url};
        return axios.request<T, R>(axiosRequestArgs);
    };
}
```

### 15.`typescript`创建二维数组

使用Array.from()方法和映射函数来创建二维数组

```tsx
let rows: number = 3;
let cols: number = 3;
let array: number[][] = Array.from({ length: rows }, () => new Array(cols));
```

以上示例代码中的 `number` 类型可以替换为你需要的其他数据类型。无论你使用哪种方法，TypeScript 将根据所提供的类型注解来进行类型检查，确保数组中的元素类型正确。

### 16.`js`判断对象中有没有值

使用 `Object.keys(obj)` 方法和 `length` 属性

使用 `Object.values(obj)` 方法和 `length` 属性

### 17.`js` 对象根据属性名删除该属性

在 JavaScript 中，你可以使用 `delete` 关键字来删除对象中的属性。只需将对象的属性名作为参数传递给 `delete` 即可删除该属性。以下是示例代码：

```
javascriptCopy code
let obj = {
  name: 'John',
  age: 30,
  city: 'New York'
};

// 删除属性
delete obj.age;

console.log(obj); // Output: { name: 'John', city: 'New York' }
```

在上面的示例中，我们使用 `delete` 关键字删除了对象 `obj` 中的 `age` 属性。结果输出中可以看到，`age` 属性已经被从对象中删除。

需要注意的是，使用 `delete` 删除属性时，属性会从对象中彻底删除，而不仅仅是赋值为 `undefined` 或 `null`。同时，对于 JavaScript 内置对象（如数组和函数）以及原型链上的属性，`delete` 操作可能无法删除或屏蔽这些属性。

### 18.es6的`Symbol`类型

ES6（ECMAScript 2015）引入了一种新的基本数据类型，称为`Symbol`。`Symbol`是一种原始值，表示独一无二的值。

以下是关于ES6中`Symbol`类型的一些重要概念：

1. **创建Symbol**:

   ```javascript
   const mySymbol = Symbol();
   ```

   `Symbol`的创建是唯一的，每次调用`Symbol()`都会返回一个新的、不同的`Symbol`值。

2. **可选的描述**:

   可以向`Symbol`传递一个可选的描述（字符串），用于标识该`Symbol`的用途。这个描述对于调试和理解代码非常有用，但并不会影响`Symbol`的唯一性。

   ```javascript
   const mySymbol = Symbol('description');
   ```

3. **Symbol的唯一性**:

   两个具有相同描述的`Symbol`也是不相等的：

   ```javascript
   const symbol1 = Symbol('description');
   const symbol2 = Symbol('description');
   console.log(symbol1 === symbol2); // false
   ```

4. **作为对象属性的键**:

   由于`Symbol`的唯一性，可以将其用作对象属性的键，以确保不会与其他属性冲突：

   ```javascript
   const mySymbol = Symbol();
   const myObject = {
     [mySymbol]: 'Hello!'
   };
   ```

5. **访问Symbol属性**:

   使用`[ ]`操作符来访问包含`Symbol`作为键的属性：

   ```javascript
   console.log(myObject[mySymbol]); // 输出: Hello!
   ```

6. **Well-known Symbols**:

   ES6引入了一些预定义的`Symbol`，称为"well-known symbols"，它们在特定的语言内部行为中起到了特殊的角色，比如`Symbol.iterator`用于定义可迭代对象的迭代行为。

   ```javascript
   const iterableObject = {
     [Symbol.iterator]: function* () {
       yield 1;
       yield 2;
       yield 3;
     }
   };
   ```

7. **Symbol的静态属性**:

   `Symbol`构造函数还具有一些预定义的静态属性，如`Symbol.iterator`和`Symbol.for`。

   ```javascript
   const mySymbol = Symbol.for('mySymbol'); // 创建一个全局Symbol
   const mySymbol2 = Symbol.for('mySymbol'); // 尝试获取全局Symbol
   console.log(mySymbol === mySymbol2); // true
   ```

   上述代码中，`Symbol.for`会在全局Symbol注册表中查找给定描述的Symbol，如果找到则返回，否则创建一个新的Symbol。

> TypeScript

## 1.范型基础

**可以将泛型理解成为把类型当作参数一样去传递**

简单的案例：

```typescript
function identity<T>(arg: T): T { 
  return arg;
}

// 调用identity时传入name，函数会自动推导出泛型T为string，自然arg类型为T，返回值类型也为T
const userName = identity('name');
// 同理，当然你也可以显示声明泛型
const id = identity<number>(1); 
```

## 2.接口范型位置

示例：

```typescript
//1.接口上定义泛型
// 定义一个泛型接口 IPerson表示一个类，它返回的实例对象取决于使用接口时传入的泛型T
interface IPerson<T> {  
  // 因为我们还没有讲到unknown 所以暂时这里使用any 代替  
  new(...args: unknown[]): T;
}

//在使用这个接口时需要传入接口的泛型类型
function getInstance<T>(Clazz: IPerson<T>) { 
  return new Clazz();
}

// use it
class Person {}

// TS推断出函数返回值是person实例类型
const person = getInstance(Person);
//--------------------------------------------
//2.接口内部定义泛型
// 声明一个接口IPerson代表函数
interface IPerson { 
  // 此时注意泛型是在函数中参数 而非在IPerson接口中 
  <T>(a: T): T;
}

// 函数接受泛型
//接口本身不需要泛型，而在实现使用接口代表的函数类型时需要声明该函数接受一个泛型参数。
const getPersonValue: IPerson = <T>(a: T): T => {  
  return a;
};

// 相当于getPersonValue<number>(2)
getPersonValue(2)
```

总结：

- **当泛型出现在接口中时，比如`interface IPerson<T>` 代表的是使用接口时需要传入泛型的类型，比如`IPerson<T>`。**
- **当泛型出现在接口内部时，比如第二个例子中的 `IPerson`接口代表一个函数，接口本身并不具备任何泛型定义。而接口代表的函数则会接受一个泛型定义。换句话说接口本身不需要泛型，而在实现使用接口代表的函数类型时需要声明该函数接受一个泛型参数。**

## 3.泛型约束

作用：**约束泛型需要满足的格式**

例如：

```typescript
// 定义方法获取传入参数的length属性
function getLength<T>(arg: T) {  
  // throw error: arr上不存在length属性 
  return arg.length;
}
```

上面案例中，会产生严重的且隐蔽的错误。如果参数`arg`中不存在`length`属性，那么就会出现错误。

想要解决这个问题，需要确保传入的参数有`length`属性

### 关键字——**extents**

```typescript
interface IHasLength { 
  length: number;
}

// 利用 extends 关键字在声明泛型时约束泛型需要满足的条件
function getLength<T extends IHasLength>(arg: T) { 
  // throw error: arr上不存在length属性 
  return arg.length;
}

getLength([1, 2, 3]); // correct
getLength('123'); // correct
getLength({ name: '19Qingfeng', length: 100 }); // correct
// error 当传入true时，TS会进行自动类型推导 相当于 getLength<boolean>(true)
// 显然 boolean 类型上并不存在拥有 length 属性的约束，所以TS会提示语法错误getLength(true); 
```

### 关键字——**keyof**

```typescript
interface IProps { 
  name: string;
  age: number; 
  sex: string;
}

// Keys 类型为 'name' | 'age' | 'sex' 组成的联合类型
type Keys = keyof IProps
```

`Keyof any`

```typescript
// Keys 类型为 string | number | symbol 组成的联合类型
type Keys = keyof any
```

使用案例：

```typescript
function getValueFromKey(obj: object, key: string) { 
  // throw error 
  // key的值为string代表它仅仅只被规定为字符串  
  // TS无法确定obj中是否存在对应的key
  return obj[key];
}

//ts会报错，因为obj中不一定有key属性
```

消除错误——

```typescript
// 函数接受两个泛型参数
// T 代表object的类型，同时T需要满足约束是一个对象
// K 代表第二个参数K的类型，同时K需要满足约束keyof T （keyof T 代表object中所有key组成的联合类型）
// 自然，我们在函数内部访问obj[key]就不会提示错误了
function getValueFromKey<T extends object, K extends keyof T>(obj: T, key: K) { 
  return obj[key];
}
```

### 关键字——`is`

 is 关键字更多用在函数的返回值上，用来表示对于函数返回值的类型保护。

```typescript
// 函数的返回值类型中 通过类型谓词 is 来保护返回值的类型
function isNumber(arg: any): arg is number { 
  return typeof arg === 'number'
}
//如果传入的参数是 number 类型的，执行函数的内容

function getTypeByVal(val:any) { 
  if (isNumber(val)) {  
    // 此时由于isNumber函数返回值根据类型谓词的保护  
    // 所以可以断定如果 isNumber 返回true 那么传入的参数 val 一定是 number 类型   
    val.toFixed() 
  }
}
```

通常我们使用 is 关键字（类型谓词）在函数的返回值中，从而对于函数传入的参数进行类型保护。

### extends的进阶用法

extents的三大用法：

- 对范型进行约束

- 接口的继承

- 条件类型（仅支持在`type`关键字中使用）

  示例：

  ```java
    // 示例1
    interface Animal {
      eat(): void
    }
    
    interface Dog extends Animal {
      bite(): void
    }
    
    // A的类型为string
    type A = Dog extends Animal ? string : number
    
    const a: A = 'this is string'
  ```

  `extends`用来条件判断的语法和JS的三元表达是很相似，如果问号前面的判断为真，则将第一个类型string赋值给A，否则将第二个类型number赋值给A。

  ```java
  type GetSomeType<T extends string | number> = T extends string ? 'a' : 'b';
  
  let someTypeOne: GetSomeType<string> // someTypeone 类型为 'a'
  
  let someTypeTwo: GetSomeType<number> // someTypeone 类型为 'b'
  
  let someTypeThree: GetSomeType<string | number>; // what ? 
  ```

### 4.ts枚举写法

在 TypeScript 中，枚举（Enum）是一种用于定义一组具名常量的数据类型。枚举可以在代码中提供更可读、更清晰的含义，用于表示一组相关的常量值。以下是 TypeScript 中枚举的写法：

1. 数字枚举：

```ts
enum Direction {
  Up,       // 默认从0开始递增，Up = 0
  Down,     // Down = 1
  Left,     // Left = 2
  Right,    // Right = 3
}

// 使用枚举成员
let direction: Direction = Direction.Up;
console.log(direction); // Output: 0

// 使用枚举成员的值
let directionValue: number = Direction.Right;
console.log(directionValue); // Output: 3
```

1. 字符串枚举：

```ts
enum Color {
  Red = 'RED',
  Green = 'GREEN',
  Blue = 'BLUE',
}

// 使用枚举成员
let color: Color = Color.Green;
console.log(color); // Output: "GREEN"

// 使用枚举成员的值
let colorValue: string = Color.Red;
console.log(colorValue); // Output: "RED"
```

1. 常量枚举：

```ts
const enum Sizes {
  Small,
  Medium,
  Large,
}

// 使用常量枚举成员的值
let size: number = Sizes.Medium;
console.log(size); // Output: 1
```

常量枚举在编译时会被完全删除，它的成员在使用时会被内联为常量值，因此只能用于数字枚举。常量枚举适用于不需要实际的对象表示的情况，例如用作数字或字符串的标识符。

## 

### 



### 异步请求顺序控制

- `promise`+`async...await`

  ```typescript
  function asyncFunction1() {
    return new Promise(resolve => {
      setTimeout(() => {
        console.log("Step 1 complete");
        resolve(1);
      }, 1000);
    });
  }
  
  function asyncFunction2() {
    return new Promise(resolve => {
      setTimeout(() => {
        console.log("Step 2 complete");
        resolve(2);
      }, 1000);
    });
  }
  
  function asyncFunction3() {
    return new Promise(resolve => {
      setTimeout(() => {
        console.log("Step 3 complete");
        resolve(3);
      }, 1000);
    });
  }
  
  async function executeAsyncFunctions() {
    // 执行第一个异步操作并等待其完成
    const result1 = await asyncFunction1();
    console.log("Result 1:", result1);
  
    // 执行第二个异步操作并等待其完成
    const result2 = await asyncFunction2();
    console.log("Result 2:", result2);
  
    // 执行第三个异步操作并等待其完成
    const result3 = await asyncFunction3();
    console.log("Result 3:", result3);
  }
  
  // 调用主函数
  executeAsyncFunctions();
  
  ```

- 生成器函数+`promise`+`async...await`

  ```typescript
  function sleep(value:number,ms:number) {// 模拟异步请求
    return new Promise(resolve => setTimeout(() => {
      resolve(value)
    }, ms));
  }
  
  function* asyncGenerator() {
    yield sleep(1, 1000); 
  
    yield sleep(2, 1000);
    
    yield sleep(3, 1000);
  }
  
  async function executeAsyncGenerator() {
    const generator = asyncGenerator();
    let result;
  
    // 执行第一个步骤
    result = generator.next(); //此时的reault是一个promise对象
    const request1 = await result.value;
    console.log("Request 1:", request1);
    
    // 执行第二个步骤
    result = generator.next();
    const request2 = await result.value;
    console.log("Request 2:", request2);
  
    // 执行第三个步骤
    result = generator.next();
    const request3 = await result.value;
    console.log("Request 3:", request3);
  }
  
  // 调用主函数
  executeAsyncGenerator();
  
  ```
