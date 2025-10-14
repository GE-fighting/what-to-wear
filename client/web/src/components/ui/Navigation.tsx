import React from 'react'
import { cn } from '@/lib/utils'

interface NavigationItem {
  label: string
  href: string
  active?: boolean
  icon?: React.ReactNode
}

interface NavigationProps {
  items: NavigationItem[]
  className?: string
}

export const Navigation: React.FC<NavigationProps> = ({
  items,
  className = '',
}) => {
  return (
    <nav className={cn('hidden md:flex items-center gap-8', className)}>
      {items.map((item) => (
        <a
          key={item.href}
          href={item.href}
          className={cn(
            'text-sm font-medium transition-colors duration-200 border-b-2',
            item.active
              ? 'text-primary border-primary'
              : 'text-text-secondary-light dark:text-text-secondary-dark border-transparent hover:text-primary'
          )}
        >
          <span className="flex items-center gap-2">
            {item.icon && <span className="text-base">{item.icon}</span>}
            {item.label}
          </span>
        </a>
      ))}
    </nav>
  )
}

interface HeaderProps {
  title?: string
  navigation?: NavigationItem[]
  actions?: React.ReactNode
  showSearch?: boolean
  userAvatar?: string
  className?: string
}

export const Header: React.FC<HeaderProps> = ({
  title = 'StyleSense',
  navigation,
  actions,
  showSearch = true,
  userAvatar,
  className = '',
}) => {
  return (
    <header className={cn(
      'bg-card-light font-display dark:bg-card-dark shadow-sm sticky top-0 z-20',
      className
    )}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center h-16 gap-4">
          <div className="flex min-w-max items-center gap-3 sm:gap-4">
            <h1 className="text-xl font-bold text-text-primary-light dark:text-text-primary-dark">
              {title}
            </h1>
          </div>

          {navigation && (
            <Navigation
              className="flex-1 justify-center"
              items={navigation}
            />
          )}

          <div className="ml-auto flex min-w-max items-center gap-4">
            {showSearch && (
              <div className="relative hidden md:block">
                <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary-light dark:text-text-secondary-dark">
                  search
                </span>
                <input
                  className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm w-48 focus:ring-primary focus:border-primary text-text-primary-light dark:text-text-primary-dark"
                  placeholder="搜索衣物"
                  type="text"
                />
              </div>
            )}

            <div className="relative">
              <button className="p-2 rounded-full text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-800">
                <span className="material-icons-outlined">notifications_none</span>
              </button>
            </div>

            {userAvatar ? (
              <button>
                <img
                  alt="User avatar"
                  className="h-8 w-8 rounded-full object-cover"
                  src={userAvatar}
                />
              </button>
            ) : (
              <button className="p-2 rounded-full text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-800">
                <span className="material-icons-outlined">account_circle</span>
              </button>
            )}

            {actions && <div>{actions}</div>}
          </div>
        </div>
      </div>
    </header>
  )
}