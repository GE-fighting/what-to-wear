'use client';

import React from 'react';
import { cn } from '@/lib/utils';

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  required?: boolean;
  icon?: React.ReactNode;
}

export const Input: React.FC<InputProps> = ({
  label,
  error,
  required = false,
  icon,
  className,
  ...props
}) => {
  const inputId = React.useId();

  return (
    <div className="space-y-2">
      {label && (
        <label
          htmlFor={inputId}
          className={`
            block text-sm font-medium
            ${error ? 'text-error' : 'text-text-primary-light dark:text-text-primary-dark'}
          `}
        >
          {label}
          {required && <span className="text-error ml-1">*</span>}
        </label>
      )}
      <div className="relative">
        {icon && (
          <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary-light dark:text-text-secondary-dark">
            {typeof icon === 'string' ? icon : icon}
          </span>
        )}
        <input
          id={inputId}
          className={`
            w-full px-3 py-2
            bg-card-light dark:bg-card-dark
            border rounded-lg
            focus:outline-none focus:ring-2 focus:ring-primary
            text-text-primary-light dark:text-text-primary-dark
            placeholder-text-secondary-light dark:placeholder-text-secondary-dark
            ${icon ? 'pl-10' : ''}
            ${error
              ? 'border-error focus:border-error'
              : 'border-border-light dark:border-border-dark focus:border-primary'
            }
            ${className}
          `}
          {...props}
        />
      </div>
      {error && (
        <p className="text-sm text-error flex items-center gap-1">
          <span className="material-icons-outlined text-xs">error</span>
          {error}
        </p>
      )}
    </div>
  );
};