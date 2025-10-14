'use client';

export default function RegisterPage() {
  return (
    <div className="register-page dark">
      <style>
        {`
          .register-page {
            --primary-color: #000000;
            --background-dark: #1a1a1a;
            --card-dark: #2a2a2a;
            --text-dark: #ffffff;
            --text-secondary-dark: #b0b0b0;
            --border-dark: #4a4a4a;
            background-color: var(--background-dark);
            color: var(--text-dark);
            font-family: 'Inter', sans-serif;
          }
          .register-page .material-icons {
            font-family: 'Material Icons';
            font-weight: normal;
            font-style: normal;
            font-size: 24px;
            line-height: 1;
            letter-spacing: normal;
            text-transform: none;
            display: inline-block;
            white-space: nowrap;
            word-wrap: normal;
            direction: ltr;
            -webkit-font-feature-settings: 'liga';
            -webkit-font-smoothing: antialiased;
          }
          .register-page #svg-background {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100vh;
            z-index: -1;
            overflow: hidden;
          }
          .register-page #svg-background svg {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
          .register-page > div:not(#svg-background) {
            position: relative;
            z-index: 1;
          }
        `}
      </style>

      <div id="svg-background">
        <svg viewBox="0 0 1920 1080" preserveAspectRatio="xMidYMid slice" xmlns="http://www.w3.org/2000/svg">
          <defs>
            <radialGradient id="gradient1" cx="20%" cy="30%">
              <stop offset="0%" stopColor="#4a4a4a" stopOpacity="1">
                <animate attributeName="stop-color" values="#4a4a4a;#5a5a5a;#4a4a4a" dur="8s" repeatCount="indefinite" />
              </stop>
              <stop offset="100%" stopColor="#1a1a1a" stopOpacity="1" />
            </radialGradient>
            <radialGradient id="gradient2" cx="80%" cy="70%">
              <stop offset="0%" stopColor="#383838" stopOpacity="1">
                <animate attributeName="stop-color" values="#383838;#484848;#383838" dur="10s" repeatCount="indefinite" />
              </stop>
              <stop offset="100%" stopColor="#151515" stopOpacity="1" />
            </radialGradient>
            <filter id="glow">
              <feGaussianBlur stdDeviation="40" result="coloredBlur" />
              <feMerge>
                <feMergeNode in="coloredBlur" />
                <feMergeNode in="SourceGraphic" />
              </feMerge>
            </filter>
          </defs>

          <rect width="100%" height="100%" fill="#1a1a1a" />
          <rect width="100%" height="100%" fill="url(#gradient1)" />
          <rect width="100%" height="100%" fill="url(#gradient2)" opacity="0.6" />

          <circle cx="30%" cy="40%" r="300" fill="#4a4a4a" opacity="0.4" filter="url(#glow)">
            <animate attributeName="cx" values="30%;35%;30%" dur="15s" repeatCount="indefinite" />
            <animate attributeName="cy" values="40%;35%;40%" dur="12s" repeatCount="indefinite" />
            <animate attributeName="r" values="300;350;300" dur="10s" repeatCount="indefinite" />
          </circle>

          <circle cx="70%" cy="60%" r="250" fill="#3a3a3a" opacity="0.5" filter="url(#glow)">
            <animate attributeName="cx" values="70%;65%;70%" dur="18s" repeatCount="indefinite" />
            <animate attributeName="cy" values="60%;65%;60%" dur="14s" repeatCount="indefinite" />
            <animate attributeName="r" values="250;300;250" dur="12s" repeatCount="indefinite" />
          </circle>

          <circle cx="50%" cy="50%" r="200" fill="#5a5a5a" opacity="0.3" filter="url(#glow)">
            <animate attributeName="r" values="200;280;200" dur="16s" repeatCount="indefinite" />
            <animate attributeName="opacity" values="0.2;0.3;0.2" dur="8s" repeatCount="indefinite" />
          </circle>
        </svg>
      </div>

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
                  <span className="px-2 bg-[var(--card-dark)] text-[var(--text-secondary-dark)]">已经有账户了？</span>
                </div>
              </div>
              <div className="mt-6">
                <a
                  className="w-full inline-flex justify-center py-3 px-4 border border-[var(--border-dark)] rounded-md shadow-sm bg-transparent text-sm font-medium text-white hover:bg-zinc-800/50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-white focus:ring-offset-zinc-900 transition-colors"
                  href="/login"
                >
                  返回登录
                </a>
              </div>
            </div>
          </div>
          <footer className="mt-8 text-center text-[var(--text-secondary-dark)]">
            <p className="text-xs">©2024 StyleSense. All rights reserved.</p>
          </footer>
        </main>
      </div>
    </div>
  );
}
