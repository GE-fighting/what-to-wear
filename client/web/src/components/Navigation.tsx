'use client';

import Link from 'next/link';

const NAV_LINKS = [
  { href: '/main/wardrobe/add', label: '记录穿搭' },
  { href: '/main/wardrobe', label: '我的衣橱' },
  { href: '#', label: '风格灵感' },
  { href: '#', label: '穿搭分析' },
];

export default function Navigation() {
  return (
    <header className="bg-card-light font-display dark:bg-card-dark shadow-sm sticky top-0 z-20">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center h-16 gap-4">
          <div className="flex min-w-max items-center gap-3 sm:gap-4">
            <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">StyleSense</h1>
          </div>

          <nav className="hidden flex-1 justify-center gap-8 md:flex">
            {NAV_LINKS.map((link) => (
              <Link
                key={link.href}
                className="nav-link text-sm font-medium text-text-light-secondary transition-colors duration-200 hover:text-primary dark:text-text-dark-secondary dark:hover:text-text-dark-primary"
                href={link.href}
              >
                {link.label}
              </Link>
            ))}
          </nav>

          <div className="ml-auto flex min-w-max items-center gap-4">
            <div className="relative hidden md:block">
              <span className="material-icons-outlined pointer-events-none absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">
                search
              </span>
              <input
                className="search-input w-48 bg-background-light py-2 pl-10 pr-4 text-sm text-text-light-primary placeholder:text-light-secondary transition-colors duration-200 focus:outline-none focus:ring-1 focus:ring-primary dark:bg-background-dark dark:text-text-dark-primary dark:placeholder:text-dark-secondary"
                placeholder="搜索衣物"
                type="text"
              />
            </div>

            <Link
              href="#"
              className="notification-btn relative inline-flex h-10 w-10 items-center justify-center rounded-full text-text-light-secondary transition-colors duration-200 hover:bg-gray-100 dark:text-text-dark-secondary dark:hover:bg-gray-800"
            >
              <span className="material-icons-outlined">notifications_none</span>
              <span className="absolute top-1 right-1 h-2 w-2 rounded-full bg-error"></span>
            </Link>

            <Link href="#" className="group relative inline-flex">
              <img
                alt="User avatar"
                className="avatar-hover h-8 w-8 cursor-pointer rounded-full object-cover transition-opacity duration-200 group-hover:opacity-80"
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL"
              />
              <span className="absolute bottom-0 right-0 h-2.5 w-2.5 rounded-full border-2 border-card-light bg-success dark:border-card-dark"></span>
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
}