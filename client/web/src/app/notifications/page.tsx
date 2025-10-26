'use client';

import { useState } from 'react';
import { Footer } from '@/components/Footer';

export default function NotificationsPage() {
  const [activeFilter, setActiveFilter] = useState('all');

  return (
    <div className="min-h-screen flex flex-col bg-background-light dark:bg-background-dark font-display">
      <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <a className="text-xl font-bold text-text-light dark:text-text-dark" href="/">StyleSense</a>
              <nav className="hidden md:flex space-x-8">
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="/record-style/my-outfit">记录穿搭</a>
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="/main/wardrobe">我的衣橱</a>
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="#">风格灵感</a>
                <a className="text-text-secondary-light dark:text-text-secondary-dark hover:text-text-light dark:hover:text-text-dark" href="#">穿搭分析</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary-light dark:text-text-secondary-dark">search</span>
                <input className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 focus:outline-none focus:ring-2 focus:ring-primary/50 text-text-light dark:text-text-dark w-64" placeholder="搜索衣物" type="text"/>
              </div>
              <button className="text-primary dark:text-white relative">
                <span className="material-icons">notifications</span>
                <span className="absolute top-0 right-0 block h-2 w-2 rounded-full bg-red-500 ring-2 ring-card-light dark:ring-card-dark"></span>
              </button>
              <button>
                <img alt="User avatar" className="h-8 w-8 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDUZrQ70mFuE94RZbqBza4cjKL55KSGqqN473iO43esR3GiQh8n5RxwtaohdLBuSLV5xqrSzkaeWiKobpKNjXtSNtJA4IqaswkZMT4bXDa5laBbWJp4hNdjEo24Tvh9wB8zk7cAQ2nNKhorjoN_JSFsZJK4QLd6yiA99c0l2dOYF0eK56jRMgRo3A5j3m8R4JYGdYTq34SKldjEBQ3a5yUMA9sg_dLbxatb4crcOQZijGPIho_bidIgOoR-sJpZwmRKlfeEGKmTjd8J"/>
              </button>
            </div>
          </div>
        </div>
      </header>
      
      <main className="flex-grow max-w-7xl w-full mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div className="max-w-3xl mx-auto">
          <div className="bg-card-light dark:bg-card-dark p-6 sm:p-8 rounded-lg shadow-sm">
            <div className="flex justify-between items-center mb-6 border-b border-border-light dark:border-border-dark pb-4">
              <h1 className="text-2xl font-bold text-text-light dark:text-text-dark">通知中心</h1>
              <button className="text-sm font-medium text-primary hover:underline dark:text-text-dark dark:hover:text-text-secondary-dark">全部标记为已读</button>
            </div>
            
            <div className="flex mb-6 space-x-2">
              <button 
                onClick={() => setActiveFilter('all')}
                className={`px-4 py-2 text-sm font-medium rounded-full ${activeFilter === 'all' ? 'bg-primary text-white' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
              >
                全部
              </button>
              <button 
                onClick={() => setActiveFilter('unread')}
                className={`px-4 py-2 text-sm font-medium rounded-full ${activeFilter === 'unread' ? 'bg-primary text-white' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
              >
                未读
              </button>
              <button 
                onClick={() => setActiveFilter('comments')}
                className={`px-4 py-2 text-sm font-medium rounded-full ${activeFilter === 'comments' ? 'bg-primary text-white' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
              >
                评论
              </button>
              <button 
                onClick={() => setActiveFilter('promotions')}
                className={`px-4 py-2 text-sm font-medium rounded-full ${activeFilter === 'promotions' ? 'bg-primary text-white' : 'text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-700'}`}
              >
                促销
              </button>
            </div>
            
            <div className="space-y-4">
              <div className="flex items-start space-x-4 p-4 rounded-lg bg-primary/5 dark:bg-primary/10 border-l-4 border-primary">
                <div className="flex-shrink-0">
                  <span className="material-icons text-primary dark:text-white mt-1">campaign</span>
                </div>
                <div className="flex-grow">
                  <p className="font-medium text-text-light dark:text-text-dark">夏季大促现已开始！</p>
                  <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-1">全场商品低至5折，快来选购你的夏日新装吧。优惠活动将于7月31日结束。</p>
                  <p className="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-2">2小时前</p>
                </div>
                <div className="flex-shrink-0 self-center">
                  <div className="w-2.5 h-2.5 rounded-full bg-primary"></div>
                </div>
              </div>
              
              <div className="flex items-start space-x-4 p-4 rounded-lg bg-primary/5 dark:bg-primary/10 border-l-4 border-primary">
                <div className="flex-shrink-0">
                  <img alt="user avatar" className="w-10 h-10 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDW2h2swZVkdE7STqWc-jkdffI2IZ8Yur5SXhOl3qkQgywypBMxrNrXsm6aQMnwGZGhnVGDFVvmH1Vg_k9s79o7UirWPfftrTwUnFm8bzL1ATxOSR1IEuSvRpXQtSCmOGIbIvOIxPKE2C9x717vnaNJcZWBWC9Ie73UvQRwan97YIRtq9njFBMzK9rRq_JHUSPGBXPLQb0CDp79XOR8BfDYJurWXMfmQXHlvQcl81y47F-CgvehizHWHoOn9c3nUzZbp1aJvllMnaQj"/>
                </div>
                <div className="flex-grow">
                  <p className="text-text-light dark:text-text-dark"><span className="font-medium">时尚小达人</span> 评论了你的穿搭 <a className="font-medium text-primary hover:underline" href="#">"夏日海滩风情"</a></p>
                  <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-1 bg-gray-100 dark:bg-gray-700 p-2 rounded-md">"这套搭配太棒了！请问这顶草帽在哪里买的？"</p>
                  <p className="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-2">昨天 18:32</p>
                </div>
                <div className="flex-shrink-0 self-center">
                  <div className="w-2.5 h-2.5 rounded-full bg-primary"></div>
                </div>
              </div>
              
              <div className="flex items-start space-x-4 p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800/50">
                <div className="flex-shrink-0">
                  <span className="material-icons text-text-secondary-light dark:text-text-secondary-dark mt-1">favorite_border</span>
                </div>
                <div className="flex-grow">
                  <p className="text-text-light dark:text-text-dark"><span className="font-medium">设计师精选</span> 喜欢了你的穿搭 <a className="font-medium text-primary hover:underline dark:text-text-secondary-dark dark:hover:text-white" href="#">"都市通勤风"</a></p>
                  <p className="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-2">3天前</p>
                </div>
              </div>
              
              <div className="flex items-start space-x-4 p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800/50">
                <div className="flex-shrink-0">
                  <span className="material-icons text-text-secondary-light dark:text-text-secondary-dark mt-1">inventory_2</span>
                </div>
                <div className="flex-grow">
                  <p className="font-medium text-text-light dark:text-text-dark">你的订单已发货！</p>
                  <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mt-1">订单 #20240520-001 已从仓库发出，预计3天内送达。 <a className="text-primary hover:underline dark:text-text-secondary-dark dark:hover:text-white" href="#">查看物流详情</a></p>
                  <p className="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-2">2024年5月20日</p>
                </div>
              </div>
              
              <div className="flex items-start space-x-4 p-4 rounded-lg hover:bg-gray-50 dark:hover:bg-gray-800/50">
                <div className="flex-shrink-0">
                  <img alt="user avatar" className="w-10 h-10 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAbimqKaSrURF6dCDH2oCz2w0cTTCWxtWWSwuBvev7mxt87rSQZNkpeBVNDrOvBEKJknt0htMZR5y44MIn7bCmFKFyI-ZccqzEPtcbyertPAsQnI1A4oXB8gkogMx-3UE-OO6JSIzg41hViTfDCdEVRnhtFbVDIe3zmMUYVE6zbmE6nZmtyip_ppKBi3qeeuw5xH2ab2owQdKW2kQUnBsvGf5XuyYEpCSt8DDb_-oZd5xrmltnGMDRRATOkR3xzNV6RwQja6iwtqqFJ"/>
                </div>
                <div className="flex-grow">
                  <p className="text-text-light dark:text-text-dark"><span className="font-medium">潮流先锋</span> 开始关注你了。</p>
                  <p className="text-xs text-text-secondary-light dark:text-text-secondary-dark mt-2">2024年5月18日</p>
                </div>
              </div>
            </div>
            
            <div className="mt-8 text-center">
              <button className="text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-primary dark:hover:text-white">
                加载更多
              </button>
            </div>
          </div>
        </div>
      </main>

      <Footer />
    </div>
  );
}
