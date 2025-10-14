'use client';

import { usePathname } from 'next/navigation';

export default function SettingsLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();

  return (
    <div className="min-h-screen flex flex-col bg-background-light dark:bg-background-dark font-display">
      <header className="bg-card-light dark:bg-card-dark shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <a className="text-xl font-bold text-text-light dark:text-text-dark" href="/">StyleSense</a>
              <nav className="hidden md:flex space-x-8">
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="/record-style/my-outfit">记录穿搭</a>
                <a className="text-text-light dark:text-text-dark font-medium border-b-2 border-primary" href="/main/wardrobe">我的衣橱</a>
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="#">风格灵感</a>
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="#">穿搭分析</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary-light dark:text-text-secondary-dark">search</span>
                <input className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 focus:outline-none focus:ring-2 focus:ring-primary/50 text-text-light dark:text-text-dark w-64" placeholder="搜索衣物" type="text"/>
              </div>
              <button className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark">
                <span className="material-icons">notifications_none</span>
              </button>
              <button>
                <img alt="User avatar" className="h-8 w-8 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDUZrQ70mFuE94RZbqBza4cjKL55KSGqqN473iO43esR3GiQh8n5RxwtaohdLBuSLV5xqrSzkaeWiKobpKNjXtSNtJA4IqaswkZMT4bXDa5laBbWJp4hNdjEo24Tvh9wB8zk7cAQ2nNKhorjoN_JSFsZJK4QLd6yiA99c0l2dOYF0eK56jRMgRo3A5j3m8R4JYGdYTq34SKldjEBQ3a5yUMA9sg_dLbxatb4crcOQZijGPIho_bidIgOoR-sJpZwmRKlfeEGKmTjd8J"/>
              </button>
            </div>
          </div>
        </div>
      </header>
      
      <main className="flex-grow max-w-7xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="flex flex-col md:flex-row gap-12">
          <aside className="md:w-1/4">
            <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm text-center">
              <img alt="User avatar" className="h-24 w-24 rounded-full mx-auto mb-4" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDW2h2swZVkdE7STqWc-jkdffI2IZ8Yur5SXhOl3qkQgywypBMxrNrXsm6aQMnwGZGhnVGDFVvmH1Vg_k9s79o7UirWPfftrTwUnFm8bzL1ATxOSR1IEuSvRpXQtSCmOGIbIvOIxPKE2C9x717vnaNJcZWBWC9Ie73UvQRwan97YIRtq9njFBMzK9rRq_JHUSPGBXPLQb0CDp79XOR8BfDYJurWXMfmQXHlvQcl81y47F-CgvehizHWHoOn9c3nUzZbp1aJvllMnaQj"/>
              <h2 className="text-xl font-bold text-text-light dark:text-text-dark">用户名</h2>
              <p className="text-text-secondary-light dark:text-text-secondary-dark mt-1">时尚爱好者</p>
              <div className="flex justify-center space-x-6 mt-4 text-text-secondary-light dark:text-text-secondary-dark">
                <div className="text-center">
                  <p className="font-bold text-lg text-text-light dark:text-text-dark">463</p>
                  <p className="text-sm">衣物</p>
                </div>
                <div className="text-center">
                  <p className="font-bold text-lg text-text-light dark:text-text-dark">128</p>
                  <p className="text-sm">穿搭</p>
                </div>
                <div className="text-center">
                  <p className="font-bold text-lg text-text-light dark:text-text-dark">32</p>
                  <p className="text-sm">灵感</p>
                </div>
              </div>
            </div>
            
            <nav className="mt-8">
              <ul className="space-y-2">
                <li>
                  <a 
                    className={`flex items-center p-3 rounded-lg ${pathname === '/settings/personal-information' ? 'bg-primary/10 text-primary font-medium' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
                    href="/settings/personal-information"
                  >
                    <span className="material-icons mr-3">account_circle</span>
                    <span>个人信息</span>
                  </a>
                </li>
                <li>
                  <a 
                    className={`flex items-center p-3 rounded-lg ${pathname === '/settings/account-setting' ? 'bg-primary/10 text-primary font-medium' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
                    href="/settings/account-setting"
                  >
                    <span className="material-icons mr-3">settings</span>
                    <span>账户设置</span>
                  </a>
                </li>
                <li>
                  <a className="flex items-center p-3 rounded-lg text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20" href="#">
                    <span className="material-icons mr-3">logout</span>
                    <span>退出登录</span>
                  </a>
                </li>
              </ul>
            </nav>
          </aside>
          
          <div className="md:w-3/4">
            {children}
          </div>
        </div>
      </main>
      
      <footer className="mt-auto py-8 text-center text-text-secondary-light dark:text-text-secondary-dark">
        <p>©2024 StyleSense. All rights reserved.</p>
      </footer>
    </div>
  );
}
