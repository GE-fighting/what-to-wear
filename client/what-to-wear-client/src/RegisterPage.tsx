import React, { useState } from 'react';
import './styles/modern.css';

interface RegisterPageProps {
  onSwitchToLogin: () => void;
  onRegisterSuccess: () => void;
}

export function RegisterPage({ onSwitchToLogin, onRegisterSuccess }: RegisterPageProps) {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    nickname: '',
    gender: '',
    birthDate: '',
    height: '',
    weight: ''
  });
  const [message, setMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const validateForm = () => {
    if (!formData.username.trim()) {
      setMessage('请输入用户名');
      return false;
    }
    if (formData.username.length < 3) {
      setMessage('用户名至少需要3个字符');
      return false;
    }
    if (!formData.email.trim()) {
      setMessage('请输入邮箱地址');
      return false;
    }
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      setMessage('请输入有效的邮箱地址');
      return false;
    }
    if (!formData.password.trim()) {
      setMessage('请输入密码');
      return false;
    }
    if (formData.password.length < 6) {
      setMessage('密码至少需要6个字符');
      return false;
    }
    if (formData.password !== formData.confirmPassword) {
      setMessage('两次输入的密码不一致');
      return false;
    }
    return true;
  };

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (!validateForm()) {
      return;
    }

    try {
      setIsLoading(true);
      setMessage('注册中...');
      
      // 准备发送到后端的数据
      const registerData = {
        username: formData.username,
        password: formData.password,
        email: formData.email,
        nickname: formData.nickname || formData.username,
        gender: formData.gender,
        birth_date: formData.birthDate,
        height: formData.height ? parseInt(formData.height) : null,
        weight: formData.weight ? parseInt(formData.weight) : null
      };
      
      const response = await fetch('http://localhost:8080/api/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(registerData),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage('注册成功！请登录');
        setTimeout(() => {
          onRegisterSuccess();
        }, 1500);
      } else {
        setMessage(`注册失败: ${data.error}`);
      }
    } catch (error) {
      setMessage('网络错误，请检查服务器连接');
      console.error('Registration failed:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="app-container">
      <div className="auth-container">
        <h1 className="app-title">今天穿什么</h1>
        <p className="page-subtitle">🎨 创建账号，开启个性化穿搭</p>

        <form onSubmit={handleRegister}>
          {/* 基本信息 */}
          <div className="form-section">
            <h3 className="section-title">基本信息</h3>

            <div className="form-group">
              <div className="form-input-wrapper">
                <input
                  type="text"
                  name="username"
                  value={formData.username}
                  onChange={handleInputChange}
                  placeholder=" "
                  className="form-input"
                  disabled={isLoading}
                  required
                  id="reg-username"
                />
                <label className="form-label" htmlFor="reg-username">用户名 *</label>
              </div>
            </div>

            <div className="form-group">
              <div className="form-input-wrapper">
                <input
                  type="email"
                  name="email"
                  value={formData.email}
                  onChange={handleInputChange}
                  placeholder=" "
                  className="form-input"
                  disabled={isLoading}
                  required
                  id="reg-email"
                />
                <label className="form-label" htmlFor="reg-email">邮箱 *</label>
              </div>
            </div>

            <div className="form-group">
              <div className="form-input-wrapper">
                <input
                  type="text"
                  name="nickname"
                  value={formData.nickname}
                  onChange={handleInputChange}
                  placeholder=" "
                  className="form-input"
                  disabled={isLoading}
                  id="reg-nickname"
                />
                <label className="form-label" htmlFor="reg-nickname">昵称</label>
              </div>
            </div>
          </div>

          {/* 密码设置 */}
          <div className="form-section">
            <h3 className="section-title">密码设置</h3>

            <div className="form-group">
              <div className="form-input-wrapper">
                <input
                  type="password"
                  name="password"
                  value={formData.password}
                  onChange={handleInputChange}
                  placeholder=" "
                  className="form-input"
                  disabled={isLoading}
                  required
                  id="reg-password"
                />
                <label className="form-label" htmlFor="reg-password">密码 *</label>
              </div>
            </div>

            <div className="form-group">
              <div className="form-input-wrapper">
                <input
                  type="password"
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleInputChange}
                  placeholder=" "
                  className="form-input"
                  disabled={isLoading}
                  required
                  id="reg-confirm-password"
                />
                <label className="form-label" htmlFor="reg-confirm-password">确认密码 *</label>
              </div>
            </div>
          </div>

          {/* 个人信息 */}
          <div className="form-section">
            <h3 className="section-title">个人信息（可选）</h3>

            <div className="form-row">
              <div className="form-group-special">
                <label className="form-label-special">性别</label>
                <select
                  name="gender"
                  value={formData.gender}
                  onChange={handleInputChange}
                  className="form-input-special"
                  disabled={isLoading}
                >
                  <option value="">请选择</option>
                  <option value="male">男</option>
                  <option value="female">女</option>
                  <option value="other">其他</option>
                </select>
              </div>

              <div className="form-group-special">
                <label className="form-label-special">生日</label>
                <input
                  type="date"
                  name="birthDate"
                  value={formData.birthDate}
                  onChange={handleInputChange}
                  className="form-input-special"
                  disabled={isLoading}
                />
              </div>
            </div>

            <div className="form-row">
              <div className="form-group-special">
                <label className="form-label-special">身高 (cm)</label>
                <input
                  type="number"
                  name="height"
                  value={formData.height}
                  onChange={handleInputChange}
                  placeholder="如：170"
                  min="100"
                  max="250"
                  className="form-input-special"
                  disabled={isLoading}
                />
              </div>

              <div className="form-group-special">
                <label className="form-label-special">体重 (kg)</label>
                <input
                  type="number"
                  name="weight"
                  value={formData.weight}
                  onChange={handleInputChange}
                  placeholder="如：65"
                  min="30"
                  max="200"
                  className="form-input-special"
                  disabled={isLoading}
                />
              </div>
            </div>
          </div>

          <button
            type="submit"
            disabled={isLoading}
            className="btn btn-primary"
          >
            {isLoading ? '注册中...' : '创建账号 🎉'}
          </button>
        </form>

        <div className="switch-page">
          <span>已有账号？</span>
          <button
            onClick={onSwitchToLogin}
            className="link-btn"
          >
            立即登录
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
