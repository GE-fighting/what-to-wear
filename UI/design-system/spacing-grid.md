# StyleSense 间距网格系统 (Spacing & Grid System)

> 间距网格系统是设计系统的基础架构，确保所有元素在空间使用上的一致性和协调性。

## 📋 目录
- [设计原则](#设计原则)
- [间距系统](#间距系统)
- [网格系统](#网格系统)
- [布局原则](#布局原则)
- [响应式原则](#响应式原则)
- [工具类](#工具类)
- [最佳实践](#最佳实践)

---

## 设计原则

### 4px 基础网格
StyleSense 采用 **4px 基础网格系统**，这是现代 UI 设计的最佳实践：

- **基础单元**: 4px
- **最小间距**: 4px
- **常用间距**: 4px 的倍数 (8px, 12px, 16px, 20px, 24px, 32px...)
- **视觉和谐**: 确保所有间距都有数学关系

### 一致性原则
1. **统一测量**: 所有间距基于 4px 网格
2. **层次清晰**: 通过间距建立视觉层次
3. **呼吸感**: 适度的留白，避免拥挤
4. **节奏感**: 重复的间距模式创造节奏

---

## 间距系统

### 基础间距 Token
```css
/* 基础间距系统 (4px 基础) */
--space-0: 0;          /* 0px */
--space-1: 0.25rem;     /* 4px */
--space-2: 0.5rem;      /* 8px */
--space-3: 0.75rem;     /* 12px */
--space-4: 1rem;        /* 16px */
--space-5: 1.25rem;     /* 20px */
--space-6: 1.5rem;      /* 24px */
--space-8: 2rem;        /* 32px */
--space-10: 2.5rem;     /* 40px */
--space-12: 3rem;       /* 48px */
--space-16: 4rem;       /* 64px */
--space-20: 5rem;       /* 80px */
--space-24: 6rem;       /* 96px */
--space-32: 8rem;       /* 128px */
--space-40: 10rem;      /* 160px */
--space-48: 12rem;      /* 192px */
--space-56: 14rem;      /* 224px */
--space-64: 16rem;      /* 256px */
```

### 语义化间距
```css
/* 极小间距 - 元素内部紧密关系 */
--spacing-xs: var(--space-1);  /* 4px */
/* 小间距 - 紧密相关元素 */
--spacing-sm: var(--space-2);  /* 8px */
/* 中等间距 - 一般元素间距 */
--spacing-md: var(--space-4);  /* 16px */
/* 大间距 - 相关元素组 */
--spacing-lg: var(--space-6);  /* 24px */
/* 超大间距 - 不相关元素组 */
--spacing-xl: var(--space-8);  /* 32px */
/* 极大间距 - 主要区域分隔 */
--spacing-2xl: var(--space-12); /* 48px */

/* 组件级间距 */
--spacing-component-padding: var(--spacing-md);     /* 16px */
--spacing-component-margin: var(--spacing-lg);      /* 24px */
--spacing-component-gap: var(--spacing-sm);         /* 8px */

/* 布局级间距 */
--spacing-layout-section: var(--spacing-xl);        /* 32px */
--spacing-layout-container: var(--spacing-md);      /* 16px */
--spacing-layout-page: var(--spacing-2xl);          /* 48px */

/* 列表间距 */
--spacing-list-item: var(--space-3);                /* 12px */
--spacing-list-group: var(--space-6);               /* 24px */

/* 表单间距 */
--spacing-form-field: var(--space-4);               /* 16px */
--spacing-form-group: var(--space-6);               /* 24px */
--spacing-form-section: var(--space-8);             /* 32px */
```

---

## 网格系统

### 12列网格布局
```html
<!-- 基础12列网格 -->
<div class="grid">
  <div class="col-12">全宽 12/12</div>
  
  <div class="col-6">半宽 6/12</div>
  <div class="col-6">半宽 6/12</div>
  
  <div class="col-4">三分之一 4/12</div>
  <div class="col-4">三分之一 4/12</div>
  <div class="col-4">三分之一 4/12</div>
  
  <div class="col-3">四分之一 3/12</div>
  <div class="col-3">四分之一 3/12</div>
  <div class="col-3">四分之一 3/12</div>
  <div class="col-3">四分之一 3/12</div>
</div>
```

### 网格变体
```html
<!-- 自定义网格 -->
<div class="grid grid-2"> <!-- 2列均分 -->
  <div>列1</div>
  <div>列2</div>
</div>

<div class="grid grid-3"> <!-- 3列均分 -->
  <div>列1</div>
  <div>列2</div>
  <div>列3</div>
</div>

<div class="grid grid-4"> <!-- 4列均分 -->
  <div>列1</div>
  <div>列2</div>
  <div>列3</div>
  <div>列4</div>
</div>

<!-- 自定义列宽 -->
<div class="grid" style="grid-template-columns: 2fr 1fr 1fr;">
  <div>2倍宽度</div>
  <div>常规宽度</div>
  <div>常规宽度</div>
</div>
```

### 网格间距
```html
<!-- 无间距网格 -->
<div class="grid grid-3">
  <div>无间距</div>
  <div>无间距</div>
  <div>无间距</div>
</div>

<!-- 小间距网格 -->
<div class="grid grid-3 grid-gap-sm">
  <div>8px间距</div>
  <div>8px间距</div>
  <div>8px间距</div>
</div>

<!-- 中等间距网格 -->
<div class="grid grid-3 grid-gap">
  <div>16px间距</div>
  <div>16px间距</div>
  <div>16px间距</div>
</div>

<!-- 大间距网格 -->
<div class="grid grid-3 grid-gap-lg">
  <div>24px间距</div>
  <div>24px间距</div>
  <div>24px间距</div>
</div>
```

---

## 布局原则

### 页面布局层次
```
页面 (Page)
├── 页面边距 (Page Margin: 24px)
├── 页面内边距 (Page Padding: 48px)
└── 页面内容区 (Content Area)
    ├── 页面标题区 (Header Area: 48px margin-bottom)
    ├── 页面主体 (Main Content)
    │   ├── 板块间距 (Section Gap: 32px)
    │   └── 板块内边距 (Section Padding: 24px)
    └── 页面底部 (Footer Area: 48px margin-top)
```

### 组件布局层次
```
组件 (Component)
├── 组件边距 (Component Margin: 24px)
├── 组件内边距 (Component Padding: 16px)
└── 组件内容
    ├── 元素间距 (Element Gap: 12px)
    ├── 元素内边距 (Element Padding: 8px)
    └── 内部元素
        ├── 文本行间距 (Line Height: 1.5)
        └── 图标间距 (Icon Gap: 4px)
```

### 典型布局模式
```html
<!-- 卡片布局 -->
<div class="card">
  <div class="card-body">
    <!-- 元素间距: 16px -->
    <h3 class="card-title" style="margin-bottom: 16px;">标题</h3>
    <p class="card-text" style="margin-bottom: 16px;">描述内容</p>
    <div class="card-actions">
      <!-- 按钮间距: 8px -->
      <button class="btn btn-secondary" style="margin-right: 8px;">取消</button>
      <button class="btn btn-primary">确认</button>
    </div>
  </div>
</div>

<!-- 表单布局 -->
<form class="form">
  <!-- 表单组间距: 24px -->
  <div class="form-group" style="margin-bottom: 24px;">
    <label class="form-label" style="margin-bottom: 8px;">标签</label>
    <input type="text" class="input">
  </div>
  
  <div class="form-group" style="margin-bottom: 24px;">
    <label class="form-label" style="margin-bottom: 8px;">标签</label>
    <select class="select"></select>
  </div>
  
  <!-- 表单操作间距: 32px -->
  <div class="form-actions" style="margin-top: 32px;">
    <button class="btn btn-secondary" style="margin-right: 8px;">取消</button>
    <button class="btn btn-primary">提交</button>
  </div>
</form>
```

---

## 响应式原则

### 断点设置
```css
/* 响应式断点 */
--breakpoint-sm: 640px;    /* 小屏幕 */
--breakpoint-md: 768px;    /* 中屏幕 */
--breakpoint-lg: 1024px;   /* 大屏幕 */
--breakpoint-xl: 1280px;   /* 超大屏幕 */
--breakpoint-2xl: 1536px;  /* 超超大屏幕 */
```

### 响应式间距
```css
/* 移动端优先的间距 */
/* 小屏幕 (默认) */
:root {
  --spacing-responsive-page: var(--spacing-lg);    /* 24px */
  --spacing-responsive-section: var(--spacing-md); /* 16px */
  --spacing-responsive-component: var(--spacing-sm); /* 8px */
}

/* 中屏幕 */
@media (min-width: 768px) {
  :root {
    --spacing-responsive-page: var(--spacing-xl);    /* 32px */
    --spacing-responsive-section: var(--spacing-lg); /* 24px */
    --spacing-responsive-component: var(--spacing-md); /* 16px */
  }
}

/* 大屏幕 */
@media (min-width: 1024px) {
  :root {
    --spacing-responsive-page: var(--spacing-2xl);   /* 48px */
    --spacing-responsive-section: var(--spacing-xl); /* 32px */
    --spacing-responsive-component: var(--spacing-lg); /* 24px */
  }
}
```

### 响应式网格
```html
<!-- 响应式列布局 -->
<div class="grid grid-responsive">
  <div class="col-12 col-md-6 col-lg-4">
    <!-- 移动端12列，平板6列，桌面4列 -->
  </div>
  <div class="col-12 col-md-6 col-lg-4">
    <!-- 移动端12列，平板6列，桌面4列 -->
  </div>
  <div class="col-12 col-md-12 col-lg-4">
    <!-- 移动端12列，平板12列，桌面4列 -->
  </div>
</div>

<!-- 响应式网格间距 -->
<div class="grid grid-3 grid-gap-sm md:grid-gap lg:grid-gap-lg">
  <!-- 移动端8px间距，平板16px间距，桌面24px间距 -->
  <div>内容1</div>
  <div>内容2</div>
  <div>内容3</div>
</div>
```

---

## 工具类

### 间距工具类
```css
/* 外边距 */
.m-0 { margin: 0; }
.m-xs { margin: var(--spacing-xs); }
.m-sm { margin: var(--spacing-sm); }
.m-md { margin: var(--spacing-md); }
.m-lg { margin: var(--spacing-lg); }
.m-xl { margin: var(--spacing-xl); }
.m-2xl { margin: var(--spacing-2xl); }

/* 单方向外边距 */
.mt-sm { margin-top: var(--spacing-sm); }
.mb-sm { margin-bottom: var(--spacing-sm); }
.ml-sm { margin-left: var(--spacing-sm); }
.mr-sm { margin-right: var(--spacing-sm); }
.mx-sm { margin-left: var(--spacing-sm); margin-right: var(--spacing-sm); }
.my-sm { margin-top: var(--spacing-sm); margin-bottom: var(--spacing-sm); }

/* 内边距 */
.p-0 { padding: 0; }
.p-xs { padding: var(--spacing-xs); }
.p-sm { padding: var(--spacing-sm); }
.p-md { padding: var(--spacing-md); }
.p-lg { padding: var(--spacing-lg); }
.p-xl { padding: var(--spacing-xl); }
.p-2xl { padding: var(--spacing-2xl); }

/* 单方向内边距 */
.pt-sm { padding-top: var(--spacing-sm); }
.pb-sm { padding-bottom: var(--spacing-sm); }
.pl-sm { padding-left: var(--spacing-sm); }
.pr-sm { padding-right: var(--spacing-sm); }
.px-sm { padding-left: var(--spacing-sm); padding-right: var(--spacing-sm); }
.py-sm { padding-top: var(--spacing-sm); padding-bottom: var(--spacing-sm); }
```

### 网格工具类
```css
/* 网格容器 */
.grid {
  display: grid;
  gap: var(--spacing-component-gap);
}

.grid-2 { grid-template-columns: repeat(2, 1fr); }
.grid-3 { grid-template-columns: repeat(3, 1fr); }
.grid-4 { grid-template-columns: repeat(4, 1fr); }
.grid-auto { grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); }

/* 网格间距 */
.grid-gap-sm { gap: var(--spacing-sm); }
.grid-gap { gap: var(--spacing-md); }
.grid-gap-lg { gap: var(--spacing-lg); }

/* 列定义 */
.col-1 { grid-column: span 1; }
.col-2 { grid-column: span 2; }
.col-3 { grid-column: span 3; }
.col-4 { grid-column: span 4; }
.col-6 { grid-column: span 6; }
.col-8 { grid-column: span 8; }
.col-9 { grid-column: span 9; }
.col-12 { grid-column: span 12; }

/* 响应式列 */
.col-sm-6 { grid-column: span 6; }
.col-md-6 { grid-column: span 6; }
.col-lg-6 { grid-column: span 6; }
.col-xl-6 { grid-column: span 6; }
```

### 弹性布局工具类
```css
/* Flex 容器 */
.flex { display: flex; }
.flex-col { flex-direction: column; }
.flex-wrap { flex-wrap: wrap; }
.flex-nowrap { flex-wrap: nowrap; }

/* 对齐 */
.justify-start { justify-content: flex-start; }
.justify-center { justify-content: center; }
.justify-end { justify-content: flex-end; }
.justify-between { justify-content: space-between; }
.justify-around { justify-content: space-around; }

.items-start { align-items: flex-start; }
.items-center { align-items: center; }
.items-end { align-items: flex-end; }
.items-stretch { align-items: stretch; }

/* 间距 */
.gap-0 { gap: 0; }
.gap-sm { gap: var(--spacing-sm); }
.gap { gap: var(--spacing-md); }
.gap-lg { gap: var(--spacing-lg); }
```

---

## 特殊布局场景

### 导航布局
```html
<!-- 顶部导航 -->
<header class="header">
  <div class="header-content">
    <div class="header-brand">Logo</div>
    <nav class="header-nav">
      <!-- 导航项间距: 32px -->
      <a href="#" style="margin-right: 32px;">首页</a>
      <a href="#" style="margin-right: 32px;">衣橱</a>
      <a href="#" style="margin-right: 32px;">穿搭</a>
    </nav>
    <div class="header-actions">
      <!-- 操作按钮间距: 16px -->
      <button style="margin-right: 16px;">搜索</button>
      <button style="margin-right: 16px;">通知</button>
      <button>用户</button>
    </div>
  </div>
</header>

<!-- 侧边导航 -->
<div class="sidebar-layout">
  <aside class="sidebar" style="width: 240px; margin-right: 24px;">
    <div class="sidebar-content">
      <!-- 侧边栏项间距: 4px -->
      <a href="#" class="sidebar-item" style="margin-bottom: 4px;">首页</a>
      <a href="#" class="sidebar-item" style="margin-bottom: 4px;">衣橱</a>
      <a href="#" class="sidebar-item" style="margin-bottom: 4px;">穿搭</a>
    </div>
  </aside>
  <main class="main-content">
    <!-- 主内容 -->
  </main>
</div>
```

### 卡片网格布局
```html
<!-- 衣物卡片网格 -->
<div class="card-grid">
  <!-- 网格间距: 16px -->
  <div class="grid grid-auto grid-gap" style="grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));">
    <!-- 卡片 -->
    <div class="card">
      <div class="card-image" style="margin-bottom: 12px;">
        <img src="clothing1.jpg" alt="衣物图片">
      </div>
      <div class="card-content">
        <h4 style="margin-bottom: 8px;">衣物名称</h4>
        <p style="margin-bottom: 12px;">描述信息</p>
        <div class="card-tags" style="margin-bottom: 16px;">
          <span class="tag" style="margin-right: 4px;">标签1</span>
          <span class="tag">标签2</span>
        </div>
      </div>
    </div>
  </div>
</div>
```

### 表单布局
```html
<!-- 表单布局 -->
<form class="form-layout" style="max-width: 600px;">
  <!-- 表单标题间距: 24px -->
  <h2 style="margin-bottom: 24px;">添加衣物</h2>
  
  <!-- 表单组间距: 20px -->
  <div class="form-group" style="margin-bottom: 20px;">
    <label style="margin-bottom: 8px;">衣物名称</label>
    <input type="text" class="input">
  </div>
  
  <div class="form-group" style="margin-bottom: 20px;">
    <label style="margin-bottom: 8px;">分类</label>
    <select class="select"></select>
  </div>
  
  <!-- 表单区段间距: 32px -->
  <div class="form-section" style="margin: 32px 0; padding: 24px; background: var(--color-gray-50);">
    <h3 style="margin-bottom: 20px;">详细信息</h3>
    
    <div class="form-group" style="margin-bottom: 20px;">
      <label style="margin-bottom: 8px;">品牌</label>
      <input type="text" class="input">
    </div>
  </div>
  
  <!-- 表单操作区间距: 32px -->
  <div class="form-actions" style="margin-top: 32px;">
    <button class="btn btn-secondary" style="margin-right: 12px;">取消</button>
    <button class="btn btn-primary">保存</button>
  </div>
</form>
```

---

## 最佳实践

### 间距使用原则
1. **保持一致性**: 使用语义化的间距 token，不要随意定义值
2. **建立层次**: 通过不同的间距建立清晰的视觉层次
3. **适度原则**: 不要过度使用间距，保持页面紧凑
4. **响应式考虑**: 在不同屏幕尺寸下调整间距

### 网格使用原则
1. **12列为基础**: 大部分布局都可以用12列实现
2. **Breaking Points**: 在关键断点调整布局结构
3. **内容为王**: 网格服务于内容，不要让内容迁就网格
4. **性能考虑**: 避免过于复杂的网格嵌套

### 常见错误避免
```css
/* ❌ 错误：随意定义间距值 */
.wrong-spacing { margin: 17px; padding: 23px; }

/* ✅ 正确：使用标准间距 token */
.correct-spacing { margin: var(--spacing-md); padding: var(--spacing-lg); }

/* ❌ 错误：过度嵌套网格 */
.wrong-nesting .grid .grid .grid { /* 太深 */ }

/* ✅ 正确：扁平化网格结构 */
.correct-nesting .grid .grid-item { /* 保持扁平 */ }

/* ❌ 错误：不响应式间距 */
.wrong-responsive { margin: 48px; /* 在小屏幕上过大 */ }

/* ✅ 正确：响应式间距 */
.correct-responsive { 
  margin: var(--spacing-responsive-page, var(--spacing-lg));
}
```

---

## Tailwind CSS 配置

```javascript
module.exports = {
  theme: {
    extend: {
      spacing: {
        // 基于 4px 的间距系统
        'xs': '0.25rem',  // 4px
        'sm': '0.5rem',   // 8px
        'md': '1rem',     // 16px
        'lg': '1.5rem',   // 24px
        'xl': '2rem',     // 32px
        '2xl': '3rem',    // 48px
        '3xl': '4rem',    // 64px
        '4xl': '6rem',    // 96px
        '5xl': '8rem',    // 128px
        
        // 自定义间距
        ' navbar': '64px',
        'sidebar': '240px',
        'form-field': '20px',
      },
      
      gridTemplateColumns: {
        // 网格变体
        'auto-fit': 'repeat(auto-fit, minmax(250px, 1fr))',
        'compact': 'repeat(auto-fill, minmax(200px, 1fr))',
        'wide': 'repeat(auto-fill, minmax(300px, 1fr))',
      },
      
      gap: {
        // 间距变体
        'xs': '0.25rem',  // 4px
        'sm': '0.5rem',   // 8px
        'md': '1rem',     // 16px
        'lg': '1.5rem',   // 24px
        'xl': '2rem',     // 32px
        'section': '2rem', // 32px
      }
    }
  }
}
```

---

**最后更新**: 2024年10月9日  
**维护者**: StyleSense 设计团队
