# StyleSense Web 前端开发指南

> 基于UI设计系统重构的Next.js 15 + React 19前端开发指南，融合黑白极简主义设计规范与现代React架构最佳实践。

本文档为 AI 助手提供全面的项目开发指南，帮助理解和生成符合统一设计规范的Next.js前端代码。

---

## 1. 项目概览

### 1.1 项目定位
StyleSense 是一个智能服饰管理平台，融合黑白极简主义UI设计系统，帮助用户管理服饰、搭配穿搭等

### 1.2 设计哲学
- **极简美学**：黑白主色调，4px网格系统，遵循UI设计系统规范
- **统一体验**：所有页面遵循统一的设计语言和交互模式
- **深色支持**：完整深色模式实现，与UI系统保持一致
- **响应式设计**：移动端优先，适配多种屏幕尺寸

### 1.3 核心功能模块
基于UI设计系统的完整功能架构：

- **用户认证**：登录/注册、Token 管理、个人资料维护
- **服饰管理**：添加、编辑、删除、分类管理衣物清单
  - 支持多维度筛选（类别、季节、颜色、品牌、穿着频率）
  - 高级筛选面板，支持组合条件筛选
  - 衣物详情展示（穿着次数、标签、图片展示）
  - 两种视图模式：网格视图和列表视图
- **穿搭记录**：记录每日穿搭、拍照上传、穿搭历史查看
  - 穿搭日历展示
  - 拍照记录和保存
  - 穿搭统计和分析
- **衣橱可视化**：衣物分类统计、智能推荐、穿搭分析
  - 衣橱概览统计（总数、价值、待清洗、新增等）
  - 最常穿衣物排行
  - 闲置衣物优化建议
  - 类别分布可视化图表
- **风格灵感**：流行趋势展示、搭配建议、风格标签
  - 流行穿搭推荐
  - 个人收藏管理
  - 风格标签系统

### 1.4 核心页面模板（基于UI系统）
根据UI项目结构，标准页面包含：

- **主页** (`/`): 穿搭推荐、今日穿搭、流行风格展示
- **我的衣橱** (`/wardrobe`): 衣物管理、筛选面板、统计概览
- **穿搭记录** (`/outfits`): 穿搭历史、日历视图、拍照记录
- **风格灵感** (`/inspiration`): 流行趋势、搭配建议、收藏管理
- **个人中心** (`/profile`): 用户信息、设置选项、数据统计
- **通知中心** (`/notifications`): 系统通知、穿搭提醒、活动消息

### 1.5 视觉特色
- **黑白极简**：主色调为纯黑(#000000)，辅以中性灰色调
- **4px网格**：所有间距基于4px倍数，确保视觉一致性
- **Noto Sans SC**：统一使用中文字体，确保品牌一致性
- **Material Icons**：图标系统与UI项目保持一致

### 1.6 技术架构
- **Next.js 15 App Router**：服务端渲染（SSR）、文件系统路由
- **React 19**：函数式组件、Hooks、客户端交互
- **TypeScript 5**：类型安全、接口定义、编译时检查
- **Tailwind CSS 4**：原子化 CSS，集成UI设计系统
- **设计系统移植**：完整实现UI项目的视觉规范
- **前后端分离**：RESTful API通信架构

---

## 2. UI设计系统集成

### 2.1 设计令牌系统
基于UI项目的完整设计令牌，建立React/Tailwind实现映射：

#### 颜色系统
```typescript
// tailwind.config.ts
colors: {
  // 主色调
  primary: '#000000',
  'primary-hover': '#333333',
  'primary-active': '#1a1a1a',
  'primary-disabled': '#666666',

  // 背景色
  'background-light': '#F8F9FA',
  'background-dark': '#121212',
  'card-light': '#FFFFFF',
  'card-dark': '#1E1E1E',

  // 边框色
  'border-light': '#E5E7EB',
  'border-dark': '#374151',

  // 文字色
  'text-primary-light': '#1F2937',
  'text-primary-dark': '#F9FAFB',
  'text-secondary-light': '#6B7280',
  'text-secondary-dark': '#9CA3AF',

  // 功能色
  success: '#10B981',
  warning: '#F59E0B',
  error: '#EF4444',
  info: '#3B82F6',
}
```

#### 间距系统（4px网格）
```typescript
// 基础间距（4px倍数）
spacing: {
  'xs': '0.25rem',  // 4px
  'sm': '0.5rem',   // 8px
  'md': '1rem',     // 16px
  'lg': '1.5rem',   // 24px
  'xl': '2rem',     // 32px
  '2xl': '3rem',    // 48px
  '3xl': '4rem',    // 64px
}
```

#### 字体系统
```typescript
// 字体配置
fontFamily: {
  display: ['"Noto Sans SC"', 'sans-serif'],
},
fontSize: {
  'xs': '0.75rem',   // 12px
  'sm': '0.875rem',  // 14px
  'base': '1rem',    // 16px
  'lg': '1.125rem',  // 18px
  'xl': '1.25rem',   // 20px
  '2xl': '1.5rem',   // 24px
}
```

### 2.2 设计令牌映射表

| UI设计令牌 | CSS变量 | Tailwind类名 | 用途 |
|------------|---------|--------------|------|
| `--color-primary` | `var(--color-primary)` | `bg-primary` | 主按钮背景 |
| `--color-background-light` | `var(--color-background-light)` | `bg-background-light` | 浅色背景 |
| `--color-card-light` | `var(--color-card-light)` | `bg-card-light` | 卡片背景 |
| `--spacing-md` | `var(--spacing-md)` | `p-4` | 标准内边距 |
| `--radius-card` | `var(--radius-card)` | `rounded-lg` | 卡片圆角 |
| `--shadow-card` | `var(--shadow-card)` | `shadow-card` | 卡片阴影 |

### 2.3 深色模式实现
```typescript
// 深色模式配置
darkMode: "class",
// 使用方式
<div className="bg-card-light dark:bg-card-dark">
  <p className="text-text-primary-light dark:text-text-primary-dark">
    适配深色模式的文字
  </p>
</div>
```

### 2.4 组件样式标准
```typescript
// 标准卡片样式
const cardStyles = `
  bg-card-light dark:bg-card-dark
  rounded-lg shadow-sm
  border border-border-light dark:border-border-dark
  p-6 hover:shadow-md transition-shadow
`;

// 标准按钮样式
const buttonStyles = `
  bg-primary hover:bg-primary-hover
  text-white px-4 py-2 rounded-lg
  transition-colors flex items-center gap-2
`;
```

---

## 3. 技术栈详解

### 2.1 核心框架
- **Next.js 15.5.3**：App Router、服务端渲染（SSR）、文件系统路由
- **React 19.1.0**：函数式组件、Hooks、客户端交互
- **TypeScript 5**：类型安全、接口定义、编译时检查

### 2.2 样式方案
- **Tailwind CSS 4**：原子化 CSS 类、响应式设计、主题定制
- **模块化 CSS**：组件专属样式文件（如 `Modal.css`、`Toast.css`）

### 2.3 开发工具
- **ESLint**：代码规范检查
- **pnpm**：依赖管理工具

### 2.4 API 通信
- **axios**（通过 `lib/api/http.ts` 封装）：HTTP 请求、错误拦截、Token 注入

---

## 3. 目录结构与文件组织

### 3.1 完整目录树
```
client/web/
├── src/
│   ├── app/                    # Next.js App Router 页面
│   │   ├── layout.tsx          # 根布局
│   │   ├── page.tsx            # 首页（重定向到 /main）
│   │   ├── login/              # 登录页面
│   │   ├── register/           # 注册页面
│   │   └── main/               # 主应用页面
│   │       ├── layout.tsx      # 主应用布局（侧边栏）
│   │       └── page.tsx        # 主控面板
│   ├── components/             # 可复用组件
│   │   ├── Modal.tsx           # 模态框容器
│   │   ├── Toast.tsx           # 提示组件
│   │   ├── ToastProvider.tsx   # Toast 上下文提供者
│   │   ├── AddClothingItem.tsx # 新增衣物表单
│   │   └── WardrobeLayout.tsx  # 衣橱布局组件
│   ├── lib/                    # 工具与 API 层
│   │   ├── api/                # API 调用封装
│   │   │   ├── http.ts         # 统一 HTTP 客户端
│   │   │   ├── auth.ts         # 认证 API
│   │   │   ├── user.ts         # 用户 API
│   │   │   ├── clothing.ts     # 衣物 API
│   │   └── config/
│   │       └── env.ts          # 环境配置
│   ├── styles/                 # 全局与组件样式
│   │   ├── globals.css         # 全局样式
│   │   ├── Modal.css           # 模态框样式
│   │   ├── Toast.css           # Toast 样式
│   │   ├── sidebar-layout.css  # 侧边栏样式
│   │   └── modern.css          # 现代化组件样式
│   └── types/                  # TypeScript 类型定义
│       ├── auth.ts             # 认证相关类型
│       ├── user.ts             # 用户类型
│       ├── clothing.ts         # 衣物类型
│       └── weather.ts          # 天气类型
├── public/                     # 静态资源
├── package.json                # 依赖配置
├── tsconfig.json               # TypeScript 配置
└── tailwind.config.ts          # Tailwind 配置
```

### 3.2 文件组织原则
- **页面文件**：放在 `src/app/` 下，遵循 Next.js App Router 约定
- **复用组件**：放在 `src/components/` 下，组件名使用 PascalCase
- **API 调用**：统一放在 `src/lib/api/` 下，按功能模块划分
- **类型定义**：放在 `src/types/` 下，与 API 模块对应
- **样式文件**：全局样式在 `src/styles/globals.css`，组件样式与组件同名放在 `src/styles/` 下

### 3.3 命名约定
- **组件文件**：`PascalCase.tsx`（如 `AddClothingItem.tsx`）
- **样式文件**：`PascalCase.css` 或 `kebab-case.css`
- **API 文件**：`lowercase.ts`（如 `auth.ts`、`clothing.ts`）
- **类型文件**：与 API 模块同名（如 `types/clothing.ts` 对应 `lib/api/clothing.ts`）
- **页面路由**：遵循 Next.js 约定，文件夹名小写（如 `login/`、`register/`）

---

## 4. UI 规范与设计系统

### 4.1 布局规范
- **侧边栏布局**：主应用使用 `src/app/main/layout.tsx` 定义的侧边栏布局
- **响应式设计**：使用 Tailwind 断点（`sm:`、`md:`、`lg:`）适配不同屏幕
- **间距系统**：使用 Tailwind 间距类（`p-4`、`m-2`、`gap-6` 等）

### 4.2 组件样式模式
- **Tailwind 优先**：大部分样式使用 Tailwind 原子类
- **复杂组件样式**：使用独立 CSS 文件（如 `Modal.css`、`sidebar-layout.css`）
- **样式导入顺序**：全局样式 → 组件样式 → Tailwind

### 4.3 颜色与主题
- **主色调**：参考 `src/styles/globals.css` 中的 CSS 变量定义
- **状态颜色**：
  - 成功：`text-green-600`、`bg-green-100`
  - 警告：`text-yellow-600`、`bg-yellow-100`
  - 错误：`text-red-600`、`bg-red-100`
  - 信息：`text-blue-600`、`bg-blue-100`

### 4.4 交互反馈
- **加载状态**：使用 loading 标志控制按钮禁用与文案变化
- **错误提示**：统一使用 Toast 组件展示错误信息
- **成功反馈**：使用 Toast 组件展示成功信息

---

## 5. 路由与导航模式

### 5.1 App Router 路由规则
- **文件系统路由**：`src/app/` 下的文件夹结构对应 URL 路径
  - `src/app/login/page.tsx` → `/login`
  - `src/app/main/page.tsx` → `/main`
- **布局嵌套**：`layout.tsx` 定义共享布局，子路由自动嵌套
- **动态路由**：使用 `[param]/page.tsx` 定义动态参数（项目当前未使用）

### 5.2 导航跳转
- **客户端跳转**：使用 `next/navigation` 的 `useRouter` Hook
  ```tsx
  'use client';
  import { useRouter } from 'next/navigation';
  
  const router = useRouter();
  router.push('/login');
  ```
- **重定向**：在 `useEffect` 中检查条件后跳转
  ```tsx
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      router.push('/login');
    }
  }, [router]);
  ```

### 5.3 认证保护
- **Token 校验**：主应用页面（`/main` 及子路由）在 `useEffect` 中检查 `localStorage.getItem('token')`
- **无 Token 重定向**：跳转到 `/login`
- **登录后跳转**：登录成功后跳转到 `/main`

---

## 6. 数据获取与状态管理

### 6.1 客户端数据获取模式
- **使用 `useEffect` + `useState`**：在客户端组件中加载数据
  ```tsx
  'use client';
  import { useEffect, useState } from 'react';
  import { getUserInfo } from '@/lib/api/user';
  
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  
  useEffect(() => {
    const fetchUser = async () => {
      try {
        const data = await getUserInfo();
        setUser(data);
      } catch (error) {
        console.error('获取用户信息失败', error);
      } finally {
        setLoading(false);
      }
    };
    fetchUser();
  }, []);
  ```

### 6.2 状态管理模式
- **组件内状态**：使用 `useState` 管理局部状态（表单输入、加载状态、错误信息）
- **跨组件共享**：使用 React Context（如 `ToastProvider`）
- **全局状态**：项目当前未使用 Redux/Zustand，通过 Context + localStorage 管理认证状态

### 6.3 加载与错误处理
- **加载状态**：使用 `loading` 状态变量控制 UI 展示
  ```tsx
  {loading ? <div>加载中...</div> : <div>{data}</div>}
  ```
- **错误处理**：在 `catch` 块中捕获错误并展示 Toast
  ```tsx
  try {
    await someAPICall();
  } catch (error: any) {
    showToast(error.response?.data?.message || '操作失败', 'error');
  }
  ```

---

## 7. API 调用规范

### 7.1 HTTP 客户端封装
- **统一入口**：`src/lib/api/http.ts` 导出封装后的 axios 实例
- **自动 Token 注入**：请求拦截器自动从 `localStorage` 读取 token 并添加到 `Authorization` 头
- **401 错误拦截**：响应拦截器检测 401 状态码，自动清除 token 并跳转到 `/login`
- **错误统一处理**：返回 `Promise.reject(error)` 供调用方捕获

**http.ts 核心逻辑**：
```typescript
import axios from 'axios';
import { API_BASE_URL } from '@/lib/config/env';

const http = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
});

// 请求拦截器
http.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// 响应拦截器
http.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default http;
```

### 7.2 API 模块划分
- **认证 API**（`src/lib/api/auth.ts`）：`login`、`register`
- **用户 API**（`src/lib/api/user.ts`）：`getUserInfo`、`updateUserInfo`
- **衣物 API**（`src/lib/api/clothing.ts`）：`getClothingList`、`addClothing`、`updateClothing`、`deleteClothing`
- **天气 API**（`src/lib/api/weather.ts`）：`getWeather`、`getWeatherSuggestion`

### 7.3 API 调用示例
**登录接口**：
```typescript
// src/lib/api/auth.ts
import http from './http';
import type { LoginRequest, LoginResponse } from '@/types/auth';

export const login = async (data: LoginRequest): Promise<LoginResponse> => {
  const response = await http.post<LoginResponse>('/auth/login', data);
  return response.data;
};
```

**使用示例**：
```tsx
// 在组件中调用
import { login } from '@/lib/api/auth';

const handleLogin = async () => {
  try {
    const result = await login({ username, password });
    localStorage.setItem('token', result.token);
    router.push('/main');
  } catch (error: any) {
    showToast(error.response?.data?.message || '登录失败', 'error');
  }
};
```

### 7.4 类型定义规范
- **请求类型**：以 `Request` 结尾（如 `LoginRequest`）
- **响应类型**：以 `Response` 结尾（如 `LoginResponse`）
- **数据模型**：使用清晰的接口名（如 `User`、`ClothingItem`、`WeatherData`）

---

## 8. 组件开发指南（基于UI设计系统）

### 8.1 设计系统优先原则
- **严格遵循UI设计令牌**：所有视觉属性必须使用设计系统定义的值
- **深色模式强制**：每个组件必须同时支持深色模式
- **4px网格对齐**：所有间距必须是4px的倍数
- **语义化命名**：使用设计系统的命名规范

### 8.2 客户端组件标记
- **必须标记 `'use client'`**：使用 Hooks、事件处理、浏览器 API 的组件必须在文件顶部添加
  ```tsx
  'use client';
  import { useState } from 'react';
  ```

### 8.3 组件样式标准

#### 8.3.1 基础组件样式模板
```tsx
// 标准按钮组件（适配UI设计系统）
const Button = ({
  variant = 'primary',
  size = 'md',
  children,
  ...props
}) => {
  const baseStyles = `
    inline-flex items-center justify-center
    font-medium transition-colors
    focus:outline-none focus:ring-2 focus:ring-primary
    disabled:opacity-50 disabled:cursor-not-allowed
  `;

  const variantStyles = {
    primary: `
      bg-primary text-white
      hover:bg-primary-hover
      active:bg-primary-active
      px-4 py-2 rounded-lg
    `,
    secondary: `
      bg-card-light dark:bg-card-dark
      text-text-primary-light dark:text-text-primary-dark
      border border-border-light dark:border-border-dark
      hover:bg-gray-50 dark:hover:bg-gray-800
      px-4 py-2 rounded-lg
    `,
  };

  const sizeStyles = {
    sm: 'h-8 px-3 text-sm',
    md: 'h-10 px-4 text-base',
    lg: 'h-12 px-6 text-lg',
  };

  return (
    <button
      className={`${baseStyles} ${variantStyles[variant]} ${sizeStyles[size]}`}
      {...props}
    >
      {children}
    </button>
  );
};
```

#### 8.3.2 卡片组件（UI设计系统适配）
```tsx
// 标准卡片组件
const Card = ({ children, className = '', hoverable = false }) => {
  return (
    <div className={`
      bg-card-light dark:bg-card-dark
      rounded-lg shadow-sm
      border border-border-light dark:border-border-dark
      p-6
      ${hoverable ? 'hover:shadow-md transition-shadow cursor-pointer' : ''}
      ${className}
    `}>
      {children}
    </div>
  );
};

// 卡片头部组件
const CardHeader = ({ title, subtitle, action }) => (
  <div className="flex items-center justify-between mb-4">
    <div>
      <h3 className="text-lg font-medium text-text-primary-light dark:text-text-primary-dark">
        {title}
      </h3>
      {subtitle && (
        <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">
          {subtitle}
        </p>
      )}
    </div>
    {action && <div>{action}</div>}
  </div>
);

// 卡片内容组件
const CardContent = ({ children, className = '' }) => (
  <div className={`text-text-primary-light dark:text-text-primary-dark ${className}`}>
    {children}
  </div>
);
```

#### 8.3.3 表单组件（UI设计系统）
```tsx
// 标准输入框组件
const Input = ({
  label,
  error,
  required = false,
  className = '',
  ...props
}) => {
  return (
    <div className="space-y-2">
      {label && (
        <label className={`
          block text-sm font-medium
          ${error ? 'text-error' : 'text-text-primary-light dark:text-text-primary-dark'}
        `}>
          {label}
          {required && <span className="text-error ml-1">*</span>}
        </label>
      )}
      <input
        className={`
          w-full px-3 py-2
          bg-card-light dark:bg-card-dark
          border rounded-lg
          focus:outline-none focus:ring-2 focus:ring-primary
          text-text-primary-light dark:text-text-primary-dark
          ${error
            ? 'border-error focus:border-error'
            : 'border-border-light dark:border-border-dark focus:border-primary'
          }
          ${className}
        `}
        {...props}
      />
      {error && (
        <p className="text-sm text-error">{error}</p>
      )}
    </div>
  );
};
```

### 8.4 复用组件清单（基于UI设计系统）

#### 8.4.1 UI基础组件（新增）
**路径**：`src/components/ui/`

基于UI设计系统新增的通用组件：

```tsx
// Button.tsx - 标准按钮
// Card.tsx, CardHeader.tsx, CardContent.tsx - 卡片组件
// Input.tsx, Textarea.tsx - 表单输入
// Badge.tsx - 标签徽章
// Avatar.tsx - 用户头像
// Divider.tsx - 分隔线
// LoadingSpinner.tsx - 加载动画
```

#### 8.4.2 Modal 组件（已适配设计系统）
**路径**：`src/components/Modal.tsx`

**设计系统适配**：
- 背景色使用 `bg-card-light dark:bg-card-dark`
- 边框使用 `border-border-light dark:border-border-dark`
- 圆角使用 `rounded-lg` (匹配UI设计)
- 阴影使用 `shadow-modal`

**使用示例**（适配UI设计系统）：
```tsx
<Modal isOpen={isOpen} onClose={onClose} title="添加新衣物">
  <div className="space-y-4">
    <!-- 模态框内容使用UI设计系统样式 -->
    <Card>
      <CardHeader title="基本信息" />
      <CardContent>
        <Input label="衣物名称" placeholder="请输入衣物名称" />
      </CardContent>
    </Card>
  </div>
</Modal>
```

#### 8.4.3 Toast 组件（已适配设计系统）
**路径**：`src/components/Toast.tsx` + `src/components/ToastProvider.tsx`

**设计系统适配**：
- 背景色适配深色模式
- 使用功能色系统（success/error/warning/info）
- 圆角和阴影符合UI设计规范

**使用示例**：
```tsx
const { showToast } = useToast();
// 使用功能色系统
showToast('操作成功', 'success');  // 绿色
showToast('操作失败', 'error');    // 红色
showToast('请注意', 'warning');    // 橙色
showToast('提示信息', 'info');     // 蓝色
```

#### 8.4.4 AddClothingItem 组件（设计系统重构版）
**路径**：`src/components/AddClothingItem.tsx`

**重构要点**：
- 使用新的UI组件库（Card、Input、Button）
- 表单布局遵循4px网格系统
- 完整深色模式支持
- 颜色选择器适配黑白极简风格

**重构示例**：
```tsx
// 原代码 - 使用内联样式
<div style={{ marginBottom: '16px' }}>
  <label>衣物名称</label>
  <input style={{ padding: '8px' }} />
</div>

// 新代码 - 使用UI设计系统
<div className="mb-4">  {/* 4px网格：mb-4 = 16px */}
  <Input
    label="衣物名称"
    placeholder="请输入衣物名称"
    className="w-full"
  />
</div>
```

### 8.5 页面布局标准（基于UI设计系统）

#### 8.5.1 标准页面结构
```tsx
// 基于UI设计系统的标准页面模板
const StandardPage = ({ children, title, subtitle }) => (
  <div className="min-h-screen bg-background-light dark:bg-background-dark">
    {/* 导航栏 - 基于UI shared/header.html */}
    <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center space-x-8">
            <h1 className="text-xl font-bold text-text-primary-light dark:text-text-primary-dark">
              StyleSense
            </h1>
            <nav className="hidden md:flex space-x-8">
              {/* 导航链接 */}
            </nav>
          </div>
          <div className="flex items-center space-x-4">
            {/* 搜索、通知、用户头像 */}
          </div>
        </div>
      </div>
    </header>

    {/* 主内容区 */}
    <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-text-primary-light dark:text-text-primary-dark">
          {title}
        </h1>
        {subtitle && (
          <p className="text-text-secondary-light dark:text-text-secondary-dark mt-2">
            {subtitle}
          </p>
        )}
      </div>
      {children}
    </main>
  </div>
);
```

#### 8.5.2 响应式布局标准
```tsx
// 移动端优先的响应式布局
const ResponsiveLayout = () => (
  <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 md:gap-6 lg:gap-8">
    {/* 卡片网格 - 1列(移动) → 2列(平板) → 3列(桌面) */}
    <Card className="col-span-1">
      <CardContent>移动端单列，桌面端三列</CardContent>
    </Card>
  </div>
);

// 侧边栏布局（参考UI设计）
const SidebarLayout = () => (
  <div className="flex flex-col lg:flex-row gap-6 lg:gap-8">
    {/* 主内容区 */}
    <div className="flex-1">
      <Card><CardContent>主内容区域</CardContent></Card>
    </div>

    {/* 侧边栏 */}
    <div className="lg:w-80">
      <Card><CardContent>侧边栏内容</CardContent></Card>
    </div>
  </div>
);
```

### 8.6 表单组件开发模式（UI设计系统版）
**模式总结**（基于UI设计系统）：
1. **组件组合**：使用UI基础组件（Input、Button、Card）
2. **布局标准**：遵循4px网格系统，使用标准间距
3. **深色模式**：所有组件自动适配深色模式
4. **验证逻辑**：统一的错误提示样式
5. **加载状态**：使用LoadingSpinner组件
6. **成功反馈**：使用功能色Toast系统

**典型代码结构**（UI设计系统版）：
```tsx
'use client';
import { useState } from 'react';
import { useToast } from '@/components/ToastProvider';
import { Card, CardHeader, CardContent } from '@/components/ui/Card';
import { Input } from '@/components/ui/Input';
import { Button } from '@/components/ui/Button';
import { LoadingSpinner } from '@/components/ui/LoadingSpinner';

export default function ClothingForm() {
  const [name, setName] = useState('');
  const [loading, setLoading] = useState(false);
  const { showToast } = useToast();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name.trim()) {
      showToast('请填写衣物名称', 'warning');
      return;
    }

    setLoading(true);
    try {
      await createClothingItem({ name });
      showToast('衣物添加成功', 'success');
      setName(''); // 重置表单
    } catch (error: any) {
      showToast(error.response?.data?.message || '添加失败', 'error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card>
      <CardHeader title="添加新衣物" subtitle="完善您的衣橱收藏" />
      <form onSubmit={handleSubmit}>
        <CardContent className="space-y-4">  {/* 4px网格间距 */}
          <Input
            label="衣物名称"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="请输入衣物名称"
            disabled={loading}
          />
          <div className="flex gap-4">  {/* 4px网格间距 */}
            <Button type="submit" disabled={loading}>
              {loading ? <LoadingSpinner /> : '添加衣物'}
            </Button>
            <Button type="button" variant="secondary">
              取消
            </Button>
          </div>
        </CardContent>
      </form>
    </Card>
  );
}
```
```tsx
import Modal from '@/components/Modal';

const [isOpen, setIsOpen] = useState(false);

<Modal isOpen={isOpen} onClose={() => setIsOpen(false)} title="标题">
  <div>模态框内容</div>
</Modal>
```

**注意事项**：
- 样式文件：`src/styles/Modal.css`
- 支持自定义宽度、层级、动画效果

#### 8.2.2 Toast 组件
**路径**：`src/components/Toast.tsx` + `src/components/ToastProvider.tsx`

**用途**：全局提示反馈

**使用示例**：
```tsx
// 1. 在根布局中包裹 ToastProvider
import { ToastProvider } from '@/components/ToastProvider';

<ToastProvider>
  {children}
</ToastProvider>

// 2. 在组件中使用
import { useToast } from '@/components/ToastProvider';

const { showToast } = useToast();
showToast('操作成功', 'success');
showToast('操作失败', 'error');
```

**类型定义**：
- `type: 'success' | 'error' | 'info' | 'warning'`
- 自动 3 秒后关闭

#### 8.2.3 AddClothingItem 组件
**路径**：`src/components/AddClothingItem.tsx`

**用途**：新增衣物复杂表单，展示表单拆分、动态属性、图片上传模式

**关键特性**：
- 表单字段分组（基础信息、分类、属性）
- 动态属性输入（根据类别展示不同属性）
- 图片预览与上传
- 表单验证与错误提示
- Toast 反馈成功/失败

**参考要点**：
- 使用多个 `useState` 管理表单字段
- `useEffect` 监听分类变化，动态调整属性字段
- API 调用后清空表单重置状态

### 8.3 表单组件开发模式
**模式总结**（参考 `AddClothingItem.tsx`）：
1. **状态管理**：每个输入字段对应一个 `useState`
2. **受控组件**：`value={state}` + `onChange={(e) => setState(e.target.value)}`
3. **验证逻辑**：提交前检查必填字段
4. **加载状态**：提交时设置 `loading` 状态，禁用按钮
5. **错误处理**：捕获 API 错误并展示 Toast
6. **成功反馈**：提交成功后展示 Toast 并重置表单

**典型代码结构**：
```tsx
'use client';
import { useState } from 'react';
import { useToast } from '@/components/ToastProvider';
import { someAPI } from '@/lib/api/some';

export default function FormComponent() {
  const [field1, setField1] = useState('');
  const [loading, setLoading] = useState(false);
  const { showToast } = useToast();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!field1) {
      showToast('请填写必填字段', 'warning');
      return;
    }

    setLoading(true);
    try {
      await someAPI({ field1 });
      showToast('提交成功', 'success');
      setField1(''); // 重置表单
    } catch (error: any) {
      showToast(error.response?.data?.message || '提交失败', 'error');
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={field1}
        onChange={(e) => setField1(e.target.value)}
        disabled={loading}
      />
      <button type="submit" disabled={loading}>
        {loading ? '提交中...' : '提交'}
      </button>
    </form>
  );
}
```

---

## 9. TypeScript 类型规范

### 9.1 类型定义位置
- **API 相关类型**：放在 `src/types/` 下，按模块划分
- **组件 Props 类型**：定义在组件文件内部，使用 `interface` 或 `type`

### 9.2 类型命名规范
- **接口**：使用 PascalCase，语义明确（如 `User`、`ClothingItem`）
- **请求类型**：`[功能]Request`（如 `LoginRequest`、`AddClothingRequest`）
- **响应类型**：`[功能]Response`（如 `LoginResponse`、`GetUserResponse`）
- **Props 类型**：`[组件名]Props`（如 `ModalProps`、`ToastProps`）

### 9.3 类型定义示例
```typescript
// src/types/clothing.ts
export interface ClothingItem {
  id: string;
  name: string;
  category: string;
  color: string;
  season: string;
  imageUrl?: string;
  createdAt: string;
}

export interface AddClothingRequest {
  name: string;
  category: string;
  color: string;
  season: string;
  image?: File;
}

export interface GetClothingListResponse {
  items: ClothingItem[];
  total: number;
}
```

### 9.4 泛型与工具类型
- **API 响应泛型**：`http.post<ResponseType>(url, data)`
- **可选字段**：使用 `?` 标记（如 `imageUrl?: string`）
- **联合类型**：使用 `|` 定义多选值（如 `type: 'success' | 'error'`）

---

## 10. 样式开发规范

### 10.1 Tailwind CSS 使用
- **优先使用 Tailwind 类**：布局、间距、颜色、文字样式
- **响应式设计**：使用断点前缀（`md:flex`、`lg:w-1/2`）
- **状态样式**：使用伪类前缀（`hover:bg-blue-500`、`focus:outline-none`）

**常用模式**：
```tsx
// 布局
<div className="flex items-center justify-between p-4 gap-2">

// 按钮
<button className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 disabled:opacity-50">

// 输入框
<input className="w-full px-3 py-2 border border-gray-300 rounded focus:border-blue-500 focus:outline-none" />

// 卡片
<div className="bg-white rounded-lg shadow p-6">
```

### 10.2 模块化 CSS
- **适用场景**：复杂动画、特定布局、跨组件复用样式
- **导入方式**：
  ```tsx
  import '@/styles/Modal.css';
  ```
- **类名约定**：使用语义化类名，避免与 Tailwind 冲突

### 10.3 样式组织原则
- **全局样式**：放在 `src/styles/globals.css`，包含 Tailwind 指令和全局 CSS 变量
- **组件样式**：放在 `src/styles/[ComponentName].css`
- **内联样式**：仅在动态计算样式值时使用

---

## 11. 错误处理与边界情况

### 11.1 API 错误处理
- **统一拦截**：`http.ts` 响应拦截器处理 401 错误
- **业务错误**：在调用方 `catch` 块中处理，展示 Toast
- **网络错误**：捕获 `error.message` 并提示网络异常

**错误处理模式**：
```tsx
try {
  const result = await someAPI();
} catch (error: any) {
  if (error.response) {
    // 服务端返回错误
    showToast(error.response.data.message || '操作失败', 'error');
  } else if (error.request) {
    // 请求发送但无响应
    showToast('网络异常，请检查连接', 'error');
  } else {
    // 其他错误
    showToast('请求失败', 'error');
  }
}
```

### 11.2 认证边界处理
- **无 Token**：在主应用页面 `useEffect` 中检测并跳转到 `/login`
- **Token 过期**：`http.ts` 拦截 401 响应，清除 Token 并跳转登录
- **重复登录**：登录成功后检查 Token 是否已存在，避免重复设置

### 11.3 表单验证
- **必填字段**：提交前检查空值
- **格式验证**：邮箱、手机号等使用正则校验
- **实时反馈**：输入时展示错误提示（可选）

### 11.4 空数据处理
- **列表为空**：展示空状态提示
- **数据加载失败**：展示错误信息和重试按钮
- **加载中**：展示骨架屏或 Loading 动画

---

## 12. 性能优化建议

### 12.1 组件优化
- **避免不必要渲染**：使用 `React.memo` 包裹纯展示组件
- **回调函数优化**：使用 `useCallback` 缓存回调函数
- **计算属性优化**：使用 `useMemo` 缓存复杂计算结果

### 12.2 图片优化
- **使用 Next.js Image 组件**：`next/image` 提供自动优化、懒加载、响应式
  ```tsx
  import Image from 'next/image';
  <Image src="/path" alt="description" width={300} height={200} />
  ```
- **图片格式**：优先使用 WebP，回退到 JPEG/PNG
- **图片尺寸**：根据展示尺寸提供合适分辨率

### 12.3 代码分割
- **动态导入**：使用 `next/dynamic` 按需加载组件
  ```tsx
  import dynamic from 'next/dynamic';
  const HeavyComponent = dynamic(() => import('./HeavyComponent'), {
    loading: () => <div>加载中...</div>,
  });
  ```

### 12.4 API 请求优化
- **避免重复请求**：使用状态标志防止重复触发
- **并发请求**：使用 `Promise.all` 并行请求
- **请求取消**：使用 `AbortController` 取消未完成的请求

---

## 13. 测试与验证

### 13.1 当前测试状态
- 项目当前 **无自动化测试覆盖**
- 建议补充单元测试（组件、工具函数）和集成测试（API 调用、页面流程）

### 13.2 手动测试清单
**功能测试**：
- [ ] 登录/注册流程正常
- [ ] Token 过期后自动跳转登录
- [ ] 衣物增删改查功能正常
- [ ] Toast 提示正确展示
- [ ] Modal 打开/关闭交互正常

**边界测试**：
- [ ] 无 Token 访问主应用自动跳转
- [ ] API 返回错误时正确展示提示
- [ ] 表单必填字段验证生效
- [ ] 加载状态按钮禁用

**兼容性测试**：
- [ ] Chrome、Firefox、Safari 浏览器测试
- [ ] 移动端响应式布局正常
- [ ] 不同屏幕尺寸下 UI 无错位

### 13.3 代码质量检查
- **运行 Lint**：`pnpm lint` 检查代码规范
- **类型检查**：`pnpm tsc --noEmit` 检查 TypeScript 类型错误
- **构建检查**：`pnpm build` 确保生产环境可正常构建

---

## 14. AI 提示词示例

### 14.1 创建新页面（基于UI设计系统）
**提示词**：
```
请在 src/app/wardrobe/page.tsx 创建衣橱展示页面，要求：
1. 使用客户端组件（'use client'）
2. 使用UI设计系统的标准页面布局（参考StandardPage模板）
3. 在 useEffect 中调用 getClothingList API 获取衣物列表
4. 使用Card组件展示衣物列表，遵循4px网格系统
5. 空数据展示"暂无衣物"提示，使用UI设计系统样式
6. 加载时使用LoadingSpinner组件
7. 错误时使用功能色Toast系统（error类型）
8. 完整深色模式支持
9. 响应式设计：移动端1列，桌面端3列
```

**实现示例**：
```tsx
'use client';
import { useEffect, useState } from 'react';
import { useToast } from '@/components/ToastProvider';
import { Card, CardHeader, CardContent } from '@/components/ui/Card';
import { LoadingSpinner } from '@/components/ui/LoadingSpinner';
import { Button } from '@/components/ui/Button';

export default function WardrobePage() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const { showToast } = useToast();

  useEffect(() => {
    fetchClothingItems();
  }, []);

  const fetchClothingItems = async () => {
    try {
      const data = await getClothingList();
      setItems(data.items);
    } catch (error) {
      showToast('获取衣物列表失败', 'error');
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <LoadingSpinner />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background-light dark:bg-background-dark">
      {/* 标准导航栏 */}
      <header className="bg-card-light dark:bg-card-dark shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h1 className="text-xl font-bold text-text-primary-light dark:text-text-primary-dark">
            我的衣橱
          </h1>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="mb-8">
          <h2 className="text-2xl font-bold text-text-primary-light dark:text-text-primary-dark">
            衣橱管理
          </h2>
        </div>

        {items.length === 0 ? (
          <Card>
            <CardContent className="text-center py-12">
              <div className="text-text-secondary-light dark:text-text-secondary-dark mb-4">
                <span className="material-icons-outlined text-6xl"> wardrobe </span>
              </div>
              <p className="text-lg mb-4">暂无衣物</p>
              <Button onClick={() => router.push('/wardrobe/add')}>
                <span className="material-icons-outlined mr-2">add</span>
                添加衣物
              </Button>
            </CardContent>
          </Card>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {items.map((item) => (
              <Card key={item.id} hoverable>
                <CardHeader title={item.name} />
                <CardContent>
                  <p className="text-text-secondary-light dark:text-text-secondary-dark">
                    {item.category} · {item.color}
                  </p>
                </CardContent>
              </Card>
            ))}
          </div>
        )}
      </main>
    </div>
  );
}
```
```

### 14.2 创建UI基础组件（基于设计系统）
**提示词**：
```
请创建 src/components/ui/Card.tsx 卡片组件，要求：
1. 严格遵循UI设计系统规范
2. 背景色使用 bg-card-light dark:bg-card-dark
3. 圆角使用 rounded-lg（匹配UI设计令牌）
4. 边框使用 border-border-light dark:border-border-dark
5. 标准内边距 p-6（24px，符合4px网格）
6. 可选的悬停效果：hover:shadow-md transition-shadow
7. 支持 className 属性透传
8. 使用 TypeScript，定义清晰的Props接口
9. 默认样式：shadow-sm，无边框透明背景
```

**实现示例**：
```tsx
import React from 'react';

interface CardProps {
  children: React.ReactNode;
  className?: string;
  hoverable?: boolean;
  bordered?: boolean;
}

export const Card: React.FC<CardProps> = ({
  children,
  className = '',
  hoverable = false,
  bordered = true,
}) => {
  return (
    <div
      className={`
        bg-card-light dark:bg-card-dark
        rounded-lg shadow-sm
        ${bordered ? 'border border-border-light dark:border-border-dark' : ''}
        p-6
        ${hoverable ? 'hover:shadow-md transition-shadow cursor-pointer' : ''}
        ${className}
      `}
    >
      {children}
    </div>
  );
};

// CardHeader 子组件
export const CardHeader: React.FC<{
  title: string;
  subtitle?: string;
  action?: React.ReactNode;
  className?: string;
}> = ({ title, subtitle, action, className = '' }) => (
  <div className={`flex items-center justify-between mb-4 ${className}`}>
    <div>
      <h3 className="text-lg font-medium text-text-primary-light dark:text-text-primary-dark">
        {title}
      </h3>
      {subtitle && (
        <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-1">
          {subtitle}
        </p>
      )}
    </div>
    {action && <div>{action}</div>}
  </div>
);

// CardContent 子组件
export const CardContent: React.FC<{
  children: React.ReactNode;
  className?: string;
}> = ({ children, className = '' }) => (
  <div className={`text-text-primary-light dark:text-text-primary-dark ${className}`}>
    {children}
  </div>
);
```

### 14.3 创建复用组件（适配设计系统）
**提示词**：
```
请创建 src/components/ClothingCard.tsx 衣物卡片组件，要求：
1. 使用UI设计系统的Card组件作为基础
2. 接收ClothingItem对象作为Props（包含id、name、category、color、imageUrl）
3. 展示图片（使用next/image，300x200尺寸）
4. 使用Material Icons图标系统
5. 布局：图片在上，信息在下，遵循4px网格间距
6. 支持深色模式切换
7. 添加hover效果（hoverable Card）
8. 使用功能色标签显示分类
9. 响应式设计：移动端全宽，桌面端网格布局
```

### 14.4 重构现有组件（设计系统适配）
**提示词**：
```
请重构 src/components/AddClothingItem.tsx 组件，要求：
1. 使用新的UI组件库（Card、Input、Button）
2. 表单布局遵循4px网格系统（space-y-4）
3. 完整深色模式支持
4. 颜色选择器适配黑白极简风格（限制颜色选项）
5. 使用功能色Toast系统
6. 表单验证样式统一
7. 加载状态使用LoadingSpinner组件
8. 移除所有内联样式，使用Tailwind类
9. 保持原有业务逻辑不变
```

### 14.5 添加设计系统支持
**提示词**：
```
请帮我更新 tailwind.config.ts 文件，添加UI设计系统的完整配置：
1. 颜色系统：primary、background-light/dark、card-light/dark、text-primary-light/dark等
2. 间距系统：基于4px网格的自定义间距
3. 字体系统：Noto Sans SC作为主要字体
4. 阴影系统：card、modal、button等级别
5. 深色模式配置：class策略
6. 扩展配置：borderRadius、boxShadow、colors完整映射
7. 参考UI项目的 design-tokens.md 文件
```

### 14.6 修复设计不一致问题
**提示词**：
```
我的 /main 页面存在以下设计问题，请修复：
1. 颜色使用不统一（使用了blue、green等非设计系统颜色）
2. 间距没有遵循4px网格系统
3. 卡片样式与UI设计系统不一致
4. 缺少深色模式支持
5. 内联样式需要替换为Tailwind类
6. 需要添加Material Icons图标
7. 布局需要适配响应式设计

请根据UI设计系统规范重构整个页面，保持功能不变。
```
### 14.7 UI到React转换指南

#### HTML → React组件映射关系
```
UI HTML结构          →    React组件结构
----------------    →    -------------------
div.bg-card-light   →    <Card> 组件
button.bg-primary   →    <Button variant="primary">
span.material-icons →    <span className="material-icons-outlined">
.dark: 类名         →    自动深色模式切换
```

#### 样式转换对照表
```css
/* UI项目CSS类 */
.bg-card-light.dark:bg-card-dark → /* React：直接使用 */
.rounded-lg → /* React：直接使用 */
.shadow-sm → /* React：直接使用 */
.p-6 → /* React：直接使用 */

/* UI项目内联样式 */
style="margin-bottom: 16px" → className="mb-4"
style="padding: 12px" → className="p-3"
style="color: #000" → className="text-primary"
```

#### 深色模式转换
```html
<!-- UI项目 -->
<div class="bg-card-light dark:bg-card-dark">
  <p class="text-text-primary-light dark:text-text-primary-dark">
    文字内容
  </p>
</div>

<!-- React项目（完全相同）-->
<div className="bg-card-light dark:bg-card-dark">
  <p className="text-text-primary-light dark:text-text-primary-dark">
    文字内容
  </p>
</div>
```

#### 组件结构转换示例
```tsx
// UI项目HTML结构转换为React组件

// UI项目 (apps/wardrobe/index.html)
/*
<div class="bg-card-light dark:bg-card-dark rounded-lg p-6">
  <h3 class="text-lg font-medium mb-2">标题</h3>
  <p class="text-sm text-gray-600 mb-4">描述</p>
  <button class="bg-primary text-white px-4 py-2 rounded">
    <span class="material-icons-outlined mr-2">add</span>
    添加
  </button>
</div>
*/

// React项目 (src/components/ui/Card.tsx)
export const CardExample = () => (
  <Card>
    <CardHeader title="标题" subtitle="描述" />
    <CardContent>
      <Button>
        <span className="material-icons-outlined mr-2">add</span>
        添加
      </Button>
    </CardContent>
  </Card>
);
```

### 14.8 设计一致性检查清单

在开发过程中，必须验证以下设计一致性要求：

#### ✅ 颜色系统检查
- [ ] 只使用设计系统定义的颜色（primary、background、card、text系列）
- [ ] 避免使用Tailwind默认颜色（blue、green、red等）
- [ ] 深色模式颜色正确切换
- [ ] 功能色使用正确（success/warning/error/info）

#### ✅ 间距系统检查
- [ ] 所有间距基于4px网格系统
- [ ] 使用标准间距类（p-4、m-6、gap-4等）
- [ ] 避免使用任意值（p-[17px]）
- [ ] 组件间距统一

#### ✅ 字体系统检查
- [ ] 使用Noto Sans SC字体家族
- [ ] 字体大小使用标准尺寸（text-sm、text-base等）
- [ ] 字体粗细使用标准值（font-medium、font-semibold等）

#### ✅ 组件一致性检查
- [ ] 使用统一的Card、Button、Input组件
- [ ] 圆角使用标准值（rounded-lg为主）
- [ ] 阴影使用标准级别（shadow-sm、shadow-md等）
- [ ] 边框样式统一

#### ✅ 深色模式检查
- [ ] 所有颜色都有对应的深色模式类
- [ ] 文字颜色适配深色背景
- [ ] 卡片背景色正确切换
- [ ] 边框颜色适配深色模式

#### ✅ 响应式设计检查
- [ ] 移动端优先的布局策略
- [ ] 断点使用标准值（sm、md、lg）
- [ ] 触摸目标大小合适（最小44px）
- [ ] 文字大小适配不同屏幕

#### ✅ 图标系统检查
- [ ] 使用Material Icons Outlined
- [ ] 图标大小统一（text-base、text-lg等）
- [ ] 图标颜色与文字一致
- [ ] 图标与文字对齐正确
### 14.9 性能优化建议（基于设计系统）

#### 组件级优化
```tsx
// 使用React.memo优化纯展示组件
const ClothingCard = React.memo(({ item, onClick }) => (
  <Card hoverable onClick={onClick}>
    <img src={item.imageUrl} alt={item.name} loading="lazy" />
    <CardHeader title={item.name} subtitle={item.category} />
  </Card>
));

// 使用useMemo缓存计算结果
const filteredItems = useMemo(() =>
  items.filter(item => item.category === selectedCategory),
  [items, selectedCategory]
);
```

#### 图片优化
```tsx
// 使用Next.js Image组件优化
import Image from 'next/image';

const OptimizedClothingImage = ({ src, alt }) => (
  <div className="relative aspect-square">
    <Image
      src={src}
      alt={alt}
      fill
      sizes="(max-width: 640px) 100vw, (max-width: 1024px) 50vw, 33vw"
      priority={false}
      className="object-cover rounded-lg"
    />
  </div>
);
```

#### 代码分割
```tsx
// 动态导入大型组件
import dynamic from 'next/dynamic';

const ClothingForm = dynamic(() => import('@/components/ClothingForm'), {
  loading: () => <LoadingSpinner />,
  ssr: false
});

const WardrobeAnalytics = dynamic(() =>
  import('@/components/WardrobeAnalytics'), {
  loading: () => <LoadingSpinner />,
}
);
```

---

## 15. 开发工具与环境

### 15.1 推荐的开发工具
- **编辑器**: VS Code + 以下扩展
  - Tailwind CSS IntelliSense
  - Material Icon Theme
  - TypeScript Hero
  - Prettier
- **浏览器**: Chrome DevTools + React Developer Tools
- **设计工具**: 直接参考UI项目的HTML实现

### 15.2 有用的VS Code扩展
- **Tailwind CSS IntelliSense**: 智能提示Tailwind类名
- **Material Icon Theme**: Material Icons图标提示
- **Thunder Client**: API测试（替代Postman）
- **GitLens**: Git历史查看

### 15.3 开发环境配置
```json
// .vscode/settings.json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "tailwindCSS.includeLanguages": {
    "typescript": "javascript",
    "typescriptreact": "javascript"
  },
  "tailwindCSS.experimental.classRegex": [
    ["clsx\\(([^)]*)\\)", "(?:'|\"|`)([^']*)(?:'|\"|`)]"
  ]
}
```

### 15.4 调试技巧
1. **深色模式测试**: 在Chrome DevTools中切换 prefers-color-scheme
2. **响应式测试**: 使用Device Mode测试不同屏幕尺寸
3. **性能分析**: 使用React DevTools Profiler分析组件渲染
4. **网络调试**: 使用Network面板检查API请求

---

## 16. 部署与最佳实践

### 16.1 构建检查清单
在部署前，确保完成以下检查：

- [ ] `pnpm build` 成功通过
- [ ] `pnpm lint` 无错误
- [ ] `pnpm tsc --noEmit` 无类型错误
- [ ] 所有环境变量已配置
- [ ] 设计一致性检查通过

### 16.2 设计系统一致性验证
```bash
# 检查Tailwind配置是否包含所有设计令牌
pnpm run validate-design-system

# 检查未使用的样式类
pnpm run analyze-css
```

### 16.3 性能监控指标
- **First Contentful Paint (FCP)**: < 1.8s
- **Largest Contentful Paint (LCP)**: < 2.5s
- **First Input Delay (FID)**: < 100ms
- **Cumulative Layout Shift (CLS)**: < 0.1

---

## 附录：快速参考

### A1. 设计令牌速查表
```tsx
// 颜色速查
primary: '#000000'                    // 主色调
background-light: '#F8F9FA'           // 浅色背景
background-dark: '#121212'            // 深色背景
card-light: '#FFFFFF'                 // 浅色卡片
card-dark: '#1E1E1E'                  // 深色卡片
text-primary-light: '#1F2937'         // 浅色主文字
text-primary-dark: '#F9FAFB'          // 深色主文字

// 间距速查（4px网格）
p-1 (4px), p-2 (8px), p-4 (16px), p-6 (24px)
m-1 (4px), m-2 (8px), m-4 (16px), m-6 (24px)
gap-1 (4px), gap-2 (8px), gap-4 (16px)
```

### A2. UI到React快速映射
```
UI HTML                          React组件
<div class="bg-card-light"/>     <Card></Card>
<button class="bg-primary"/>   <Button variant="primary"></Button>
<input class="border"/>        <Input />
<span class="material-icons"/> <span className="material-icons-outlined">
```

### A3. 常见问题快速解决
**Q: 深色模式不生效？**
A: 检查是否添加了 `dark:` 前缀，以及是否在 `html` 标签上有 `dark` 类

**Q: 样式与设计系统不一致？**
A: 检查是否使用了正确的Tailwind配置，避免使用默认颜色

**Q: 组件间距不符合4px网格？**
A: 使用标准间距类（p-4、m-6等），避免任意值

**Q: Material Icons不显示？**
A: 确保在 `app/layout.tsx` 中引入了Material Icons CSS

---

**文档版本**：v2.0（基于UI设计系统重构）
**最后更新**：2025-01-12
**维护者**：StyleSense 开发团队
**设计系统来源**：StyleSense UI 项目

---

*本文档融合了StyleSense UI项目的黑白极简主义设计规范与Next.js/React现代前端架构，确保视觉一致性与技术先进性的完美统一。*
```

### 14.3 添加 API 接口
**提示词**：
```
请在 src/lib/api/clothing.ts 添加 deleteClothing 接口，要求：
1. 使用 http.delete 方法
2. 接收参数：itemId (string)
3. 请求路径：/clothing/{itemId}
4. 返回类型：Promise<void>
5. 在 src/types/clothing.ts 中无需新增类型（使用 void）
```

### 14.4 修复认证问题
**提示词**：
```
我的 /wardrobe 页面没有检查 Token，请帮我：
1. 在 useEffect 中添加 Token 检查逻辑
2. 无 Token 时跳转到 /login
3. 使用 useRouter 进行跳转
4. 确保 useEffect 依赖项包含 router
```

### 14.5 优化表单组件
**提示词**：
```
请优化 src/components/EditClothingForm.tsx 组件：
1. 添加表单验证（必填字段：name、category）
2. 提交时显示加载状态，按钮文案改为"保存中..."
3. 成功后展示 Toast 并关闭表单
4. 失败时展示错误 Toast
5. 参考 AddClothingItem.tsx 的模式
```

---

## 附录：快速参考

### A1. 常用导入语句
```tsx
// Next.js
import { useRouter } from 'next/navigation';
import Image from 'next/image';
import dynamic from 'next/dynamic';

// React
import { useState, useEffect, useCallback, useMemo } from 'react';

// 项目内部
import { useToast } from '@/components/ToastProvider';
import Modal from '@/components/Modal';
import { someAPI } from '@/lib/api/some';
import type { SomeType } from '@/types/some';
```

### A2. 常见问题排查
**问题：页面无法跳转**
- 检查是否使用了 `'use client'` 标记
- 确认使用 `next/navigation` 而非 `next/router`

**问题：API 调用 401 错误**
- 检查 `localStorage` 中是否存在 Token
- 确认 Token 未过期
- 检查 `http.ts` 拦截器是否正常工作

**问题：Toast 不显示**
- 确认根布局包裹了 `ToastProvider`
- 检查是否正确调用 `showToast` 函数

**问题：样式不生效**
- 检查 Tailwind 类名是否正确
- 确认模块化 CSS 已正确导入
- 检查是否被其他样式覆盖

### A3. 环境配置
- **API 地址**：在 `.env.local` 中设置 `NEXT_PUBLIC_API_BASE_URL`
- **开发服务器**：`pnpm dev`，访问 `http://localhost:3000`
- **生产构建**：`pnpm build` + `pnpm start`

---

**文档版本**：v1.0  
**最后更新**：2025-01-12  
**维护者**：StyleSense 开发团队
