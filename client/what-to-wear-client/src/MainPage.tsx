import React, { useState, useEffect } from 'react';
import './styles/modern.css';

interface MainPageProps {
  onLogout: () => void;
}

export function MainPage({ onLogout }: MainPageProps) {
  const [username, setUsername] = useState('');
  const [userProfile, setUserProfile] = useState<any>(null);
  const [weather, setWeather] = useState<any>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const storedUsername = localStorage.getItem('username');
    if (storedUsername) {
      setUsername(storedUsername);
    }

    // 获取用户资料和天气信息
    fetchUserData();
  }, []);

  const fetchUserData = async () => {
    const token = localStorage.getItem('token');
    if (!token) {
      onLogout();
      return;
    }

    try {
      // 获取用户资料
      const profileResponse = await fetch('http://localhost:8080/api/user/profile', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (profileResponse.ok) {
        const profileData = await profileResponse.json();
        setUserProfile(profileData);
      }

      // 获取天气信息
      const weatherResponse = await fetch('http://localhost:8080/api/weather/current', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
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
    onLogout();
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
    <div className="main-container">
      {/* 顶部导航栏 */}
      <header className="main-header">
        <div className="header-content">
          <div className="header-left">
            <div className="app-logo">
              <span className="logo-icon">👗</span>
              <h1 className="app-name">What to Wear</h1>
            </div>
            <div className="user-greeting">
              <span>Hi, {username}</span>
            </div>
          </div>
          <div className="header-right">
            <button className="header-btn notification-btn">
              <span>🔔</span>
            </button>
            <button className="header-btn profile-btn">
              <span>👤</span>
            </button>
            <button onClick={handleLogout} className="logout-btn">
              退出
            </button>
          </div>
        </div>
      </header>

      <main className="main-content">
        <div className="content-wrapper">
          {/* 今日概览 */}
          <section className="today-overview">
            <div className="overview-header">
              <h2>今日概览</h2>
              <span className="date">{new Date().toLocaleDateString('zh-CN', {
                year: 'numeric',
                month: 'long',
                day: 'numeric',
                weekday: 'long'
              })}</span>
            </div>

            {/* 天气卡片 */}
            <div className="weather-card">
              <div className="weather-header">
                <h3>天气状况</h3>
                <span className="weather-icon">🌤️</span>
              </div>
              {weather ? (
                <div className="weather-content">
                  <div className="weather-main">
                    <span className="temperature">{weather.temperature}°</span>
                    <div className="weather-details">
                      <span className="condition">{weather.condition}</span>
                      <span className="humidity">湿度 {weather.humidity}%</span>
                    </div>
                  </div>
                </div>
              ) : (
                <div className="weather-placeholder">
                  <span>暂无天气数据</span>
                </div>
              )}
            </div>
          </section>

          {/* 穿搭建议 */}
          <section className="outfit-suggestion">
            <div className="suggestion-header">
              <h3>今日穿搭建议</h3>
              <button className="refresh-btn">🔄</button>
            </div>
            <div className="suggestion-content">
              <div className="suggestion-text">
                <p>根据今日天气，为您推荐以下搭配：</p>
                <ul className="outfit-list">
                  <li>轻薄长袖衬衫</li>
                  <li>休闲长裤</li>
                  <li>舒适运动鞋</li>
                </ul>
                <p className="tip">💡 建议携带薄外套备用</p>
              </div>
            </div>
          </section>

          {/* 快捷功能 */}
          <section className="quick-actions">
            <h3>快捷功能</h3>
            <div className="actions-grid">
              <button className="action-card primary">
                <span className="action-icon">📸</span>
                <span className="action-title">记录穿搭</span>
                <span className="action-desc">拍照记录今日搭配</span>
              </button>

              <button className="action-card">
                <span className="action-icon">👗</span>
                <span className="action-title">我的衣橱</span>
                <span className="action-desc">管理服装单品</span>
              </button>

              <button className="action-card">
                <span className="action-icon">🎨</span>
                <span className="action-title">风格推荐</span>
                <span className="action-desc">发现新风格</span>
              </button>

              <button className="action-card">
                <span className="action-icon">📊</span>
                <span className="action-title">穿搭统计</span>
                <span className="action-desc">查看数据分析</span>
              </button>
            </div>
          </section>

          {/* 最近活动 */}
          <section className="recent-activity">
            <div className="section-header">
              <h3>最近活动</h3>
              <button className="view-all-btn">查看全部</button>
            </div>
            <div className="activity-list">
              <div className="activity-item">
                <div className="activity-icon">📸</div>
                <div className="activity-content">
                  <span className="activity-title">记录了今日穿搭</span>
                  <span className="activity-time">2小时前</span>
                </div>
              </div>

              <div className="activity-item">
                <div className="activity-icon">👗</div>
                <div className="activity-content">
                  <span className="activity-title">添加了新的衣服</span>
                  <span className="activity-time">昨天</span>
                </div>
              </div>

              <div className="activity-item">
                <div className="activity-icon">🎨</div>
                <div className="activity-content">
                  <span className="activity-title">收藏了穿搭灵感</span>
                  <span className="activity-time">3天前</span>
                </div>
              </div>
            </div>
          </section>

          {/* 个人信息卡片 */}
          {userProfile && (
            <section className="profile-summary">
              <h3>个人信息</h3>
              <div className="profile-card">
                <div className="profile-avatar">
                  <span>👤</span>
                </div>
                <div className="profile-info">
                  <h4>{userProfile.username}</h4>
                  <span className="profile-id">ID: {userProfile.user_id}</span>
                </div>
                <button className="edit-profile-btn">编辑</button>
              </div>
            </section>
          )}
        </div>
      </main>
    </div>
  );
}
