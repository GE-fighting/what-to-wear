# StyleSense UI 共享组件

此目录包含所有页面共享的主题配置和 UI 组件。

## 📁 文件说明

### `theme.html`
统一的主题配置文件，包括：
- 字体引入（Noto Sans SC + Material Icons）
- Tailwind CSS 配置
- 颜色系统
- 全局样式

### `header.html`
统一的顶部导航栏组件，包括：
- Logo
- 主导航（记录穿搭、我的衣橱、风格灵感、穿搭分析）
- 搜索框
- 收藏、通知、用户头像

## 🎨 主题颜色系统

```css
primary: #000000           /* 主色调 - 黑色 */
primary-hover: #333333     /* 主色调悬停态 */
background-light: #F8F9FA  /* 浅色模式背景 */
background-dark: #121212   /* 深色模式背景 */
card-light: #FFFFFF        /* 浅色模式卡片 */
card-dark: #1E1E1E         /* 深色模式卡片 */
text-light-primary: #1F2937    /* 浅色模式主要文本 */
text-dark-primary: #F9FAFB     /* 深色模式主要文本 */
text-light-secondary: #6B7280  /* 浅色模式次要文本 */
text-dark-secondary: #9CA3AF   /* 深色模式次要文本 */
border-light: #E5E7EB      /* 浅色模式边框 */
border-dark: #374151       /* 深色模式边框 */
```

## 📖 使用方法

### 方法一：直接复制（推荐用于静态 HTML）

在每个页面的 `<head>` 中复制 `theme.html` 的内容：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>页面标题</title>
  
  <!-- 复制 theme.html 的全部内容到这里 -->
  <link href="https://fonts.googleapis.com/css2?family=Material+Icons+Outlined" rel="stylesheet"/>
  <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+SC:wght@400;500;700&display=swap" rel="stylesheet"/>
  <script src="https://cdn.tailwindcss.com?plugins=forms,typography,container-queries"></script>
  <script>
    tailwind.config = { ... };
  </script>
  <style>
    body { font-family: 'Noto Sans SC', sans-serif; }
    /* ... */
  </style>
</head>
```

在 `<body>` 开头复制 `header.html` 的内容。

### 方法二：使用 JavaScript 动态加载（适用于开发环境）

```html
<head>
  <script>
    // 加载主题配置
    fetch('/UI/shared/theme.html')
      .then(res => res.text())
      .then(html => {
        document.head.innerHTML += html;
      });
  </script>
</head>
<body>
  <div id="header"></div>
  <script>
    // 加载导航栏
    fetch('/UI/shared/header.html')
      .then(res => res.text())
      .then(html => {
        document.getElementById('header').innerHTML = html;
      });
  </script>
</body>
```

### 方法三：使用模板引擎（适用于构建工具）

如果使用构建工具（如 Vite、Webpack），可以使用模板引擎如 EJS、Handlebars 等来引入共享组件。

## 🔄 更新策略

当需要修改主题或导航栏时：
1. 只需修改 `shared/` 目录下的文件
2. 使用脚本批量更新所有页面（或手动更新）

## ⚠️ 注意事项

1. **导航链接**: 根据页面的实际位置调整 `header.html` 中的链接路径
2. **当前页面高亮**: 在具体页面中添加 `border-b-2 border-primary` 类来高亮当前导航项
3. **响应式**: 所有组件已支持响应式设计
4. **深色模式**: 使用 `class="dark"` 在 `<html>` 标签上切换深色模式

## 📝 示例完整页面结构

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>页面标题 - StyleSense</title>
  
  <!-- 主题配置 -->
  <link href="https://fonts.googleapis.com/css2?family=Material+Icons+Outlined" rel="stylesheet"/>
  <link href="https://fonts.googleapis.com/css2?family=Noto+Sans+SC:wght@400;500;700&display=swap" rel="stylesheet"/>
  <script src="https://cdn.tailwindcss.com?plugins=forms,typography,container-queries"></script>
  <script>
    tailwind.config = {
      darkMode: "class",
      theme: {
        extend: {
          colors: {
            primary: "#000000",
            // ... 其他颜色配置
          },
          fontFamily: {
            display: ["'Noto Sans SC'", "sans-serif"],
          },
        },
      },
    };
  </script>
  <style>
    body { font-family: 'Noto Sans SC', sans-serif; }
  </style>
</head>
<body class="bg-background-light dark:bg-background-dark font-display">
  <!-- 导航栏 -->
  <header class="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
    <!-- header.html 内容 -->
  </header>
  
  <!-- 主要内容 -->
  <main class="flex-grow max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full py-8">
    <!-- 页面内容 -->
  </main>
  
  <!-- 页脚 -->
  <footer class="bg-card-light dark:bg-card-dark mt-auto">
    <div class="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
      <p class="text-center text-sm text-text-light-secondary dark:text-text-dark-secondary">
        @2024 StyleSense. All rights reserved.
      </p>
    </div>
  </footer>
</body>
</html>
```
