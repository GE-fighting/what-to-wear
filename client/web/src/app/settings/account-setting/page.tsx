'use client';

export default function AccountSettingPage() {
  return (
    <div className="space-y-8">
      <div className="bg-card-light dark:bg-card-dark p-8 rounded-lg shadow-sm">
        <h2 className="text-2xl font-bold text-text-light dark:text-text-dark mb-6">账户设置</h2>

        <div className="border-b border-border-light dark:border-border-dark pb-6 mb-6">
          <h3 className="text-lg font-medium text-text-light dark:text-text-dark mb-4">更改密码</h3>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="current-password">当前密码</label>
              <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="current-password" type="password"/>
            </div>
            <div>
              <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="new-password">新密码</label>
              <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="new-password" type="password"/>
            </div>
          </div>
          <div className="mt-4 flex justify-end">
            <button className="rounded-md border border-transparent bg-primary py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" type="button">更新密码</button>
          </div>
        </div>

        <div className="border-b border-border-light dark:border-border-dark pb-6 mb-6">
          <h3 className="text-lg font-medium text-text-light dark:text-text-dark mb-4">更改邮箱</h3>
          <div className="flex items-center space-x-4">
            <div className="flex-grow">
              <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="email">邮箱地址</label>
              <input className="mt-1 block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="email" type="email" defaultValue="user@example.com"/>
            </div>
            <button className="self-end rounded-md border border-transparent bg-primary py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" type="button">更新邮箱</button>
          </div>
        </div>

        <div className="border-b border-border-light dark:border-border-dark pb-6 mb-6">
          <h3 className="text-lg font-medium text-text-light dark:text-text-dark mb-4">通知设置</h3>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-text-light dark:text-text-dark">新品推荐</p>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">获取最新单品和品牌合作的通知。</p>
              </div>
              <label className="switch">
                <input type="checkbox" defaultChecked/>
                <span className="slider"></span>
              </label>
            </div>

            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-text-light dark:text-text-dark">促销活动</p>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">接收特别优惠和折扣信息。</p>
              </div>
              <label className="switch">
                <input type="checkbox" defaultChecked/>
                <span className="slider"></span>
              </label>
            </div>

            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-text-light dark:text-text-dark">社区动态</p>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">当有人评论或喜欢你的穿搭时收到通知。</p>
              </div>
              <label className="switch">
                <input type="checkbox"/>
                <span className="slider"></span>
              </label>
            </div>
          </div>
        </div>

        <div>
          <h3 className="text-lg font-medium text-text-light dark:text-text-dark mb-4">隐私设置</h3>
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-text-light dark:text-text-dark">公开个人资料</p>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">允许其他用户查看你的个人资料和穿搭。</p>
              </div>
              <label className="switch">
                <input type="checkbox" defaultChecked/>
                <span className="slider"></span>
              </label>
            </div>

            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-text-light dark:text-text-dark">允许被搜索</p>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">允许其他用户通过你的用户名搜索到你。</p>
              </div>
              <label className="switch">
                <input type="checkbox"/>
                <span className="slider"></span>
              </label>
            </div>
          </div>
        </div>
      </div>

      <div className="bg-red-50 dark:bg-red-900/20 p-6 rounded-lg shadow-sm">
        <h3 className="text-lg font-medium text-red-700 dark:text-red-400 mb-2">删除账户</h3>
        <p className="text-sm text-red-600 dark:text-red-300 mb-4">一旦删除你的账户，所有数据将被永久移除。此操作无法撤销。</p>
        <button className="w-full sm:w-auto rounded-md border border-transparent bg-red-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-red-50 focus:ring-red-500">
          删除我的账户
        </button>
      </div>
    </div>
  );
}