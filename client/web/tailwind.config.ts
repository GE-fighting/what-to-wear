import type { Config } from 'tailwindcss'

const config: Config = {
  content: [
    './src/pages/**/*.{js,ts,jsx,tsx,mdx}',
    './src/components/**/*.{js,ts,jsx,tsx,mdx}',
    './src/app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  darkMode: 'class',
  theme: {
    extend: {
      colors: {
        // 主色调 - 黑白极简主义
        primary: '#000000',
        'primary-hover': '#333333',
        'primary-active': '#1a1a1a',
        'primary-disabled': '#666666',

        // 背景色
        'background-light': '#F8F9FA',
        'background-dark': '#121212',

        // 卡片色
        'card-light': '#FFFFFF',
        'card-dark': '#1E1E1E',

        // 边框色
        'border-light': '#E5E7EB',
        'border-dark': '#374151',

        // 文字色 - 浅色模式
        'text-primary-light': '#1F2937',
        'text-secondary-light': '#6B7280',
        'text-tertiary-light': '#9CA3AF',
        'text-disabled-light': '#D1D5DB',

        // 文字色 - 深色模式
        'text-primary-dark': '#F9FAFB',
        'text-secondary-dark': '#9CA3AF',
        'text-tertiary-dark': '#6B7280',
        'text-disabled-dark': '#4B5563',

        // 简化的文字色命名 (与 HTML 一致)
        'text-light': '#1F2937',
        'text-dark': '#F9FAFB',
        'text-light-primary': '#1F2937',
        'text-light-secondary': '#6B7280',
        'text-dark-primary': '#F9FAFB',
        'text-dark-secondary': '#9CA3AF',
        'subtext-light': '#6B7280',
        'subtext-dark': '#9CA3AF',

        // 功能色
        success: '#10B981',
        'success-light': '#D1FAE5',
        'success-dark': '#065F46',

        warning: '#F59E0B',
        'warning-light': '#FEF3C7',
        'warning-dark': '#92400E',

        error: '#EF4444',
        'error-light': '#FEE2E2',
        'error-dark': '#991B1B',

        info: '#3B82F6',
        'info-light': '#DBEAFE',
        'info-dark': '#1E3A8A',
      },
      fontFamily: {
        display: ['"Noto Sans SC"', 'sans-serif'],
        sans: ['"Noto Sans SC"', 'sans-serif'],
      },
      fontSize: {
        'xs': '0.75rem',    // 12px
        'sm': '0.875rem',   // 14px
        'base': '1rem',     // 16px
        'lg': '1.125rem',   // 18px
        'xl': '1.25rem',    // 20px
        '2xl': '1.5rem',    // 24px
        '3xl': '1.875rem',  // 30px
        '4xl': '2.25rem',   // 36px
        '5xl': '3rem',      // 48px
      },
      spacing: {
        // 4px网格系统
        '0.5': '0.125rem',  // 2px
        '1': '0.25rem',     // 4px
        '1.5': '0.375rem',  // 6px
        '2': '0.5rem',      // 8px
        '2.5': '0.625rem',  // 10px
        '3': '0.75rem',     // 12px
        '3.5': '0.875rem',  // 14px
        '4': '1rem',        // 16px
        '5': '1.25rem',     // 20px
        '6': '1.5rem',      // 24px
        '7': '1.75rem',     // 28px
        '8': '2rem',        // 32px
        '9': '2.25rem',     // 36px
        '10': '2.5rem',     // 40px
        '11': '2.75rem',    // 44px
        '12': '3rem',       // 48px
        '14': '3.5rem',     // 56px
        '16': '4rem',       // 64px
        '20': '5rem',       // 80px
        '24': '6rem',       // 96px
        '28': '7rem',       // 112px
        '32': '8rem',       // 128px
      },
      borderRadius: {
        'none': '0',
        'sm': '0.125rem',    // 2px
        'DEFAULT': '0.5rem', // 8px - 默认圆角
        'md': '0.375rem',    // 6px
        'lg': '0.5rem',      // 8px
        'xl': '0.75rem',     // 12px
        '2xl': '1rem',       // 16px
        '3xl': '1.5rem',     // 24px
        'full': '9999px',
      },
      boxShadow: {
        'sm': '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
        'DEFAULT': '0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px 0 rgba(0, 0, 0, 0.06)',
        'md': '0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)',
        'lg': '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
        'xl': '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
        '2xl': '0 25px 50px -12px rgba(0, 0, 0, 0.25)',

        // 特殊阴影
        'card': '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
        'button': '0 1px 2px 0 rgba(0, 0, 0, 0.05)',
        'dropdown': '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
        'modal': '0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)',
      },
      animation: {
        'fade-in': 'fadeIn 0.3s ease-in-out',
        'slide-in': 'slideIn 0.3s ease-out',
        'scale-in': 'scaleIn 0.2s ease-out',
      },
      keyframes: {
        fadeIn: {
          '0%': { opacity: '0' },
          '100%': { opacity: '1' },
        },
        slideIn: {
          '0%': { transform: 'translateY(-10px)', opacity: '0' },
          '100%': { transform: 'translateY(0)', opacity: '1' },
        },
        scaleIn: {
          '0%': { transform: 'scale(0.95)', opacity: '0' },
          '100%': { transform: 'scale(1)', opacity: '1' },
        },
      },
      transitionDuration: {
        'DEFAULT': '250ms',
        'fast': '150ms',
        'slow': '350ms',
      },
      transitionTimingFunction: {
        'DEFAULT': 'cubic-bezier(0.4, 0, 0.2, 1)',
        'in': 'cubic-bezier(0.4, 0, 1, 1)',
        'out': 'cubic-bezier(0, 0, 0.2, 1)',
      },
    },
  },
  plugins: [],
}

export default config