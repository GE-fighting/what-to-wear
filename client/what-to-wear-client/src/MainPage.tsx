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
      <div className="loading">
        加载中
      </div>
    );
  }

  return (
    <div className="main-container">
      {/* 头部 */}
      <div className="main-header">
        <div className="header-content">
          <div className="header-info">
            <h1>今天穿什么</h1>
            <p>欢迎回来，{username}！✨</p>
          </div>
          <button
            onClick={handleLogout}
            className="logout-btn"
          >
            退出登录
          </button>
        </div>
      </div>

      <div style={{ maxWidth: '1200px', margin: '0 auto', padding: '32px 24px' }}>

        {/* 天气信息 */}
        <div className="card" style={{ marginBottom: '24px' }}>
          <h2 className="card-title">
            🌤️ 今日天气
          </h2>
          {weather ? (
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(120px, 1fr))', gap: '16px' }}>
              <div style={{ textAlign: 'center', padding: '16px', background: 'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)', borderRadius: '12px', color: 'white' }}>
                <div style={{ fontSize: '24px', fontWeight: '700' }}>{weather.temperature}°C</div>
                <div style={{ fontSize: '14px', opacity: 0.9 }}>温度</div>
              </div>
              <div style={{ textAlign: 'center', padding: '16px', background: 'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)', borderRadius: '12px', color: 'white' }}>
                <div style={{ fontSize: '18px', fontWeight: '600' }}>{weather.condition}</div>
                <div style={{ fontSize: '14px', opacity: 0.9 }}>天气</div>
              </div>
              <div style={{ textAlign: 'center', padding: '16px', background: 'linear-gradient(135deg, #fa709a 0%, #fee140 100%)', borderRadius: '12px', color: 'white' }}>
                <div style={{ fontSize: '24px', fontWeight: '700' }}>{weather.humidity}%</div>
                <div style={{ fontSize: '14px', opacity: 0.9 }}>湿度</div>
              </div>
            </div>
          ) : (
            <p style={{ color: '#666', textAlign: 'center', padding: '20px' }}>暂无天气信息</p>
          )}
        </div>

        {/* 穿衣建议 */}
        <div className="card" style={{ marginBottom: '24px' }}>
          <h2 className="card-title">
            ✨ 今日穿搭建议
          </h2>
          <div style={{
            padding: '20px',
            background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
            borderRadius: '12px',
            color: 'white'
          }}>
            <p style={{ marginBottom: '16px', fontSize: '16px' }}>🌤️ 今天天气晴朗，温度适中</p>
            <div style={{ marginBottom: '16px' }}>
              <p style={{ fontWeight: '600', marginBottom: '8px' }}>👕 建议穿着：</p>
              <ul style={{ paddingLeft: '20px', lineHeight: '1.6' }}>
                <li>轻薄的长袖衬衫或T恤</li>
                <li>休闲裤或牛仔裤</li>
                <li>运动鞋或休闲鞋</li>
              </ul>
            </div>
            <p style={{ fontSize: '14px', opacity: 0.9' }}>💡 小贴士：可以准备一件薄外套，以防晚上降温</p>
          </div>
        </div>

        {/* 用户信息 */}
        <div className="card" style={{ marginBottom: '32px' }}>
          <h2 className="card-title">
            👤 个人信息
          </h2>
          {userProfile ? (
            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))', gap: '16px' }}>
              <div style={{ padding: '16px', background: '#f8f9fa', borderRadius: '8px' }}>
                <div style={{ fontSize: '14px', color: '#666', marginBottom: '4px' }}>用户ID</div>
                <div style={{ fontWeight: '600' }}>{userProfile.user_id}</div>
              </div>
              <div style={{ padding: '16px', background: '#f8f9fa', borderRadius: '8px' }}>
                <div style={{ fontSize: '14px', color: '#666', marginBottom: '4px' }}>用户名</div>
                <div style={{ fontWeight: '600' }}>{userProfile.username}</div>
              </div>
            </div>
          ) : (
            <p style={{ color: '#666', textAlign: 'center', padding: '20px' }}>暂无用户信息</p>
          )}
        </div>

        {/* 功能区域 */}
        <div className="feature-grid">
          <div className="feature-card">
            <span className="feature-icon">📸</span>
            <div className="feature-title">记录今日穿搭</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>拍照记录每日搭配</p>
          </div>

          <div className="feature-card">
            <span className="feature-icon">👗</span>
            <div className="feature-title">我的衣橱</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>管理服装单品</p>
          </div>

          <div className="feature-card">
            <span className="feature-icon">📊</span>
            <div className="feature-title">穿搭统计</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>查看穿搭数据</p>
          </div>

          <div className="feature-card">
            <span className="feature-icon">⚙️</span>
            <div className="feature-title">个人设置</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>偏好和配置</p>
          </div>

          <div className="feature-card">
            <span className="feature-icon">🎨</span>
            <div className="feature-title">风格推荐</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>发现新风格</p>
          </div>

          <div className="feature-card">
            <span className="feature-icon">🌟</span>
            <div className="feature-title">搭配灵感</div>
            <p style={{ color: '#666', fontSize: '14px', marginTop: '8px' }}>获取搭配灵感</p>
          </div>
        </div>
      </div>
    </div>
  );
}
