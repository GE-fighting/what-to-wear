import { useState, useEffect } from 'react';
import { LoginPage } from './LoginPage';
import { RegisterPage } from './RegisterPage';
import { MainPage } from './MainPage';
import "./App.css";

type PageType = 'login' | 'register' | 'main';

function App() {
  const [currentPage, setCurrentPage] = useState<PageType>('login');

  useEffect(() => {
    // 检查是否已经登录
    const token = localStorage.getItem('token');
    if (token) {
      setCurrentPage('main');
    }
  }, []);

  const handleLoginSuccess = () => {
    setCurrentPage('main');
  };

  const handleLogout = () => {
    setCurrentPage('login');
  };

  const handleSwitchToRegister = () => {
    setCurrentPage('register');
  };

  const handleSwitchToLogin = () => {
    setCurrentPage('login');
  };

  const handleRegisterSuccess = () => {
    setCurrentPage('login');
  };

  const renderCurrentPage = () => {
    switch (currentPage) {
      case 'login':
        return (
          <LoginPage
            onSwitchToRegister={handleSwitchToRegister}
            onLoginSuccess={handleLoginSuccess}
          />
        );
      case 'register':
        return (
          <RegisterPage
            onSwitchToLogin={handleSwitchToLogin}
            onRegisterSuccess={handleRegisterSuccess}
          />
        );
      case 'main':
        return (
          <MainPage
            onLogout={handleLogout}
          />
        );
      default:
        return (
          <LoginPage
            onSwitchToRegister={handleSwitchToRegister}
            onLoginSuccess={handleLoginSuccess}
          />
        );
    }
  };

  return (
    <div className="container">
      {renderCurrentPage()}
    </div>
  );
}

export default App;
