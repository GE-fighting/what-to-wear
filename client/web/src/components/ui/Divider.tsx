'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface DividerProps extends React.HTMLAttributes<HTMLDivElement> {
  orientation?: 'horizontal' | 'vertical';
  variant?: 'solid' | 'dashed';
}

export const Divider: React.FC<DividerProps> = ({
  orientation = 'horizontal',
  variant = 'solid',
  className,
  ...props
}) => {
  const baseStyles = `
    ${orientation === 'horizontal' ? 'w-full h-px' : 'w-px h-full'}
  `;

  const variantStyles = {
    solid: 'bg-border-light dark:bg-border-dark',
    dashed: 'border border-border-light dark:border-border-dark',
  };

  if (variant === 'dashed') {
    return (
      <div
        className={cn(
          baseStyles,
          orientation === 'horizontal'
            ? 'border-t border-dashed'
            : 'border-l border-dashed',
          className
        )}
        {...props}
      />
    );
  }

  return (
    <div
      className={cn(
        baseStyles,
        variantStyles[variant],
        className
      )}
      {...props}
    />
  );
};