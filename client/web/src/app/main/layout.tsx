"use client";
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import "@/styles/sidebar-layout.css";

export default function MainLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();

  const isActive = (path: string) => {
    return pathname.startsWith(path);
  };

  return (
    <div className="app-layout">
      {/* 侧边导航栏 */}
      <aside className="sidebar">
        <div className="sidebar-header">
          <div className="app-logo">
            <div className="logo-icon">👗</div>
            <div className="app-name">What to Wear</div>
          </div>
          
          <div className="user-profile">
            <div className="user-avatar">👤</div>
            <div className="user-info">
              <h4>用户名</h4>
              <span>在线</span>
            </div>
          </div>
        </div>

        <nav className="sidebar-nav">
          <div className="nav-section">
            <div className="nav-section-title">主要功能</div>
            <Link href="/main" className={`nav-item ${pathname === '/main' ? 'active' : ''}`}>
              <span className="nav-icon">🏠</span>
              <span className="nav-text">今日概览</span>
            </Link>
            <Link href="/main/record" className={`nav-item ${isActive('/main/record') ? 'active' : ''}`}>
              <span className="nav-icon">📸</span>
              <span className="nav-text">记录穿搭</span>
              <span className="nav-badge">3</span>
            </Link>
            <Link href="/main/wardrobe" className={`nav-item ${isActive('/main/wardrobe') ? 'active' : ''}`}>
              <span className="nav-icon">👗</span>
              <span className="nav-text">我的衣橱</span>
            </Link>
            <Link href="/main/style" className={`nav-item ${isActive('/main/style') ? 'active' : ''}`}>
              <span className="nav-icon">🎨</span>
              <span className="nav-text">风格推荐</span>
            </Link>
          </div>

          <div className="nav-section">
            <div className="nav-section-title">数据分析</div>
            <Link href="/main/stats" className={`nav-item ${isActive('/main/stats') ? 'active' : ''}`}>
              <span className="nav-icon">📊</span>
              <span className="nav-text">穿搭统计</span>
            </Link>
            <Link href="/main/trends" className={`nav-item ${isActive('/main/trends') ? 'active' : ''}`}>
              <span className="nav-icon">📈</span>
              <span className="nav-text">趋势分析</span>
            </Link>
            <Link href="/main/inspiration" className={`nav-item ${isActive('/main/inspiration') ? 'active' : ''}`}>
              <span className="nav-icon">🌟</span>
              <span className="nav-text">搭配灵感</span>
            </Link>
          </div>

          <div className="nav-section">
            <div className="nav-section-title">设置</div>
            <Link href="/main/settings" className={`nav-item ${isActive('/main/settings') ? 'active' : ''}`}>
              <span className="nav-icon">⚙️</span>
              <span className="nav-text">个人设置</span>
            </Link>
            <Link href="/main/notifications" className={`nav-item ${isActive('/main/notifications') ? 'active' : ''}`}>
              <span className="nav-icon">🔔</span>
              <span className="nav-text">通知设置</span>
            </Link>
          </div>
        </nav>

        <div className="sidebar-footer">
          <button className="sidebar-logout-btn">
            <span>🚪</span>
            <span>退出登录</span>
          </button>
        </div>
      </aside>

      {/* 主内容区域 */}
      <main className="main-content">
        {children}
      </main>
    </div>
  );
}
