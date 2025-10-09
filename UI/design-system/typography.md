# StyleSense 字体系统 (Typography System)

> 字体系统定义了所有文本样式的基础规范，确保信息层次清晰、易读性强、品牌一致。

## 📋 目录
- [字体族选择](#字体族选择)
- [字体大小层级](#字体大小层级)
- [字体粗细规范](#字体粗细规范)
- [行高规范](#行高规范)
- [颜色规范](#颜色规范)
- [应用场景](#应用场景)
- [响应式字体](#响应式字体)
- [最佳实践](#最佳实践)

---

## 字体族选择

### 主要字体族
```css
--font-family-primary: 'Noto Sans SC', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
```

**选择理由**：
- **Noto Sans SC**: Google 开发的高质量中文字体，支持完整的中文字符集
- **系统字体备用**: 确保在不同平台下都有良好的显示效果
- **现代无衬线**: 简洁、现代，适合数字产品界面

### 辅助字体族
```css
--font-family-mono: 'Fira Code', 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, 'Courier New', monospace;
--font-family-icons: 'Material Icons Outlined', sans-serif;
```

**使用场景**：
- **等宽字体**: 代码、数字显示、终端界面
- **图标字体**: 所有图标的统一字体

---

## 字体大小层级

### 显示字体 (Display)
用于页面主标题、 hero 区域

| 层级 | 大小 | 行高 | 字重 | 使用场景 |
|------|------|------|------|----------|
| Display 1 | 3rem (48px) | 1.2 | 700 | 页面主标题 |
| Display 2 | 2.25rem (36px) | 1.25 | 700 | 区域主标题 |
| Display 3 | 1.875rem (30px) | 1.3 | 700 | 卡片主标题 |

### 标题字体 (Headings)
用于各级标题

| 层级 | 大小 | 行高 | 字重 | Tailwind | 使用场景 |
|------|------|------|------|---------|----------|
| H1 | 1.5rem (24px) | 1.33 | 700 | text-2xl | 页面一级标题 |
| H2 | 1.25rem (20px) | 1.4 | 600 | text-xl | 页面二级标题 |
| H3 | 1.125rem (18px) | 1.44 | 600 | text-lg | 页面三级标题 |
| H4 | 1rem (16px) | 1.5 | 600 | text-base | 页面四级标题 |
| H5 | 0.875rem (14px) | 1.43 | 600 | text-sm | 小标题 |
| H6 | 0.75rem (12px) | 1.33 | 600 | text-xs | 微标题 |

### 正文字体 (Body)
用于主要内容

| 层级 | 大小 | 行高 | 字重 | Tailwind | 使用场景 |
|------|------|------|------|---------|----------|
| Body Large | 1.125rem (18px) | 1.56 | 400 | text-lg | 重要正文 |
| Body | 1rem (16px) | 1.5 | 400 | text-base | 普通正文 |
| Body Small | 0.875rem (14px) | 1.43 | 400 | text-sm | 辅助正文 |

### 辅助字体 (Support)
用于次要信息

| 层级 | 大小 | 行高 | 字重 | Tailwind | 使用场景 |
|------|------|------|------|---------|----------|
| Caption | 0.75rem (12px) | 1.33 | 400 | text-xs | 图注、标签 |
| Overline | 0.625rem (10px) | 1.6 | 500 | text-[10px] | 导航标识 |

---

## 字体粗细规范

### 字重层级
```css
--font-weight-extralight: 100;  /* 极细 - 很少使用 */
--font-weight-light: 300;       /* 细体 - 大段文字 */
--font-weight-normal: 400;      /* 正常 - 正文 */
--font-weight-medium: 500;      /* 中等 - 强调文字 */
--font-weight-semibold: 600;    /* 半粗 - 次要标题 */
--font-weight-bold: 700;        /* 粗体 - 主要标题 */
--font-weight-extrabold: 800;   /* 极粗 - 特殊强调 */
--font-weight-black: 900;       /* 黑体 - 很少使用 */
```

### 使用原则
1. **标题**: 使用 600-700 字重，确保层次分明
2. **正文**: 主要使用 400 字重，保证易读性
3. **强调**: 使用 500 字重进行适度强调
4. **避免**: 不要过多使用不同字重，控制在 2-3 种内

---

## 行高规范

### 行高计算
行高应该是字体大小的 1.2-2.0 倍，具体根据使用场景调整：

```css
--line-height-tight: 1.2;      /* 紧凑 - 大标题 */
--line-height-snug: 1.375;     /* 适中 - 小标题 */
--line-height-normal: 1.5;     /* 标准 - 正文 */
--line-height-relaxed: 1.625;  /* 宽松 - 长文 */
--line-height-loose: 2;        /* 很宽松 - 阅读密集内容 */
```

### 使用场景
- **大标题**: 1.2-1.3，保持紧凑感
- **小标题**: 1.3-1.4，适中间距
- **正文**: 1.5-1.6，保证易读性
- **表单**: 1.4-1.5，适合输入框
- **长文阅读**: 1.6-1.8，减少阅读疲劳

---

## 颜色规范

### 浅色模式
```css
/* 主要文本 */
--text-color-primary: #1F2937;      /* 主要信息 */
--text-color-secondary: #6B7280;    /* 次要信息 */
--text-color-tertiary: #9CA3AF;     /* 辅助信息 */
--text-color-disabled: #D1D5DB;     /* 禁用状态 */

/* 链接文本 */
--text-color-link: #000000;         /* 链接主色 */
--text-color-link-visited: #333333; /* 访问过的链接 */
--text-color-link-hover: #555555;   /* 悬停状态 */
```

### 深色模式
```css
/* 主要文本 */
--text-color-primary-dark: #F9FAFB;    /* 主要信息 */
--text-color-secondary-dark: #9CA3AF;  /* 次要信息 */
--text-color-tertiary-dark: #6B7280;   /* 辅助信息 */
--text-color-disabled-dark: #4B5563;   /* 禁用状态 */

/* 链接文本 */
--text-color-link-dark: #FFFFFF;       /* 链接主色 */
--text-color-link-visited-dark: #CCCCCC; /* 访问过的链接 */
--text-color-link-hover-dark: #E0E0E0; /* 悬停状态 */
```

---

## 应用场景

### 1. 页面层级
```css
/* 页面主标题 */
.page-title {
  font-size: 3rem;
  line-height: 1.2;
  font-weight: 700;
  color: var(--text-color-primary);
}

/* 区域标题 */
.section-title {
  font-size: 1.5rem;
  line-height: 1.33;
  font-weight: 600;
  color: var(--text-color-primary);
  margin-bottom: 1rem;
}

/* 卡片标题 */
.card-title {
  font-size: 1.125rem;
  line-height: 1.44;
  font-weight: 600;
  color: var(--text-color-primary);
}
```

### 2. 内容层级
```css
/* 正文内容 */
.body-text {
  font-size: 1rem;
  line-height: 1.5;
  font-weight: 400;
  color: var(--text-color-primary);
}

/* 辅助文本 */
.supporting-text {
  font-size: 0.875rem;
  line-height: 1.43;
  font-weight: 400;
  color: var(--text-color-secondary);
}

/* 标签文本 */
.label-text {
  font-size: 0.75rem;
  line-height: 1.33;
  font-weight: 500;
  color: var(--text-color-tertiary);
}
```

### 3. 交互元素
```css
/* 按钮文字 */
.button-text {
  font-size: 1rem;
  line-height: 1.5;
  font-weight: 500;
  color: var(--color-white);
}

/* 输入框文字 */
.input-text {
  font-size: 0.875rem;
  line-height: 1.43;
  font-weight: 400;
  color: var(--text-color-primary);
}

/* 链接文字 */
.link-text {
  font-size: 0.875rem;
  line-height: 1.43;
  font-weight: 500;
  color: var(--text-color-link);
  text-decoration: none;
}
```

---

## 响应式字体

### 断点设置
```css
/* 移动端 */
@media (max-width: 767px) {
  :root {
    --font-size-h1: 1.5rem;
    --font-size-h2: 1.25rem;
    --font-size-h3: 1.125rem;
    --font-size-body: 1rem;
  }
}

/* 平板端 */
@media (min-width: 768px) and (max-width: 1023px) {
  :root {
    --font-size-h1: 1.75rem;
    --font-size-h2: 1.5rem;
    --font-size-h3: 1.25rem;
    --font-size-body: 1rem;
  }
}

/* 桌面端 */
@media (min-width: 1024px) {
  :root {
    --font-size-h1: 3rem;
    --font-size-h2: 1.5rem;
    --font-size-h3: 1.125rem;
    --font-size-body: 1rem;
  }
}
```

### 流式缩放
使用 clamp() 实现流畅的字体缩放：

```css
/* 主标题：从 1.5rem 到 3rem */
.text-fluid-h1 {
  font-size: clamp(1.5rem, 2.5vw + 1rem, 3rem);
}

/* 正文：从 0.875rem 到 1rem */
.text-fluid-body {
  font-size: clamp(0.875rem, 1vw + 0.8rem, 1rem);
}
```

---

## 特殊文本样式

### 强调文本
```css
/* 粗体强调 */
.text-emphasis-bold {
  font-weight: 600;
  color: var(--text-color-primary);
}

/* 颜色强调 */
.text-emphasis-color {
  font-weight: 500;
  color: var(--color-primary);
}

/* 斜体强调（慎用） */
.text-emphasis-italic {
  font-style: italic;
  font-weight: 500;
}
```

### 状态文本
```css
/* 成功状态 */
.text-success {
  color: var(--color-success);
}

/* 警告状态 */
.text-warning {
  color: var(--color-warning);
}

/* 错误状态 */
.text-error {
  color: var(--color-error);
}

/* 信息状态 */
.text-info {
  color: var(--color-info);
}
```

### 特殊用途
```css
/* 价格显示 */
.text-price {
  font-family: var(--font-family-mono);
  font-weight: 600;
  font-variant-numeric: tabular-nums;
}

/* 数字显示 */
.text-number {
  font-family: var(--font-family-mono);
  font-variant-numeric: tabular-nums;
}

/* 代码显示 */
.text-code {
  font-family: var(--font-family-mono);
  font-size: 0.875em;
  background-color: var(--color-gray-100);
  padding: 0.125rem 0.25rem;
  border-radius: 4px;
}
```

---

## 最佳实践

### 1. 层次清晰
- 每个页面最多使用 3-4 种不同的字体大小
- 标题和正文要有明显的视觉对比
- 通过字重和颜色区分重要性

### 2. 易读性优先
- 正文字体不小于 14px
- 行高适中，避免过于紧凑
- 确保足够的颜色对比度

### 3. 一致性
- 在整个产品中使用相同的字体规范
- 使用设计令牌而不是硬编码值
- 保持响应式字体的一致性

### 4. 性能考虑
- 优先系统字体，减少加载时间
- 使用 font-display: swap 优化字体加载
- 限制使用字体文件的数量

### 5. 可访问性
- 尊重用户的字体大小设置
- 确保文本在放大时仍然可用
- 提供足够的颜色对比度

---

## 实现示例

### Tailwind CSS 配置
```javascript
module.exports = {
  theme: {
    extend: {
      fontFamily: {
        display: ['Noto Sans SC', 'sans-serif'],
        body: ['Noto Sans SC', 'sans-serif'],
        mono: ['Fira Code', 'monospace'],
      },
      fontSize: {
        'display-1': ['3rem', { lineHeight: '1.2', letterSpacing: '-0.02em' }],
        'display-2': ['2.25rem', { lineHeight: '1.25', letterSpacing: '-0.015em' }],
        'display-3': ['1.875rem', { lineHeight: '1.3', letterSpacing: '-0.01em' }],
        'h1': ['1.5rem', { lineHeight: '1.33', letterSpacing: '-0.01em' }],
        'h2': ['1.25rem', { lineHeight: '1.4', letterSpacing: '-0.005em' }],
        'h3': ['1.125rem', { lineHeight: '1.44' }],
        'h4': ['1rem', { lineHeight: '1.5' }],
        'h5': ['0.875rem', { lineHeight: '1.43' }],
        'h6': ['0.75rem', { lineHeight: '1.33' }],
      },
      fontWeight: {
        'extralight': '100',
        'light': '300',
        'normal': '400',
        'medium': '500',
        'semibold': '600',
        'bold': '700',
        'extrabold': '800',
      },
      lineHeight: {
        'tight': '1.2',
        'snug': '1.375',
        'normal': '1.5',
        'relaxed': '1.625',
        'loose': '2',
      },
    },
  },
}
```

### CSS 变量定义
```css
:root {
  /* 字体族 */
  --font-primary: 'Noto Sans SC', -apple-system, BlinkMacSystemFont, sans-serif;
  --font-mono: 'Fira Code', 'SF Mono', Monaco, monospace;
  --font-icons: 'Material Icons Outlined', sans-serif;

  /* 字体大小 */
  --text-display-1: 3rem;
  --text-display-2: 2.25rem;
  --text-display-3: 1.875rem;
  --text-h1: 1.5rem;
  --text-h2: 1.25rem;
  --text-h3: 1.125rem;
  --text-body: 1rem;
  --text-body-small: 0.875rem;
  --text-caption: 0.75rem;

  /* 字重 */
  --font-weight-light: 300;
  --font-weight-normal: 400;
  --font-weight-medium: 500;
  --font-weight-semibold: 600;
  --font-weight-bold: 700;

  /* 行高 */
  --line-height-tight: 1.2;
  --line-height-normal: 1.5;
  --line-height-relaxed: 1.6;
}
```

---

**最后更新**: 2024年10月9日  
**维护者**: StyleSense 设计团队
