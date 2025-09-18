"use client";
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { Modal } from '@/components/Modal';
import { AddClothingItem } from '@/components/AddClothingItem';
import { getUserProfile } from '@/lib/api/user';
import { getCurrentWeather } from '@/lib/api/weather';
import { createClothingItem } from '@/lib/api/clothing';
import type { UserProfile } from '@/types/user';
import type { Weather } from '@/types/weather';
import type { ClothingItemData } from '@/types/clothing';

export default function MainPage() {
  const router = useRouter();
  const [username, setUsername] = useState('');
  const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
  const [weather, setWeather] = useState<Weather | null>(null);
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
      const [profileData, weatherData] = await Promise.all([
        getUserProfile().catch(() => null),
        getCurrentWeather().catch(() => null)
      ]);
      
      if (profileData) setUserProfile(profileData);
      if (weatherData) setWeather(weatherData);
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

  const handleAddClothingSubmit = async (data: ClothingItemData) => {
    const token = localStorage.getItem('token');
    if (!token) { router.replace('/login'); return; }
    if (token === 'demo-token') { setIsAddClothingModalOpen(false); return; }
    
    try {
      await createClothingItem(data);
      setIsAddClothingModalOpen(false);
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
    <>
      {/* 主页面内容，现在由 layout.tsx 提供导航栏 */}
      <main>
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
                  <p style={{ marginBottom: '16px', color: '#64748b' }}>根据{weather?.temperature || '25'}°天气，为您推荐以下搭配：</p>
                  <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: '12px', marginBottom: '16px' }}>
                    <div style={{ padding: '16px', background: '#f8fafc', borderRadius: '12px', textAlign: 'center', border: '1px solid #e2e8f0' }}>
                      <div style={{ fontSize: '24px', marginBottom: '8px' }}>👕</div>
                      <div style={{ fontSize: '14px', fontWeight: '600', marginBottom: '4px' }}>轻薄长袖衬衫</div>
                      <div style={{ fontSize: '12px', color: '#64748b' }}>透气舒适</div>
                    </div>
                    <div style={{ padding: '16px', background: '#f8fafc', borderRadius: '12px', textAlign: 'center', border: '1px solid #e2e8f0' }}>
                      <div style={{ fontSize: '24px', marginBottom: '8px' }}>👖</div>
                      <div style={{ fontSize: '14px', fontWeight: '600', marginBottom: '4px' }}>休闲长裤</div>
                      <div style={{ fontSize: '12px', color: '#64748b' }}>百搭款式</div>
                    </div>
                    <div style={{ padding: '16px', background: '#f8fafc', borderRadius: '12px', textAlign: 'center', border: '1px solid #e2e8f0' }}>
                      <div style={{ fontSize: '24px', marginBottom: '8px' }}>👟</div>
                      <div style={{ fontSize: '14px', fontWeight: '600', marginBottom: '4px' }}>舒适运动鞋</div>
                      <div style={{ fontSize: '12px', color: '#64748b' }}>全天候</div>
                    </div>
                  </div>
                  <div style={{ padding: '12px', background: '#eef2ff', borderRadius: '8px', border: '1px solid #c7d2fe' }}>
                    <p style={{ margin: 0, fontSize: '14px', color: '#4f46e5' }}>💡 建议携带薄外套备用，下午可能转凉</p>
                  </div>
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
                    <button className="action-btn">
                      <span className="action-icon">🌤️</span>
                      <span className="action-title">天气提醒</span>
                      <span className="action-desc">设置穿搭提醒</span>
                    </button>
                    <Link href="/main/wardrobe" className="action-btn">
                      <span className="action-icon">👗</span>
                      <span className="action-title">我的衣橱</span>
                      <span className="action-desc">管理衣橱物品</span>
                    </Link>
                    <button className="action-btn">
                      <span className="action-icon">🎯</span>
                      <span className="action-title">今日任务</span>
                      <span className="action-desc">完成穿搭挑战</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <div className="side-column">
              <div className="card">
                <div className="card-header">
                  <h3 className="card-title">今日概况</h3>
                </div>
                <div className="card-body">
                  <div style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
                    <div style={{ padding: '12px', background: '#f0f9ff', borderRadius: '8px', border: '1px solid #bae6fd' }}>
                      <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '4px' }}>
                        <span style={{ fontSize: '16px' }}>☀️</span>
                        <span style={{ fontSize: '14px', fontWeight: '600', color: '#0369a1' }}>今日适宜外出</span>
                      </div>
                      <div style={{ fontSize: '12px', color: '#0c4a6e' }}>紫外线指数适中，建议防晒</div>
                    </div>
                    
                    <div style={{ display: 'flex', flexDirection: 'column', gap: '12px' }}>
                      <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                        <div style={{ width: '32px', height: '32px', background: '#dcfce7', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>✅</div>
                        <div>
                          <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>已完成今日穿搭</div>
                          <div style={{ fontSize: '12px', color: '#64748b' }}>休闲舒适风格</div>
                        </div>
                      </div>
                      <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                        <div style={{ width: '32px', height: '32px', background: '#fef3c7', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>📷</div>
                        <div>
                          <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>待拍照记录</div>
                          <div style={{ fontSize: '12px', color: '#64748b' }}>记录今日搭配效果</div>
                        </div>
                      </div>
                      <div style={{ display: 'flex', gap: '12px', alignItems: 'center' }}>
                        <div style={{ width: '32px', height: '32px', background: '#e0e7ff', borderRadius: '8px', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>🎯</div>
                        <div>
                          <div style={{ fontSize: '14px', fontWeight: '500', marginBottom: '2px' }}>穿搭挑战</div>
                          <div style={{ fontSize: '12px', color: '#64748b' }}>尝试新的颜色搭配</div>
                        </div>
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
    </>
  );
}
