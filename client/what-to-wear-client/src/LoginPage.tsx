import React, { useState } from 'react';
import './styles/modern.css';

interface LoginPageProps {
  onSwitchToRegister: () => void;
  onLoginSuccess: (token: string) => void;
}

export function LoginPage({ onSwitchToRegister, onLoginSuccess }: LoginPageProps) {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!username.trim() || !password.trim()) {
      setMessage('请填写用户名和密码');
      return;
    }

    try {
      setIsLoading(true);
      setMessage('登录中...');
      
      const response = await fetch('http://localhost:8080/api/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage('登录成功！');
        localStorage.setItem('token', data.token);
        localStorage.setItem('username', username);
        onLoginSuccess(data.token);
      } else {
        setMessage(`登录失败: ${data.error}`);
      }
    } catch (error) {
      setMessage('网络错误，请检查服务器连接');
      console.error('Login failed:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="app-container">
      <div className="auth-container">
        <h1 className="app-title">今天穿什么</h1>
        <p className="page-subtitle">✨ 智能穿搭，从登录开始</p>

        <form onSubmit={handleLogin}>
          <div className="form-group">
            <div className="form-input-wrapper">
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder=" "
                className="form-input"
                disabled={isLoading}
                id="username"
              />
              <label className="form-label" htmlFor="username">用户名</label>
            </div>
          </div>

          <div className="form-group">
            <div className="form-input-wrapper">
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder=" "
                className="form-input"
                disabled={isLoading}
                id="password"
              />
              <label className="form-label" htmlFor="password">密码</label>
            </div>
          </div>

          <button
            type="submit"
            disabled={isLoading}
            className="btn btn-primary"
          >
            {isLoading ? '登录中...' : '开始穿搭之旅 →'}
          </button>
        </form>

        <div className="switch-page">
          <span>还没有账号？</span>
          <button
            onClick={onSwitchToRegister}
            className="link-btn"
          >
            立即注册
          </button>
        </div>

        {message && (
          <div className={`message ${message.includes('成功') ? 'message-success' : 'message-error'}`}>
            {message}
          </div>
        )}
      </div>
    </div>
  );
}
