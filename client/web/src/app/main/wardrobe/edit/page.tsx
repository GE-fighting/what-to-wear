'use client';

export default function WardrobeEditPage() {
  return (
    <div className="min-h-screen flex flex-col bg-background-light dark:bg-background-dark">
      <header className="bg-card-light dark:bg-card-dark shadow-sm">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">StyleSense</h1>
              <nav className="hidden md:flex space-x-8">
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="/record-style/my-outfit">记录穿搭</a>
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="/main/wardrobe">我的衣橱</a>
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">风格灵感</a>
                <a className="text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary hover:text-primary" href="#">穿搭分析</a>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative">
                <span className="material-icons absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">search</span>
                <input className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm text-text-light-primary dark:text-text-dark-primary focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent" placeholder="搜索衣物" type="text" />
              </div>
              <button className="text-text-light-secondary dark:text-text-dark-secondary hover:text-primary">
                <span className="material-icons">notifications_none</span>
              </button>
              <button className="text-text-light-secondary dark:text-text-dark-secondary hover:text-primary">
                <span className="material-icons">favorite_border</span>
              </button>
              <img alt="User avatar" className="w-8 h-8 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuDB2hXkp-JgqjfVeRvnaUKkI0-DAM3BOkXy4WToB6rhOQgqu3zgJdjpeJ4zY47-fRqDE4FFYulRATIEzEb19Nk9sWowqQ8uFd1s9yEJi8Jh_sE_fiUQE0X4rqOdlPGvrzsoD5OEZ58IqMU8pdUikB06l1Phd7Au2jMKUVL0hG2I-oe-RSf8uhLAjPF6IqcV6szP6MQuxhc3d5wIz1hRFEOFuX5Q5v68NRFYoaRlsi9OcAaqSujJITyzdmzjVSzXAzST8CleGE-CaAyM" />
            </div>
          </div>
        </div>
      </header>

      <main className="flex-grow container mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-text-light-primary dark:text-text-dark-primary">编辑衣物</h2>
          <div className="flex space-x-4">
            <a href="/main/wardrobe/detail" className="px-4 py-2 text-sm font-medium text-text-light-secondary dark:text-text-secondary-dark bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-md hover:bg-gray-50 dark:hover:bg-gray-700">取消</a>
            <button className="px-4 py-2 text-sm font-medium text-white bg-black rounded-md hover:bg-gray-800 flex items-center space-x-2">
              <span className="material-icons text-base">save</span>
              <span>保存衣物</span>
            </button>
          </div>
        </div>

        <div className="space-y-8">
          <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm">
            <h3 className="text-lg font-medium text-text-light-primary dark:text-text-dark-primary mb-4">衣物图片</h3>
            <div className="grid grid-cols-2 sm:grid-cols-4 md:grid-cols-6 gap-4">
              <div className="relative aspect-square border-2 border-dashed border-border-light dark:border-border-dark rounded-lg flex items-center justify-center">
                <span className="material-icons text-text-light-secondary dark:text-text-dark-secondary">add_photo_alternate</span>
              </div>
              <div className="aspect-square rounded-lg overflow-hidden">
                <img alt="Front" className="w-full h-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuCY2Nw73oScRhrbxbU0mXQ-D63FxrAHoK3cM_KC_uvRNHBBeXuGuly6ZgSeU1j8GvNrkw6zmpz7stsXMMJiPcm4XT-D86pdJD7s_QI1Rb9Hc4yAg1RqOitSJBHa3gnCU4GJ9Jq31svy5AnZVexrCQj7xYwRAaRhy3xxjk-i06t55rNcdmXD2Emt6uNX1b5YhbUIgZJt2yJupUgQNFIV18Ypf0nuiRdkVaJUc8hdTw-gC4lZH_BIASHjMvU_vU-SpJaOPRUIsfH6dbjl" />
              </div>
              <div className="aspect-square rounded-lg overflow-hidden">
                <img alt="Back" className="w-full h-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuAhww5X1RVqQN8csE48lPcAc-HO1vsUOhABHNvn45Ohd8RGy5b9UYip4s32hRHdCw6YxLAVCCnWIb0A0ZIDbgnsXhd-ILgD_ttGsR-TPR7yTOAJyVHgsjtX-2CfkMCsPGnIgwNriNKaJc5bD0zmvUgvUX3OzD_L_mDqt22a533f-mUuwJSCF-cHFRdnkCWF5WhZ83Oo78qG3WVWB_5SZuOqNWxeFsFbZMXTQ0B8YMEN1rsFc1W1lAqI4cq3sFDDMBHP1RQKeCAeA7Hk" />
              </div>
              <div className="aspect-square rounded-lg overflow-hidden">
                <img alt="Fabric" className="w-full h-full object-cover" src="https://lh3.googleusercontent.com/aida-public/AB6AXuCaSsedotpw-dgqAmU9ChWvPZOuY83-NmfXnZb88Bs2XjN-ufhGmbPWCU22CXHPm18EYym59fz8cnIoHVd3aqqYMB22Kcp2n6dJB0EAoHkPrwGrohwIimF8Mhckc_b9q9nfgDNL_626lr-5FKKDtgFOkIAQFAZkkePiFP0DE8OvNTLm5mYXBLn8HXT0gfKcBViJqw-O0jzKYKZasVtUcHDUVtkz0WkgIy4RH4bdfTNvn9J50AUWK5o-SunKmt9nusQgi2CR_bi-4RCF" />
              </div>
              <div className="relative aspect-square border-2 border-dashed border-border-light dark:border-border-dark rounded-lg flex items-center justify-center text-text-light-secondary dark:text-text-dark-secondary">
                <div className="text-center">
                  <span className="material-icons">add_a_photo</span>
                  <p className="text-xs mt-1">添加图片</p>
                </div>
              </div>
            </div>
          </div>

          <div className="bg-green-50 dark:bg-green-900/20 p-6 rounded-lg">
            <h3 className="text-lg font-medium text-green-800 dark:text-green-300 mb-4 flex items-center">
              <span className="material-icons mr-2">auto_awesome</span>
              AI 标签建议
            </h3>
            <p className="text-sm text-green-700 dark:text-green-400 mb-4">根据您上传的图片，我们为您推荐了以下标签。您可以直接使用或进行修改。</p>
            <div className="flex flex-wrap gap-2 mb-4">
              {['颜色: 黑色', '款式: T恤', '材质: 棉', '风格: 简约'].map((t) => (
                <span key={t} className="inline-flex items-center px-2 py-1 bg-card-light dark:bg-card-dark text-sm font-medium text-text-light-primary dark:text-text-dark-primary rounded-full border border-border-light dark:border-border-dark">{t}
                  <button className="ml-1.5 text-text-light-secondary hover:text-text-light-primary dark:text-text-dark-secondary dark:hover:text-text-dark-primary">
                    <span className="material-icons text-sm">close</span>
                  </button>
                </span>
              ))}
            </div>
            <div className="flex space-x-4">
              <button className="px-4 py-2 text-sm font-medium text-white bg-green-600 rounded-md hover:bg-green-700 flex items-center space-x-2">
                <span className="material-icons text-base">check</span>
                <span>全部应用</span>
              </button>
              <button className="px-4 py-2 text-sm font-medium text-green-700 dark:text-green-300 bg-green-100 dark:bg-green-900/30 border border-green-200 dark:border-green-700 rounded-md hover:bg-green-200 dark:hover:bg-green-900/50 flex items-center space-x-2">
                <span className="material-icons text-base">edit</span>
                <span>编辑</span>
              </button>
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm space-y-6">
              <h3 className="text-lg font-medium text-text-light-primary dark:text-text-dark-primary">基本信息</h3>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
                {[
                  ['衣物名称', '黑色T恤'],
                  ['分类', '上装'],
                  ['品牌', 'Uniqlo'],
                  ['颜色', '黑色'],
                  ['材质', '棉'],
                  ['风格', '简约'],
                  ['尺码', 'M / 170'],
                ].map(([label, value]) => (
                  <div key={label} className="sm:col-span-1">
                    <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">{label}</label>
                    <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" defaultValue={value} />
                  </div>
                ))}
              </div>
            </div>
            <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm space-y-6">
              <h3 className="text-lg font-medium text-text-light-primary dark:text-text-dark-primary">适用场景</h3>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">季节</label>
                  <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" defaultValue="春夏" />
                </div>
                <div>
                  <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">场合</label>
                  <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" defaultValue="日常通勤, 周末出游" />
                </div>
              </div>
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm space-y-6">
              <h3 className="text-lg font-medium text-text-light-primary dark:text-text-dark-primary">购买信息</h3>
              <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
                <div>
                  <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">价格</label>
                  <div className="relative mt-1 rounded-md shadow-sm">
                    <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                      <span className="text-text-light-secondary dark:text-text-dark-secondary sm:text-sm">¥</span>
                    </div>
                    <input className="block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark pl-7 pr-12 focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" defaultValue="0.00" />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">商店</label>
                  <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" placeholder="例如: 淘宝官方旗舰店" />
                </div>
                <div className="sm:col-span-2">
                  <label className="block text-sm font-medium text-text-light-secondary dark:text-text-dark-secondary">购买日期</label>
                  <div className="relative mt-1">
                    <input className="block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm pl-3 pr-10 text-text-light-primary dark:text-text-dark-primary" placeholder="mm/dd/yyyy" />
                    <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                      <span className="material-icons text-text-light-secondary dark:text-text-dark-secondary">calendar_today</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="bg-card-light dark:bg-card-dark p-6 rounded-lg shadow-sm space-y-6">
              <h3 className="text-lg font-medium text-text-light-primary dark:text-text-dark-primary">其他备注</h3>
              <textarea className="block w-full rounded-md border-border-light dark:border-border-dark bg-background-light dark:bg-background-dark shadow-sm focus:border-primary focus:ring-primary sm:text-sm text-text-light-primary dark:text-text-dark-primary" rows={4} placeholder="可以记录洗护方式、搭配建议等..." />
            </div>
          </div>

          <div className="flex justify-end space-x-4 mt-8">
            <a href="/main/wardrobe/detail" className="px-4 py-2 text-sm font-medium text-text-light-secondary dark:text-text-secondary-dark bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-md hover:bg-gray-50 dark:hover:bg-gray-700">取消</a>
            <button className="px-4 py-2 text-sm font-medium text-white bg-black rounded-md hover:bg-gray-800 flex items-center space-x-2">
              <span className="material-icons text-base">save</span>
              <span>保存衣物</span>
            </button>
          </div>
        </div>
      </main>

      <footer className="bg-card-light dark:bg-card-dark mt-auto">
        <div className="container mx-auto py-4 px-4 sm:px-6 lg:px-8">
          <p className="text-center text-sm text-text-light-secondary dark:text-text-secondary-dark">©2024 StyleSense. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}
