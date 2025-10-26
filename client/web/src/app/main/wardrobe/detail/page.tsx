'use client';

import React, { useState } from 'react';
import { Footer } from '@/components/Footer';

const IMAGES = [
  'https://lh3.googleusercontent.com/aida-public/AB6AXuAXLlX9_wDDbI1JL-XWOD2ZzHctF1C2qgXOKY6rFySp-D5_oW7il56harRlIlVMm03T2AmlAh8Opr3WV5KOtTzA7XaxBZ1r3elJ4tt44AQeoBxntu0L1FNKNk4svMzRDsiLL5Fs37HdiM20WXgZBVxn2yHjDnsrb_4alz15IzpFel9by6Jd-mPwKwGaShtyXUhiGfUNOMeypYitGpQettbrOOG94YqlwmHen3dRw02omW7iMLy8yZPCIid6LErstr7BlzPXsbFSx6M',
  'https://lh3.googleusercontent.com/aida-public/AB6AXuAkHfG8nBe7u5ky9SU6swqaWeW1-vlQh4NGnzLV5wcO1HGZ8UrMgwB1I8etuA48UQjqQTaC7FyDzurm8Yg_tjUXEXX0K9HqsQk4J16OZbuegSNbXf35h2CzqYrbzyReOOQQeD5llB_Kxkk_FzpBnnFtvBm6D9IfD9JV4vIepeh28KsFUbiffkw7PAya8M5RvMmByPuoK1cGq4zL6mSQqiZaJVIU67cpamB-8RDI3B3pv7kb0_HOsBXODmEfF0y4IEcft4h388KEs9Yd',
  'https://lh3.googleusercontent.com/aida-public/AB6AXuBpKRSkI0k_rv3-f3pkIHx8gSRFlXMKWkz5pMMTOl91whfAGmxZ7W_Wvp_51sCXqPS05Ghwswh09y79iYU2YbqzgVI4qFgBJ2Mczj3QLPKX7k1sxRZ3sMniVjXctOedrgdGbFGiCcgRUuiEp4TgBZ2UcKsMvm9ii-cWu6FEeSZ_3WqYfeCv_IZ8HWbm70HZy2d4gHO0jqVJ3ZoyD639pZGAXsiMRunOgs_pq5LK-DtbU-il8n_jKcAe8zuVzcacFavxvNxlkCyNSpRJ',
];

export default function WardrobeDetailPage() {
  const [current, setCurrent] = useState(0);
  const [fullscreen, setFullscreen] = useState(false);

  return (
    <div className="min-h-screen bg-background-light dark:bg-background-dark font-display">
      <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">StyleSense</h1>
              <nav className="hidden md:flex space-x-8">
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="/record-style/my-outfit">记录穿搭</a>
                <a className="text-sm font-medium text-primary border-b-2 border-primary" href="/main/wardrobe">我的衣橱</a>
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">风格灵感</a>
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">穿搭分析</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
                <input className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm w-48 focus:ring-primary focus:border-primary text-text-light-primary dark:text-text-dark-primary" placeholder="搜索衣物" type="text" />
              </div>
              <button className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800">
                <span className="material-icons-outlined">notifications_none</span>
              </button>
              <img alt="User avatar" className="h-8 w-8 rounded-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL" />
            </div>
          </div>
        </div>
      </header>

      <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex items-center justify-between mb-6">
          <a className="flex items-center text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary" href="/main/wardrobe">
            <span className="material-icons mr-1">arrow_back_ios</span>
            返回衣物列表
          </a>
          <div className="flex items-center space-x-2">
            <button className="flex items-center justify-center px-4 py-2 bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark text-sm font-medium rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 text-text-light-primary dark:text-text-dark-primary">
              <span className="material-icons-outlined mr-2 text-base">favorite_border</span>
              收藏
            </button>
            <button className="flex items-center justify-center px-4 py-2 bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark text-sm font-medium rounded-lg hover:bg-gray-50 dark:hover:bg-gray-700 text-text-light-primary dark:text-text-dark-primary">
              <span className="material-icons-outlined mr-2 text-base">share</span>
              分享
            </button>
            <a href="/main/wardrobe/edit" className="flex items-center justify-center px-4 py-2 bg-primary text-white dark:bg-text-dark-primary dark:text-background-dark text-sm font-medium rounded-lg hover:bg-primary-hover">
              <span className="material-icons-outlined mr-2 text-base">edit</span>
              编辑
            </a>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          <div className="space-y-4">
            <div className="relative bg-card-light dark:bg-card-dark rounded-lg overflow-hidden shadow-sm max-h-[600px]">
              <div className="max-h-[600px] aspect-square">
                <img src={IMAGES[current]} alt="机能风冲锋衣" className="w-full h-full object-cover" />
              </div>
              <button onClick={() => setFullscreen(true)} className="absolute top-4 right-4 bg-black/20 text-white p-2 rounded-full hover:bg-black/40 transition-opacity">
                <span className="material-icons-outlined">fullscreen</span>
              </button>
            </div>
            <div className="flex space-x-2 overflow-x-auto pb-2">
              {IMAGES.map((img, idx) => (
                <button key={idx} onClick={() => setCurrent(idx)} className={`flex-shrink-0 w-20 h-20 rounded-lg overflow-hidden transition-all ${current === idx ? 'ring-2 ring-primary' : 'opacity-70 hover:opacity-100'}`}>
                  <img src={img} alt="thumb" className="w-full h-full object-cover" />
                </button>
              ))}
            </div>
          </div>

          <div className="bg-card-light dark:bg-card-dark rounded-lg shadow-sm p-6">
            <h1 className="text-3xl font-bold mb-2 text-text-light-primary dark:text-text-dark-primary">机能风冲锋衣</h1>
            <div className="flex items-center space-x-2 mb-4">
              <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200">户外</span>
              <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200">春秋</span>
              <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200">深灰</span>
            </div>
            <div className="space-y-4 text-sm text-text-light-secondary dark:text-text-dark-secondary">
              {[
                ['分类', '外套'],
                ['品牌', 'Urban Explorer'],
                ['尺码', 'L'],
                ['材质', '聚酯纤维'],
                ['购买日期', '2024-03-15'],
                ['购入价格', '¥ 899.00'],
                ['购买渠道', '线上商城'],
              ].map(([k, v]) => (
                <div key={k} className="flex justify-between">
                  <span>{k}:</span>
                  <span className="font-medium text-text-light-primary dark:text-text-dark-primary">{v}</span>
                </div>
              ))}
            </div>
            <div className="my-6 border-t border-border-light dark:border-border-dark" />
            <div>
              <h2 className="text-base font-semibold mb-2 text-text-light-primary dark:text-text-dark-primary">备注</h2>
              <p className="text-sm text-text-light-secondary dark:text-text-dark-secondary">防水性能优秀，适合户外活动。内胆可拆卸，春秋两季都可穿。</p>
            </div>
            <div className="my-6 border-t border-border-light dark:border-border-dark" />
            <div>
              <h2 className="text-lg font-semibold mb-3 text-text-light-primary dark:text-text-dark-primary">穿着统计</h2>
              <p className="text-sm text-text-light-secondary dark:text-text-dark-secondary">累计穿着 <span className="font-bold text-primary">6</span> 次 · 本月 <span className="font-bold text-primary">2</span> 次 · 最后一次 <span className="font-bold text-primary">3天前</span></p>
              <div className="mt-4">
                <h2 className="text-lg font-semibold mb-2 text-text-light-primary dark:text-text-dark-primary">耐久度</h2>
                <div className="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2.5">
                  <div className="bg-green-500 h-2.5 rounded-full" style={{ width: '90%' }} />
                </div>
                <p className="text-xs text-right mt-1 text-text-light-secondary dark:text-text-dark-secondary">预计剩余寿命: 2 年</p>
              </div>
            </div>
          </div>
        </div>

        <div className="mt-12">
          <h2 className="text-xl font-bold mb-6 text-text-light-primary dark:text-text-dark-primary">穿着历史</h2>
          <div className="bg-card-light dark:bg-card-dark rounded-lg shadow-sm p-6">
            {[['周末出行', '3天前', '晴天 • 15°C', '8小时'], ['通勤穿搭', '1周前', '多云 • 12°C', '10小时'], ['户外徒步', '2周前', '阴天 • 10°C', '6小时']].map(([title, time, weather, duration], idx) => (
              <div key={idx} className={`flex items-start space-x-4 pb-4 ${idx < 2 ? 'border-b border-border-light dark:border-border-dark' : ''}`}>
                <div className="flex-shrink-0 text-center">
                  <div className="text-sm font-semibold text-text-light-primary dark:text-text-dark-primary">{idx === 2 ? '11月' : '12月'}</div>
                  <div className="text-2xl font-bold text-primary">{idx === 0 ? '15' : idx === 1 ? '08' : '28'}</div>
                </div>
                <div className="flex-grow">
                  <div className="flex items-center justify-between mb-2">
                    <span className="text-sm font-medium text-text-light-primary dark:text-text-dark-primary">{title}</span>
                    <span className="text-xs text-text-light-secondary dark:text-text-dark-secondary">{time}</span>
                  </div>
                  <div className="flex items-center space-x-2 text-xs text-text-light-secondary dark:text-text-dark-secondary">
                    <span className="material-icons-outlined text-sm">wb_sunny</span>
                    <span>{weather}</span>
                    <span className="ml-2">穿着时长: {duration}</span>
                  </div>
                </div>
                <img src={IMAGES[0]} alt="穿搭照片" className="w-16 h-16 rounded-lg object-cover" />
              </div>
            ))}
            <button className="w-full mt-6 py-2 text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary border border-border-light dark:border-border-dark rounded-lg hover:bg-background-light dark:hover:bg-background-dark transition-colors">查看全部历史记录</button>
          </div>
        </div>
      </main>

      {fullscreen && (
        <div className="fixed inset-0 bg-black z-50 flex items-center justify-center">
          <button onClick={() => setFullscreen(false)} className="absolute top-4 right-4 text-white p-2 rounded-full hover:bg-white/20">
            <span className="material-icons-outlined text-3xl">close</span>
          </button>
          <img src={IMAGES[current]} alt="full" className="max-h-[90vh] max-w-[90vw] object-contain" />
        </div>
      )}

      <Footer />
    </div>
  );
}
