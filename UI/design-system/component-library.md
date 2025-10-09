# StyleSense 组件库 (Component Library)

> 组件库定义了 StyleSense 所有可复用的 UI 组件，确保整个产品的视觉一致性和交互统一性。

## 📋 目录
- [组件概览](#组件概览)
- [基础组件](#基础组件)
- [表单组件](#表单组件)
- [展示组件](#展示组件)
- [导航组件](#导航组件)
- [反馈组件](#反馈组件)
- [布局组件](#布局组件)
- [使用指南](#使用指南)

---

## 组件概览

### 组件分类
```
基础组件 (Base)
├── Button          # 按钮组件
├── Input           # 输入框组件
├── Icon            # 图标组件
├── Text            # 文本组件
└── Divider         # 分割线组件

表单组件 (Forms)
├── Form            # 表单容器
├── Select          # 下拉选择
├── Checkbox        # 复选框
├── Radio           # 单选框
├── Switch          # 开关
└── Textarea        # 文本域

展示组件 (Display)
├── Card            # 卡片组件
├── Avatar          # 头像组件
├── Badge           # 徽章组件
├── Tag             # 标签组件
├── List            # 列表组件
└── Table           # 表格组件

导航组件 (Navigation)
├── Header          # 顶部导航
├── Sidebar         # 侧边栏
├── Breadcrumbs     # 面包屑
├── Tabs            # 标签页
└── Pagination      # 分页

反馈组件 (Feedback)
├── Alert           # 警告提示
├── Toast           # 消息提示
├── Modal           # 模态框
├── Tooltip         # 工具提示
└── Loading         # 加载状态

布局组件 (Layout)
├── Container       # 容器组件
├── Grid            # 网格系统
├── Flex            # 弹性布局
├── Stack           # 堆叠布局
└── Spacer          # 间距组件
```

---

## 基础组件

### Button 按钮组件

#### 变体类型
```html
<!-- 主要按钮 -->
<button class="btn btn-primary">
  <span class="material-icons-outlined">add</span>
  添加衣物
</button>

<!-- 次要按钮 -->
<button class="btn btn-secondary">
  取消
</button>

<!-- 文字按钮 -->
<button class="btn btn-ghost">
  编辑
</button>

<!-- 危险按钮 -->
<button class="btn btn-danger">
  <span class="material-icons-outlined">delete</span>
  删除
</button>
```

#### 尺寸规格
```html
<!-- 小尺寸 -->
<button class="btn btn-primary btn-sm">
  小按钮
</button>

<!-- 中等尺寸（默认） -->
<button class="btn btn-primary">
  中按钮
</button>

<!-- 大尺寸 -->
<button class="btn btn-primary btn-lg">
  大按钮
</button>
```

#### 状态样式
```html
<!-- 默认状态 -->
<button class="btn btn-primary">默认</button>

<!-- 悬停状态 -->
<button class="btn btn-primary btn-hover">悬停</button>

<!-- 激活状态 -->
<button class="btn btn-primary btn-active">激活</button>

<!-- 禁用状态 -->
<button class="btn btn-primary" disabled>禁用</button>

<!-- 加载状态 -->
<button class="btn btn-primary btn-loading">
  <span class="loading-spinner"></span>
  加载中...
</button>
```

#### 完整样式定义
```css
/* 按钮基础样式 */
.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-xs);
  padding: var(--spacing-sm) var(--spacing-md);
  border: none;
  border-radius: var(--radius-button);
  font-family: var(--font-primary);
  font-size: var(--font-size-body);
  font-weight: var(--font-weight-medium);
  line-height: var(--line-height-normal);
  cursor: pointer;
  transition: var(--transition-button);
  text-decoration: none;
  outline: none;
  user-select: none;
}

/* 尺寸变体 */
.btn-sm {
  padding: var(--spacing-xs) var(--spacing-sm);
  font-size: var(--font-size-body-small);
  height: var(--size-button-sm);
}

.btn-lg {
  padding: var(--spacing-md) var(--spacing-lg);
  font-size: var(--font-size-body-large);
  height: var(--size-button-lg);
}

/* 颜色变体 */
.btn-primary {
  background-color: var(--color-primary);
  color: var(--color-white);
  box-shadow: var(--shadow-button);
}

.btn-primary:hover {
  background-color: var(--color-primary-hover);
  box-shadow: var(--shadow-button-hover);
}

.btn-secondary {
  background-color: var(--color-surface-primary);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border-primary);
}

.btn-ghost {
  background-color: transparent;
  color: var(--color-text-secondary);
  padding: var(--spacing-xs) var(--spacing-sm);
}

.btn-danger {
  background-color: var(--color-error);
  color: var(--color-white);
}
```

### Input 输入框组件

#### 基础输入框
```html
<!-- 标准输入框 -->
<input type="text" class="input" placeholder="请输入衣物名称">

<!-- 带标签 -->
<div class="form-group">
  <label class="form-label">衣物名称</label>
  <input type="text" class="input" placeholder="例如：白色T恤">
</div>

<!-- 带图标 -->
<div class="input-group">
  <span class="input-icon">
    <span class="material-icons-outlined">search</span>
  </span>
  <input type="text" class="input input-with-icon" placeholder="搜索衣物">
</div>
```

#### 输入框状态
```html
<!-- 默认状态 -->
<input type="text" class="input" placeholder="默认状态">

<!-- 聚焦状态 -->
<input type="text" class="input focused" placeholder="聚焦状态">

<!-- 错误状态 -->
<div class="form-group">
  <input type="text" class="input input-error" placeholder="输入有误">
  <span class="form-error">请输入有效的衣物名称</span>
</div>

<!-- 禁用状态 -->
<input type="text" class="input" placeholder="禁用状态" disabled>

<!-- 成功状态 -->
<input type="text" class="input input-success" placeholder="输入正确">
```

#### 样式定义
```css
/* 输入框基础样式 */
.input {
  width: 100%;
  padding: var(--spacing-sm) var(--spacing-md);
  border: 1px solid var(--color-border-primary);
  border-radius: var(--radius-input);
  background-color: var(--color-surface-primary);
  color: var(--color-text-primary);
  font-family: var(--font-primary);
  font-size: var(--font-size-body);
  line-height: var(--line-height-normal);
  transition: var(--transition-colors);
}

.input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 2px rgba(0, 0, 0, 0.1);
}

.input-error {
  border-color: var(--color-error);
}

.input-error:focus {
  box-shadow: 0 0 0 2px rgba(239, 68, 68, 0.2);
}

.input-success {
  border-color: var(--color-success);
}

/* 输入框组 */
.input-group {
  position: relative;
  display: flex;
  align-items: center;
}

.input-with-icon {
  padding-left: var(--spacing-xl);
}

.input-icon {
  position: absolute;
  left: var(--spacing-sm);
  color: var(--color-text-tertiary);
  pointer-events: none;
}
```

---

## 表单组件

### Form 表单容器
```html
<form class="form">
  <div class="form-section">
    <h3 class="form-section-title">基本信息</h3>
    
    <div class="form-group">
      <label class="form-label" for="name">衣物名称 *</label>
      <input type="text" id="name" class="input" required>
    </div>

    <div class="form-group">
      <label class="form-label" for="category">分类</label>
      <select id="category" class="select">
        <option value="">请选择分类</option>
        <option value="tops">上装</option>
        <option value="bottoms">下装</option>
        <option value="shoes">鞋子</option>
      </select>
    </div>
  </form-section>

  <div class="form-actions">
    <button type="button" class="btn btn-secondary">取消</button>
    <button type="submit" class="btn btn-primary">保存</button>
  </div>
</form>
```

### Select 下拉选择
```html
<!-- 标准选择器 -->
<select class="select">
  <option value="">请选择</option>
  <option value="option1">选项一</option>
  <option value="option2">选项二</option>
</select>

<!-- 多选 -->
<select class="select" multiple size="4">
  <option value="casual">休闲</option>
  <option value="formal">正装</option>
  <option value="sport">运动</option>
  <option value="vintage">复古</option>
</select>

<!-- 禁用状态 -->
<select class="select" disabled>
  <option value="">禁用状态</option>
</select>
```

### Checkbox 复选框
```html
<!-- 单个复选框 -->
<label class="checkbox">
  <input type="checkbox" class="checkbox-input">
  <span class="checkbox-label">接受用户协议</span>
</label>

<!-- 复选框组 -->
<div class="checkbox-group">
  <label class="checkbox">
    <input type="checkbox" name="style" value="casual" class="checkbox-input">
    <span class="checkbox-label">休闲风格</span>
  </label>
  <label class="checkbox">
    <input type="checkbox" name="style" value="formal" class="checkbox-input">
    <span class="checkbox-label">正装风格</span>
  </label>
  <label class="checkbox">
    <input type="checkbox" name="style" value="sport" class="checkbox-input">
    <span class="checkbox-label">运动风格</span>
  </label>
</div>
```

---

## 展示组件

### Card 卡片组件
```html
<!-- 基础卡片 -->
<div class="card">
  <div class="card-header">
    <h3 class="card-title">衣物信息</h3>
  </div>
  <div class="card-body">
    <p class="card-text">这是衣物的详细描述信息</p>
  </div>
  <div class="card-footer">
    <button class="btn btn-ghost">查看详情</button>
  </div>
</div>

<!-- 图片卡片 -->
<div class="card card-with-image">
  <div class="card-image">
    <img src="clothing-image.jpg" alt="衣物图片" class="card-img">
  </div>
  <div class="card-body">
    <h4 class="card-title">白色T恤</h4>
    <p class="card-text">纯棉材质，适合夏季穿着</p>
  </div>
</div>

<!-- 可交互卡片 -->
<div class="card card-interactive">
  <div class="card-body">
    <h4 class="card-title">夏季穿搭</h4>
    <p class="card-text">清爽舒适的夏日搭配</p>
    <div class="card-tags">
      <span class="tag">夏季</span>
      <span class="tag">休闲</span>
    </div>
  </div>
</div>
```

### Avatar 头像组件
```html
<!-- 图片头像 -->
<div class="avatar avatar-md">
  <img src="user-avatar.jpg" alt="用户头像" class="avatar-img">
</div>

<!-- 文字头像 -->
<div class="avatar avatar-lg">
  <span class="avatar-text">JD</span>
</div>

<!-- 状态头像 -->
<div class="avatar avatar-sm avatar-with-status">
  <img src="user-avatar.jpg" alt="用户头像" class="avatar-img">
  <span class="avatar-status avatar-status-online"></span>
</div>

<!-- 头像组 -->
<div class="avatar-group">
  <div class="avatar avatar-sm">
    <img src="user1.jpg" alt="用户1" class="avatar-img">
  </div>
  <div class="avatar avatar-sm">
    <img src="user2.jpg" alt="用户2" class="avatar-img">
  </div>
  <div class="avatar avatar-sm">
    <img src="user3.jpg" alt="用户3" class="avatar-img">
  </div>
  <div class="avatar avatar-sm avatar-more">
    <span class="avatar-text">+3</span>
  </div>
</div>
```

### Badge 徽章组件
```html
<!-- 数字徽章 -->
<span class="badge badge-primary">5</span>
<span class="badge badge-success">12</span>
<span class="badge badge-warning">3</span>
<span class="badge badge-error">1</span>

<!-- 文字徽章 -->
<span class="badge badge-ghost">新</span>
<span class="badge badge-info">推荐</span>
<span class="badge badge-success">已认证</span>

<!-- 状态徽章 -->
<span class="badge badge-success">在线</span>
<span class="badge badge-warning">忙碌</span>
<span class="badge badge-error">离线</span>

<!-- 圆点徽章 -->
<span class="badge badge-dot badge-success"></span>
<span class="badge badge-dot badge-warning"></span>
<span class="badge badge-dot badge-error"></span>
```

---

## 导航组件

### Header 顶部导航
```html
<header class="header">
  <div class="header-container">
    <div class="header-brand">
      <h1 class="header-logo">StyleSense</h1>
    </div>
    
    <nav class="header-nav">
      <a href="#home" class="header-nav-link header-nav-link-active">首页</a>
      <a href="#wardrobe" class="header-nav-link">我的衣橱</a>
      <a href="#outfit" class="header-nav-link">穿搭记录</a>
      <a href="#inspiration" class="header-nav-link">风格灵感</a>
    </nav>
    
    <div class="header-actions">
      <div class="header-search">
        <input type="text" class="header-search-input" placeholder="搜索衣物">
        <span class="header-search-icon">
          <span class="material-icons-outlined">search</span>
        </span>
      </div>
      
      <button class="header-btn">
        <span class="material-icons-outlined">notifications_none</span>
        <span class="badge badge-error header-badge">3</span>
      </button>
      
      <div class="header-avatar">
        <img src="user-avatar.jpg" alt="用户头像" class="avatar avatar-sm">
      </div>
    </div>
  </div>
</header>
```

### Tabs 标签页
```html
<!-- 标准标签页 -->
<div class="tabs">
  <div class="tabs-list">
    <button class="tabs-trigger tabs-trigger-active">概览</button>
    <button class="tabs-trigger">详情</button>
    <button class="tabs-trigger">数据分析</button>
    <button class="tabs-trigger">设置</button>
  </div>
  
  <div class="tabs-content">
    <div class="tabs-panel tabs-panel-active">
      <h3>概览内容</h3>
      <p>这里是概览标签页的内容</p>
    </div>
  </div>
</div>

<!-- 垂直标签页 -->
<div class="tabs tabs-vertical">
  <div class="tabs-list">
    <button class="tabs-trigger tabs-trigger-active">基础信息</button>
    <button class="tabs-trigger">分类标签</button>
    <button class="tabs-trigger">图片管理</button>
  </div>
  
  <div class="tabs-content">
    <div class="tabs-panel tabs-panel-active">
      <!-- 内容区域 -->
    </div>
  </div>
</div>
```

---

## 反馈组件

### Alert 警告提示
```html
<!-- 成功提示 -->
<div class="alert alert-success">
  <span class="alert-icon">
    <span class="material-icons-outlined">check_circle</span>
  </span>
  <div class="alert-content">
    <div class="alert-title">操作成功</div>
    <div class="alert-description">衣物信息已保存</div>
  </div>
  <button class="alert-close">
    <span class="material-icons-outlined">close</span>
  </button>
</div>

<!-- 警告提示 -->
<div class="alert alert-warning">
  <span class="alert-icon">
    <span class="material-icons-outlined">warning</span>
  </span>
  <div class="alert-content">
    <div class="alert-title">请注意</div>
    <div class="alert-description">某些信息可能不完整</div>
  </div>
</div>

<!-- 错误提示 -->
<div class="alert alert-error">
  <span class="alert-icon">
    <span class="material-icons-outlined">error</span>
  </span>
  <div class="alert-content">
    <div class="alert-title">操作失败</div>
    <div class="alert-description">网络连接异常，请重试</div>
  </div>
</div>

<!-- 信息提示 -->
<div class="alert alert-info">
  <span class="alert-icon">
    <span class="material-icons-outlined">info</span>
  </span>
  <div class="alert-content">
    <div class="alert-description">这是一条信息提示</div>
  </div>
</div>
```

### Toast 消息提示
```html
<!-- 成功消息 -->
<div class="toast toast-success">
  <span class="toast-icon">
    <span class="material-icons-outlined">check_circle</span>
  </span>
  <div class="toast-content">
    <div class="toast-message">保存成功</div>
  </div>
</div>

<!-- 加载消息 -->
<div class="toast toast-loading">
  <span class="toast-icon toast-icon-spin">
    <span class="material-icons-outlined">hourglass_empty</span>
  </span>
  <div class="toast-content">
    <div class="toast-message">正在处理中...</div>
  </div>
</div>
```

### Modal 模态框
```html
<!-- 基础模态框 -->
<div class="modal">
  <div class="modal-backdrop"></div>
  <div class="modal-content">
    <div class="modal-header">
      <h3 class="modal-title">确认删除</h3>
      <button class="modal-close">
        <span class="material-icons-outlined">close</span>
      </button>
    </div>
    
    <div class="modal-body">
      <p>确定要删除这个衣物吗？此操作不可恢复。</p>
    </div>
    
    <div class="modal-footer">
      <button class="btn btn-secondary">取消</button>
      <button class="btn btn-danger">确认删除</button>
    </div>
  </div>
</div>

<!-- 全屏模态框 -->
<div class="modal modal-fullscreen">
  <div class="modal-backdrop"></div>
  <div class="modal-content">
    <div class="modal-header">
      <h3 class="modal-title">编辑衣物</h3>
      <button class="modal-close">
        <span class="material-icons-outlined">close</span>
      </button>
    </div>
    
    <div class="modal-body">
      <!-- 表单内容 -->
    </div>
  </div>
</div>
```

---

## 布局组件

### Container 容器组件
```html
<!-- 响应式容器 -->
<div class="container">
  <h1>标题</h1>
  <p>内容</p>
</div>

<!-- 最大宽度容器 -->
<div class="container container-xl">
  <h1>大容器</h1>
</div>

<!-- 流体容器 -->
<div class="container container-fluid">
  <h1>流体容器</h1>
</div>
```

### Grid 网格系统
```html
<!-- 12列网格系统 -->
<div class="grid">
  <div class="col-12 col-md-6 col-lg-4">
    3列布局
  </div>
  <div class="col-12 col-md-6 col-lg-4">
    3列布局
  </div>
  <div class="col-12 col-md-12 col-lg-4">
    3列布局
  </div>
</div>

<!-- 自定义网格 -->
<div class="grid grid-3">
  <div>均等3列</div>
  <div>均等3列</div>
  <div>均等3列</div>
</div>

<!-- 间隔网格 -->
<div class="grid grid-2 grid-gap-4">
  <div>有间隔</div>
  <div>有间隔</div>
</div>
```

---

## 使用指南

### 组件命名规范
```css
/* BEM 命名规范 */
.block { /* 块 */ }
.block__element { /* 元素 */ }
.block--modifier { /* 修饰符 */ }

/* 示例 */
.card { }
.card__header { }
.card__title { }
.card--interactive { }
```

### 组件状态类
```css
/* 状态类 */
.is-active    /* 激活状态 */
.is-disabled  /* 禁用状态 */
.is-loading   /* 加载状态 */
.is-expanded  /* 展开状态 */
.is-collapsed /* 收起状态 */
.is-focused   /* 聚焦状态 */
.is-error     /* 错误状态 */
.is-success   /* 成功状态 */
```

### 响应式类
```css
/* 响应式前缀 */
.sm- /* 小屏幕: ≥640px */
.md- /* 中屏幕: ≥768px */
.lg- /* 大屏幕: ≥1024px */
.xl- /* 超大屏幕: ≥1280px */

/* 示例 */
.grid { /* 移动端单列 */ }
.md\:grid-2 { /* 平板端2列 */ }
.lg\:grid-3 { /* 桌面端3列 */ }
```

### 主题类
```css
/* 主题类 */
.light { /* 浅色主题 */ }
.dark { /* 深色主题 */ }
/* 通过 JavaScript 切换 class 来改变主题 */
```

---

## 性能优化建议

### CSS 优化
1. **使用 CSS 变量**：便于主题切换和维护
2. **避免过度嵌套**：选择器层级不超过3层
3. **使用原子化类**：减少重复代码
4. **懒加载组件**：按需加载非关键组件

### JavaScript 优化
1. **按需导入**：只导入使用的组件
2. **代码分割**：将不同功能的组件分开打包
3. **缓存组件**：复用组件实例
4. **虚拟滚动**：处理大量数据时使用虚拟列表

### 图片优化
1. **WebP 格式**：使用现代图片格式
2. **响应式图片**：根据设备加载合适尺寸
3. **懒加载**：视口外的图片延迟加载
4. **占位符**：使用低质量图片占位

---

**最后更新**: 2024年10月9日  
**维护者**: StyleSense 设计团队
