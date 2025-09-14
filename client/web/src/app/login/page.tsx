"use client";
import React, { useState } from "react";
import "@/styles/modern.css";
import { useRouter } from "next/navigation";
import { login } from "@/lib/api/auth";
import type { LoginResponse } from "@/types/auth";

export default function LoginPage() {
  const router = useRouter();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!username.trim() || !password.trim()) {
      setMessage("请填写用户名和密码");
      return;
    }
    try {
      setIsLoading(true);
      setMessage("登录中...");
      const data: LoginResponse = await login({ username, password });
      setMessage('登录成功！');
      localStorage.setItem('token', data.token);
      localStorage.setItem('username', username);
      router.replace('/main');
    } catch (error: any) {
      setMessage(`登录失败: ${error?.message || '未知错误'}`);
      console.error('Login failed:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSkipLogin = () => {
    localStorage.setItem('username', '演示用户');
    localStorage.setItem('token', 'demo-token');
    router.replace('/main');
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
            onClick={() => router.push('/register')}
            className="link-btn"
          >
            立即注册
          </button>
        </div>

        <div className="demo-section">
          <div className="demo-divider">
            <span>演示模式</span>
          </div>
          <button
            type="button"
            onClick={handleSkipLogin}
            className="btn btn-demo"
          >
            🚀 跳过登录，直接体验
          </button>
          <p className="demo-hint">
            * 演示模式下部分功能可能无法正常使用
          </p>
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
