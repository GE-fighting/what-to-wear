'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface PageTransitionProps {
  children: React.ReactNode;
  className?: string;
  duration?: 'fast' | 'normal' | 'slow';
}

export const PageTransition: React.FC<PageTransitionProps> = ({
  children,
  className,
  duration = 'normal',
}) => {
  const durationClasses = {
    fast: 'duration-150',
    normal: 'duration-250',
    slow: 'duration-350',
  };

  return (
    <div
      className={cn(
        'animate-fade-in',
        durationClasses[duration],
        className
      )}
    >
      {children}
    </div>
  );
};

// 加载骨架屏组件
export const SkeletonCard: React.FC<{
  className?: string;
  lines?: number;
}> = ({ className, lines = 3 }) => (
  <div className={cn('animate-pulse', className)}>
    <div className="h-4 bg-gray-200 dark:bg-gray-700 rounded w-3/4 mb-2"></div>
    {Array.from({ length: lines - 1 }).map((_, i) => (
      <div
        key={i}
        className="h-3 bg-gray-200 dark:bg-gray-700 rounded mb-2"
        style={{ width: `${Math.random() * 40 + 60}%` }}
      ></div>
    ))}
  </div>
);

// 悬浮按钮组件
export const FloatingActionButton: React.FC<{
  icon: string;
  onClick: () => void;
  className?: string;
  variant?: 'primary' | 'secondary';
}> = ({ icon, onClick, className, variant = 'primary' }) => {
  const variantStyles = {
    primary: 'bg-primary hover:bg-primary-hover text-white',
    secondary: 'bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark text-text-primary-light dark:text-text-primary-dark',
  };

  return (
    <button
      onClick={onClick}
      className={cn(
        'fixed bottom-6 right-6 w-14 h-14 rounded-full shadow-lg',
        'flex items-center justify-center',
        'transition-all duration-300 transform hover:scale-110',
        'focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-2',
        variantStyles[variant],
        className
      )}
    >
      <span className="material-icons-outlined text-2xl">{icon}</span>
    </button>
  );
};