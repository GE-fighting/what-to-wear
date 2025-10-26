'use client';

import { FormEvent, useCallback, useMemo, useState } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';

import AuthBackground from '@/components/AuthBackground';
import { Footer } from '@/components/Footer';
import { login } from '@/lib/api/auth';
import '@/styles/auth.css';

export default function LoginPage() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [pending, setPending] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const next = useMemo(() => searchParams?.get('next') ?? '/', [searchParams]);

  const handleSubmit = useCallback(
    async (event: FormEvent<HTMLFormElement>) => {
      event.preventDefault();
      if (pending) return;

      setPending(true);
      setError(null);

      try {
        const token = await login({ username, password });
        if (typeof window !== 'undefined') {
          window.localStorage.setItem('token', token);
        }
        router.replace(next);
      } catch (err) {
        const message = err instanceof Error ? err.message : '登录失败，请稍后再试。';
        setError(message);
      } finally {
        setPending(false);
      }
    },
    [username, password, next, router, pending],
  );

  return (
    <div className="auth-page dark">
      <AuthBackground />

      <div className="relative min-h-screen flex items-center justify-center isolate py-12">
        <div className="absolute inset-0 bg-gradient-to-br from-black/10 via-transparent to-black/10" />
        <main className="w-full max-w-md px-4 z-10 my-auto">
          <div className="bg-[var(--card-dark)]/80 backdrop-blur-lg p-8 sm:p-10 rounded-2xl border border-[var(--border-dark)] shadow-2xl shadow-black/50">
            <div className="text-center mb-8">
              <a className="text-3xl font-black text-white drop-shadow-lg" href="#">
                StyleSense
              </a>
              <h1 className="text-3xl font-bold text-white mt-6">欢迎回来</h1>
              <p className="text-[var(--text-secondary-dark)] mt-2">登录以继续您的时尚之旅。</p>
            </div>
            <form className="space-y-6" onSubmit={handleSubmit}>
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="identifier">
                  用户名或邮箱
                </label>
                <div className="mt-1">
                  <input
                    autoComplete="email"
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="identifier"
                    name="identifier"
                    placeholder="you@example.com"
                    required
                    type="text"
                    value={username}
                    onChange={(event) => setUsername(event.target.value)}
                  />
                </div>
              </div>
              <div>
                <label className="block text-sm font-medium text-[var(--text-secondary-dark)]" htmlFor="password">
                  密码
                </label>
                <div className="mt-1">
                  <input
                    autoComplete="current-password"
                    className="block w-full rounded-md border-[var(--border-dark)] shadow-sm focus:border-white focus:ring-white bg-zinc-800/50 text-white placeholder-[var(--text-secondary-dark)] py-3 px-4"
                    id="password"
                    name="password"
                    placeholder="••••••••"
                    required
                    type="password"
                    value={password}
                    onChange={(event) => setPassword(event.target.value)}
                  />
                </div>
              </div>
              {error ? (
                <p className="text-sm text-red-400" role="alert">
                  {error}
                </p>
              ) : null}
              <div className="flex items-center justify-end">
                <div className="text-sm">
                  <a className="font-medium text-[var(--text-secondary-dark)] hover:text-white" href="#">
                    忘记密码？
                  </a>
                </div>
              </div>
              <div>
                <button
                  className="w-full flex justify-center py-3 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-[var(--primary-color)] hover:bg-zinc-900 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white focus:ring-offset-zinc-900 transition-all duration-300 transform hover:scale-105 disabled:opacity-70 disabled:hover:scale-100"
                  disabled={pending}
                  type="submit"
                >
                  {pending ? '登录中…' : '登录'}
                </button>
              </div>
            </form>
            <div className="mt-6">
              <div className="relative">
                <div className="absolute inset-0 flex items-center">
                  <div className="w-full border-t border-[var(--border-dark)]" />
                </div>
                <div className="relative flex justify-center text-sm">
                  <span className="px-2 bg-[var(--card-dark)] text-[var(--text-secondary-dark)]">还没有账户？</span>
                </div>
              </div>
              <div className="mt-6">
                <a
                  className="w-full inline-flex justify-center py-3 px-4 border border-[var(--border-dark)] rounded-md shadow-sm bg-transparent text-sm font-medium text-white hover:bg-zinc-800/50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white focus:ring-offset-zinc-900 transition-colors"
                  href="/auth/register"
                >
                  注册
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
