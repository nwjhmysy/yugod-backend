## 1.css事件穿透

```css
pointer-events: none;//穿透
pointer-events: auto;//默认
```

效果：如果某一元素带有该css属性，那么点击该元素**不触发**该元素点击事件，而是**触发**父级元素的点击事件。

## 2.css阻止事件冒泡

```
onClick={(e) => e.stopPropagation()}
```

效果：子元素的点击事件函数中如果添加了该语句，那么不会触发外层父元素的点击事件。

## 3.react 滚动到指定元素位置

类似于锚点。

```javascript
let element = document.getElementById(elementID);
if(element) { 
  element.scrollIntoView();
}
```

`scrollIntoView(option?)`

```
option:
{
     behavior?: 'auto' | 'instant' | 'smooth', //设置过度效果
     block?: 'start' | 'end' ｜ 'center' | 'nearest', //定位至元素所在屏幕上的位置
}
```

## 2.flex布局

### justify-content,justify-items和justify-self的区别

ChatGPT



`justify-content`, `justify-items` 和 `justify-self` 是 Flexbox 布局中用于对齐子元素的属性，但它们作用于不同的层级和对象上。

1. **`justify-content`**:
   - 作用对象：父容器（Flex容器）。
   - 作用：控制Flex容器内的所有子元素在主轴上的对齐方式。
   - 值：
     - `flex-start`：子元素向父容器的起始位置对齐。
     - `flex-end`：子元素向父容器的结束位置对齐。
     - `center`：子元素在父容器中居中对齐。
     - `space-between`：子元素在父容器中平均分布，首尾没有间隙。
     - `space-around`：子元素在父容器中平均分布，首尾有间隙。
     - `space-evenly`：子元素在父容器中平均分布，每个元素两侧间隔相等。
2. **`justify-items`**:
   - 作用对象：父容器（Flex容器）。
   - 作用：设置Flex容器内所有子元素的默认对齐方式（在没有单独设置 `align-self` 的情况下）。
   - 值：
     - `stretch`：子元素将会拉伸以填满交叉轴。
     - `flex-start`：子元素向交叉轴的起始位置对齐。
     - `flex-end`：子元素向交叉轴的结束位置对齐。
     - `center`：子元素在交叉轴上居中对齐。
     - `baseline`：子元素在基线上对齐。
3. **`justify-self`**:
   - 作用对象：单个子元素（Flex项目）。
   - 作用：设置单个子元素在主轴上的对齐方式，会覆盖 `justify-items` 的设置。
   - 值：
     - `stretch`：子元素将会拉伸以填满主轴。
     - `flex-start`：子元素向主轴的起始位置对齐。
     - `flex-end`：子元素向主轴的结束位置对齐。
     - `center`：子元素在主轴上居中对齐。

简而言之：

- `justify-content` 控制整个 Flex 容器内的所有项目在主轴上的对齐方式。
- `justify-items` 设置所有项目的默认对齐方式（如果没有单独设置 `align-self`）。
- `justify-self` 控制单个项目在主轴上的对齐方式，会覆盖 `justify-items` 的设置。

需要注意的是，这些属性只在 Flexbox 布局中生效，而在传统的块级布局（非 Flexbox）中将不会起作用。



### align-content，align-self和align-self的区别

ChatGPT



在 Flexbox 布局中，`align-content`、`align-items` 和 `align-self` 是用于控制项目在交叉轴上的对齐方式的属性。

1. **`align-content`**：
   - 作用对象：父容器（Flex容器）。
   - 作用：当交叉轴上有多条轴线时，`align-content` 控制这些轴线的对齐方式。
   - 值：
     - `flex-start`：所有轴线在交叉轴的起始位置对齐。
     - `flex-end`：所有轴线在交叉轴的结束位置对齐。
     - `center`：所有轴线在交叉轴上居中对齐。
     - `space-between`：所有轴线在交叉轴上平均分布，首尾没有间隙。
     - `space-around`：所有轴线在交叉轴上平均分布，首尾有间隙。
     - `stretch`：轴线将会被拉伸以填满交叉轴。
2. **`align-items`**：
   - 作用对象：父容器（Flex容器）。
   - 作用：控制Flex容器内所有项目在交叉轴上的对齐方式（在没有单独设置 `align-self` 的情况下）。
   - 值：
     - `stretch`：项目将会被拉伸以填满交叉轴。
     - `flex-start`：项目向交叉轴的起始位置对齐。
     - `flex-end`：项目向交叉轴的结束位置对齐。
     - `center`：项目在交叉轴上居中对齐。
     - `baseline`：项目在基线上对齐。
3. **`align-self`**：
   - 作用对象：单个项目（Flex项目）。
   - 作用：控制单个项目在交叉轴上的对齐方式，会覆盖 `align-items` 的设置。
   - 值：
     - `stretch`：项目将会被拉伸以填满交叉轴。
     - `flex-start`：项目向交叉轴的起始位置对齐。
     - `flex-end`：项目向交叉轴的结束位置对齐。
     - `center`：项目在交叉轴上居中对齐。
     - `baseline`：项目在基线上对齐。

简而言之：

- `align-content` 控制所有轴线的对齐方式（在有多条轴线时）。
- `align-items` 控制所有项目的对齐方式（在没有单独设置 `align-self` 的情况下）。
- `align-self` 控制单个项目的对齐方式，会覆盖 `align-items` 的设置。

需要注意的是，这些属性只在 Flexbox 布局中生效，而在传统的块级布局（非 Flexbox）中将不会起作用。
