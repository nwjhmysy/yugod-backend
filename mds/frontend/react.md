### 1.筛选对象，只传递子组件需要的

```javascript
const Button = props => {
  const { kind, ...other } = props;
  const className = kind === "primary" ? "PrimaryButton" : "SecondaryButton";
  return <button className={className} {...other} />;
};

const App = () => {
  return (
    <div>
      <Button kind="primary" onClick={() => console.log("clicked!")}>
        Hello World!
      </Button>
    </div>
  );
};
```

### 2.memo()、useCallback()、useMemo()使用场景

#### 问题一:

React 中当组件的 props 或 state 变化时，会重新渲染，实际开发会遇到不必要的渲染场景。比如：
父组件：

```javascript
父组件：
import { useState } from "react";
import { Child } from "./child";

export const Parent = () => {
  const [count, setCount] = useState(0);
  const increment = () => setCount(count + 1);

  return (
    <div>
      <button onClick={increment}>点击次数：{count}</button>
      <Child />
    </div>
  );
};

子组件：
export const Child = ({}) => {
  console.log("渲染了");
  return <div>子组件</div>;
};
每次改变父组件中的count,子组件都会被重新渲染,但是子组件的 props 和 state 没有变化，我们并不希望它重现渲染。
```

**如何解决：**

```javascript
使用memo将函数式子组件包裹
React.memo()是React v16.6引进来的新属性，用来控制函数组件的重新渲染。
React.memo()将组件作为函数(memo)的参数，函数的返回值(Child)是一个新的组件。
子组件：
import { memo } from "react";

export const Child = memo(() => {
  console.log("渲染了");
  return <div>子组件</div>;
});
显而易见,改变父组件中的count并没有重新渲染子组件。
```

#### 问题二：

在问题一中,父组件只是简单调用子组件，并未给子组件传递任何属性。我们传值看看：

```javascript
父组件：
import { useState } from "react";
import { Child } from "./child";

export const Parent = () => {
  const [count, setCount] = useState(0);
  const [name, setName] = useState("小明");
  const increment = () => setCount(count + 1);

  const onClick = (name: string) => {
    setName(name);
  };

  return (
    <div>
      <button onClick={increment}>点击次数：{count}</button>
      <Child name={name} onClick={onClick} />
    </div>
  );
};

子组件:
import { memo } from "react";

export const Child = memo(
  (props: { name: string; onClick: (value: any) => void }) => {
    const { name, onClick } = props;
    console.log("渲染了", name, onClick);
    return (
      <>
        <div>子组件</div>
        <button onClick={() => onClick("小红")}>改变 name 值</button>
      </>
    );
  }
);
父组件向子组件传递了状态name,函数onClick。
父组件只改变了count,但是依然引起了子组件的重新渲染,这显然是不合理的
---------------------------------------------------------
分析下原因：

1.点击父组件按钮，改变了父组件中 count 变量，进而导致父组件重新渲染；
2.父组件重新渲染时，会重新创建 onClick 函数，即传给子组件的 onClick 属性发生了变化，导致子组件渲染；
3.如果传给子组件的props只有基本数据类型将不会重新渲染。
```

**如何解决：**

```javascript
把内联回调函数及依赖项数组作为参数传入 useCallback，它将返回该回调函数的memoized回调函数，该回调函数仅在某个依赖项改变时才会更新。
当你把回调函数传递给经过优化的并使用引用相等性去避免非必要渲染（例如 shouldComponentUpdate）的子组件时，它将非常有用。
import { useCallback, useState } from "react";
import { Child } from "./child";

export const Parent = () => {
  const [count, setCount] = useState(0);
  const [name, setName] = useState("小明");
  const increment = () => setCount(count + 1);

  const onClick = useCallback((name: string) => {
    setName(name);
  }, []);

  return (
    <div>
      <button onClick={increment}>点击次数：{count}</button>
      <Child name={name} onClick={onClick} />
    </div>
  );
};
点击父组件count，子组件将不会重新渲染。
```

- `memoized`回调函数: 使用一组参数初次调用函数时，缓存参数和计算结果，当再次使用相同的参数调用该函数时，直接返回相应的缓存结果。(返回对应饮用，所以恒等于 ===)

#### 问题三：

在问题二中,name的属性是字符串(基本数据类型)，如果换成引用数据类型会怎样呢？

```javascript
父组件：
import { useCallback, useState } from "react";
import { Child } from "./child";

export const Parent = () => {
  const [count, setCount] = useState(0);
  // const [userInfo, setUserInfo] = useState({ name: "小明", age: 18 });
  const increment = () => setCount(count + 1);
  const userInfo = { name: "小明", age: 18 };

  return (
    <div>
      <button onClick={increment}>点击次数：{count}</button>
      <Child userInfo={userInfo} />
    </div>
  );
};

子组件：
import { memo } from "react";

export const Child = memo(
  (props: { userInfo: { name: string; age: number } }) => {
    const { userInfo } = props;
    console.log("渲染了", userInfo);
    return (
      <>
        <div>名字： {userInfo.name}</div>
        <div>年龄：{userInfo.age}</div>
      </>
    );
  }
);
结果：
点击父组件count，看到子组件每次都重新渲染了。
```

分析：

- 点击父组件按钮，触发父组件重新渲染。
- 父组件渲染，`const userInfo = { name: "小明", age: 18 };` 一行会重新生成一个新对象，导致传递给子组件的 userInfo 属性值变化，进而导致子组件重新渲染。
- 注意: 如果使用`useState`解构的userInfo, 子组件将不会重复渲染，因为解构出来的是一个memoized 值。

**如何解决？**

使用 useMemo 将对象属性包一层。

```javascript
import { useMemo, useState } from "react";
import { Child } from "./child";

export const Parent = () => {
  const [count, setCount] = useState(0);
  // const [userInfo, setUserInfo] = useState({ name: "小明", age: 18 });
  const increment = () => setCount(count + 1);
  const userInfo = useMemo(() => ({ name: "小明", age: 18 }), []);

  return (
    <div>
      <button onClick={increment}>点击次数：{count}</button>
      <Child userInfo={userInfo} />
    </div>
  );
};

```

useMemo()返回一个 memoized 值。

把“创建”函数和依赖项数组作为参数传入 `useMemo`，它仅会在某个依赖项改变时才重新计算 memoized 值。这种优化有助于避免在每次渲染时都进行高开销的计算。

记住，传入 `useMemo` 的函数会在渲染期间执行。请不要在这个函数内部执行与渲染无关的操作，诸如副作用这类的操作属于 `useEffect` 的适用范畴，而不是 `useMemo`。

如果没有提供依赖项数组，`useMemo` 在每次渲染时都会计算新的值。

----------------------------------------------------------------------------------------------------------------------------------

#### 注意：

使用memo包裹子组件时以下四种情况子组件不会被重新渲染（不改变传递的参数时，如果改变了传递的参数子组件肯定会被重新渲染）

- 没有给子组件传递props时，子组件不会重复渲染。

- 如果直接使用useState解构的**setState**传给子组件, 子组件将不会重复渲染，因为解构出来的是一个**memoized 函数**。

- 如果使用`useState`解构的**state**, 子组件将不会重复渲染，因为解构出来的是一个**memoized** 值。

- 如果传给子组件的props只有**基本数据类型**将不会重新渲染。

### 3.在开发环境中react子组件会被渲染两次

```
这个原因是 react刻意导致的，并发模式下在dev 时render-phase会执行两次

在react渲染组件中  ， 初始化的index.tsx文件里，存在一个<React.StrictMode>标签，
这个是react的一个用来突出显示应用程序中潜在问题的工具（严格模式）
有一项检测意外的副作用，严格模式不能自动检测到你的副作用，但它可以帮助你发现它们，使它们更具确定性。通过故意重复调用以下函数来实现的该操作。

注意：这仅适用于开发模式。生产模式下生命周期不会被调用两次。
```

### 4.useCall()的基本使用

组件父子关系：A--->B--->C

不使用该hook时

```javascript
//组件A
import React from "react";
import B from "./B";

export const nameContext = React.createContext("");
export default function App() {
  return (
    <nameContext.Provider value={"ys"}>
      大家好，
      <B />
    </nameContext.Provider>
  );
}
​//组件B
import C from "./C";

export default function B() {
  return (
    <>
      我是今天的分享者，
      <C />
    </>
  );
}
​//组件C
import React from "react";
import { nameContext } from "./App";

export default function C() {
  return (
    <nameContext.Consumer>
      {(name) => <span>我叫{name}</span>}
    </nameContext.Consumer>
  );
}
```

使用hook：

```javascript
//组件A
import React from "react";
import B from "./B";

export const nameContext = React.createContext("");
export const titleContext = React.createContext("");
export default function App() {
  return (
    <nameContext.Provider value={"我是ys，"}>
      <titleContext.Provider value={"今天的主题是Hooks使用"}>
        大家好，
        <B />
      </titleContext.Provider>
    </nameContext.Provider>
  );
}
​//组件C：
import React, { useContext } from "react"
import { nameContext, titleContext } from "./App"

export default function C() {
  const name = useContext(nameContext)
  const title = useContext(titleContext)

  return (
    <span>
      {name}
      {title}
      （useContext方式）
    </span>
  )
}
```

注意：

调用了 `useContext` 的组件总会在 context 值变化时重新渲染。如果重渲染组件的开销较大，你可以 [通过使用 memoization 来优化](https://github.com/facebook/react/issues/15156#issuecomment-474590693)。

### 5.`ref`与`useCallback`结合使用获取元素

```javascript
import React, {useCallback} from "react";

const Child = () => {
  const myInput=useCallback(node=>{
    node.defaultValue="123"
  },[]);

  return (
    <div>
      我是child组件:
      <input type={"text"} ref={myInput}/>
    </div>
  )
};
export default Child;

```

使用场景：

每次初始化渲染dom后获取dom的一些信息，或者对dom做一些初始化的操作

优势(好处)：

我们传递了 `[]` 作为 `useCallback` 的依赖列表。这确保了 ref callback 不会在再次渲染时改变，因此 React 不会在非必要的时候调用它。

### 6.深入useEffect原理

**`useEffect` 做了什么？** 通过使用这个 Hook，你可以告诉 React 组件需要在渲染后执行某些操作（**副作用**）。React 会保存你传递的函数（我们将它称之为 “effect”），并且在执行 DOM 更新之后调用它。

Hook 使用了 JavaScript 的闭包机制， 当然**useEffect**也不例外

传递给 `useEffect` 的函数在每次渲染中都会有所不同，这是刻意为之的。事实上这正是我们可以在 effect 中获取最新的 `count` 的值，而不用担心其过期的原因。每次我们重新渲染，都会生成*新的* effect，替换掉之前的。

所以useEffect所依赖的state，props都是最新的。

分析：

```javascript
function Counter() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    const id = setInterval(() => {
      setCount(c => c + 1); // ✅ 在这不依赖于外部的 `count` 变量
    }, 1000);
    return () => clearInterval(id);
  }, []); // ✅ 我们的 effect 不使用组件作用域中的任何变量

  return <h1>{count}</h1>;
}
为什么可以？：
```

```javascript
function Counter() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    const id = setInterval(() => {
      setCount(count + 1); // 这个 effect 依赖于 `count` state
    }, 1000);
    return () => clearInterval(id);
  }, []); // 🔴 Bug: `count` 没有被指定为依赖

  return <h1>{count}</h1>;
}

为什么不可以？：
```

### 7.`clsx`实现动态样式类

安装：

```
$ npm install --save clsx
```

使用规则：

```javascript
通过clsx函数返回类名列表

import clsx from 'clsx';
// or
import { clsx } from 'clsx';

// Strings (variadic)
clsx('foo', true && 'bar', 'baz');
//=> 'foo bar baz'

// Objects
clsx({ foo:true, bar:false, baz:isTrue() });
//=> 'foo baz'

// Objects (variadic)
clsx({ foo:true }, { bar:false }, null, { '--foobar':'hello' });
//=> 'foo --foobar'

// Arrays
clsx(['foo', 0, false, 'bar']);
//=> 'foo bar'

// Arrays (variadic)
clsx(['foo'], ['', 0, false, 'bar'], [['baz', [['hello'], 'there']]]);
//=> 'foo bar baz hello there'

// Kitchen sink (with nesting)
clsx('foo', [1 && 'bar', { baz:false, bat:null }, ['hello', ['world']]], 'cya');
//=> 'foo bar hello world cya'
```

### 8.`recoil`的使用（重）

`recoil`作为一个状态管理工具，只能用在`react`项目的函数式组件中

#### 一.安装

```
npm install recoil
```

#### 二.基本使用

传统状态管理：集中式存储，由根组件进行分发，一级一级的传递。

`recoil`：离散型存储，进行状态的分发，不会导致无关子节点被重新渲染。

1. `atom`(原子状态)

   ```javascript
   有两种：
   atom——是存储状态的最小单位
   atomFamil——允许传参
   
   function atom<>({
   	//唯一建
   	key:string,
   	//默认值
   	default?:T|Promise<T>|Loadable<T>|WrappedValue<T>|RecoilValue<T>
   	//副作用
   	effects?:
     //取消Immutable 设置为true时表示为一个可变的状态
     dangerourslyAllowMutability?:boolean
   }):RecoilState<T>
     
   //immutable的优势
     降低Mutable带来的复杂度
   	节省内存空间
     随意穿越（Undo/Redo，Copy/Paste）
   	拥抱函数式编程
   ```

2. `selector`（衍生状态｜计算状态）类似于计算属性

   ```javascript
   有四种：
   selector——以其他状态（atom｜selector）为参数的纯函数
   selectorFamily——允许传参
   constSelector——常亮选择器
   errorSelector——错误选择器
   
   function selector<T>({
   	key:string,
     get:({
       get:GetRecoilValue,
     })=>T|Promise<T>|Loadable<T>|WrappedValue<T>|RecoilValue<T>,
     
     set?:({
     	get:GetRecoilValue,
     	set:SetRecoilState,
   	})=>void,
     
     dangerourslyAllowMutability?:boolean,
     cachePolicy_UNSTABLE?:CachePolicy,//制定缓存策略
   })
   ```

   `atom`和`selector`的两个实践案例

   ```javascript
   //用户信息
   export const userInfoAtom = atom({
     key:'userInfoAtom',
     default: {
       userName: '张三',
       score: 10
     },
     effects: [
       ({node,onSet}) => {
         //设置数据时监测atom的变化
         onSet((newValue,oldValue)=>{
           //...
         })
       }
     ]
   })
   ==============================
   //字体大小原子状态
   export const fontSizeAtom = atom({
     key: 'fontSizeAtom',
     default: 20
   });
   //页面字体大小
   export const fontSizeState = selector({
     key: 'fontSizeStat',
     get: ({get})=>{
       const fontSizeNum = get(fontSizeAtom);
       return `${fontSizeNum}px`
     }
   });
   ```
   
   

   #### 三.`RecoilHooks`
   
   1. 同步(类似useState)
   
      ```javascript
      先声明状态
      const recoilState = atom | atomFamily | selector | selectorFamily
      //钩子一：读和写
      const [stateValue,setStateValue] = useRecoilState(recoilState)
      //钩子二：读
      const stateValue = useRecoilValue(recoilState)
      //钩子三：写
      const setStateValue = useSetRecoilState(recoilState)
      //钩子四：重置状态
      const resetStateValue = useResetRecoilState(recoilState)
      //钩子五：查看状态-不成熟
      const [] = useGetRecoilValueInfo_UNSTABLE()
      //钩子六：刷新状态-不成熟
      const stateValue = useRecoilRefresher_UNSTABLE()
      
      //注意，使用useRecoilState(),会导致页面的重新渲染
      //因此尽量分开使用
      ```
   
   2. 异步
   
      ```javascript
      loadable
      	loadable.state——('loading' | 'hasValue' | 'hasError')
      	loadable.contents——数据
      	
      //读和写
      const [loadable,setState] = useRecoilStateLoadable(recoilState)
      //读
      const loadable = useRecoilValueLoadable(recoilState)
      ```

### 9.`recoil`中`atom`和`selector`的细节

1.atom

规定默认值（包括数值类型），是recoil中的最小单位。

2.atomFamily（函数）

与对象相似，是保存对应atom的集合。

使用：

```javascript
atomFamily<value, key>(options: AtomFamilyOptions<value, key>): (param: key) => RecoilState<value>
import atomFamily
//获取某一key对应的value

//例
export const testAtomFamily = atomFamily<value,key>({
  key:"testAtomFamily",
  default:"",
})
get(testAtomFamily(key1))
await snapshot.getPromise(testAtomFamily(key1))
```

3.selector

真正使用recoil时是使用的**selector**

作用：类似于计算变量

一，get（外）

```javascript
get:({get,getCallback})=>{return}
```

这个**get**返回计算后的最终结果或getCallback方法

二，get（内）

这个get作为get（外）参数对象的一个属性。

作用：获取其他recoil的值（包括atom，atomFamily,selector）,将这些recoil的值作为该selector的依赖。

三，getCallback

```javascript
getCallback(({set,snapshot})=>{
	//set:改变atom或atomFamily的值
	//snapshot：获取其他recoil的值（包括atom，atomFamily,selector）,将这些recoil的值作为该getCallback的依赖
	return //计算后的值
})
```

作用：可以看作是子**selector**

注意：也可以讲**getCallback**从selector中分离出来——useRecoilCallback(({set,snapshot})=>{return})

四，注意点

getCallback中的方法是**async**方法，**recoil**中的异步操作一般在getCallback中进行。

get（内）尽量不获取getCallback的返回值。

get（内）是整个selector的依赖；snapshot是getCallback的依赖。

### 10.接口`ReadOnlySelectorOptions`

```javascript
export interface ReadOnlySelectorOptions<T> {
    key: string;
    get: (opts: {
      get: GetRecoilValue,
      getCallback: GetCallback,
    }) => Promise<T> | RecoilValue<T> | Loadable<T> | WrappedValue<T> | T;
    dangerouslyAllowMutability?: boolean;
    cachePolicy_UNSTABLE?: CachePolicyWithoutEquality; // TODO: using the more restrictive CachePolicyWithoutEquality while we discuss long term API
 }
```

### 11.接口`CallbackInterface`

```javascript
export interface CallbackInterface {
  set: <T>(recoilVal: RecoilState<T>, valOrUpdater: ((currVal: T) => T) | T) => void;
  reset: (recoilVal: RecoilState<any>) => void; // eslint-disable-line @typescript-eslint/no-explicit-any
  refresh: (recoilValue: RecoilValue<any>) => void;
  snapshot: Snapshot;
  gotoSnapshot: (snapshot: Snapshot) => void;
  transact_UNSTABLE: (cb: (i: TransactionInterface_UNSTABLE) => void) => void;
 }
```

### 12.`react-router-dom v6`完整使用示例

当我回答你的问题时，`react-router-dom`的最新版本是`v6.0.0-beta.6`。以下是一个完整的示例，展示如何在React应用中使用`react-router-dom` v6。

首先，确保你的项目已经安装了`react-router-dom`。你可以使用以下命令进行安装：

```
npm install react-router-dom@next
```

或者，如果你使用`yarn`作为包管理器，可以运行以下命令：

```
yarn add react-router-dom@next
```

接下来，让我们创建一个简单的React组件来设置路由和页面：

```javascript
// App.js

import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';

// 导入页面组件
import Home from './Home';
import About from './About';
import NotFound from './NotFound';

function App() {
  return (
    <BrowserRouter>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/about">About</Link>
          </li>
        </ul>
      </nav>

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
```

在上面的示例中，我们导入了`BrowserRouter`、`Routes`、`Route`和`Link`组件。我们使用`BrowserRouter`作为我们的根路由器，并使用`Routes`组件来定义不同的路由。在`Routes`中，我们使用`Route`组件来设置特定路径的页面组件。

注意，在v6中，我们使用`element`属性而不是`component`属性来指定路由的组件。另外，我们在`Route`组件中使用的`path`属性也不再以斜杠开头。

现在，我们来创建一些页面组件来匹配我们的路由：

```javascript
// Home.js

function Home() {
  return <h1>Home Page</h1>;
}

export default Home;
jsxCopy code
// About.js

function About() {
  return <h1>About Page</h1>;
}

export default About;
jsxCopy code
// NotFound.js

function NotFound() {
  return <h1>Page Not Found</h1>;
}

export default NotFound;
```

在上面的示例中，我们简单地返回了包含标题的组件。你可以根据需要进行更改和扩展。

现在，你可以在你的应用中使用`App`组件作为根组件：

```javascript
// index.js

import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

ReactDOM.render(<App />, document.getElementById('root'));
```

现在，当你在浏览器中运行应用时，你将看到一个包含"Home"和"About"链接的导航栏。点击链接将会呈现相应的页面组件。如果访问了一个未定义的路由，将会呈现"Not Found"页面。

这就是一个使用`react-router-dom` v6的基本示例

### 13.`react-router-dom v6`常用的`hooks`

1. `useNavigate`：用于编程式导航，可以在组件中通过调用`navigate`函数进行页面跳转。它返回一个函数，你可以在需要导航到其他页面时调用它。示例用法如下：

   ```javascript
   jsxCopy code
   import { useNavigate } from 'react-router-dom';
   
   function MyComponent() {
     const navigate = useNavigate();
   
     const handleClick = () => {
       navigate('/about'); // 导航到 '/about' 页面
     };
   
     return (
       <button onClick={handleClick}>Go to About</button>
     );
   }
   ```

2. `useParams`：用于获取路由参数。如果你在路由中定义了参数占位符，例如`/users/:id`，那么可以使用`useParams`来获取该参数的值。示例用法如下：

   ```javascript
   jsxCopy code
   import { useParams } from 'react-router-dom';
   
   function UserPage() {
     const { id } = useParams(); // 获取路由参数 'id'
   
     return (
       <h1>User ID: {id}</h1>
     );
   }
   ```

3. `useLocation`：用于获取当前页面的位置信息，包括路径、查询参数等。示例用法如下：

   ```javascript
   jsxCopy code
   import { useLocation } from 'react-router-dom';
   
   function MyComponent() {
     const location = useLocation();
   
     return (
       <div>
         <h1>Current Path: {location.pathname}</h1>
         <p>Query Params: {location.search}</p>
       </div>
     );
   }
   ```

4. `useMatch`：用于在嵌套路由中匹配当前路径。它可以帮助你在嵌套路由结构中找到与当前路径匹配的路由配置。示例用法如下：

   ```javascript
   jsxCopy code
   import { useMatch } from 'react-router-dom';
   
   function NestedComponent() {
     const match = useMatch('/users/:id');
   
     if (match) {
       // 当前路径与 '/users/:id' 匹配
       return (
         <h1>Matched User ID: {match.params.id}</h1>
       );
     }
   
     return (
       <h1>No Match</h1>
     );
   }
   ```

这些是`react-router-dom` v6中一些常用的钩子。它们可以帮助你处理导航、路由参数和当前位置等方面的逻辑。请根据你的具体需求选择合适的钩子进行使用。

### 14.`react-router-dom v6`路由间传递数据的三种方式

在 `react-router-dom` 版本 6 中，有几种方式可以在路由之间传递数据。以下是其中几种常用的方式：

1. URL 参数（URL Parameters）：你可以将数据作为 URL 的一部分来传递。在路由声明中，可以通过使用冒号 `:` 来定义参数，然后在组件中使用 `useParams` 钩子来获取参数的值。例如：

```javascript
jsxCopy code
import { BrowserRouter as Router, Route, Link, useParams } from 'react-router-dom';

function App() {
  return (
    <Router>
      <nav>
        <ul>
          <li>
            <Link to="/users/1">User 1</Link>
          </li>
          <li>
            <Link to="/users/2">User 2</Link>
          </li>
        </ul>
      </nav>

      <Route path="/users/:id" exact>
        <User />
      </Route>
    </Router>
  );
}

function User() {
  const { id } = useParams();

  return <h1>User {id}</h1>;
}
```

在上面的示例中，当你点击链接时，会导航到对应的用户页面，并在页面上显示相应的用户 ID。

1. 查询字符串（Query Parameters）：你可以使用查询字符串来传递数据。在路由声明中，可以使用 `useLocation` 钩子来获取当前路径的查询参数，并在组件中解析参数的值。例如：

```javascript
jsxCopy code
import { BrowserRouter as Router, Route, Link, useLocation } from 'react-router-dom';

function App() {
  return (
    <Router>
      <nav>
        <ul>
          <li>
            <Link to="/user?id=1">User 1</Link>
          </li>
          <li>
            <Link to="/user?id=2">User 2</Link>
          </li>
        </ul>
      </nav>

      <Route path="/user" exact>
        <User />
      </Route>
    </Router>
  );
}

function User() {
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const id = searchParams.get('id');

  return <h1>User {id}</h1>;
}
```

在上面的示例中，当你点击链接时，会导航到对应的用户页面，并在页面上显示相应的用户 ID。

1. 状态（State）：使用 `react-router-dom` 提供的 `useNavigate` 钩子，你可以在路由之间传递状态数据。你可以将状态数据作为对象传递给 `navigate` 方法的第二个参数。例如：

```javascript
jsxCopy code
import { BrowserRouter as Router, Route, Link, useNavigate } from 'react-router-dom';

function App() {
  return (
    <Router>
      <nav>
        <ul>
          <li>
            <Link to="/users" state={{ from: 'home' }}>Users</Link>
          </li>
        </ul>
      </nav>

      <Route path="/users" exact>
        <Users />
      </Route>
    </Router>
  );
}

function Users() {
  const navigate = useNavigate();

  const handleClick = () => {
    navigate('/users', { state: { from: 'about' } });
  };

  return (
    <div>
      <h1>Users</h1>
      <button onClick={handleClick}>Go to Users (from about)</button>
    </div>
  );
}
```

在上面的示例中，当你点击按钮时，会导航到 `/users` 路径，并在页面上显示传递的状态数据。

这些是在 `react-router-dom` 版本 6 中常用的几种方式来在路由之间传递数据。你可以根据你的需求选择适合的方式来实现数据的传递。

### 15.`axios`初步封装

Axios 是一个基于 Promise 的 JavaScript HTTP 客户端，用于在浏览器和 Node.js 中发送 HTTP 请求。它提供了许多强大的功能，如拦截请求和响应、转换请求和响应数据、取消请求、自动转换 JSON 数据等。

要使用 Axios，首先需要在项目中安装它。可以通过 npm 或 yarn 进行安装：

使用 npm：

```
Copy code
npm install axios
```

使用 yarn：

```
csharpCopy code
yarn add axios
```

安装完成后，可以在 JavaScript 代码中导入 Axios：

```
javascriptCopy code
import axios from 'axios';
```

现在，我们来看一下如何封装和使用 Axios。

封装 Axios：

通常情况下，我们会对 Axios 进行一些封装，以方便在整个项目中使用。以下是一个简单的封装示例：

```
javascriptCopy code
import axios from 'axios';

const instance = axios.create({
  baseURL: 'https://api.example.com', // 设置基础URL
  timeout: 5000, // 请求超时时间
});

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // 在发送请求之前做些处理，例如添加请求头等
    return config;
  },
  error => {
    // 处理请求错误
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  response => {
    // 对响应数据进行处理
    return response.data;
  },
  error => {
    // 处理响应错误
    return Promise.reject(error);
  }
);

export default instance;
```

在这个封装中，我们创建了一个 Axios 实例，设置了基础 URL 和请求超时时间。还定义了请求拦截器和响应拦截器，以便在发送请求和处理响应时进行一些处理。

使用 Axios：

封装完成后，就可以在项目中使用 Axios 进行 HTTP 请求了。以下是一些简单的示例：

发送 GET 请求：

```
javascriptCopy code
import axios from './axiosInstance'; // 导入封装的 Axios 实例

axios.get('/users')
  .then(response => {
    console.log(response);
  })
  .catch(error => {
    console.error(error);
  });
```

发送 POST 请求：

```
javascriptCopy code
import axios from './axiosInstance';

axios.post('/users', { name: 'John Doe', age: 30 })
  .then(response => {
    console.log(response);
  })
  .catch(error => {
    console.error(error);
  });
```

上述示例中，我们通过调用 Axios 实例的 `get` 和 `post` 方法来发送 GET 和 POST 请求。可以通过 `.then` 来处理成功响应的数据，通过 `.catch` 来处理请求或响应的错误。

这只是 Axios 的基本用法，你可以根据实际需求进行更多的配置和使用。Axios 提供了丰富的功能和选项，可以查阅官方文档以了解更多详情：[Axios GitHub 仓库](https://github.com/axios/axios)

### 16.函数组件规范

- 子组件

  ```typescript
  type ChildCompProps = {
    ...
  }
  
  const ChildComp: React.FC<ChildCompProps> = ({
    ...
  })=>{
  		return(<> </>)
  }
                                               
  export default ChildComp
  ```

- 页面组件

  ```typescript
  const IndexPage = () => {
  		return (<></>)
  }
  
  export default IndexPage
  ```

- hook

  ```typescript
  export const useRequest = () => {
  
  }
  ```

### 17.使用`rollup`打包`react-ts`项目的自定义`hooks`并发布在`npm`

1. 创建react-ts项目

   ```shell
   npx create-react-app 项目名字 --template typescript
   # 或
   npm init react-app 项目名字 --template typescript
   # 或
   yarn create react-app 项目名字 --template typescript
   ```

2. 安装`rollup`

   为什么不选择`webpack`进行打包？

   rollup相对webpack更轻量，其构建的代码并不会像webpack一样被注入大量的webpack内部结构，而是尽量的精简保持代码原有的状态。

   如果你要开发js库，那webpack的繁琐和打包后的文件体积就不太适用了。有需求就有工具，所以rollup的产生就是针对开发js库的。

   ```shell
   npm install rollup --save-dev
   
   # 或者
   
   yarn add rollup --dev
   ```

   rollup的核心包既包括核心代码也包括cli指令工具集，所以他不需要像webpack或gulp一样安装webpack和webpack-cli。

3. 安装`@rollup/plugin-typescript`

   该插件的作用是用来处理`.ts`或`.tsx`文件

   ```shell
   npm install @rollup/plugin-typescript --save-dev
   
   # 或者
   
   yarn add @rollup/plugin-typescript --dev
   ```

4. 创建配置文件`rollup.config.mjs`

   ```js
   import typescript from '@rollup/plugin-typescript';
   
   export default {
     input: 'src/index.tsx',  // 入口 TypeScript 文件路径
     output: {
       file: 'dist/bundle.js',  // 输出文件路径和文件名
       format: 'es',  // 输出模块格式，例如 CommonJS (cjs) 或 ES 模块 (es)
     },
     plugins: [
       typescript(),  // 使用 @rollup/plugin-typescript 插件处理 TypeScript 文件
     ],
   };
   ```

5. 可能会出现的报错

   报错 Node tried to load your configuration file as CommonJS even though it is likely an ES module.

   原因：

   这个错误通常出现在你的配置文件被错误地识别为 CommonJS（即旧的 Node.js 模块系统）而不是 ES 模块（即 ECMAScript 模块）的情况下。从 Node.js 版本 13 开始，Node.js 默认支持 ES 模块。因此，当你的配置文件被识别为 ES 模块时，Node.js 会尝试以 CommonJS 的方式加载它，这就会导致出现该错误。

   要解决这个问题，你可以尝试以下几种方法：

   1. 确保你的配置文件是一个有效的 ES 模块。确保文件的扩展名是 `.mjs`（例如 `config.mjs`），或者在你的配置文件中使用 ES 模块的语法（例如 `import` 和 `export` 关键字）。
   2. 如果你的配置文件是一个 CommonJS 模块（使用 `require` 和 `module.exports`），你可以尝试将它转换为一个 ES 模块。可以通过更改文件的扩展名为 `.mjs` 或者在文件中使用 Babel 等工具进行转换。
   3. 如果你的 Node.js 版本较旧，不支持 ES 模块的话，你可以尝试升级到较新的版本。Node.js 14+ 版本支持原生的 ES 模块系统。
   4. 如果你希望继续使用 CommonJS 模块，可以在你的配置文件中添加以下代码，明确告诉 Node.js 使用 CommonJS 加载你的配置文件：

6. 创建打包的入口文件

   如：

   ```ts
   // useCount.ts
   import {useState} from "react";
   
   export const useCount = () => {
     const [count,setCount] = useState<number>(0)
     const addOneCount = ()=>{
       let num = count+1
       setCount(num)
     }
     return {
       count,
       addOneCount,
     }
   }
   
   // main.ts
   import {useCount} from './useCount'
   
   export {useCount}
   ```

7. 发布在`npm`库

   1. 初始化项目：在您的项目目录中使用终端或命令提示符窗口运行以下命令来初始化一个新的npm包：

      ```
      npm init
      ```

      这将引导您完成一系列问题，例如包名称、版本、描述等。您可以根据需要提供相关信息，或者按回车键接受默认值。

   2. 创建文件结构：在您的项目目录中创建必要的文件和文件夹结构。通常，一个npm包至少应包含一个入口文件（例如`index.js`），以及其他您认为必要的文件和文件夹。

   3. 实现功能：根据您的包的目标和用途，编写实现功能的代码。您可以使用任何适合的编程语言和框架。

   4. 定义包依赖：如果您的包依赖于其他npm包，可以使用`npm install`命令安装它们。在项目根目录下运行以下命令：

      ```
      npm install <package-name>
      ```

      该命令将自动将依赖项添加到您的`package.json`文件的`dependencies`部分。

   5. 编写文档：为了帮助其他开发者正确使用您的npm包，编写详细的文档是一个好习惯。您可以创建一个`README.md`文件来描述如何安装、使用和配置您的包。

   6. 测试包：编写单元测试和集成测试来验证您的npm包的功能和稳定性。您可以使用适合您选择的编程语言和框架的测试工具。

   7. 发布包：当您准备好发布您的npm包时，首先需要在[npm官网](https://www.npmjs.com/)上注册一个账号。登录后，您可以使用以下命令发布您的包：

      ```
      npm publish
      ```

      这将将您的包上传到npm注册表，并使其可供其他人使用。

   请注意，发布npm包是一个重要的过程，因此在发布之前确保您的包是稳定和可靠的。同时，遵循最佳实践和安全原则来保护您的代码和用户的安全。

   这只是一个简单的概述，帮助您开始开发一个npm包。具体的开发过程可能会根据您的项目需求和选择的技术栈而有所不同。

在npm包中，`package.json`文件包含了描述包的元数据和配置信息。以下是`package.json`中常见的属性值：

1. `name`：包的名称，必须是唯一的。*
2. `version`：包的版本号，遵循语义化版本规范。*
3. `description`：包的简要描述。
4. `keywords`：关键字数组，用于描述包的特性和功能。*
5. `author`：包的作者信息。
6. `license`：包的许可证信息。
7. `repository`：包的代码仓库信息，包括类型（`type`）和URL（`url`）。
8. `bugs`：报告问题的URL或邮箱地址。
9. `homepage`：包的主页URL。
10. `dependencies`：指定包的生产环境依赖项及其版本号。*
11. `devDependencies`：指定包的开发环境依赖项及其版本号。
12. `peerDependencies`：指定包的对等依赖项及其版本号。*
13. `scripts`：定义可以通过`npm run`命令执行的脚本命令。*一般定义打包的命令
14. `main`：指定包的入口文件。*
15. `module`：指定包的ES模块入口文件。*
16. `typings`：指定包的TypeScript类型声明文件。*
17. `files`：定义包发布时需要包含在内的文件和目录。*
18. `engines`：指定包所需的Node.js版本范围。*
19. `peerDependenciesMeta`：用于定义对等依赖的元数据。
20. `publishConfig`：用于配置发布包时的行为，如访问权限和发布标签。

这些是`package.json`中最常见的属性值，你可以根据自己的需求和项目要求进行适当的配置。

### 18.`map`类型的`useState`

在 TypeScript 中，你可以使用泛型来为 `useState` 声明一个 `Map` 类型的 `useState`。首先，你需要定义一个 `Map` 类型，然后在 `useState` 的泛型参数中使用这个类型。假设我们要在 React 组件中使用 `useState` 来维护一个键值对的 `Map`，以下是实现的示例代码：

```tsx
import React, { useState } from 'react';

// 定义 Map 类型
type MyMap = Map<string, number>;

const MyComponent = () => {
  // 使用 Map 类型的 useState，初始值为空 Map
  const [myMap, setMyMap] = useState<MyMap>(new Map());

  // 添加键值对的处理函数
  const addKeyValuePair = (key: string, value: number) => {
    const newMap = new Map(myMap);
    newMap.set(key, value);
    setMyMap(newMap);
  };

  // 示例：添加一个键值对
  addKeyValuePair('key1', 10);

  return (
    <div>
      {Array.from(myMap).map(([key, value]) => (
        <div key={key}>
          {key}: {value}
        </div>
      ))}
    </div>
  );
};

export default MyComponent;
```

在上面的示例中，我们首先定义了一个 `MyMap` 类型，它是一个键为字符串，值为数字的 `Map`。然后，我们使用 `useState` 来声明 `myMap` 和 `setMyMap` 这两个状态。初始状态为空 `Map`，并通过 `addKeyValuePair` 函数来添加键值对。

需要注意的是，`Map` 的特性决定它是一个引用类型，所以在更新状态时，我们需要先创建一个新的 `Map` 对象，并通过新的 `Map` 对象来更新状态，而不是直接在现有的 `Map` 上进行更改。这是为了确保 React 可以正确检测状态的变化，并进行渲染更新。

### 19.不确定属性个数的对象类型的`useState`——Record

当你在 TypeScript 中使用 `useState` 来维护一个不确定属性个数的对象类型时，你可以使用泛型，并为 `useState` 声明一个合适的对象类型。在这种情况下，你可以使用 TypeScript 的 `Record` 类型，它允许你定义一个具有指定键类型的对象，并将它们映射到特定的值类型。

假设我们要在 React 组件中使用 `useState` 来维护一个属性个数不确定的对象类型，以下是一个示例代码：

```tsx
import React, { useState } from 'react';

// 定义 Record 类型，其中键为字符串类型，值为数字类型
type MyObject = Record<string, number>;

const MyComponent = () => {
  // 使用 Record 类型的 useState，初始值为空对象
  const [myObject, setMyObject] = useState<MyObject>({});

  // 添加属性的处理函数
  const addProperty = (key: string, value: number) => {
    setMyObject((prevObject) => ({ ...prevObject, [key]: value }));
  };

  // 示例：添加一个属性
  addProperty('property1', 10);

  return (
    <div>
      {Object.entries(myObject).map(([key, value]) => (
        <div key={key}>
          {key}: {value}
        </div>
      ))}
    </div>
  );
};

export default MyComponent;
```

在上面的示例中，我们定义了一个 `MyObject` 类型，它是一个键为字符串类型，值为数字类型的对象。然后，我们使用 `useState` 来声明 `myObject` 和 `setMyObject` 这两个状态。初始状态为空对象，并通过 `addProperty` 函数来添加属性。

在更新状态时，我们使用了函数形式的 `setMyObject`，确保我们可以正确地更新之前的状态，并添加新的属性而不是覆盖原有的属性。

需要注意的是，由于 `Record` 类型允许键为任意字符串类型，所以你可以添加任意数量和任意名称的属性。但是在使用时请注意确保键的唯一性，以避免不必要的覆盖或混淆。

### 20.`react`中如果只改变数组中的一个元素，如何避免遍历该数组后其他子组件的不必要的重新渲染

在 React 中，如果你只改变数组中的一个元素，并且希望避免其他子组件的不必要重新渲染，可以使用不可变更新的方式来更新数组。不可变更新是指创建一个新的数组副本，而不是直接在原始数组上进行修改。这样可以确保 React 可以正确检测到状态的变化，并只重新渲染受影响的组件。

你可以使用 `map()` 方法来遍历数组，并在需要修改的元素上执行更新操作。同时，使用 `...` 展开运算符（spread operator）来创建新的数组副本。

下面是一个示例代码：

```jsx
import React, { useState } from 'react';

const MyComponent = () => {
  const [myArray, setMyArray] = useState([1, 2, 3, 4, 5]);

  const updateElement = (index, newValue) => {
    setMyArray((prevArray) => {
      // 创建新的数组副本
      const newArray = [...prevArray];
      // 更新指定索引的元素
      newArray[index] = newValue;
      return newArray;
    });
  };

  return (
    <div>
      {myArray.map((item, index) => (
        <div key={index}>
          {/* 假设点击按钮后，只更新数组中的第一个元素 */}
          {index === 0 ? (
            <button onClick={() => updateElement(index, item + 1)}>Update</button>
          ) : (
            <span>{item}</span>
          )}
        </div>
      ))}
    </div>
  );
};

export default MyComponent;
```

在上面的示例中，我们使用 `map()` 方法遍历数组，并在需要更新的元素（第一个元素）上放置一个按钮。当按钮被点击时，调用 `updateElement` 函数来更新数组中的第一个元素，并且只重新渲染受影响的组件。其他未更改的元素将保持不变，不会触发重新渲染。

这种方式可以有效地避免不必要的重新渲染，并提高应用性能。

### 21.`memo`的作用和用法

`React.memo` 是 React 中的高阶组件（Higher-Order Component，HOC），用于优化函数组件的性能。它可以在组件渲染时对组件的 props 进行浅层比较，如果前后 props 没有发生变化，则跳过重新渲染，从而避免不必要的渲染和提高性能。

`React.memo` 的作用是用于记忆组件的渲染结果。当组件被包裹在 `React.memo` 中时，它会将组件的输出结果缓存起来，并在下一次渲染时，比较新的 props 是否与之前的 props 相同。如果相同，`React.memo` 将直接返回之前缓存的渲染结果，而不会重新渲染组件。

用法：

```tsx
import React from 'react';

// 普通的函数组件
const MyComponent = ({ value }) => {
  console.log('MyComponent is re-rendered');
  return <div>{value}</div>;
};

// 使用 React.memo 包裹组件
const MemoizedComponent = React.memo(MyComponent);

const ParentComponent = () => {
  const [count, setCount] = React.useState(0);

  const handleClick = () => {
    setCount(count + 1);
  };

  return (
    <div>
      <MemoizedComponent value={count} />
      <button onClick={handleClick}>Increment</button>
    </div>
  );
};

export default ParentComponent;
```

在上面的示例中，`MyComponent` 是一个普通的函数组件，它会在每次渲染时打印一条日志。然后，我们使用 `React.memo` 包裹 `MyComponent`，创建了一个 `MemoizedComponent`。`ParentComponent` 中使用了 `MemoizedComponent`，每次点击 "Increment" 按钮时，`ParentComponent` 的状态 `count` 会增加，但 `MemoizedComponent` 并不会重新渲染，因为 `value` 的值没有发生变化。

`React.memo` 默认会对所有的 props 进行浅层比较。如果需要自定义比较逻辑，可以使用第二个参数作为比较函数。这个比较函数接收两个参数 `prevProps` 和 `nextProps`，并返回一个布尔值，表示是否相等。如果返回 `true`，则表示前后 props 相等，组件将跳过重新渲染。

请注意，尽管 `React.memo` 可以优化性能，但不是在所有情况下都需要使用。只有在确实存在性能问题，并且组件渲染较为频繁时，才值得使用 `React.memo` 进行性能优化。

### 22.`useCallback`的作用

`useCallback` 是 React 中的一个 Hook，它用于优化函数的性能，特别是在处理子组件的渲染过程中。当父组件的状态或属性改变时，子组件通常会重新渲染，如果父组件传递给子组件的函数没有经过优化，可能会导致不必要的重新渲染，从而影响应用性能。

`useCallback` 的作用是用于缓存函数，确保在组件重新渲染时，不会重新创建相同的函数实例。它接收一个函数和一个依赖数组，然后返回一个经过缓存的函数。只有依赖数组中的依赖发生变化时，才会重新创建函数实例。

使用 `useCallback` 可以避免在父组件重新渲染时，将相同的函数实例传递给子组件，从而减少子组件的不必要重新渲染。

示例：

```tsx
import React, { useState, useCallback } from 'react';

const ChildComponent = React.memo(({ onClick }) => {
  console.log('ChildComponent is re-rendered');
  return <button onClick={onClick}>Click Me</button>;
});

const ParentComponent = () => {
  const [count, setCount] = useState(0);

  // 使用 useCallback 缓存回调函数
  const handleClick = useCallback(() => {
    setCount((prevCount) => prevCount + 1);
  }, []);

  return (
    <div>
      <ChildComponent onClick={handleClick} />
      <p>Count: {count}</p>
    </div>
  );
};

export default ParentComponent;
```

在上面的示例中，当 `ParentComponent` 重新渲染时，`handleClick` 函数不会重新创建，因为它经过了 `useCallback` 的缓存处理。这样，即使 `ParentComponent` 发生重新渲染，也不会导致 `ChildComponent` 的重新渲染，除非 `handleClick` 的依赖发生变化（目前依赖数组为空，表示没有依赖）。

请注意，尽管 `useCallback` 可以优化性能，但不是在所有情况下都需要使用。只有在确实存在性能问题，并且函数被频繁传递给子组件时，才值得使用 `useCallback` 进行性能优化。

### 23.`react ts `项目的`manifest.json`文件的作用

在 React TypeScript 项目中，`manifest.json` 文件是一个用于配置 Progressive Web Apps (PWA) 的清单文件。PWA 是一种使用现代 Web 技术来提供类似原生应用的体验的 Web 应用程序。`manifest.json` 文件允许你在移动设备上添加应用图标、设置应用的名称、主题颜色、启动方式等信息，使得用户可以将网站添加到主屏幕并在离线状态下访问。

主要作用包括：

1. **应用图标和启动画面**：通过 `manifest.json` 文件，你可以指定不同尺寸的应用图标和启动画面，这样在添加到主屏幕时，应用会有一个类似原生应用的图标，以及在启动时显示的画面。
2. **应用名称**：通过 `manifest.json` 文件，你可以指定应用的名称，这样在添加到主屏幕时，用户会看到你指定的名称而不是默认的网站名称。
3. **主题颜色**：通过 `manifest.json` 文件，你可以指定应用的主题颜色，这将影响浏览器工具栏、地址栏和导航条等的颜色，提供更加一致的用户体验。
4. **启动方式**：通过 `manifest.json` 文件，你可以指定应用的启动方式，比如全屏模式、独立窗口模式等。
5. **离线访问**：通过使用 Service Worker 和 `manifest.json` 文件，你可以实现应用的离线访问，使得用户在没有网络连接的情况下仍然能够访问应用的内容。

示例 `manifest.json` 文件：

```json
{
  "name": "My React App",
  "short_name": "React App",
  "start_url": "/",
  "display": "standalone",
  "theme_color": "#ffffff",
  "background_color": "#ffffff",
  "icons": [
    {
      "src": "icon-192x192.png",
      "type": "image/png",
      "sizes": "192x192"
    },
    {
      "src": "icon-512x512.png",
      "type": "image/png",
      "sizes": "512x512"
    }
  ]
}
```

要使 `manifest.json` 文件生效，你需要在 HTML 页面中添加一个链接标签，指向该文件：

```html
<link rel="manifest" href="/manifest.json" />
```

请注意，PWA 的支持和表现方式可能会因浏览器和设备而异。在使用 `manifest.json` 文件前，建议先了解目标浏览器的兼容性和支持情况。
