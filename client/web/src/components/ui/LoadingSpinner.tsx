'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface LoadingSpinnerProps extends React.HTMLAttributes<HTMLDivElement> {
  size?: 'sm' | 'md' | 'lg';
  color?: 'primary' | 'secondary';
}

export const LoadingSpinner: React.FC<LoadingSpinnerProps> = ({
  size = 'md',
  color = 'primary',
  className,
  ...props
}) => {
  const sizeStyles = {
    sm: 'w-4 h-4',
    md: 'w-6 h-6',
    lg: 'w-8 h-8',
  };

  const colorStyles = {
    primary: 'border-primary',
    secondary: 'border-text-secondary-light dark:border-text-secondary-dark',
  };

  return (
    <div
      className={cn(
        'animate-spin rounded-full border-2 border-t-transparent',
        sizeStyles[size],
        colorStyles[color],
        className
      )}
      {...props}
    />
  );
};