# StyleSense 颜色系统 (Color System)

> 颜色系统定义了 StyleSense 品牌的视觉语言，通过系统化的颜色规范确保产品界面的一致性、易用性和美观性。

## 📋 目录
- [设计理念](#设计理念)
- [调色板概览](#调色板概览)
- [主色调系统](#主色调系统)
- [中性色调系统](#中性色调系统)
- [功能色彩系统](#功能色彩系统)
- [语义化颜色](#语义化颜色)
- [深色模式适配](#深色模式适配)
- [颜色应用规范](#颜色应用规范)
- [可访问性指南](#可访问性指南)

---

## 设计理念

### 核心原则
1. **简洁优雅** - 采用极简的黑白主色调，体现现代时尚感
2. **层次清晰** - 通过灰度变化建立信息层级
3. **功能明确** - 颜色服务于功能，不干扰内容
4. **品牌一致** - 统一的色彩语言，强化品牌识别

### 设计哲学
StyleSense 采用**黑白灰极简主义**风格，这种设计选择基于：
- **时尚属性** - 黑白色调最符合时尚调性
- **内容聚焦** - 让用户的穿搭内容成为视觉焦点
- **易读性** - 高对比度确保信息清晰
- **可持续性** - 简单的色彩体系易于维护和扩展

---

## 调色板概览

### 完整色板
```
主色调:
Black #000000 ━━━━━━━━━━━━━━━━━━━

中性色调:
Gray 50   #F9FAFB ██████████
Gray 100  #F3F4F6 ████████
Gray 200  #E5E7EB ██████
Gray 300  #D1D5DB ████
Gray 400  #9CA3AF ██
Gray 500  #6B7280 █
Gray 600  #4B5563 ▓
Gray 700  #374151 ▒
Gray 800  #1F2937 ░
Gray 900  #111827 

功能色:
Success  #10B981 (绿)
Warning  #F59E0B (橙)
Error    #EF4444 (红)
Info     #3B82F6 (蓝)
```

---

## 主色调系统

### 黑色主调
```css
--color-primary: #000000;           /* 主色 */
--color-primary-hover: #333333;     /* 悬停态 - 20% 透明度 */
--color-primary-active: #1a1a1a;    /* 激活态 - 10% 黑色 */
--color-primary-light: #666666;     /* 浅色态 - 60% 黑色 */
--color-primary-disabled: #999999;  /* 禁用态 - 40% 透明度 */
```

### 品牌应用场景
- **#000000**: 主要交互元素（按钮、链接、重点文字）
- **#333333**: 次要交互元素（次要按钮、副标题）
- **#666666**: 辅助元素（图标、分隔线）
- **#999999**: 禁用状态或占位符

### 合成色（基于主色）
```css
/* 渐变背景 */
-gradient-primary: linear-gradient(135deg, #000000 0%, #333333 100%);

/* 阴影色 */
-shadow-primary: rgba(0, 0, 0, 0.1);   /* 浅阴影 */
-shadow-primary-hover: rgba(0, 0, 0, 0.2); /* 悬停阴影 */
-shadow-primary-active: rgba(0, 0, 0, 0.3); /* 激活阴影 */
```

---

## 中性色调系统

### 浅色模式中性色
```css
/* 从白到灰的渐进 */
--color-white: #FFFFFF;        /* 纯白 - 高亮背景 */
--color-gray-50: #F9FAFB;      /* 极浅灰 - 次要背景 */
--color-gray-100: #F3F4F6;     /* 浅灰 - 卡片背景 */
--color-gray-200: #E5E7EB;     /* 边框、分隔线 */
--color-gray-300: #D1D5DB;     /* 禁用边框 */
--color-gray-400: #9CA3AF;     /* 占位符文本 */
--color-gray-500: #6B7280;     /* 次要文本 */
--color-gray-600: #4B5563;     /* 中等强调文本 */
--color-gray-700: #374151;     /* 强调文本 */
--color-gray-800: #1F2937;     /* 主要文本 */
--color-gray-900: #111827;     /* 标题文本 */
```

### 深色模式中性色
```css
/* 从黑到灰的渐进 */
--color-dark-950: #030712;     /* 纯黑 - 深度背景 */
--color-dark-900: #111827;     /* 极深灰 - 主背景 */
--color-dark-800: #1F2937;     /* 深灰 - 卡片背景 */
--color-dark-700: #374151;     /* 边框、分隔线 */
--color-dark-600: #4B5563;     /* 禁用边框 */
--color-dark-500: #6B7280;     /* 占位符文本 */
--color-dark-400: #9CA3AF;     /* 次要文本 */
--color-dark-300: #D1D5DB;     /* 中等强调文本 */
--color-dark-200: #E5E7EB;     /* 强调文本 */
--color-dark-100: #F3F4F6;     /* 主要文本 */
--color-dark-50: #F9FAFB;      /* 高亮文本 */
```

### 背景色层映射
```css
/* 浅色模式 */
--color-background-primary: var(--color-white);      /* 主背景 */
--color-background-secondary: var(--color-gray-50);   /* 次要背景 */
--color-background-tertiary: var(--color-gray-100);  /* 三级背景 */
--color-surface-primary: var(--color-white);         /*卡片背景 */
--color-surface-secondary: var(--color-gray-50);     /* 次要卡片 */

/* 深色模式 */
--color-background-primary-dark: var(--color-dark-950); // #030712
--color-background-secondary-dark: var(--color-dark-900); // #111827
--color-background-tertiary-dark: var(--color-dark-800); // #1F2937
--color-surface-primary-dark: var(--color-dark-900); // #111827
--color-surface-secondary-dark: var(--color-dark-800); // #1F2937
```

---

## 功能色彩系统

### 成功色系 (绿色)
```css
/* 主色 */
--color-success: #10B981;
--color-success-light: #34D399;
--color-success-dark: #059669;

/* 背景色 */
--color-success-bg: #D1FAE5;
--color-success-bg-light: #ECFDF5;
--color-success-bg-dark: #064E3B;

/* 文本色 */
--color-success-text: #065F46;
--color-success-text-light: #047857;
```

### 警告色系 (橙色)
```css
/* 主色 */
--color-warning: #F59E0B;
--color-warning-light: #FBBF24;
--color-warning-dark: #D97706;

/* 背景色 */
--color-warning-bg: #FEF3C7;
--color-warning-bg-light: #FFFBEB;
--color-warning-bg-dark: #78350F;

/* 文本色 */
--color-warning-text: #92400E;
--color-warning-text-light: #B45309;
```

### 错误色系 (红色)
```css
/* 主色 */
--color-error: #EF4444;
--color-error-light: #F87171;
--color-error-dark: #DC2626;

/* 背景色 */
--color-error-bg: #FEE2E2;
--color-error-bg-light: #FEF2F2;
--color-error-bg-dark: #7F1D1D;

/* 文本色 */
--color-error-text: #991B1B;
--color-error-text-light: #B91C1C;
```

### 信息色系 (蓝色)
```css
/* 主色 */
--color-info: #3B82F6;
--color-info-light: #60A5FA;
--color-info-dark: #2563EB;

/* 背景色 */
--color-info-bg: #DBEAFE;
--color-info-bg-light: #EFF6FF;
--color-info-bg-dark: #1E3A8A;

/* 文本色 */
--color-info-text: #1E3A8A;
--color-info-text-light: #1D4ED8;
```

---

## 语义化颜色

### 状态颜色映射
```css
/* 通用状态 */
--color-status-online: var(--color-success);      /* 在线 */
--color-status-offline: var(--color-gray-400);    /* 离线 */
--color-status-busy: var(--color-warning);        /* 忙碌 */
--color-status-error: var(--color-error);         /* 错误 */

/* 用户等级 */
--color-level-beginner: var(--color-gray-400);     /* 初级 */
--color-level-intermediate: var(--color-info);     /* 中级 */
--color-level-advanced: var(--color-warning);     /* 高级 */
--color-level-expert: var(--color-success);       /* 专家 */

/* 内容类型 */
--color-content-article: var(--color-primary);     /* 文章 */
--color-content-video: var(--color-error);         /* 视频 */
--color-content-image: var(--color-success);       /* 图片 */
--color-content-link: var(--color-info);           /* 链接 */
```

### 分类标签颜色
```css
/* 衣物分类 */
--color-category-tops: #8B5CF6;     /* 上装 */
--color-category-bottoms: #3B82F6;  /* 下装 */
--color-category-shoes: #10B981;    /* 鞋子 */
--color-category-accessories: #F59E0B; /* 饰品 */

/* 季节标签 */
--color-season-spring: #10B981;     /* 春季 */
--color-season-summer: #F59E0B;     /* 夏季 */
--color-season-autumn: #DC2626;     /* 秋季 */
--color-season-winter: #3B82F6;     /* 冬季 */

/* 风格标签 */
--color-style-casual: #6B7280;      /* 休闲 */
--color-style-formal: #000000;      /* 正装 */
--color-style-sport: #10B981;       /* 运动 */
--color-style-vintage: #92400E;     /* 复古 */
```

---

## 深色模式适配

### 自动切换逻辑
```css
/* CSS 檔测深色模式 */
@media (prefers-color-scheme: dark) {
  :root {
    --color-text-primary: var(--color-text-primary-dark);
    --color-background: var(--color-background-primary-dark);
    --color-surface: var(--color-surface-primary-dark);
    --color-border: var(--color-border-dark);
  }
}

/* 手动切换深色模式 */
.dark {
  --color-text-primary: var(--color-text-primary-dark);
  --color-background: var(--color-background-primary-dark);
  --color-surface: var(--color-surface-primary-dark);
  --color-border: var(--color-border-dark);
}
```

### 深色模式特殊处理
```css
/* 在深色模式下调整功能色饱和度 */
.dark {
  --color-success: #22D3EE;    /* 更亮的青绿色 */
  --color-warning: #FCD34D;    /* 更亮的黄色 */
  --color-error: #F87171;      /* 更亮的红色 */
  --color-info: #60A5FA;       /* 更亮的蓝色 */
}

/* 深色模式阴影调整 */
.dark {
  --shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.3);
  --shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.4), 0 2px 4px -1px rgba(0, 0, 0, 0.3);
  --shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.5), 0 4px 6px -2px rgba(0, 0, 0, 0.4);
}
```

---

## 颜色应用规范

### 基础颜色使用
```css
/* 文本颜色 */
.text-primary { color: var(--color-text-primary); }
.text-secondary { color: var(--color-text-secondary); }
.text-tertiary { color: var(--color-text-tertiary); }
.text-disabled { color: var(--color-text-disabled); }

/* 背景颜色 */
.bg-primary { background-color: var(--color-background-primary); }
.bg-surface { background-color: var(--color-surface-primary); }
.bg-surface-secondary { background-color: var(--color-surface-secondary); }

/* 边框颜色 */
.border-primary { border-color: var(--color-border-primary); }
.border-secondary { border-color: var(--color-border-secondary); }
```

### 主色调应用
```css
/* 主要按钮 */
.btn-primary {
  background-color: var(--color-primary);
  color: var(--color-white);
}

.btn-primary:hover {
  background-color: var(--color-primary-hover);
  box-shadow: var(--shadow-primary-hover);
}

/* 主要链接 */
.link-primary {
  color: var(--color-primary);
  text-decoration: none;
}

.link-primary:hover {
  color: var(--color-primary-hover);
  text-decoration: underline;
}
```

### 功能色应用
```css
/* 状态指示器 */
.status-success { color: var(--color-success); }
.status-warning { color: var(--color-warning); }
.status-error { color: var(--color-error); }
.status-info { color: var(--color-info); }

/* 状态背景 */
.bg-success-subtle { background-color: var(--color-success-bg); }
.bg-warning-subtle { background-color: var(--color-warning-bg); }
.bg-error-subtle { background-color: var(--color-error-bg); }
.bg-info-subtle { background-color: var(--color-info-bg); }

/* 状态边框 */
.border-success { border-color: var(--color-success); }
.border-warning { border-color: var(--color-warning); }
.border-error { border-color: var(--color-error); }
.border-info { border-color: var(--color-info); }
```

---

## 可访问性指南

### 对比度要求
- **WCAG AA 标准**: 正常文字至少 4.5:1，大文字至少 3:1
- **WCAG AAA 标准**: 正常文字至少 7:1，大文字至少 4.5:1

### 推荐对比度组合
```css
/* 高对比度组合（推荐） */
.text-on-primary {
  color: var(--color-white);           /* 白色文字 */
  background-color: var(--color-black); /* 黑色背景 */
  /* 对比度: 21:1 (AAA) */
}

.text-on-surface {
  color: var(--color-gray-900);        /* 深灰文字 */
  background-color: var(--color-white); /* 白色背景 */
  /* 对比度: 15.8:1 (AAA) */
}

.text-on-success {
  color: var(--color-white);
  background-color: var(--color-success);
  /* 对比度: 3.5:1 (AA) */
}
```

### 避免的问题组合
```css
/* 避免：对比度不足 */
.text-low-contrast {
  color: var(--color-gray-400); /* #9CA3AF */
  background-color: var(--color-white); /* #FFFFFF */
  /* 对比度: 3.0:1 (不满足 AA 标准) */
}

/* 避免：颜色仅传递信息 */
.color-only-indicator {
  color: var(--color-error); /* 仅用红色表示错误 */
  /* 应该同时使用图标、文本等辅助手段 */
}
```

### 色盲友好设计
```css
/* 不仅依赖颜色区分 */
.indicator-success::before {
  content: "✓";
  color: var(--color-success);
  background-color: var(--color-success-bg);
  border-radius: 50%;
  /* 同时使用图标和颜色 */
}

.indicator-error::before {
  content: "✗";
  color: var(--color-error);
  background-color: var(--color-error-bg);
  border-radius: 50%;
  /* 使用不同形状避免纯颜色区分 */
}
```

---

## Tailwind CSS 配置示例

```javascript
module.exports = {
  theme: {
    extend: {
      colors: {
        // 主色调
        primary: {
          DEFAULT: '#000000',
          hover: '#333333',
          active: '#1a1a1a',
          light: '#666666',
          disabled: '#999999',
        },
        
        // 中性色
        gray: {
          50: '#F9FAFB',
          100: '#F3F4F6',
          200: '#E5E7EB',
          300: '#D1D5DB',
          400: '#9CA3AF',
          500: '#6B7280',
          600: '#4B5563',
          700: '#374151',
          800: '#1F2937',
          900: '#111827',
        },
        
        // 功能色
        success: {
          DEFAULT: '#10B981',
          light: '#34D399',
          dark: '#059669',
          bg: '#D1FAE5',
          text: '#065F46',
        },
        
        // 语义化颜色
        'status-online': '#10B981',
        'status-offline': '#9CA3AF',
        'status-busy': '#F59E0B',
        'status-error': '#EF4444',
        
        // 深色模式
        dark: {
          950: '#030712',
          900: '#111827',
          800: '#1F2937',
          700: '#374151',
          600: '#4B5563',
          500: '#6B7280',
          400: '#9CA3AF',
          300: '#D1D5DB',
          200: '#E5E7EB',
          100: '#F3F4F6',
          50: '#F9FAFB',
        }
      }
    }
  }
}
```

---

## 使用检查清单

### 设计时检查
- [ ] 使用了正确的颜色令牌，没有硬编码色值
- [ ] 文本颜色和背景色有足够的对比度
- [ ] 功能色使用符合语义（绿色=成功，红色=错误等）
- [ ] 深色模式下颜色仍然清晰可辨
- [ ] 主色调使用适量，避免视觉疲劳

### 开发时检查
- [ ] CSS 变量名称符合命名规范
- [ ] 响应式颜色设置正确
- [ ] 状态颜色（hover、active、disabled）完整
- [ ] 深色模式切换功能正常
- [ ] 可访问性对比度通过检测

### 测试时检查
- [ ] 不同设备显示效果一致
- [ ] 深色/浅色模式切换顺畅
- [ ] 色盲用户可正常使用
- [ ] 高对比度模式可用
- [ ] 打印模式下颜色合适

---

**最后更新**: 2024年10月9日  
**维护者**: StyleSense 设计团队
