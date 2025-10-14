'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'ghost';
  size?: 'sm' | 'md' | 'lg';
  loading?: boolean;
  icon?: React.ReactNode;
  children: React.ReactNode;
}

export const Button: React.FC<ButtonProps> = ({
  variant = 'primary',
  size = 'md',
  loading = false,
  icon,
  children,
  className,
  disabled,
  ...props
}) => {
  const baseStyles = `
    inline-flex items-center justify-center
    font-medium transition-all duration-200
    focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-2
    disabled:opacity-50 disabled:cursor-not-allowed
    rounded-lg
  `;

  const variantStyles = {
    primary: `
      bg-primary text-white
      hover:bg-primary-hover
      active:bg-primary-active
      shadow-button
      focus:ring-primary
    `,
    secondary: `
      bg-card-light dark:bg-card-dark
      text-text-primary-light dark:text-text-primary-dark
      border border-border-light dark:border-border-dark
      hover:bg-gray-50 dark:hover:bg-gray-800
      focus:ring-primary
    `,
    ghost: `
      bg-transparent
      text-text-primary-light dark:text-text-primary-dark
      hover:bg-gray-100 dark:hover:bg-gray-800
      focus:ring-primary
    `,
  };

  const sizeStyles = {
    sm: 'h-8 px-3 text-sm gap-2',
    md: 'h-10 px-4 text-base gap-2',
    lg: 'h-12 px-6 text-lg gap-3',
  };

  return (
    <button
      className={cn(
        baseStyles,
        variantStyles[variant],
        sizeStyles[size],
        className
      )}
      disabled={disabled || loading}
      {...props}
    >
      {loading && (
        <div className="w-4 h-4 border-2 border-current border-t-transparent rounded-full animate-spin" />
      )}
      {!loading && icon && (
        <span className="material-icons-outlined text-current">
          {typeof icon === 'string' ? icon : icon}
        </span>
      )}
      {children}
    </button>
  );
};