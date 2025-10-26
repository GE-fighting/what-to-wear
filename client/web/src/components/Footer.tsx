'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface FooterProps {
  className?: string;
  variant?: 'default' | 'auth' | 'dark' | 'card';
}

export const Footer: React.FC<FooterProps> = ({ className, variant = 'default' }) => {
  // 动态获取当前年份
  const currentYear = new Date().getFullYear();

  const variantStyles = {
    default: 'mt-auto py-8 text-center text-text-secondary-light dark:text-text-secondary-dark',
    auth: 'mt-8 text-center text-[var(--text-secondary-dark)]',
    dark: 'mt-auto pt-8 text-center text-subtext-light dark:text-subtext-dark text-sm',
    card: 'bg-card-light dark:bg-card-dark mt-auto',
  };

  return (
    <footer className={cn(variantStyles[variant], className)}>
      <div className={cn(
        variant === 'card' ? 'max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8' : '',
        'text-sm'
      )}>
        <p>
          ©{currentYear} StyleSense. All rights reserved.
        </p>
      </div>
    </footer>
  );
};
