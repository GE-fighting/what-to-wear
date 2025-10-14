'use client';

import React, { useState } from 'react';

export default function WardrobePage() {
  const [activeTab, setActiveTab] = useState<'list' | 'overview'>('list');
  const [isFilterOpen, setIsFilterOpen] = useState(false);

  return (
    <div className="min-h-screen flex flex-col bg-background-light dark:bg-background-dark font-display">
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
              <div className="relative">
                <button className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800">
                  <span className="material-icons-outlined">notifications_none</span>
                </button>
              </div>
              <button>
                <img alt="User avatar" className="h-8 w-8 rounded-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL" />
              </button>
            </div>
          </div>
        </div>
      </header>

      <main className="flex-grow max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 w-full py-8">
        <div className="border-b border-border-light dark:border-border-dark mb-6">
          <nav className="-mb-px flex space-x-8">
            <button
              onClick={() => setActiveTab('list')}
              className={`py-4 px-1 border-b-2 border-transparent font-medium text-sm transition-colors ${
                activeTab === 'list'
                  ? 'tab-active'
                  : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'
              }`}
            >
              衣物列表
            </button>
            <button
              onClick={() => setActiveTab('overview')}
              className={`py-4 px-1 border-b-2 border-transparent font-medium text-sm transition-colors ${
                activeTab === 'overview'
                  ? 'tab-active'
                  : 'text-text-light-secondary dark:text-text-dark-secondary hover:text-text-light-primary dark:hover:text-text-dark-primary'
              }`}
            >
              衣橱分析
            </button>
          </nav>
        </div>

        {activeTab === 'list' && (
          <div>
            <div className="flex justify-end items-center mb-6">
              <div className="flex items-center space-x-4">
                <a href="/main/wardrobe/add" className="bg-primary text-white px-4 py-2 rounded-lg flex items-center space-x-2 hover:bg-primary-hover transition-colors dark:bg-text-dark-primary dark:text-background-dark dark:hover:bg-gray-200">
                  <span className="material-icons-outlined text-base">add</span>
                  <span>添加衣物</span>
                </a>
                <button className="bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark px-4 py-2 rounded-lg text-sm text-text-light-primary dark:text-text-dark-primary hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors">查看全部</button>
              </div>
            </div>

            <div className="mb-6 bg-card-light dark:bg-card-dark p-4 rounded-lg shadow-sm flex flex-wrap items-center justify-between gap-4">
              <div className="flex items-center space-x-4">
                <div className="relative">
                  <select className="appearance-none bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-md py-2 pl-3 pr-8 text-sm text-text-light-primary dark:text-text-dark-primary focus:ring-primary focus:border-primary">
                    <option>所有类别</option>
                    <option>上装</option>
                    <option>下装</option>
                    <option>外套</option>
                    <option>鞋履</option>
                    <option>配饰</option>
                  </select>
                  <span className="material-icons-outlined absolute right-2 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary text-base pointer-events-none">expand_more</span>
                </div>
                <div className="relative">
                  <select className="appearance-none bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-md py-2 pl-3 pr-8 text-sm text-text-light-primary dark:text-text-dark-primary focus:ring-primary focus:border-primary">
                    <option>所有季节</option>
                    <option>春</option>
                    <option>夏</option>
                    <option>秋</option>
                    <option>冬</option>
                  </select>
                  <span className="material-icons-outlined absolute right-2 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary text-base pointer-events-none">expand_more</span>
                </div>
                <button onClick={() => setIsFilterOpen(true)} className="text-sm text-text-light-secondary dark:text-text-dark-secondary hover:text-primary flex items-center space-x-1">
                  <span className="material-icons-outlined text-base">filter_list</span>
                  <span>更多筛选</span>
                </button>
              </div>
              <div className="flex items-center space-x-2">
                <div className="relative">
                  <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
                  <input className="w-full max-w-xs bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm focus:ring-primary focus:border-primary text-text-light-primary dark:text-text-dark-primary" placeholder="搜索名称、品牌、标签..." type="text" />
                </div>
                <button className="p-2 rounded-full text-white bg-primary dark:bg-text-dark-primary dark:text-background-dark border border-transparent">
                  <span className="material-icons-outlined">grid_view</span>
                </button>
                <button className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800">
                  <span className="material-icons-outlined">view_list</span>
                </button>
              </div>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-6">
              {/* 简约白T */}
              <div className="group relative overflow-hidden rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 bg-card-light dark:bg-card-dark">
                <button className="absolute top-2 right-2 z-10 bg-white/80 dark:bg-black/80 p-2 rounded-full hover:bg-white dark:hover:bg-black transition-colors" aria-label="编辑衣物">
                  <a href="/main/wardrobe/edit" className="flex items-center justify-center">
                    <span className="material-icons-outlined text-sm text-text-light-primary dark:text-text-dark-primary">edit</span>
                  </a>
                </button>
                <a href="/main/wardrobe/detail" className="block group relative overflow-hidden rounded-lg">
                  <div className="w-full overflow-hidden aspect-square">
                    <img alt="Green T-shirt" className="w-full h-full object-contain p-8 transform group-hover:scale-105 transition-transform duration-300" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBpKRSkI0k_rv3-f3pkIHx8gSRFlXMKWkz5pMMTOl91whfAGmxZ7W_Wvp_51sCXqPS05Ghwswh09y79iYU2YbqzgVI4qFgBJ2Mczj3QLPKX7k1sxRZ3sMniVjXctOedrgdGbFGiCcgRUuiEp4TgBZ2UcKsMvm9ii-cWu6FEeSZ_3WqYfeCv_IZ8HWbm70HZy2d4gHO0jqVJ3ZoyD639pZGAXsiMRunOgs_pq5LK-DtbU-il8n_jKcAe8zuVzcacFavxvNxlkCyNSpRJ" />
                  </div>
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <div className="absolute bottom-0 left-0 p-4 w-full">
                    <h3 className="font-bold text-white text-lg">简约白T</h3>
                    <p className="text-sm text-gray-300">上装, 夏</p>
                    <div className="flex items-center justify-between mt-2">
                      <div className="flex space-x-2">
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">通勤</span>
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">百搭</span>
                      </div>
                      <span className="text-xs text-gray-300">穿着 8次</span>
                    </div>
                  </div>
                </a>
              </div>

              {/* 直筒牛仔裤 */}
              <div className="group relative overflow-hidden rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 bg-card-light dark:bg-card-dark">
                <button className="absolute top-2 right-2 z-10 bg-white/80 dark:bg-black/80 p-2 rounded-full hover:bg-white dark:hover:bg-black transition-colors" aria-label="编辑衣物">
                  <a href="/main/wardrobe/edit" className="flex items-center justify-center">
                    <span className="material-icons-outlined text-sm text-text-light-primary dark:text-text-dark-primary">edit</span>
                  </a>
                </button>
                <a href="/main/wardrobe/detail" className="block group relative overflow-hidden rounded-lg">
                  <div className="w-full overflow-hidden aspect-square">
                    <img alt="Blue Jeans" className="w-full h-full object-contain p-8 transform group-hover:scale-105 transition-transform duration-300" src="https://lh3.googleusercontent.com/aida-public/AB6AXuC1CcPAuYZDAzOoS09aTYninYVkA0F2bZ6MQS1XvIVykKnhdjXc3ILR5f1bddH6AKyZii9O-OWLbeYZdr5JF5c0jgotXh3R0Hduw2ZjFcjneACgQHUmJPTsB-w-bezs3EzNP9D6lDQIA3d5eCo1LWojxp7-mkYLjJme5Xrah723APDeA0vX2eD6n2vS4l3r4FQBpYvXVRPkp0ZzRO0f70ezn9BaACBMIfvjgDn-THgQ-sLvJjcAijZoDaqZ2xj2qpx16du5mQTaOANO" />
                  </div>
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <div className="absolute bottom-0 left-0 p-4 w-full">
                    <h3 className="font-bold text-white text-lg">直筒牛仔裤</h3>
                    <p className="text-sm text-gray-300">下装, 春秋</p>
                    <div className="flex items-center justify-between mt-2">
                      <div className="flex space-x-2">
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">休闲</span>
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">耐穿</span>
                      </div>
                      <span className="text-xs text-gray-300">穿着 12次</span>
                    </div>
                  </div>
                </a>
              </div>

              {/* 机能风冲锋衣 */}
              <div className="group relative overflow-hidden rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 bg-card-light dark:bg-card-dark">
                <button className="absolute top-2 right-2 z-10 bg-white/80 dark:bg-black/80 p-2 rounded-full hover:bg-white dark:hover:bg-black transition-colors" aria-label="编辑衣物">
                  <a href="/main/wardrobe/edit" className="flex items-center justify-center">
                    <span className="material-icons-outlined text-sm text-text-light-primary dark:text-text-dark-primary">edit</span>
                  </a>
                </button>
                <a href="/main/wardrobe/detail" className="block group relative overflow-hidden rounded-lg">
                  <div className="w-full overflow-hidden aspect-square">
                    <img alt="Beige Jacket" className="w-full h-full object-contain p-8 transform group-hover:scale-105 transition-transform duration-300" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAkHfG8nBe7u5ky9SU6swqaWeW1-vlQh4NGnzLV5wcO1HGZ8UrMgwB1I8etuA48UQjqQTaC7FyDzurm8Yg_tjUXEXX0K9HqsQk4J16OZbuegSNbXf35h2CzqYrbzyReOOQQeD5llB_Kxkk_FzpBnnFtvBm6D9IfD9JV4vIepeh28KsFUbiffkw7PAya8M5RvMmByPuoK1cGq4zL6mSQqiZaJVIU67cpamB-8RDI3B3pv7kb0_HOsBXODmEfF0y4IEcft4h388KEs9Yd" />
                  </div>
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <div className="absolute bottom-0 left-0 p-4 w-full">
                    <h3 className="font-bold text-white text-lg">机能风冲锋衣</h3>
                    <p className="text-sm text-gray-300">外套, 春秋</p>
                    <div className="flex items-center justify-between mt-2">
                      <div className="flex space-x-2">
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">户外</span>
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">防风</span>
                      </div>
                      <span className="text-xs text-gray-300">穿着 6次</span>
                    </div>
                  </div>
                </a>
              </div>

              {/* 白色运动鞋 */}
              <div className="group relative overflow-hidden rounded-lg shadow-md hover:shadow-xl transition-shadow duration-300 bg-card-light dark:bg-card-dark">
                <button className="absolute top-2 right-2 z-10 bg-white/80 dark:bg-black/80 p-2 rounded-full hover:bg-white dark:hover:bg-black transition-colors" aria-label="编辑衣物">
                  <a href="/main/wardrobe/edit" className="flex items-center justify-center">
                    <span className="material-icons-outlined text-sm text-text-light-primary dark:text-text-dark-primary">edit</span>
                  </a>
                </button>
                <a href="/main/wardrobe/detail" className="block group relative overflow-hidden rounded-lg">
                  <div className="w-full overflow-hidden aspect-square">
                    <img alt="White Sneakers" className="w-full h-full object-contain p-8 transform group-hover:scale-105 transition-transform duration-300" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBDGHZOG1AVpehpXck7NDvsGZfMo9PMolJN4mPrh3Qw9f47KWQniwwoRyFEOOoIZ-MmJzuHCGjaW1eUC5ix2wdCz5t2taSiXbrfHTGnRrEBKZaX4zL3vYa5ttzh0MA_FOYAiW9kDO6s8SRWVfCwAlTjeHcu5SoAsMlbcOjgkizcEMVl4Q8PPe7wfF3rl0UPRmpa23iHIaKr8jPVC3OjxAjqodqcwh5slG7xQLzUgbvkYNLhHp5IAvuDlHHHdAySC0o0vN0sxYWXQ5gf" />
                  </div>
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <div className="absolute bottom-0 left-0 p-4 w-full">
                    <h3 className="font-bold text-white text-lg">白色运动鞋</h3>
                    <p className="text-sm text-gray-300">鞋履, 全季</p>
                    <div className="flex items-center justify-between mt-2">
                      <div className="flex space-x-2">
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">百搭</span>
                        <span className="text-xs bg-gray-700/50 text-white py-1 px-2 rounded-full">舒适</span>
                      </div>
                      <span className="text-xs text-gray-300">穿着 22次</span>
                    </div>
                  </div>
                </a>
              </div>
            </div>
          </div>
        )}

        {activeTab === 'overview' && (
          <div>
            <div className="flex flex-wrap gap-4 mb-6">
              {[
                ['总衣物数', '250'],
                ['总价值', '¥15,000'],
                ['待清洗', '25'],
                ['本周新增', '5'],
                ['30天未穿', '50'],
                ['平均单价', '¥60'],
              ].map(([label, value]) => (
                <div key={label} className="flex min-w-[158px] flex-1 flex-col gap-2 rounded-lg p-6 border border-border-light dark:border-border-dark bg-card-light dark:bg-card-dark">
                  <p className="text-text-light-primary dark:text-text-dark-primary text-base font-medium leading-normal">{label}</p>
                  <p className="text-text-light-primary dark:text-text-dark-primary tracking-light text-2xl font-bold leading-tight">{value}</p>
                </div>
              ))}
            </div>

            <h2 className="text-xl font-bold leading-tight text-text-light-primary dark:text-text-dark-primary pb-3 pt-5">最常穿衣物</h2>
            <div className="py-3">
              <div className="flex overflow-hidden rounded-lg border border-border-light dark:border-border-dark bg-card-light dark:bg-card-dark">
                <table className="flex-1">
                  <thead>
                    <tr className="bg-card-light dark:bg-card-dark">
                      <th className="px-4 py-3 text-left text-text-light-primary dark:text-text-dark-primary w-14 text-sm font-medium leading-normal">衣物</th>
                      <th className="px-4 py-3 text-left text-text-light-primary dark:text-text-dark-primary w-[400px] text-sm font-medium leading-normal">类别</th>
                      <th className="px-4 py-3 text-left text-text-light-primary dark:text-text-dark-primary w-[400px] text-sm font-medium leading-normal">穿着次数</th>
                      <th className="px-4 py-3 text-left text-text-light-primary dark:text-text-dark-primary w-[400px] text-sm font-medium leading-normal">最后穿着</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr className="border-t border-t-border-light dark:border-t-border-dark">
                      <td className="h-[72px] px-4 py-2 w-14 text-sm font-normal leading-normal">
                        <div className="bg-center bg-no-repeat aspect-square bg-cover rounded-full w-10" style={{ backgroundImage: "url('https://lh3.googleusercontent.com/aida-public/AB6AXuBA9qc8f02Eb6pBlPnZ5hjd6ahswtmo-54Vi_A-_DoiLttKhpLw52_eCG9vCzw204Kuk2YiLrA927morLcz5X24_mq6djakfPIxx6Tbty7YMcfwqAwnqvOhN8F8xiGCeEh-4O1G17oh_UrkedeWe94zHUqWfGo7-LMdiFDg2RTljtmo3o4TgyT5hoLNcR7UI55u4C2Obn9ZPcnNjNalQ2-06NDZvkuGmDbybWHZkAfxXkjJhhw3v57or4LeEoJMpLnrokIHuBd4If4Q')" }}></div>
                      </td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-primary dark:text-text-dark-primary text-sm font-normal leading-normal">牛仔裤</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">30</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">2周前</td>
                    </tr>
                    <tr className="border-t border-t-border-light dark:border-t-border-dark">
                      <td className="h-[72px] px-4 py-2 w-14 text-sm font-normal leading-normal">
                        <div className="bg-center bg-no-repeat aspect-square bg-cover rounded-full w-10" style={{ backgroundImage: "url('https://lh3.googleusercontent.com/aida-public/AB6AXuApOC0h0aJyNcT5K8inpPoB1vvNVJAj1NwCg0vinH5PvwCE0UvVJbq06w19puQhHn3Yb4SJli7P5Hy62k65r1IWu4cJVcR8dtlxhLw2P1bfs4f811AL7n4ScKCGhZ4EJTFQ-RmJuzf5O0_2BNuxnMZ8CiXrkymqHF4WiYpoajHq0imSmWEyLBnP9c-j8rfLSdEPR26iyXQLg8r7RCnTqD2VSO0eCXXau-TfLIC6Og_ciImaaYEJMvznTnmwr6hUh0xY4p1cx8uYC5gK')" }}></div>
                      </td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-primary dark:text-text-dark-primary text-sm font-normal leading-normal">T恤</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">25</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">1周前</td>
                    </tr>
                    <tr className="border-t border-t-border-light dark:border-t-border-dark">
                      <td className="h-[72px] px-4 py-2 w-14 text-sm font-normal leading-normal">
                        <div className="bg-center bg-no-repeat aspect-square bg-cover rounded-full w-10" style={{ backgroundImage: "url('https://lh3.googleusercontent.com/aida-public/AB6AXuCOehIoB8rdT88gyyfJl47f8TcCbMLnqSO9wdb975Wf4DEGf9qjhPWEUDSf5t1-8TsLdJYeygs3at1s4WvW1xy14rzdZHqzISlhi2pSnaREdhG_fG0bLggItR1-CRouluBb894-nmAb_GNR9KA1nShB6xQVOYeoymoG9r6PTjI44eMZuqTY9-Ae7ab1tbT38nbDeJkaW6I2BvFQHCB9ljng8MFANSsCNSDuaNd_fq6aYud__vWzNDtK32bYQSIvMtsGkNgbcUjcyYuQ')" }}></div>
                      </td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-primary dark:text-text-dark-primary text-sm font-normal leading-normal">夹克</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">20</td>
                      <td className="h-[72px] px-4 py-2 w-[400px] text-text-light-secondary dark:text-text-dark-secondary text-sm font-normal leading-normal">3周前</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <h2 className="text-xl font-bold leading-tight text-text-light-primary dark:text-text-dark-primary pb-3 pt-5">类别分布</h2>
            <div className="flex flex-wrap gap-4 py-6">
              <div className="flex min-w-72 flex-1 flex-col gap-2 rounded-lg border border-border-light dark:border-border-dark p-6 bg-card-light dark:bg-card-dark">
                <p className="text-text-light-primary dark:text-text-dark-primary text-base font-medium leading-normal">类别分布</p>
                <div className="grid min-h-[180px] gap-x-4 gap-y-6 grid-cols-[auto_1fr] items-center py-3">
                  {[
                    ['上装', '80%'],
                    ['下装', '30%'],
                    ['外套', '40%'],
                    ['鞋履', '50%'],
                    ['配饰', '50%'],
                  ].map(([label, width]) => (
                    <React.Fragment key={label}>
                      <p className="text-text-light-secondary dark:text-text-dark-secondary text-[13px] font-bold leading-normal tracking-[0.015em]">{label}</p>
                      <div className="h-full flex-1"><div className="border-primary bg-background-light dark:bg-background-dark border-r-2 h-full" style={{ width }}></div></div>
                    </React.Fragment>
                  ))}
                </div>
              </div>
            </div>
          </div>
        )}
      </main>

      {/* 遮罩层 */}
      {isFilterOpen && (
        <div className="fixed inset-0 bg-black bg-opacity-50 z-30" onClick={() => setIsFilterOpen(false)} />
      )}

      {/* 侧边筛选 */}
      <aside className={`fixed top-0 right-0 h-full w-full max-w-sm bg-card-light dark:bg-card-dark shadow-lg z-40 flex flex-col no-scrollbar transform transition-transform duration-300 ${
        isFilterOpen ? 'translate-x-0' : 'translate-x-full'
      }`}>
        <div className="flex items-center justify-between p-6 border-b border-border-light dark:border-border-dark">
          <h3 className="text-xl font-semibold text-text-light-primary dark:text-text-dark-primary">高级筛选</h3>
          <button onClick={() => setIsFilterOpen(false)} className="p-2 rounded-full text-text-light-secondary dark:text-text-dark-secondary hover:bg-gray-100 dark:hover:bg-gray-800">
            <span className="material-icons-outlined">close</span>
          </button>
        </div>
        <div className="flex-grow p-6 space-y-8 overflow-y-auto no-scrollbar">
          <div>
            <h4 className="font-semibold mb-3 text-text-light-primary dark:text-text-dark-primary">类别</h4>
            <div className="grid grid-cols-3 gap-2">
              {['所有', '上装', '下装', '外套', '鞋履', '配饰'].map((it, idx) => (
                <button key={it} className={`px-3 py-2 text-sm rounded-lg text-center ${idx === 0 ? 'bg-primary text-white dark:bg-text-dark-primary dark:text-background-dark' : 'bg-background-light dark:bg-background-dark text-text-light-primary dark:text-text-dark-primary border border-border-light dark:border-border-dark hover:border-primary'}`}>{it}</button>
              ))}
            </div>
          </div>
          <div>
            <h4 className="font-semibold mb-3 text-text-light-primary dark:text-text-dark-primary">季节</h4>
            <div className="flex space-x-2">
              {['春', '夏', '秋', '冬'].map((s, idx) => (
                <button key={s} className={`flex-1 px-3 py-2 text-sm rounded-lg text-center ${idx === 1 ? 'bg-primary text-white dark:bg-text-dark-primary dark:text-background-dark' : 'bg-background-light dark:bg-background-dark text-text-light-primary dark:text-text-dark-primary border border-border-light dark:border-border-dark hover:border-primary'}`}>{s}</button>
              ))}
            </div>
          </div>
          <div>
            <h4 className="font-semibold mb-3 text-text-light-primary dark:text-text-dark-primary">颜色</h4>
            <div className="flex flex-wrap gap-3">
              <button className="w-8 h-8 rounded-full bg-black border-2 border-primary ring-2 ring-offset-2 ring-primary dark:ring-offset-background-dark"></button>
              <button className="w-8 h-8 rounded-full bg-white border border-border-light dark:border-border-dark"></button>
              <button className="w-8 h-8 rounded-full bg-gray-400"></button>
              <button className="w-8 h-8 rounded-full bg-red-500"></button>
              <button className="w-8 h-8 rounded-full bg-blue-500"></button>
              <button className="w-8 h-8 rounded-full bg-green-500"></button>
              <button className="w-8 h-8 rounded-full bg-yellow-400"></button>
              <button className="w-8 h-8 rounded-full bg-purple-500"></button>
              <button className="w-8 h-8 flex items-center justify-center rounded-full bg-gradient-to-br from-red-500 to-blue-500 text-white">
                <span className="material-icons-outlined text-sm">palette</span>
              </button>
            </div>
          </div>
          <div>
            <h4 className="font-semibold mb-3 text-text-light-primary dark:text-text-dark-primary">品牌</h4>
            <div className="relative">
              <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
              <input className="w-full bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-lg py-2 pl-10 pr-4 text-sm focus:ring-primary focus:border-primary text-text-light-primary dark:text-text-dark-primary" placeholder="搜索品牌" type="text" />
            </div>
            <div className="mt-3 space-y-2 max-h-32 overflow-y-auto no-scrollbar">
              {['无品牌', 'Uniqlo', 'Nike', 'Zara'].map((b, idx) => (
                <label key={b} className="flex items-center space-x-3 cursor-pointer">
                  <input defaultChecked={b === 'Uniqlo'} className="rounded text-primary focus:ring-primary dark:bg-background-dark dark:border-border-dark" type="checkbox" />
                  <span className="text-sm text-text-light-primary dark:text-text-dark-primary">{b}</span>
                </label>
              ))}
            </div>
          </div>
          <div>
            <h4 className="font-semibold mb-4 text-text-light-primary dark:text-text-dark-primary">穿着频率</h4>
            <div className="relative">
              <input className="w-full h-2 bg-background-light dark:bg-background-dark rounded-lg appearance-none cursor-pointer accent-primary" max={50} min={0} type="range" defaultValue={10} />
              <div className="flex justify-between text-xs text-text-light-secondary dark:text-text-dark-secondary mt-2">
                <span>0次</span>
                <span>25次</span>
                <span>50+次</span>
              </div>
            </div>
          </div>
        </div>
        <div className="p-6 border-t border-border-light dark:border-border-dark bg-card-light dark:bg-card-dark">
          <div className="flex space-x-4">
            <button className="flex-1 px-4 py-2 text-sm rounded-lg border border-border-light dark:border-border-dark text-text-light-primary dark:text-text-dark-primary bg-card-light dark:bg-card-dark hover:bg-gray-50 dark:hover:bg-gray-700">重置</button>
            <button onClick={() => setIsFilterOpen(false)} className="flex-1 px-4 py-2 text-sm rounded-lg bg-primary text-white dark:bg-text-dark-primary dark:text-background-dark hover:bg-primary-hover">应用</button>
          </div>
        </div>
      </aside>

      <footer className="bg-card-light dark:bg-card-dark mt-auto">
        <div className="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8">
          <p className="text-center text-sm text-text-light-secondary dark:text-text-dark-secondary">@2024 StyleSense. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}
