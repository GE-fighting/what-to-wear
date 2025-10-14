import React from 'react'

interface CardProps {
  children: React.ReactNode
  className?: string
  hoverable?: boolean
  bordered?: boolean
}

interface CardHeaderProps {
  title: string
  subtitle?: string
  action?: React.ReactNode
  className?: string
}

interface CardContentProps {
  children: React.ReactNode
  className?: string
}

export const Card: React.FC<CardProps> = ({
  children,
  className = '',
  hoverable = false,
  bordered = true,
}) => {
  return (
    <div
      className={`
        bg-card-light dark:bg-card-dark
        rounded-lg shadow-sm
        ${bordered ? 'border border-border-light dark:border-border-dark' : ''}
        p-6
        ${hoverable ? 'hover:shadow-md transition-shadow cursor-pointer' : ''}
        ${className}
      `}
    >
      {children}
    </div>
  )
}

export const CardHeader: React.FC<CardHeaderProps> = ({
  title,
  subtitle,
  action,
  className = '',
}) => (
  <div className={`flex items-center justify-between mb-4 ${className}`}>
    <div>
      <h3 className="text-lg font-medium text-text-primary-light dark:text-text-primary-dark">
        {title}
      </h3>
      {subtitle && (
        <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-1">
          {subtitle}
        </p>
      )}
    </div>
    {action && <div>{action}</div>}
  </div>
)

export const CardContent: React.FC<CardContentProps> = ({
  children,
  className = '',
}) => (
  <div className={`text-text-primary-light dark:text-text-primary-dark ${className}`}>
    {children}
  </div>
)