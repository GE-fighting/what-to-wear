export default function PersonalInformationPage() {
  return (
    <div className="bg-card-light dark:bg-card-dark p-8 rounded-lg shadow-sm">
      <h2 className="text-2xl font-bold text-text-light dark:text-text-dark mb-6">个人信息</h2>
      <form action="#" method="POST">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div className="col-span-2">
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="username">用户名</label>
            <div className="mt-1">
              <input className="block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="username" name="username" type="text" defaultValue="用户名"/>
            </div>
          </div>
          
          <div className="col-span-2">
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="bio">个人简介</label>
            <div className="mt-1">
              <textarea className="block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="bio" name="bio" rows={3} defaultValue="时尚爱好者"/>
            </div>
            <p className="mt-2 text-sm text-text-secondary-light dark:text-text-secondary-dark">简单介绍一下你自己。</p>
          </div>
          
          <div className="col-span-2">
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark">头像</label>
            <div className="mt-2 flex items-center space-x-5">
              <img alt="Current user avatar" className="h-16 w-16 rounded-full" src="https://lh3.googleusercontent.com/aida-public/AB6AXuBCbrb1iMiE1nW4fXi9ho0FZqEkffLllkk5nFsZdniknGfbwoo4qaYif1WRnHGCZZEQwcfYU5wm8dd58QNBysdyxyaM1nPgfThjyI3Qi0oTUHV3XHduzE4Y8lApVAfTXlM1MBqOQbFXF8YQqLdeZlYd5mrDXJI8Tx1rWzmq3XhlQjLScSFZeQiKo8ddLgJKhEM6nnxwDot4KxVmhYUlCKUulGYFX1J5TXV512BtSS_PkJFa4Q3J9rkQDNoLvcRDu8i1TVyiSCS5UOwP"/>
              <button className="rounded-md border border-border-light dark:border-border-dark bg-card-light dark:bg-card-dark py-2 px-3 text-sm font-medium leading-4 text-text-secondary-light dark:text-text-secondary-dark shadow-sm hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" type="button">更换</button>
            </div>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="email">邮箱地址</label>
            <div className="mt-1">
              <input className="block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="email" name="email" type="email" defaultValue="user@example.com"/>
            </div>
          </div>
          
          <div>
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="phone">手机号码</label>
            <div className="mt-1">
              <input className="block w-full rounded-md border-border-light dark:border-border-dark shadow-sm focus:border-primary focus:ring-primary bg-background-light dark:bg-background-dark text-text-light dark:text-text-dark" id="phone" name="phone" type="tel" defaultValue="+86 123 4567 8901"/>
            </div>
          </div>
          
          <div className="col-span-2">
            <label className="block text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark" htmlFor="style-preference">风格偏好</label>
            <div className="mt-2 flex flex-wrap gap-2">
              <span className="inline-flex items-center rounded-full bg-primary/10 px-3 py-1 text-sm font-medium text-primary">简约</span>
              <span className="inline-flex items-center rounded-full bg-primary/10 px-3 py-1 text-sm font-medium text-primary">休闲</span>
              <span className="inline-flex items-center rounded-full bg-primary/10 px-3 py-1 text-sm font-medium text-primary">通勤</span>
              <button className="inline-flex items-center rounded-full border border-dashed border-border-light dark:border-border-dark px-3 py-1 text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:border-primary hover:text-primary" type="button">
                <span className="material-icons text-base mr-1">add</span>
                <span>添加</span>
              </button>
            </div>
          </div>
        </div>
        
        <div className="pt-8 flex justify-end">
          <button className="rounded-md border border-border-light dark:border-border-dark bg-card-light dark:bg-card-dark py-2 px-4 text-sm font-medium text-text-light dark:text-text-dark shadow-sm hover:bg-gray-50 dark:hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" type="button">取消</button>
          <button className="ml-3 inline-flex justify-center rounded-md border border-transparent bg-primary py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary" type="submit">保存</button>
        </div>
      </form>
    </div>
  );
}
