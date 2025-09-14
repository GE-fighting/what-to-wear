"use client";
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { Modal } from '@/components/Modal';
import { AddClothingItem } from '@/components/AddClothingItem';
import '@/styles/sidebar-layout.css';

export default function MainPage() {
  const router = useRouter();
  const [username, setUsername] = useState('');
  const [userProfile, setUserProfile] = useState<any>(null);
  const [weather, setWeather] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [activeNav, setActiveNav] = useState('overview');
  const [isAddClothingModalOpen, setIsAddClothingModalOpen] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    const storedUsername = localStorage.getItem('username');
    if (!token) {
      router.replace('/login');
      return;
    }
    if (storedUsername) setUsername(storedUsername);

    fetchUserData(token);
  }, [router]);

  const fetchUserData = async (token: string) => {
    if (token === 'demo-token') {
      setUserProfile({ username: '演示用户', user_id: 'demo-001' });
      setWeather({ temperature: 25, condition: '晴朗', humidity: 60 });
      setLoading(false);
      return;
    }

    try {
      const profileResponse = await fetch('http://localhost:8080/api/user/profile', {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (profileResponse.ok) {
        const profileData = await profileResponse.json();
        setUserProfile(profileData);
      }

      const weatherResponse = await fetch('http://localhost:8080/api/weather/current', {
        headers: { 'Authorization': `Bearer ${token}` }
      });
      if (weatherResponse.ok) {
        const weatherData = await weatherResponse.json();
        setWeather(weatherData);
      }
    } catch (error) {
      console.error('Failed to fetch user data:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('username');
    router.replace('/login');
  };

  const handleNavClick = (navId: string) => setActiveNav(navId);
  const handleAddClothingClick = () => setIsAddClothingModalOpen(true);
  const handleCloseAddClothingModal = () => setIsAddClothingModalOpen(false);

  const handleAddClothingSubmit = async (data: any) => {
    const token = localStorage.getItem('token');
    if (!token) { router.replace('/login'); return; }
    if (token === 'demo-token') { setIsAddClothingModalOpen(false); return; }
    try {
      const response = await fetch('http://localhost:8080/api/clothing-items', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(data)
      });
      if (response.ok) {
        setIsAddClothingModalOpen(false);
      } else {
        console.error('添加衣物失败');
      }
    } catch (error) {
      console.error('添加衣物时发生错误:', error);
    }
  };

  if (loading) {
    return (
      <div className="main-loading">
        <div className="loading-spinner"></div>
        <p>加载中...</p>
      </div>
    );
  }

  return (
    <div className="app-layout">
      <aside className="sidebar">
        <div className="sidebar-header">
          <div className="app-logo">
            <div className="logo-icon">👗</div>
            <div className="app-name">What to Wear</div>
          </div>
          
          <div className="user-profile">
            <div className="user-avatar">👤</div>
            <div className="user-info">
              <h4>{username || '用户名'}</h4>
              <span>在线</span>
            </div>
          </div>
        </div>

        <nav className="sidebar-nav">
          <div className="nav-section">
            <div className="nav-section-title">主要功能</div>
            <div className={`nav-item ${activeNav === 'overview' ? 'active' : ''}`} onClick={() => handleNavClick('overview')}>
              <span className="nav-icon">🏠</span>
              <span className="nav-text">今日概览</span>
            </div>
            <div className={`nav-item ${activeNav === 'record' ? 'active' : ''}`} onClick={() => handleNavClick('record')}>
              <span className="nav-icon">📸</span>
              <span className="nav-text">记录穿搭</span>
              <span className="nav-badge">3</span>
            </div>
            <div className={`nav-item ${activeNav === 'wardrobe' ? 'active' : ''}`} onClick={() => handleNavClick('wardrobe')}>
              <span className="nav-icon">👗</span>
              <span className="nav-text">我的衣橱</span>
            </div>
            <div className={`nav-item ${activeNav === 'style' ? 'active' : ''}`} onClick={() => handleNavClick('style')}>
              <span className="nav-icon">🎨</span>
              <span className="nav-text">风格推荐</span>
            </div>
          </div>

          <div className="nav-section">
            <div className="nav-section-title">数据分析</div>
            <div className={`nav-item ${activeNav === 'stats' ? 'active' : ''}`} onClick={() => handleNavClick('stats')}>
              <span className="nav-icon">📊</span>
              <span className="nav-text">穿搭统计</span>
            </div>
            <div className={`nav-item ${activeNav === 'trends' ? 'active' : ''}`} onClick={() => handleNavClick('trends')}>
              <span className="nav-icon">📈</span>
              <span className="nav-text">趋势分析</span>
            </div>
            <div className={`nav-item ${activeNav === 'inspiration' ? 'active' : ''}`} onClick={() => handleNavClick('inspiration')}>
              <span className="nav-icon">🌟</span>
              <span className="nav-text">搭配灵感</span>
            </div>
          </div>

          <div className="nav-section">
            <div className="nav-section-title">设置</div>
            <div className={`nav-item ${activeNav === 'settings' ? 'active' : ''}`} onClick={() => handleNavClick('settings')}>
              <span className="nav-icon">⚙️</span>
              <span className="nav-text">个人设置</span>
            </div>
            <div className={`nav-item ${activeNav === 'notifications' ? 'active' : ''}`} onClick={() => handleNavClick('notifications')}>
              <span className="nav-icon">🔔</span>
              <span className="nav-text">通知设置</span>
            </div>
          </div>
        </nav>

        <div className="sidebar-footer">
          <button className="sidebar-logout-btn" onClick={handleLogout}>
            <span>🚪</span>
            <span>退出登录</span>
          </button>
        </div>
      </aside>

      <main className="main-content">
        <header className="content-header">
          <div className="header-top">
            <div>
              <h1 className="page-title">今日概览</h1>
              <p className="page-subtitle">{new Date().toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric', weekday: 'long' })} · 北京</p>
            </div>
            <div className="header-actions">
              <button className="header-btn">🔄 刷新</button>
              <button className="header-btn" onClick={handleAddClothingClick}>👗 添加衣物</button>
              <button className="header-btn primary">📸 记录穿搭</button>
            </div>
          </div>
        </header>

        <div className="content-body">
          <div className="content-grid">
            <div className="main-column">
              <div className="weather-card">
                <div className="weather-main">
                  <div className="temperature">{weather?.temperature || '25'}°</div>
                  <div className="weather-details">
                    <h3>{weather?.condition || '晴朗'}</h3>
                    <p>今天是个好天气，适合外出</p>
                  </div>
                </div>
                <div className="weather-stats">
                  <div className="weather-stat">
                    <div className="weather-stat-value">{weather?.humidity || '60'}%</div>
                    <div className="weather-stat-label">湿度</div>
                  </div>
                  <div className="weather-stat">
                    <div className="weather-stat-value">5km/h</div>
                    <div className="weather-stat-label">风速</div>
                  </div>
                  <div className="weather-stat">
                    <div className="weather-stat-value">良好</div>
                    <div className="weather-stat-label">空气质量</div>
                  </div>
                </div>
              </div>

              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">今日穿搭建议</h3>
                  <button className="header-btn">🔄</button>
                </div>
                <div className="card-body">
                  <p style={{ marginBottom: '16px', color: '#64748b' }}>根据今日天气，为您推荐以下搭配：</p>
                  <ul style={{ listStyle: 'none', padding: 0 }}>
                    <li style={{ padding: '8px 0', borderLeft: '3px solid #667eea', paddingLeft: '12px', marginBottom: '8px' }}>轻薄长袖衬衫</li>
                    <li style={{ padding: '8px 0', borderLeft: '3px solid #667eea', paddingLeft: '12px', marginBottom: '8px' }}>休闲长裤</li>
                    <li style={{ padding: '8px 0', borderLeft: '3px solid #667eea', paddingLeft: '12px', marginBottom: '8px' }}>舒适运动鞋</li>
                  </ul>
                  <p style={{ marginTop: '16px', fontSize: '14px', color: '#64748b' }}>💡 建议携带薄外套备用</p>
                </div>
              </div>

              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">快捷操作</h3>
                </div>
                <div className="card-body">
                  <div className="quick-actions">
                    <button className="action-btn primary">
                      <span className="action-icon">📸</span>
                      <span className="action-title">记录穿搭</span>
                      <span className="action-desc">拍照记录今日搭配</span>
                    </button>
                    <button className="action-btn" onClick={handleAddClothingClick}>
                      <span className="action-icon">👗</span>
                      <span className="action-title">添加衣物</span>
                      <span className="action-desc">管理衣橱物品</span>
                    </button>
                    <button className="action-btn">
                      <span className="action-icon">🎨</span>
                      <span className="action-title">风格推荐</span>
                      <span className="action-desc">发现新搭配</span>
                    </button>
                    <button className="action-btn">
                      <span className="action-icon">📊</span>
                      <span className="action-title">查看统计</span>
                      <span className="action-desc">分析穿搭数据</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div className="side-column">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">最近活动</h3>
                  <button className="header-btn">查看全部</button>
                </div>
                <div className="card-body">
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
                    <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                      <div style={{ width: '32px', height: '32px', background: '#f1f5f9', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>📸</div>
                      <div>
                        <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>记录了今日穿搭</div>
                        <div style={{ fontSize: '12px', color: '#64748b' }}>2小时前</div>
                      </div>
                    </div>
                    <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                      <div style={{ width: '32px', height: '32px', background: '#f1f5f9', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>👗</div>
                      <div>
                        <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>添加了新衣服</div>
                        <div style={{ fontSize: '12px', color: '#64748b' }}>昨天</div>
                      </div>
                    </div>
                    <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                      <div style={{ width: '32px', height: '32px', background: '#f1f5f9', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>🎨</div>
                      <div>
                        <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>收藏了搭配灵感</div>
                        <div style={{ fontSize: '12px', color: '#64748b' }}>3天前</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              {userProfile && (
                <div className="card">
                  <div className="card-header">
                    <h3 className="card-title">个人信息</h3>
                  </div>
                  <div className="card-body">
                    <div style={{ display: 'flex', alignItems: 'center', gap: '16px', padding: '16px', background: '#f8fafc', borderRadius: '12px' }}>
                      <div style={{ width: '48px', height: '48px', background: '#e2e8f0', borderRadius: '50%', display: 'flex', alignItems: 'center', justifyContent: 'center', fontSize: '20px' }}>
                        👤
                      </div>
                      <div style={{ flex: 1, display: 'flex', flexDirection: 'column', gap: '4px' }}>
                        <h4 style={{ fontSize: '16px', fontWeight: '600', color: '#1e293b', margin: 0 }}>{userProfile.username}</h4>
                        <span style={{ fontSize: '12px', color: '#64748b' }}>ID: {userProfile.user_id}</span>
                      </div>
                      <button style={{ background: 'none', border: '1px solid #e2e8f0', color: '#1e293b', padding: '8px 16px', borderRadius: '8px', fontSize: '14px', fontWeight: '500', cursor: 'pointer' }}>编辑</button>
                    </div>
                  </div>
                </div>
              )}
            </div>
          </div>
        </div>
      </main>

      <Modal isOpen={isAddClothingModalOpen} onClose={handleCloseAddClothingModal} title="添加新衣物" size="large">
        <AddClothingItem onSubmit={handleAddClothingSubmit} onCancel={handleCloseAddClothingModal} />
      </Modal>
    </div>
  );
}
