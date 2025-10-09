# StyleSense 设计令牌 (Design Tokens)

> 设计令牌是设计系统的原子单位，定义了所有视觉设计的基础决策。它们在代码中被复用，确保设计的一致性和可维护性。

## 📋 目录
- [颜色 tokens](#颜色-tokens)
- [字体 tokens](#字体-tokens)  
- [间距 tokens](#间距-tokens)
- [尺寸 tokens](#尺寸-tokens)
- [阴影 tokens](#阴影-tokens)
- [圆角 tokens](#圆角-tokens)
- [动画 tokens](#动画-tokens)
- [Z-index tokens](#z-index-tokens)

---

## 颜色 Tokens

### 主色调
```css
--color-primary: #000000;           /* 主色 - 纯黑 */
--color-primary-hover: #333333;     /* 主色悬停态 */
--color-primary-active: #1a1a1a;    /* 主色激活态 */
--color-primary-disabled: #666666;  /* 主色禁用态 */
```

### 中性色调
```css
/* 浅色模式 */
--color-white: #FFFFFF;
--color-gray-50: #F9FAFB;
--color-gray-100: #F3F4F6;
--color-gray-200: #E5E7EB;
--color-gray-300: #D1D5DB;
--color-gray-400: #9CA3AF;
--color-gray-500: #6B7280;
--color-gray-600: #4B5563;
--color-gray-700: #374151;
--color-gray-800: #1F2937;
--color-gray-900: #111827;

/* 深色模式别名 */
--color-background-light: #F8F9FA;
--color-background-dark: #121212;
--color-card-light: #FFFFFF;
--color-card-dark: #1E1E1E;
--color-border-light: #E5E7EB;
--color-border-dark: #374151;
```

### 功能色
```css
/* 成功色 */
--color-success: #10B981;
--color-success-light: #D1FAE5;
--color-success-dark: #065F46;

/* 警告色 */
--color-warning: #F59E0B;
--color-warning-light: #FEF3C7;
--color-warning-dark: #92400E;

/* 错误色 */
--color-error: #EF4444;
--color-error-light: #FEE2E2;
--color-error-dark: #991B1B;

/* 信息色 */
--color-info: #3B82F6;
--color-info-light: #DBEAFE;
--color-info-dark: #1E3A8A;
```

### 文本色
```css
/* 浅色模式文本 */
--color-text-primary-light: #1F2937;      /* 主要文本 */
--color-text-secondary-light: #6B7280;    /* 次要文本 */
--color-text-tertiary-light: #9CA3AF;     /* 三级文本 */
--color-text-disabled-light: #D1D5DB;     /* 禁用文本 */

/* 深色模式文本 */
--color-text-primary-dark: #F9FAFB;       /* 主要文本 */
--color-text-secondary-dark: #9CA3AF;     /* 次要文本 */
--color-text-tertiary-dark: #6B7280;      /* 三级文本 */
--color-text-disabled-dark: #4B5563;      /* 禁用文本 */
```

---

## 字体 Tokens

### 字体族
```css
--font-family-primary: 'Noto Sans SC', sans-serif;  /* 主要字体 */
--font-family-mono: 'Fira Code', 'Consolas', monospace; /* 等宽字体 */
--font-family-icons: 'Material Icons Outlined';     /* 图标字体 */
```

### 字体大小
```css
--font-size-xs: 0.75rem;    /* 12px */
--font-size-sm: 0.875rem;   /* 14px */
--font-size-base: 1rem;     /* 16px */
--font-size-lg: 1.125rem;   /* 18px */
--font-size-xl: 1.25rem;    /* 20px */
--font-size-2xl: 1.5rem;    /* 24px */
--font-size-3xl: 1.875rem;  /* 30px */
--font-size-4xl: 2.25rem;   /* 36px */
--font-size-5xl: 3rem;      /* 48px */
```

### 字体粗细
```css
--font-weight-light: 300;
--font-weight-normal: 400;
--font-weight-medium: 500;
--font-weight-semibold: 600;
--font-weight-bold: 700;
--font-weight-extrabold: 800;
```

### 行高
```css
--line-height-tight: 1.25;
--line-height-normal: 1.5;
--line-height-relaxed: 1.75;
--line-height-loose: 2;
```

### 字母间距
```css
--letter-spacing-tight: -0.025em;
--letter-spacing-normal: 0;
--letter-spacing-wide: 0.025em;
--letter-spacing-wider: 0.05em;
--letter-spacing-widest: 0.1em;
```

---

## 间距 Tokens

基于 4px 网格系统

### 基础间距
```css
--space-0: 0;
--space-1: 0.25rem;  /* 4px */
--space-2: 0.5rem;   /* 8px */
--space-3: 0.75rem;  /* 12px */
--space-4: 1rem;     /* 16px */
--space-5: 1.25rem;  /* 20px */
--space-6: 1.5rem;   /* 24px */
--space-8: 2rem;     /* 32px */
--space-10: 2.5rem;  /* 40px */
--space-12: 3rem;    /* 48px */
--space-16: 4rem;    /* 64px */
--space-20: 5rem;    /* 80px */
--space-24: 6rem;    /* 96px */
--space-32: 8rem;    /* 128px */
```

### 语义化间距
```css
--spacing-xs: var(--space-1);      /* 4px - 极小间距 */
--spacing-sm: var(--space-2);      /* 8px - 小间距 */
--spacing-md: var(--space-4);      /* 16px - 中间距 */
--spacing-lg: var(--space-6);      /* 24px - 大间距 */
--spacing-xl: var(--space-8);      /* 32px - 超大间距 */
--spacing-2xl: var(--space-12);    /* 48px - 极大间距 */

/* 组件内间距 */
--spacing-component-padding: var(--space-4);   /* 16px */
--spacing-component-margin: var(--space-6);    /* 24px */

/* 布局间距 */
--spacing-section-gap: var(--space-16);         /* 64px */
--spacing-container-padding: var(--space-4);    /* 16px */

列表间距
--spacing-list-item: var(--space-3);            /* 12px */
```

---

## 尺寸 Tokens

### 容器尺寸
```css
--size-container-sm: 640px;      /* 小容器 */
--size-container-md: 768px;      /* 中容器 */
--size-container-lg: 1024px;     /* 大容器 */
--size-container-xl: 1280px;     /* 超大容器 */
--size-container-2xl: 1536px;    /* 最大容器 */
```

### 组件尺寸
```css
/* 按钮高度 */
--size-button-sm: 2rem;          /* 32px */
--size-button-md: 2.5rem;        /* 40px */
--size-button-lg: 3rem;          /* 48px */

/* 输入框高度 */
--size-input-sm: 2rem;           /* 32px */
--size-input-md: 2.5rem;         /* 40px */
--size-input-lg: 3rem;           /* 48px */

/* 头像尺寸 */
--size-avatar-xs: 1.5rem;        /* 24px */
--size-avatar-sm: 2rem;          /* 32px */
--size-avatar-md: 2.5rem;        /* 40px */
--size-avatar-lg: 3rem;          /* 48px */
--size-avatar-xl: 4rem;          /* 64px */

/* 图标尺寸 */
--size-icon-xs: 1rem;            /* 16px */
--size-icon-sm: 1.25rem;         /* 20px */
--size-icon-md: 1.5rem;          /* 24px */
--size-icon-lg: 1.875rem;        /* 30px */
--size-icon-xl: 2.25rem;         /* 36px */
```

---

## 阴影 Tokens

```css
--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
--shadow-DEFAULT: 0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06);
--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
--shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
--shadow-xl: 0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04);
--shadow-2xl: 0 25px 50px -12px rgba(0, 0, 0, 0.25);

/* 特殊阴影 */
--shadow-card: var(--shadow-sm);
--shadow-button: var(--shadow-sm);
--shadow-dropdown: var(--shadow-lg);
--shadow-modal: var(--shadow-xl);

深色模式阴影
--shadow-dark-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.2);
--shadow-dark-md: 0 4px 6px -1px rgba(0, 0, 0, 0.3), 0 2px 4px -1px rgba(0, 0, 0, 0.2);
--shadow-dark-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.4), 0 4px 6px -2px rgba(0, 0, 0, 0.3);
```

---

## 圆角 Tokens

```css
--radius-none: 0;
--radius-sm: 0.125rem;    /* 2px */
--radius-DEFAULT: 0.5rem; /* 8px - 默认圆角 */
--radius-md: 0.375rem;    /* 6px */
--radius-lg: 0.5rem;      /* 8px */
--radius-xl: 0.75rem;     /* 12px */
--radius-2xl: 1rem;       /* 16px */
--radius-3xl: 1.5rem;     /* 24px */
--radius-full: 9999px;    /* 完全圆角 */

/* 组件圆角 */
--radius-button: var(--radius-md);         /* 6px */
--radius-input: var(--radius-md);          /* 6px */
--radius-card: var(--radius-lg);           /* 8px */
--radius-modal: var(--radius-xl);          /* 12px */
--radius-avatar: var(--radius-full);       /* 完全圆角 */
```

---

## 动画 Tokens

### 过渡时间
```css
--duration-fast: 150ms;
--duration-normal: 250ms;
--duration-slow: 350ms;
--duration-slower: 500ms;
```

### 缓动函数
```css
--ease-linear: linear;
--ease-in: cubic-bezier(0.4, 0, 1, 1);
--ease-out: cubic-bezier(0, 0, 0.2, 1);
--ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);
```

### 过渡属性
```css
--transition-colors: color var(--duration-normal) var(--ease-in-out);
--transition-opacity: opacity var(--duration-normal) var(--ease-in-out);
--transition-shadow: box-shadow var(--duration-normal) var(--ease-in-out);
--transition-transform: transform var(--duration-normal) var(--ease-in-out);
--transition-all: all var(--duration-normal) var(--ease-in-out);

/* 常用组合 */
--transition-button: var(--transition-colors), var(--transition-shadow);
--transition-card: var(--transition-shadow), var(--transition-transform);
--transition-modal: opacity var(--duration-normal) var(--ease-in-out), transform var(--duration-normal) var(--ease-in-out);
```

---

## Z-Index Tokens

```css
--z-index-dropdown: 1000;
--z-index-sticky: 1020;
--z-index-fixed: 1030;
--z-index-modal-backdrop: 1040;
--z-index-modal: 1050;
--z-index-popover: 1060;
--z-index-tooltip: 1070;
--z-index-toast: 1080;
```

---

## 使用指南

### 在 CSS 中使用
```css
.my-component {
  color: var(--color-text-primary-light);
  background-color: var(--color-background-light);
  padding: var(--spacing-md);
  border-radius: var(--radius-card);
  box-shadow: var(--shadow-card);
  transition: var(--transition-all);
}
```

### 在 Tailwind CSS 中配置
```javascript
module.exports = {
  theme: {
    extend: {
      colors: {
        primary: 'var(--color-primary)',
        'text-primary-light': 'var(--color-text-primary-light)',
        // ... 其他颜色
      },
      spacing: {
        'xs': 'var(--spacing-xs)',
        'sm': 'var(--spacing-sm)',
        // ... 其他间距
      },
      boxShadow: {
        'card': 'var(--shadow-card)',
        // ... 其他阴影
      }
    }
  }
}
```

### 响应式设计令牌
```css
@media (prefers-color-scheme: dark) {
  :root {
    --color-text-primary: var(--color-text-primary-dark);
    --color-background: var(--color-background-dark);
    --color-card: var(--color-card-dark);
    --color-border: var(--color-border-dark);
  }
}
```

---

## 维护指南

1. **新增设计令牌时**：
   - 评估是否真的需要新的令牌
   - 遵循命名约定
   - 在所有相关文档中更新

2. **修改现有令牌时**：
   - 评估对现有组件的影响
   - 考虑向后兼容性
   - 通知开发团队

3. **定期审查**：
   - 检查未使用的令牌
   - 评估令牌的使用频率
   - 优化令牌的组织结构

---

**最后更新**: 2024年10月9日  
**维护者**: StyleSense 设计团队
