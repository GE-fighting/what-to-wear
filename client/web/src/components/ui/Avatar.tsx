'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface AvatarProps extends React.ImgHTMLAttributes<HTMLImageElement> {
  src?: string;
  alt?: string;
  size?: 'sm' | 'md' | 'lg' | 'xl';
  fallback?: string;
  children?: React.ReactNode;
}

export const Avatar: React.FC<AvatarProps> = ({
  src,
  alt = '',
  size = 'md',
  fallback,
  children,
  className,
  ...props
}) => {
  const [imageError, setImageError] = React.useState(false);

  const sizeStyles = {
    sm: 'w-8 h-8 text-xs',
    md: 'w-10 h-10 text-sm',
    lg: 'w-12 h-12 text-base',
    xl: 'w-16 h-16 text-lg',
  };

  const showFallback = !src || imageError;

  if (showFallback) {
    return (
      <div
        className={cn(
          'inline-flex items-center justify-center rounded-full bg-gray-200 dark:bg-gray-700 text-text-primary-light dark:text-text-primary-dark font-medium',
          sizeStyles[size],
          className
        )}
      >
        {children || fallback || (
          <span className="material-icons-outlined">
            person
          </span>
        )}
      </div>
    );
  }

  return (
    <img
      src={src}
      alt={alt}
      className={cn(
        'rounded-full object-cover',
        sizeStyles[size],
        className
      )}
      onError={() => setImageError(true)}
      {...props}
    />
  );
};