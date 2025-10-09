# StyleSense UI - AI 助手开发指南

> 本文档专门为 AI 助手编写，用于快速理解和生成 StyleSense UI 代码。

## 🎯 快速开始

### 你需要知道的
- **技术栈**: HTML + Tailwind CSS + Alpine.js + Material Icons
- **字体**: Noto Sans SC (中文字体)
- **设计风格**: 黑白极简主义，基于 4px 网格系统
- **主题**: 支持深色/浅色模式切换

### 文件结构
```
UI/
├── agent.md              # 📋 本文件 - AI 助手指南
├── design-system/        # 🎨 设计系统 (必读)
│   ├── design-tokens.md  # 设计令牌参考
│   ├── typography.md     # 字体规范
│   ├── color-system.md   # 颜色系统
│   ├── component-library.md # 组件库
│   └── spacing-grid.md   # 间距网格
├── demo/                 # 📺 UI 实例演示
├── shared/               # 🔧 共享组件和主题
│   ├── theme.html        # 主题配置
│   └── header.html       # 导航栏
├── apps/                 # 📱 业务功能模块
│   ├── auth/             # 认证相关
│   ├── index.html        # 主页应用
│   ├── wardrobe/         # 衣橱管理
│   ├── record-style/     # 穿搭记录
│   ├── settings/         # 设置中心
│   └── notifications/    # 通知中心
```

## 📁 目录结构详解

## 🎨 核心设计规范

### 颜色系统
```css
--color-primary: #000000;           /* 主色 - 纯黑 */
--color-primary-hover: #333333;     /* 主色悬停 */
--color-background-light: #F8F9FA;  /* 浅色背景 */
--color-background-dark: #121212;   /* 深色背景 */
--color-text-primary-light: #1F2937; /* 浅色主文字 */
--color-text-primary-dark: #F9FAFB;  /* 深色主文字 */
```

### 间距系统 (4px 基础)
```css
--spacing-sm: 8px;     /* 小间距 */
--spacing-md: 16px;    /* 中间距 */
--spacing-lg: 24px;    /* 大间距 */
--spacing-xl: 32px;    /* 超大间距 */
```

### 字体系统
```css
font-family: 'Noto Sans SC', sans-serif;
font-size: 16px (BASE), 14px (SM), 18px (LG)
font-weight: 400 (normal), 500 (medium), 600 (semibold), 700 (bold)
```

## 🧩 常用组件模板

### 按钮
```html
<!-- 主要按钮 -->
<button class="bg-primary hover:bg-primary-hover text-white px-4 py-2 rounded-lg transition-colors flex items-center gap-2">
  <span class="material-icons-outlined">add</span>
  添加
</button>

<!-- 次要按钮 -->
<button class="bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark hover:bg-gray-50 dark:hover:bg-gray-800 text-text-light-primary dark:text-text-dark-primary px-4 py-2 rounded-lg transition-colors">
  取消
</button>
```

### 表单
```html
<div class="space-y-6">
  <div>
    <label class="block text-sm font-medium text-text-light-primary dark:text-text-dark-primary mb-2">
      标签名称
    </label>
    <input type="text" class="w-full px-3 py-2 bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-lg focus:outline-none focus:ring-2 focus:ring-primary text-text-light-primary dark:text-text-dark-primary">
  </div>
</div>
```

### 卡片
```html
<div class="bg-card-light dark:bg-card-dark rounded-lg shadow-sm border border-border-light dark:border-border-dark p-6 hover:shadow-md transition-shadow">
  <h3 class="text-lg font-medium text-text-light-primary dark:text-text-dark-primary mb-2">卡片标题</h3>
  <p class="text-text-light-secondary dark:text-text-dark-secondary mb-4">卡片内容描述</p>
  <div class="flex gap-2">
    <button class="btn btn-primary btn-sm">查看详情</button>
    <button class="btn btn-secondary btn-sm">编辑</button>
  </div>
</div>
```

### 导航高亮
```html
<!-- 当前页面高亮 -->
<a href="/wardrobe" class="px-3 py-2 text-sm font-medium text-text-light-primary dark:text-text-dark-primary border-b-2 border-primary">
  我的衣橱
</a>
<!-- 其他页面 -->
<a href="/outfit" class="px-3 py-2 text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary">
  穿搭记录
</a>
```

## 💡 AI 助手使用指南

### 代码生成提示词模板
```
请为 StyleSense UI 创建一个 [页面类型] 页面，要求：
1. 使用 shared/theme.html 的主题配置
2. 使用 shared/header.html 的导航栏结构
3. 遵循黑白极简设计风格
4. 支持深色/浅色模式切换
5. 基于 4px 网格系统的间距
6. 保存路径为 UI/[模块名]/index.html
```

### 关键要点
1. **复用 shared 组件** - 优先使用 shared/ 中的主题和导航
2. **颜色规范** - 主色调为纯黑 (#000000)，避免使用其他颜色
3. **间距基准** - 所有间距基于 4px 倍数
4. **字体统一** - 一律使用 Noto Sans SC
5. **响应式设计** - 移动端优先，支持多种屏幕尺寸
6. **深色模式** - 使用 `dark` class 切换主题

### 常见任务模板

#### 创建新页面
```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>页面标题 - StyleSense</title>
  
  <!-- 复制 shared/theme.html 的内容 -->
  
</head>
<body class="bg-background-light dark:bg-background-dark font-display min-h-screen flex flex-col">
  <!-- 复制 shared/header.html 的内容 -->
  
  <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <!-- 页面内容 -->
  </main>
</body>
</html>
```

#### 添加新组件样式
```css
/* 在 style 标签中添加 */
.new-component {
  @apply bg-card-light dark:bg-card-dark rounded-lg p-4;
}
```

### 文件路径规范
- 页面文件: `UI/apps/[模块]/index.html`
- 子页面: `UI/apps/[模块]/[功能]/index.html`
- 组件复用: `../shared/theme.html`, `../shared/header.html`
- 相对链接: 根据文件位置调整路径（从 apps/ 目录需使用 `../../shared/`）

## 🔧 开发规范

## 📋 快速参考

### 页面模板
```html
<body class="bg-background-light dark:bg-background-dark font-display min-h-screen flex flex-col">
  <!-- shared/header.html -->
  <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
    <!-- 内容 -->
  </main>
</body>
```

### 卡片组件
```html
<div class="bg-card-light dark:bg-card-dark rounded-lg p-6 hover:shadow-md">
  <h3 class="text-text-light-primary dark:text-text-dark-primary">标题</h3>
  <p class="text-text-light-secondary dark:text-text-dark-secondary">内容</p>
</div>
```

### 3. 按钮组件模式
```html
<!-- 主要按钮 -->
<button class="bg-primary hover:bg-primary-hover text-white px-4 py-2 rounded-lg transition-colors flex items-center gap-2">
  <span class="material-icons-outlined">add</span>
  添加
</button>

<!-- 次要按钮 -->
<button class="bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark hover:bg-gray-50 dark:hover:bg-gray-800 text-text-light-primary dark:text-text-dark-primary px-4 py-2 rounded-lg transition-colors">
  取消
</button>
```

### 4. 表单组件模式
```html
<form class="space-y-6">
  <div>
    <label class="block text-sm font-medium text-text-light-primary dark:text-text-dark-primary mb-2">
      标签名称
    </label>
    <input type="text" class="w-full px-3 py-2 bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-lg focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent text-text-light-primary dark:text-text-dark-primary">
  </div>
</form>
```

### 5. 导航高亮模式
```html
<!-- 在导航链接中添加当前页面高亮 -->
<a href="/wardrobe" class="px-3 py-2 text-sm font-medium text-text-light-primary dark:text-text-dark-primary border-b-2 border-primary">
  我的衣橱
</a>
```

## 🔄 组件复用策略

### 1. 使用 shared 组件
```html
<!-- 复制 shared/theme.html 内容到 <head> -->
<!-- 复制 shared/header.html 内容到页面开头 -->
```

### 2. 组件参数化
使用 Alpine.js 的 `x-data` 和模板语法：
```html
<div x-data="{ isOpen: false }">
  <button @click="isOpen = !isOpen">切换</button>
  <div x-show="isOpen" x-transition>
    可切换内容
  </div>
</div>
```

### 3. 样式复用
使用 Tailwind CSS 的 `@apply` 指令：
```html
<style>
  .btn-primary {
    @apply bg-primary hover:bg-primary-hover text-white px-4 py-2 rounded-lg transition-colors;
  }
</style>
```

## 🚀 UI 到 React 转换指南

### 1. 组件映射关系
```
HTML 页面          →    React 组件
----------------→-------------------
wardrobe/index.html → src/pages/Wardrobe/index.tsx
shared/header.html  → src/components/Header.tsx
```

### 2. CSS 样式迁移
```typescript
// Tailwind CSS 类名保持不变
const Card = ({ children, className = "" }) => (
  <div className={`bg-card-light dark:bg-card-dark rounded-lg shadow-sm border border-border-light dark:border-border-dark p-6 hover:shadow-md transition-shadow ${className}`}>
    {children}
  </div>
);
```

### 3. 交互逻辑转换
```html
<!-- HTML + Alpine.js -->
<div x-data="{ count: 0 }">
  <button @click="count++">点击次数: <span x-text="count"></span></button>
</div>
```

```typescript
// React + useState
const [count, setCount] = useState(0);
<button onClick={() => setCount(count + 1)}>点击次数: {count}</button>
```

### 4. 状态管理
```typescript
// 使用 Context 或状态管理库
const ThemeContext = createContext();

const useTheme = () => {
  const context = useContext(ThemeContext);
  if (!context) throw new Error('useTheme must be used within ThemeProvider');
  return context;
};
```

## ⚠️ 注意事项

### 1. 文件路径
- 所有链接使用相对路径
- 共享组件路径: `../shared/`
- 资源文件路径: `../assets/`

### 2. 性能优化
- 使用 CDN 加载外部资源
- 图片使用 WebP 格式
- 合理使用 Alpine.js，避免过度复杂化

### 3. 兼容性
- 支持现代浏览器 (Chrome 90+, Firefox 88+, Safari 14+)
- 使用语义化 HTML 标签
- 提供键盘导航支持

### 4. 测试策略
- 在不同屏幕尺寸下测试响应式
- 测试深色模式切换
- 验证表单提交和交互

## 🛠️ 开发工具

### 推荐的开发工具
- **编辑器**: VS Code + Tailwind CSS IntelliSense
- **浏览器**: Chrome DevTools
- **设计工具**: Figma (用于设计系统维护)

### 有用的浏览器扩展
- Tailwind CSS DevTools
- ColorZilla (颜色选择器)
- Responsive Viewer (响应式测试)

## 📝 AI 开发助手使用指南

当为 StyleSense UI 生成代码时，请遵循以下原则：

1. **优先使用已有的颜色系统和设计令牌**
2. **保持 HTML 结构的语义化和可访问性**
3. **使用一致的前缀和命名约定**
4. **确保响应式设计和深色模式支持**
5. **复用 shared 目录中的组件和样式**
6. **保持与设计系统的一致性**

### 代码生成提示词示例
```
请为 StyleSense UI 创建一个添加衣物的页面，要求：
- 使用 UI/shared/theme.html 中的主题配置
- 使用 UI/shared/header.html 的导航栏结构
- 遵循设计系统的颜色和字体规范
- 包含表单验证和响应式设计
- 支持深色模式切换
- 保存路径为 UI/apps/wardrobe/add/index.html
```

## 📚 参考资源

### 设计系统文档
- `design-system/design-tokens.md` - 完整设计令牌
- `design-system/typography.md` - 字体系统详情  
- `design-system/color-system.md` - 颜色使用规则
- `design-system/component-library.md` - 组件库规范
- `design-system/spacing-grid.md` - 间距网格系统

### 示例页面
- `apps/index.html` - 主应用页面示例
- `demo/index.html` - 演示中心导航
- `apps/wardrobe/index.html` - 衣橱页面示例
- `apps/auth/login/index.html` - 登录页面示例

---

**文档维护**: 当 UI 系统更新时，请同步更新此 agent.md 文件，确保 AI 助手始终获得最新的开发指南。
