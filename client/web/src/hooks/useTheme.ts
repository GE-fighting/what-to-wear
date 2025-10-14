'use client';

import { useState, useEffect } from 'react';

export type Theme = 'light' | 'dark' | 'system';

export const useTheme = () => {
  const [theme, setTheme] = useState<Theme>(() => {
    // 检查localStorage中保存的主题设置
    if (typeof window !== 'undefined') {
      const savedTheme = localStorage.getItem('theme') as Theme;
      if (savedTheme && ['light', 'dark', 'system'].includes(savedTheme)) {
        return savedTheme;
      }

      // 检查系统主题偏好
      if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        return 'dark';
      }
    }
    return 'light';
  });

  // 应用主题到DOM
  useEffect(() => {
    const root = document.documentElement;

    const applyTheme = (themeValue: Theme) => {
      if (themeValue === 'system') {
        const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
        root.classList.toggle('dark', systemTheme === 'dark');
      } else {
        root.classList.toggle('dark', themeValue === 'dark');
      }
    };

    applyTheme(theme);

    // 监听系统主题变化
    if (theme === 'system') {
      const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
      const handleChange = () => applyTheme(theme);
      mediaQuery.addEventListener('change', handleChange);
      return () => mediaQuery.removeEventListener('change', handleChange);
    }
  }, [theme]);

  const toggleTheme = () => {
    setTheme(prevTheme => {
      let newTheme: Theme;

      switch (prevTheme) {
        case 'light':
          newTheme = 'dark';
          break;
        case 'dark':
          newTheme = 'system';
          break;
        case 'system':
        default:
          newTheme = 'light';
          break;
      }

      // 保存到localStorage
      localStorage.setItem('theme', newTheme);
      return newTheme;
    });
  };

  const setThemeMode = (newTheme: Theme) => {
    setTheme(newTheme);
    localStorage.setItem('theme', newTheme);
  };

  return {
    theme,
    isDark: theme === 'dark' || (theme === 'system' &&
      (typeof window !== 'undefined' ? window.matchMedia('(prefers-color-scheme: dark)').matches : false)),
    isLight: theme === 'light' || (theme === 'system' &&
      (typeof window !== 'undefined' ? window.matchMedia('(prefers-color-scheme: light)').matches : false)),
    toggleTheme,
    setThemeMode,
  };
};