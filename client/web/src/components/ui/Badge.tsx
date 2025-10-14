'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface BadgeProps extends React.HTMLAttributes<HTMLSpanElement> {
  variant?: 'default' | 'primary' | 'secondary' | 'success' | 'warning' | 'error';
  size?: 'sm' | 'md';
  children: React.ReactNode;
}

export const Badge: React.FC<BadgeProps> = ({
  variant = 'default',
  size = 'md',
  children,
  className,
  ...props
}) => {
  const baseStyles = `
    inline-flex items-center justify-center
    font-medium
    rounded-full
  `;

  const variantStyles = {
    default: `
      bg-gray-100 dark:bg-gray-800
      text-text-primary-light dark:text-text-primary-dark
    `,
    primary: `
      bg-primary text-white
    `,
    secondary: `
      bg-card-light dark:bg-card-dark
      text-text-primary-light dark:text-text-primary-dark
      border border-border-light dark:border-border-dark
    `,
    success: `
      bg-success-light dark:bg-success-dark
      text-success dark:text-success-light
    `,
    warning: `
      bg-warning-light dark:bg-warning-dark
      text-warning dark:text-warning-light
    `,
    error: `
      bg-error-light dark:bg-error-dark
      text-error dark:text-error-light
    `,
  };

  const sizeStyles = {
    sm: 'px-2 py-1 text-xs',
    md: 'px-3 py-1 text-sm',
  };

  return (
    <span
      className={cn(
        baseStyles,
        variantStyles[variant],
        sizeStyles[size],
        className
      )}
      {...props}
    >
      {children}
    </span>
  );
};