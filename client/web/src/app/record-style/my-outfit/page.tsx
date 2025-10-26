'use client';

import { useState } from 'react';
import { Footer } from '@/components/Footer';

export default function MyOutfitPage() {
  const [activeTab, setActiveTab] = useState('my-room');
  const [showOutfitModal, setShowOutfitModal] = useState(false);
  const [analysisStep, setAnalysisStep] = useState('upload');
  const [currentCategory, setCurrentCategory] = useState('tops');

  return (
    <>
      <style jsx>{`
        .tab-active {
          border-bottom: 2px solid #000000;
          color: #000000;
        }
      `}</style>
      
      <div className="min-h-screen bg-background-light dark:bg-background-dark font-display">
        <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex justify-between items-center h-16">
              <div className="flex items-center space-x-8">
                <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">StyleSense</h1>
                <nav className="hidden md:flex space-x-8">
                  <a className="text-sm font-medium text-primary border-b-2 border-primary" href="/record-style/my-outfit">记录穿搭</a>
                  <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="/main/wardrobe">我的衣橱</a>
                  <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">风格灵感</a>
                  <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">穿搭分析</a>
                </nav>
              </div>
              <div className="flex items-center space-x-4">
                <div className="relative hidden md:block">
                  <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
                  <input className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm w-48 focus:ring-primary focus:border-primary text-text-light-primary dark:text-text-dark-primary" placeholder="搜索衣物" type="text"/>
                </div>
                <div className="relative">
                  <button className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800">
                    <span className="material-icons-outlined">notifications_none</span>
                  </button>
                </div>
                <button>
                  <img alt="User avatar" className="h-8 w-8 rounded-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL"/>
                </button>
              </div>
            </div>
          </div>
        </header>

        <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
          <div className="border-b border-border-light dark:border-border-dark mb-6">
            <nav className="-mb-px flex space-x-8">
              <button 
                onClick={() => setActiveTab('my-room')}
                className={`py-4 px-1 border-b-2 border-transparent font-medium text-sm transition-colors ${activeTab === 'my-room' ? 'tab-active' : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'}`}
              >
                我的搭间
              </button>
              <button 
                onClick={() => setActiveTab('mirror-inspiration')}
                className={`py-4 px-1 border-b-2 border-transparent font-medium text-sm transition-colors ${activeTab === 'mirror-inspiration' ? 'tab-active' : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'}`}
              >
                镜像灵感
              </button>
            </nav>
          </div>

          {activeTab === 'my-room' && (
            <div>
              <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
                <div className="lg:col-span-2 bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm border border-border-light dark:border-border-dark">
                  <h2 className="text-xl font-semibold mb-6 text-text-light-primary dark:text-text-dark-primary">设计我的穿搭</h2>
                  <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <div onClick={() => setShowOutfitModal(true)} className="flex flex-col items-center justify-center h-56 bg-background-light dark:bg-background-dark rounded-lg border-2 border-dashed border-border-light dark:border-border-dark cursor-pointer hover:border-primary/70 transition-all duration-200">
                      <span className="material-icons-outlined text-4xl text-text-light-secondary dark:text-text-dark-secondary">checkroom</span>
                      <p className="mt-2 text-text-light-secondary dark:text-text-dark-secondary">选择上衣</p>
                    </div>
                    <div onClick={() => setShowOutfitModal(true)} className="flex flex-col items-center justify-center h-56 bg-background-light dark:bg-background-dark rounded-lg border-2 border-dashed border-border-light dark:border-border-dark cursor-pointer hover:border-primary/70 transition-all duration-200">
                      <span className="material-icons-outlined text-4xl text-text-light-secondary dark:text-text-dark-secondary">king_bed</span>
                      <p className="mt-2 text-text-light-secondary dark:text-text-dark-secondary">选择下装</p>
                    </div>
                    <div onClick={() => setShowOutfitModal(true)} className="flex flex-col items-center justify-center h-56 bg-background-light dark:bg-background-dark rounded-lg border-2 border-dashed border-border-light dark:border-border-dark cursor-pointer hover:border-primary/70 transition-all duration-200">
                      <span className="material-icons-outlined text-4xl text-text-light-secondary dark:text-text-dark-secondary">ice_skating</span>
                      <p className="mt-2 text-text-light-secondary dark:text-text-dark-secondary">选择鞋子</p>
                    </div>
                  </div>
                </div>
                <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm border border-border-light dark:border-border-dark flex flex-col justify-between">
                  <div>
                    <h2 className="text-xl font-semibold mb-4 text-text-light-primary dark:text-text-dark-primary">AI 助手</h2>
                    <div className="space-y-4">
                      <div>
                        <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary mb-1" htmlFor="style-input">输入风格或关键词</label>
                        <input className="w-full bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-md py-2 px-3 focus:outline-none focus:ring-2 focus:ring-primary/50 text-text-light-primary dark:text-text-dark-primary" id="style-input" placeholder="例如：都市休闲、复古、Y2K" type="text"/>
                      </div>
                      <button className="w-full bg-primary text-white py-2.5 rounded-md font-semibold hover:bg-primary-hover transition-colors duration-200 flex items-center justify-center space-x-2">
                        <span className="material-icons-outlined text-xl">auto_awesome</span>
                        <span>AI 智能生成</span>
                      </button>
                    </div>
                  </div>
                  <div className="mt-8 space-y-4">
                    <button className="w-full bg-green-500 text-white py-2.5 rounded-md font-semibold hover:bg-green-600 transition-colors duration-200 flex items-center justify-center space-x-2">
                      <span className="material-icons-outlined text-xl">camera_alt</span>
                      <span>AI 试穿</span>
                    </button>
                    <button className="w-full border border-border-light dark:border-border-dark text-text-light-secondary dark:text-text-dark-secondary py-2.5 rounded-md font-semibold hover:bg-background-light dark:hover:bg-background-dark hover:text-text-light-primary dark:hover:text-text-dark-primary transition-colors duration-200">
                      保存穿搭
                    </button>
                  </div>
                </div>
              </div>
            </div>
          )}

          {activeTab === 'mirror-inspiration' && (
            <div>
              <div className="flex justify-between items-center mb-6">
                <h1 className="text-3xl font-bold text-text-light-primary dark:text-text-dark-primary">镜像灵感</h1>
                {analysisStep !== 'upload' && (
                  <button onClick={() => setAnalysisStep('upload')} className="text-sm text-primary font-medium hover:underline flex items-center space-x-1">
                    <span className="material-icons-outlined text-base">arrow_back</span>
                    <span>重新上传</span>
                  </button>
                )}
              </div>
              
              <div className="bg-card-light dark:bg-card-dark p-8 rounded-lg shadow-sm border border-border-light dark:border-border-dark">
                {analysisStep === 'upload' && (
                  <div className="text-center">
                    <h2 className="text-2xl font-semibold mb-2 text-text-light-primary dark:text-text-dark-primary">上传穿搭图片</h2>
                    <p className="text-text-light-secondary dark:text-text-dark-secondary mb-8 max-w-2xl mx-auto">上传一张您喜欢的穿搭照片，AI 将为您解析并从您的衣橱中推荐相似单品，助您轻松复刻心仪造型。</p>
                    <div onClick={() => setAnalysisStep('analysis')} className="relative border-2 border-dashed border-border-light dark:border-border-dark rounded-lg p-10 text-center cursor-pointer hover:border-primary/70 transition-colors duration-200 max-w-lg mx-auto">
                      <input className="absolute inset-0 w-full h-full opacity-0 cursor-pointer" type="file"/>
                      <div className="flex flex-col items-center">
                        <span className="material-icons-outlined text-5xl text-text-light-secondary dark:text-text-dark-secondary mb-4">cloud_upload</span>
                        <p className="font-semibold text-text-light-primary dark:text-text-dark-primary">点击或拖拽图片到这里</p>
                        <p className="text-sm text-text-light-secondary dark:text-text-dark-secondary">支持 PNG, JPG, GIF (最大 5MB)</p>
                      </div>
                    </div>
                  </div>
                )}

                {analysisStep === 'analysis' && (
                  <div>
                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
                      <div className="flex flex-col items-center">
                        <h2 className="text-xl font-semibold mb-4 text-text-light-primary dark:text-text-dark-primary">灵感来源</h2>
                        <img alt="Uploaded Outfit" className="rounded-lg w-full max-w-md object-cover aspect-[3/4] shadow-md" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw"/>
                      </div>
                      <div className="flex flex-col">
                        <h2 className="text-xl font-semibold mb-4 text-text-light-primary dark:text-text-dark-primary">AI 解析与单品选择</h2>
                        <div className="border-b border-border-light dark:border-border-dark mb-4">
                          <nav className="-mb-px flex space-x-6">
                            <button 
                              onClick={() => setCurrentCategory('tops')}
                              className={`py-3 px-1 border-b-2 border-transparent font-medium transition-colors ${currentCategory === 'tops' ? 'tab-active' : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'}`}
                            >
                              上装
                            </button>
                            <button 
                              onClick={() => setCurrentCategory('bottoms')}
                              className={`py-3 px-1 border-b-2 border-transparent font-medium transition-colors ${currentCategory === 'bottoms' ? 'tab-active' : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'}`}
                            >
                              下装
                            </button>
                            <button 
                              onClick={() => setCurrentCategory('shoes')}
                              className={`py-3 px-1 border-b-2 border-transparent font-medium transition-colors ${currentCategory === 'shoes' ? 'tab-active' : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'}`}
                            >
                              鞋子
                            </button>
                          </nav>
                        </div>
                        <div className="flex-grow overflow-y-auto pr-2">
                          <div className="bg-background-light dark:bg-background-dark p-4 rounded-lg mb-4">
                            <p className="text-sm font-medium text-text-light-primary dark:text-text-dark-primary">AI 分析结果:</p>
                            <div className="flex flex-wrap gap-2 mt-2">
                              <span className="bg-blue-100 text-blue-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-blue-900 dark:text-blue-300">颜色: 白色</span>
                              <span className="bg-gray-100 text-gray-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-gray-700 dark:text-gray-300">风格: 简约</span>
                              <span className="bg-green-100 text-green-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-green-900 dark:text-green-300">款式: T恤</span>
                              <span className="bg-yellow-100 text-yellow-800 text-xs font-medium px-2.5 py-0.5 rounded-full dark:bg-yellow-900 dark:text-yellow-300">厚度: 薄款</span>
                            </div>
                          </div>
                          <p className="text-text-light-secondary dark:text-text-dark-secondary mb-4 text-sm">根据AI分析，为你从衣橱中筛选出以下相似单品：</p>
                          <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
                            <div className="group relative cursor-pointer border-2 border-primary/80 rounded-lg overflow-hidden shadow-md">
                              <div className="absolute top-2 right-2 bg-primary text-white rounded-full h-6 w-6 flex items-center justify-center">
                                <span className="material-icons-outlined text-base">check</span>
                              </div>
                              <img alt="简约白T" className="h-40 w-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw"/>
                              <div className="p-2 bg-card-light dark:bg-card-dark">
                                <h4 className="text-sm font-medium truncate text-text-light-primary dark:text-text-dark-primary">简约白T</h4>
                              </div>
                            </div>
                            <div className="group relative cursor-pointer border-2 border-transparent hover:border-primary/50 rounded-lg overflow-hidden transition">
                              <img alt="棉质圆领T恤" className="h-40 w-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDH-L7FdmL3xHhlueKM1mE__6QUrDRhY3Lpu69oN3tqwTZgTzqhjiMG4KKmHqm3hfhYFuE3IyfN4wqsbV4uXFepSq3L3VbONrlxEC2SPr1rSxTwDle1CL6nu92HiusbkGSQUU-h41V2a9UnJBIxHa1b65CaD5f04CLPvuSQ5Ua1zMRvMv5GB-17Oy7kw3oEsVihT9dgmO03x4zPGUnCEGbYpjcNsOD4sj38ph_VyANpYycuQSJgI6MYMLtNTevwKLM0iRuXyxpisQ"/>
                              <div className="p-2 bg-card-light dark:bg-card-dark">
                                <h4 className="text-sm font-medium truncate text-text-light-primary dark:text-text-dark-primary">棉质圆领T恤</h4>
                              </div>
                            </div>
                            <div className="group relative cursor-pointer border-2 border-transparent hover:border-primary/50 rounded-lg overflow-hidden transition">
                              <img alt="基础款白T" className="h-40 w-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDU-ygsALDYwS7zNZWUs9pZDlDvvy_Wq8126XFo7nZ-V69gRjCJXJdHExRtPVzzzd12cUiT3m74IQOCcSreUxzm1Mpf5v-PmquoM8yuCHU63SpnaNCBkRpqpVeiBSrKZda3FlXe9oPWDBHK3uCWEdp6_5ZBlfxMn-YsBzChoFHOF___qI7oODtntM3hbmA6JTJObhYACPN8ClNjeq54cpGvroi1HCWaHTv-Z8Dq3YxhPEo2bdWCeJ6wFfOC4_dP49DoJT8chSwbCg"/>
                              <div className="p-2 bg-card-light dark:bg-card-dark">
                                <h4 className="text-sm font-medium truncate text-text-light-primary dark:text-text-dark-primary">基础款白T</h4>
                              </div>
                            </div>
                          </div>
                        </div>
                        <div className="mt-auto pt-6 flex justify-end">
                          <button onClick={() => setAnalysisStep('result')} className="w-full lg:w-auto bg-primary text-white py-2.5 px-6 rounded-md font-semibold hover:bg-primary-hover transition-colors duration-200 flex items-center justify-center space-x-2">
                            <span>生成我的穿搭</span>
                            <span className="material-icons-outlined text-xl">arrow_forward</span>
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                )}

                {analysisStep === 'result' && (
                  <div>
                    <div className="text-center mb-8">
                      <h2 className="text-2xl font-semibold mb-2 text-text-light-primary dark:text-text-dark-primary">我的镜像穿搭</h2>
                      <p className="text-text-light-secondary dark:text-text-dark-secondary max-w-2xl mx-auto">基于你的选择，这是为你生成的穿搭。你可以保存或进行AI试穿。</p>
                    </div>
                    <div className="grid grid-cols-1 lg:grid-cols-2 gap-12">
                      <div className="bg-background-light dark:bg-background-dark p-6 rounded-lg flex flex-col items-center justify-center">
                        <h3 className="font-semibold text-lg mb-6 text-text-light-primary dark:text-text-dark-primary">灵感来源</h3>
                        <div className="grid grid-cols-3 gap-2 w-full max-w-sm">
                          <img alt="Inspiration Top" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw"/>
                          <img alt="Inspiration Bottom" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDH-L7FdmL3xHhlueKM1mE__6QUrDRhY3Lpu69oN3tqwTZgTzqhjiMG4KKmHqm3hfhYFuE3IyfN4wqsbV4uXFepSq3L3VbONrlxEC2SPr1rSxTwDle1CL6nu92HiusbkGSQUU-h41V2a9UnJBIxHa1b65CaD5f04CLPvuSQ5Ua1zMRvMv5GB-17Oy7kw3oEsVihT9dgmO03x4zPGUnCEGbYpjcNsOD4sj38ph_VyANpYycuQSJgI6MYMLtNTevwKLM0iRuXyxpisQ"/>
                          <img alt="Inspiration Shoes" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDGzrvTvoWWB_5H4heTler19abv4xv0FsGLBMon5u3nr8duZlOsHKN-oc9Gt38i5wxIPV4cUWxuP9MXzGXsb88m8zh07YrOaTG1iPr7BwLH0rYJDe9UIdrZKXg2vgSUsO3bxAHlExprC0JNdQNabNnoZes8G3GSYK-XYB2tOvt8cPD_vRRBIJjIxtPDlVkA9QwsTFlxFuhqxcAHcx87m-b5bjoUpH4RGzUB0i6NDuE49NJ7dzGjjTrusd5HtjsZW65tc90XmD_Cnw"/>
                        </div>
                      </div>
                      <div className="bg-background-light dark:bg-background-dark p-6 rounded-lg flex flex-col items-center justify-center">
                        <h3 className="font-semibold text-lg mb-6 text-text-light-primary dark:text-text-dark-primary">我的衣橱复刻</h3>
                        <div className="grid grid-cols-3 gap-2 w-full max-w-sm">
                          <img alt="My Top" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw"/>
                          <img alt="My Bottom" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDH-L7FdmL3xHhlueKM1mE__6QUrDRhY3Lpu69oN3tqwTZgTzqhjiMG4KKmHqm3hfhYFuE3IyfN4wqsbV4uXFepSq3L3VbONrlxEC2SPr1rSxTwDle1CL6nu92HiusbkGSQUU-h41V2a9UnJBIxHa1b65CaD5f04CLPvuSQ5Ua1zMRvMv5GB-17Oy7kw3oEsVihT9dgmO03x4zPGUnCEGbYpjcNsOD4sj38ph_VyANpYycuQSJgI6MYMLtNTevwKLM0iRuXyxpisQ"/>
                          <img alt="My Shoes" className="rounded-md object-cover aspect-square" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDGzrvTvoWWB_5H4heTler19abv4xv0FsGLBMon5u3nr8duZlOsHKN-oc9Gt38i5wxIPV4cUWxuP9MXzGXsb88m8zh07YrOaTG1iPr7BwLH0rYJDe9UIdrZKXg2vgSUsO3bxAHlExprC0JNdQNabNnoZes8G3GSYK-XYB2tOvt8cPD_vRRBIJjIxtPDlVkA9QwsTFlxFuhqxcAHcx87m-b5bjoUpH4RGzUB0i6NDuE49NJ7dzGjjTrusd5HtjsZW65tc90XmD_Cnw"/>
                        </div>
                      </div>
                    </div>
                    <div className="mt-8 flex flex-col sm:flex-row justify-center gap-4">
                      <button className="w-full sm:w-auto bg-green-500 text-white py-2.5 px-6 rounded-md font-semibold hover:bg-green-600 transition-colors duration-200 flex items-center justify-center space-x-2">
                        <span className="material-icons-outlined text-xl">camera_alt</span>
                        <span>AI 试穿</span>
                      </button>
                      <button className="w-full sm:w-auto border border-border-light dark:border-border-dark text-text-light-secondary dark:text-text-dark-secondary py-2.5 px-6 rounded-md font-semibold hover:bg-background-light dark:hover:bg-background-dark hover:text-text-light-primary dark:hover:text-text-dark-primary transition-colors duration-200">
                        保存穿搭
                      </button>
                    </div>
                  </div>
                )}
              </div>
            </div>
          )}
        </main>

        {showOutfitModal && (
          <div className="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4">
            <div className="bg-card-light dark:bg-card-dark rounded-lg shadow-xl w-full max-w-4xl max-h-[90vh] flex flex-col">
              <div className="flex justify-between items-center p-4 border-b border-border-light dark:border-border-dark">
                <h3 className="text-lg font-semibold text-text-light-primary dark:text-text-dark-primary">从衣橱选择单品</h3>
                <button onClick={() => setShowOutfitModal(false)} className="text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary">
                  <span className="material-icons-outlined">close</span>
                </button>
              </div>
              <div className="p-6 overflow-y-auto">
                <div className="flex flex-wrap gap-4 mb-6">
                  <select className="rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark focus:ring-primary/50">
                    <option>所有类别</option>
                    <option>上衣</option>
                    <option>下装</option>
                    <option>鞋子</option>
                  </select>
                  <select className="rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark focus:ring-primary/50">
                    <option>所有季节</option>
                    <option>春</option>
                    <option>夏</option>
                    <option>秋</option>
                    <option>冬</option>
                  </select>
                  <div className="relative flex-grow">
                    <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
                    <input className="w-full bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-md py-2 pl-10 pr-4 focus:outline-none focus:ring-2 focus:ring-primary/50 text-text-light-primary dark:text-text-dark-primary" placeholder="搜索名称、品牌..." type="text"/>
                  </div>
                </div>
                <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-4">
                  <div onClick={() => setShowOutfitModal(false)} className="group relative cursor-pointer">
                    <div className="aspect-w-1 aspect-h-1 w-full overflow-hidden rounded-md bg-background-light dark:bg-background-dark">
                      <img alt="简约白T" className="h-full w-full object-cover object-center group-hover:opacity-75 transition-opacity" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw"/>
                    </div>
                    <h4 className="mt-2 text-sm text-text-light-primary dark:text-text-dark-primary font-medium">简约白T</h4>
                    <p className="text-xs text-text-light-secondary dark:text-text-dark-secondary">上装, 夏</p>
                  </div>
                  <div onClick={() => setShowOutfitModal(false)} className="group relative cursor-pointer">
                    <div className="aspect-w-1 aspect-h-1 w-full overflow-hidden rounded-md bg-background-light dark:bg-background-dark">
                      <img alt="直筒牛仔裤" className="h-full w-full object-cover object-center group-hover:opacity-75 transition-opacity" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDH-L7FdmL3xHhlueKM1mE__6QUrDRhY3Lpu69oN3tqwTZgTzqhjiMG4KKmHqm3hfhYFuE3IyfN4wqsbV4uXFepSq3L3VbONrlxEC2SPr1rSxTwDle1CL6nu92HiusbkGSQUU-h41V2a9UnJBIxHa1b65CaD5f04CLPvuSQ5Ua1zMRvMv5GB-17Oy7kw3oEsVihT9dgmO03x4zPGUnCEGbYpjcNsOD4sj38ph_VyANpYycuQSJgI6MYMLtNTevwKLM0iRuXyxpisQ"/>
                    </div>
                    <h4 className="mt-2 text-sm text-text-light-primary dark:text-text-dark-primary font-medium">直筒牛仔裤</h4>
                    <p className="text-xs text-text-light-secondary dark:text-text-dark-secondary">下装, 春秋</p>
                  </div>
                  <div onClick={() => setShowOutfitModal(false)} className="group relative cursor-pointer">
                    <div className="aspect-w-1 aspect-h-1 w-full overflow-hidden rounded-md bg-background-light dark:bg-background-dark">
                      <img alt="白色运动鞋" className="h-full w-full object-cover object-center group-hover:opacity-75 transition-opacity" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDGzrvTvoWWB_5H4heTler19abv4xv0FsGLBMon5u3nr8duZlOsHKN-oc9Gt38i5wxIPV4cUWxuP9MXzGXsb88m8zh07YrOaTG1iPr7BwLH0rYJDe9UIdrZKXg2vgSUsO3bxAHlExprC0JNdQNabNnoZes8G3GSYK-XYB2tOvt8cPD_vRRBIJjIxtPDlVkA9QwsTFlxFuhqxcAHcx87m-b5bjoUpH4RGzUB0i6NDuE49NJ7dzGjjTrusd5HtjsZW65tc90XmD_Cnw"/>
                    </div>
                    <h4 className="mt-2 text-sm text-text-light-primary dark:text-text-dark-primary font-medium">白色运动鞋</h4>
                    <p className="text-xs text-text-light-secondary dark:text-text-dark-secondary">鞋履, 全季</p>
                  </div>
                </div>
              </div>
              <div className="p-4 border-t border-border-light dark:border-border-dark flex justify-end">
                <button onClick={() => setShowOutfitModal(false)} className="bg-primary text-white px-4 py-2 rounded-md font-semibold hover:bg-primary-hover">确认选择</button>
              </div>
            </div>
          </div>
        )}

        <Footer />
      </div>
    </>
  );
}
