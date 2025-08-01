# 🎨 What to Wear - UI 设计预览

这个目录包含了项目的所有UI设计预览文件，用于设计开发和团队协作。

## 📁 目录结构

```
UI/
├── README.md                    # 本文件
├── index.html                   # UI预览导航页面
├── pages/                       # 页面级UI预览
│   ├── main-page-top-nav.html   # 主页面 - 顶部导航版本
│   ├── main-page-sidebar.html   # 主页面 - 侧边导航版本 (推荐)
│   ├── login-page.html          # 登录页面
│   └── register-page.html       # 注册页面
├── components/                  # 组件级UI预览
│   ├── navigation.html          # 导航组件
│   ├── weather-card.html        # 天气卡片
│   ├── outfit-suggestion.html   # 穿搭建议
│   └── quick-actions.html       # 快捷操作
├── assets/                      # UI资源文件
│   ├── styles/                  # 样式文件
│   │   ├── modern.css           # 现代化样式
│   │   ├── components.css       # 组件样式
│   │   └── variables.css        # CSS变量
│   ├── images/                  # 图片资源
│   └── icons/                   # 图标资源
└── design-tokens/               # 设计令牌
    ├── colors.json              # 颜色系统
    ├── typography.json          # 字体系统
    └── spacing.json             # 间距系统
```

## 🎯 使用方式

### 1. 快速预览
打开 `index.html` 查看所有UI预览的导航页面

### 2. 页面预览
- **主页面 (推荐)**: `pages/main-page-sidebar.html`
- **主页面 (备选)**: `pages/main-page-top-nav.html`
- **登录页面**: `pages/login-page.html`
- **注册页面**: `pages/register-page.html`

### 3. 组件预览
在 `components/` 目录下查看独立的组件设计

## 🎨 设计系统

### 配色方案
- **主色**: #667eea (紫蓝色)
- **辅助色**: #764ba2 (深紫色)
- **中性色**: #1e293b, #64748b, #f8fafc
- **功能色**: 成功#10b981, 警告#f59e0b, 错误#ef4444

### 字体系统
- **主字体**: Inter (现代无衬线字体)
- **标题**: 700/600 字重
- **正文**: 400/500 字重
- **辅助**: 300 字重

### 间距系统
- **基础单位**: 4px
- **常用间距**: 8px, 12px, 16px, 20px, 24px, 32px
- **大间距**: 48px, 64px

## 🚀 开发流程

### 1. 设计阶段
1. 在UI目录中创建HTML预览
2. 使用设计令牌确保一致性
3. 团队review和确认设计

### 2. 开发阶段
1. 参考UI预览实现React组件
2. 复用CSS样式和设计令牌
3. 确保与预览效果一致

### 3. 测试阶段
1. 对比预览效果和实际效果
2. 调整和优化
3. 更新预览文件

## 📝 贡献指南

### 添加新的UI预览
1. 在对应目录创建HTML文件
2. 使用统一的样式系统
3. 更新导航页面链接
4. 添加响应式设计

### 修改现有预览
1. 保持向后兼容
2. 更新相关文档
3. 通知团队成员

## 🔗 相关链接

- [项目主目录](../README.md)
- [前端代码](../client/what-to-wear-client/)
- [后端代码](../server/)
- [设计文档](../docs/design.md)

## 📱 响应式支持

所有UI预览都支持以下断点：
- **桌面端**: 1024px+
- **平板端**: 768px - 1023px  
- **手机端**: < 768px

## 🎭 主题支持

目前支持：
- ✅ 浅色主题 (默认)
- 🚧 深色主题 (计划中)
- 🚧 高对比度主题 (计划中)

---

**最后更新**: 2024年8月1日  
**维护者**: What to Wear 开发团队
