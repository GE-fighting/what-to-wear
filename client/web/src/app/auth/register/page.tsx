'use client';

import AuthBackground from '@/components/AuthBackground';
import { Footer } from '@/components/Footer';
import '@/styles/auth.css';

export default function RegisterPage() {
  return (
    <div className="auth-page dark">
      <AuthBackground />

      <div className="relative min-h-screen flex items-center justify-center isolate py-12">
        <div className="absolute inset-0 bg-gradient-to-br from-black/10 via-transparent to-black/10"></div>
        <main className="w-full max-w-md px-4 z-10 my-auto">
          <div className="bg-[var(--card-dark)]/80 backdrop-blur-lg p-8 sm:p-10 rounded-2xl border border-[var(--border-dark)] shadow-2xl shadow-black/50">
            <div className="text-center mb-8">
              <a className="text-3xl font-black text-white drop-shadow-lg" href="#">StyleSense</a>
              <h1 className="text-3xl font-bold text-white mt-6">创建您的账户</h1>
              <p className="text-[var(--text-secondary-dark)] mt-2">开启您的时尚之旅</p>
            </div>
            <form className="space-y-6" action="#" method="POST">
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="username">
                  用户名
                </label>
                <div className="mt-1">
                  <input
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="username"
                    name="username"
                    placeholder="输入您的用户名"
                    required
                    type="text"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="email">
                  邮箱地址
                </label>
                <div className="mt-1">
                  <input
                    autoComplete="email"
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="email"
                    name="email"
                    placeholder="you@example.com"
                    required
                    type="email"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="password">
                  密码
                </label>
                <div className="mt-1">
                  <input
                    autoComplete="new-password"
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="password"
                    name="password"
                    placeholder="输入您的密码"
                    required
                    type="password"
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="confirm-password">
                  确认密码
                </label>
                <div className="mt-1">
                  <input
                    autoComplete="new-password"
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="confirm-password"
                    name="confirm-password"
                    placeholder="再次输入您的密码"
                    required
                    type="password"
                  />
                </div>
              </div>
              <div>
                <button
                  className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[var(--primary-color)] hover:bg-zinc-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white focus:ring-offset-zinc-900 transition-all duration-300 transform hover:scale-105"
                  type="submit"
                >
                  注册
                </button>
              </div>
            </form>
            <div className="mt-6">
              <div className="relative">
                <div className="absolute inset-0 flex items-center">
                  <div className="w-full border-t border-[var(--border-dark)]"></div>
                </div>
                <div className="relative flex justify-center text-sm">
                  <span className="px-2 bg-[var(--card-dark)] text-[var(--text-secondary-dark)]">已经有账户了吗？</span>
                </div>
              </div>
              <div className="mt-6">
                <a
                  className="w-full inline-flex justify-center py-3 px-4 border border-[var(--border-dark)] rounded-md shadow-sm bg-transparent text-sm font-medium text-white hover:bg-zinc-800/50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white focus:ring-offset-zinc-900 transition-colors"
                  href="/auth/login"
                >
                  返回登录
                </a>
              </div>
            </div>
          </div>
          <Footer variant="auth" />
        </main>
      </div>
    </div>
  );
}
