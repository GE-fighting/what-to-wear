'use client';

import React, { useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { Button, Card, CardHeader, CardContent, Badge, Avatar } from '@/components/ui';
import { useToast } from '@/components/ToastProvider';
import { Footer } from '@/components/Footer';

// 模拟衣物详情数据
const mockClothingDetail = {
  id: '1',
  name: '机能风冲锋衣',
  category: '外套',
  brand: 'Urban Explorer',
  size: 'L',
  material: '聚酯纤维',
  color: '深灰',
  price: '¥899.00',
  purchaseDate: '2024-03-15',
  purchaseChannel: '线上商城',
  notes: '防水性能优秀，适合户外活动。内胆可拆卸，春秋两季都可穿。',
  tags: ['户外', '春秋', '深灰'],
  images: [
    'https://lh3.googleusercontent.com/aida-public/AB6AXuAXLlX9_wDDbI1JL-XWOD2ZzHctF1C2qgXOKY6rFySp-D5_oW7il56harRlIlVMm03T2AmlAh8Opr3WV5KOtTzA7XaxBZ1r3elJ4tt44AQeoBxntu0L1FNKNk4svMzRDsiLL5Fs37HdiM20WXgZBVxn2yHjDnsrb_4alz15IzpFel9by6Jd-mPwKwGaShtyXUhiGfUNOMeypYitGpQettbrOOG94YqlwmHen3dRw02omW7iMLy8yZPCIid6LErstr7BlzPXsbFSx6M',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuAkHfG8nBe7u5ky9SU6swqaWeW1-vlQh4NGnzLV5wcO1HGZ8UrMgwB1I8etuA48UQjqQTaC7FyDzurm8Yg_tjUXEXX0K9HqsQk4J16OZbuegSNbXf35h2CzqYrbzyReOOQQeD5llB_Kxkk_FzpBnnFtvBm6D9IfD9JV4vIepeh28KsFUbiffkw7PAya8M5RvMmByPuoK1cGq4zL6mSQqiZaJVIU67cpamB-8RDI3B3pv7kb0_HOsBXODmEfF0y4IEcft4h388KEs9Yd',
    'https://lh3.googleusercontent.com/aida-public/AB6AXuBpKRSkI0k_rv3-f3pkIHx8gSRFlXMKWkz5pMMTOl91whfAGmxZ7W_Wvp_51sCXqPS05Ghwswh09y79iYU2YbqzgVI4qFgBJ2Mczj3QLPKX7k1sxRZ3sMniVjXctOedrgdGbFGiCcgRUuiEp4TgBZ2UcKsMvm9ii-cWu6FEeSZ_3WqYfeCv_IZ8HWbm70HZy2d4gHO0jqVJ3ZoyD639pZGAXsiMRunOgs_pq5LK-DtbU-il8n_jKcAe8zuVzcacFavxvNxlkCyNSpRJ',
  ],
  wearStats: {
    total: 6,
    monthly: 2,
    lastWorn: '3天前',
  },
  durability: {
    percentage: 90,
    estimatedLife: '2年',
  },
  wearHistory: [
    {
      id: '1',
      date: '12月15日',
      daysAgo: '3天前',
      occasion: '周末出行',
      weather: { condition: '晴天', temperature: '15°C', icon: 'wb_sunny' },
      duration: '8小时',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuAXLlX9_wDDbI1JL-XWOD2ZzHctF1C2qgXOKY6rFySp-D5_oW7il56harRlIlVMm03T2AmlAh8Opr3WV5KOtTzA7XaxBZ1r3elJ4tt44AQeoBxntu0L1FNKNk4svMzRDsiLL5Fs37HdiM20WXgZBVxn2yHjDnsrb_4alz15IzpFel9by6Jd-mPwKwGaShtyXUhiGfUNOMeypYitGpQettbrOOG94YqlwmHen3dRw02omW7iMLy8yZPCIid6LErstr7BlzPXsbFSx6M',
    },
    {
      id: '2',
      date: '12月08日',
      daysAgo: '1周前',
      occasion: '通勤穿搭',
      weather: { condition: '多云', temperature: '12°C', icon: 'cloud' },
      duration: '10小时',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuAXLlX9_wDDbI1JL-XWOD2ZzHctF1C2qgXOKY6rFySp-D5_oW7il56harRlIlVMm03T2AmlAh8Opr3WV5KOtTzA7XaxBZ1r3elJ4tt44AQeoBxntu0L1FNKNk4svMzRDsiLL5Fs37HdiM20WXgZBVxn2yHjDnsrb_4alz15IzpFel9by6Jd-mPwKwGaShtyXUhiGfUNOMeypYitGpQettbrOOG94YqlwmHen3dRw02omW7iMLy8yZPCIid6LErstr7BlzPXsbFSx6M',
    },
    {
      id: '3',
      date: '11月28日',
      daysAgo: '2周前',
      occasion: '户外徒步',
      weather: { condition: '阴天', temperature: '10°C', icon: 'wb_cloudy' },
      duration: '6小时',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuAXLlX9_wDDbI1JL-XWOD2ZzHctF1C2qgXOKY6rFySp-D5_oW7il56harRlIlVMm03T2AmlAh8Opr3WV5KOtTzA7XaxBZ1r3elJ4tt44AQeoBxntu0L1FNKNk4svMzRDsiLL5Fs37HdiM20WXgZBVxn2yHjDnsrb_4alz15IzpFel9by6Jd-mPwKwGaShtyXUhiGfUNOMeypYitGpQettbrOOG94YqlwmHen3dRw02omW7iMLy8yZPCIid6LErstr7BlzPXsbFSx6M',
    },
  ],
  relatedOutfits: [
    {
      id: '1',
      title: '都市休闲风',
      date: '11月28日',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuDZDhxirXA366O_DiSOoHAEEaXEEZl52tKVDCulW78Jn8rXFke-wFY-bhLXzU92MuCC5fEPXK9M8mmylm2lEMDK5VNUZLPjSfHsdHObV7yx2gz7J50f56OZvBcn-YkoVl8PBL6Y8sbRVM9w45LydSqfN6CnlhYkBvG8AXwJT98XFFAZXx9mjRcrHgsZsQmDhfMb4Esv4zydxO9AS1Xr5Ni8WqlzAp-smDDXfaRX9buyS4fUrPDuAiWUSyOYmTDl5A6xtmeRScB7cw',
    },
    {
      id: '2',
      title: '户外运动风',
      date: '12月08日',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuDH-L7FdmL3xHhlueKM1mE__6QUrDRhY3Lpu69oN3tqwTZgTzqhjiMG4KKmHqm3hfhYFuE3IyfN4wqsbV4uXFepSq3L3VbONrlxEC2SPr1rSxTwDle1CL6nu92HiusbkGSQUU-h41V2a9UnJBIxHa1b65CaD5f04CLPvuSQ5Ua1zMRvMv5GB-17Oy7kw3oEsVihT9dgmO03x4zPGUnCEGbYpjcNsOD4sj38ph_VyANpYycuQSJgI6MYMLtNTevwKLM0iRuXyxpisQ',
    },
    {
      id: '3',
      title: '极简通勤',
      date: '12月15日',
      image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBpKRSkI0k_rv3-f3pkIHx8gSRFlXMKWkz5pMMTOl91whfAGmxZ7W_Wvp_51sCXqPS05Ghwswh09y79iYU2YbqzgVI4qFgBJ2Mczj3QLPKX7k1sxRZ3sMniVjXctOedrgdGbFGiCcgRUuiEp4TgBZ2UcKsMvm9ii-cWu6FEeSZ_3WqYfeCv_IZ8HWbm70HZy2d4gHO0jqVJ3ZoyD639pZGAXsiMRunOgs_pq5LK-DtbU-il8n_jKcAe8zuVzcacFavxvNxlkCyNSpRJ',
    },
  ],
};

export default function ClothingDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { showToast } = useToast();
  const [currentImage, setCurrentImage] = useState(0);
  const [showFullscreen, setShowFullscreen] = useState(false);
  const [showDeleteModal, setShowDeleteModal] = useState(false);

  const clothing = mockClothingDetail; // 在实际项目中应该根据 id 从 API 获取

  const handlePreviousImage = () => {
    setCurrentImage((prev) => (prev - 1 + clothing.images.length) % clothing.images.length);
  };

  const handleNextImage = () => {
    setCurrentImage((prev) => (prev + 1) % clothing.images.length);
  };

  const handleDelete = () => {
    showToast('衣物已删除', 'success');
    router.push('/main');
  };

  const handleEdit = () => {
    router.push(`/main/wardrobe/edit/${params.id}`);
  };

  const handleShare = () => {
    showToast('分享功能开发中', 'info');
  };

  const handleFavorite = () => {
    showToast('已添加到收藏', 'success');
  };

  return (
    <div className="min-h-screen bg-background-light dark:bg-background-dark font-display">
      {/* 统一导航栏 */}
      <header className="bg-card-light dark:bg-card-dark shadow-sm sticky top-0 z-20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-8">
              <h1 className="text-xl font-bold text-text-primary-light dark:text-text-primary-dark">StyleSense</h1>
              <nav className="hidden md:flex space-x-8">
                <button className="text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-primary transition-colors">
                  记录穿搭
                </button>
                <button className="text-sm font-medium text-primary border-b-2 border-primary">
                  我的衣橱
                </button>
                <button className="text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-primary transition-colors">
                  风格灵感
                </button>
                <button className="text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-primary transition-colors">
                  穿搭分析
                </button>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-secondary-light dark:text-text-secondary-dark">
                  search
                </span>
                <input
                  type="text"
                  className="bg-background-light dark:bg-background-dark border border-border-light dark:border-border-dark rounded-full py-2 pl-10 pr-4 text-sm w-48 focus:ring-primary focus:border-primary text-text-primary-light dark:text-text-primary-dark"
                  placeholder="搜索衣物"
                />
              </div>
              <button className="p-2 rounded-full text-text-secondary-light dark:text-text-secondary-dark hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors">
                <span className="material-icons-outlined">notifications_none</span>
              </button>
              <Avatar
                src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL"
                alt="用户头像"
                size="sm"
              />
            </div>
          </div>
        </div>
      </header>

      <main className="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="flex items-center justify-between mb-6">
          <button
            onClick={() => router.push('/main')}
            className="flex items-center text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-text-primary-light dark:hover:text-text-primary-dark transition-colors"
          >
            <span className="material-icons mr-1">arrow_back_ios</span>
            返回衣物列表
          </button>
          <div className="flex items-center space-x-2">
            <Button variant="secondary" onClick={handleFavorite} icon="favorite_border">
              收藏
            </Button>
            <Button variant="secondary" onClick={handleShare} icon="share">
              分享
            </Button>
            <Button onClick={handleEdit} icon="edit">
              编辑
            </Button>
            <Button
              onClick={() => setShowDeleteModal(true)}
              className="bg-red-500 hover:bg-red-600 text-white"
              icon="delete"
            >
              删除
            </Button>
          </div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* 图片展示区 */}
          <div className="lg:col-span-1 space-y-4">
            <div className="relative bg-card-light dark:bg-card-dark rounded-lg overflow-hidden shadow-sm max-h-[600px]">
              <div
                onClick={() => setShowFullscreen(true)}
                className="aspect-w-1 aspect-h-1 cursor-zoom-in max-h-[600px]"
              >
                <img
                  src={clothing.images[currentImage]}
                  alt={clothing.name}
                  className="w-full h-full object-cover"
                />
              </div>
              <button
                onClick={handlePreviousImage}
                className="absolute top-1/2 left-4 -translate-y-1/2 bg-black bg-opacity-20 text-white p-2 rounded-full hover:bg-opacity-40 transition-opacity"
              >
                <span className="material-icons">chevron_left</span>
              </button>
              <button
                onClick={handleNextImage}
                className="absolute top-1/2 right-4 -translate-y-1/2 bg-black bg-opacity-20 text-white p-2 rounded-full hover:bg-opacity-40 transition-opacity"
              >
                <span className="material-icons">chevron_right</span>
              </button>
              <button
                onClick={() => setShowFullscreen(true)}
                className="absolute top-4 right-4 bg-black bg-opacity-20 text-white p-2 rounded-full hover:bg-opacity-40 transition-opacity"
              >
                <span className="material-icons-outlined">fullscreen</span>
              </button>
            </div>

            {/* 缩略图 */}
            <div className="flex space-x-2 overflow-x-auto pb-2">
              {clothing.images.map((image, index) => (
                <button
                  key={index}
                  onClick={() => setCurrentImage(index)}
                  className={`flex-shrink-0 w-20 h-20 rounded-lg overflow-hidden transition-all ${
                    currentImage === index ? 'ring-2 ring-primary' : 'opacity-60 hover:opacity-100'
                  }`}
                >
                  <img src={image} alt="" className="w-full h-full object-cover" />
                </button>
              ))}
            </div>
          </div>

          {/* 详情信息区 */}
          <div className="lg:col-span-1">
            <Card className="p-6">
              <h1 className="text-3xl font-bold mb-2 text-text-primary-light dark:text-text-primary-dark">
                {clothing.name}
              </h1>
              <div className="flex items-center space-x-2 mb-4">
                {clothing.tags.map((tag, index) => (
                  <Badge
                    key={index}
                    variant="secondary"
                    className="bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200"
                  >
                    {tag}
                  </Badge>
                ))}
              </div>

              <div className="space-y-4 text-sm text-text-secondary-light dark:text-text-secondary-dark">
                <div className="flex justify-between">
                  <span>分类:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.category}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>品牌:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.brand}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>尺码:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.size}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>材质:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.material}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>购买日期:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.purchaseDate}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>购入价格:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.price}
                  </span>
                </div>
                <div className="flex justify-between">
                  <span>购买渠道:</span>
                  <span className="font-medium text-text-primary-light dark:text-text-primary-dark">
                    {clothing.purchaseChannel}
                  </span>
                </div>
              </div>

              <div className="my-6 border-t border-border-light dark:border-border-dark"></div>

              <div>
                <h2 className="text-base font-semibold mb-2 text-text-primary-light dark:text-text-primary-dark">
                  备注
                </h2>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">
                  {clothing.notes}
                </p>
              </div>

              <div className="my-6 border-t border-border-light dark:border-border-dark"></div>

              <div>
                <h2 className="text-lg font-semibold mb-3 text-text-primary-light dark:text-text-primary-dark">
                  穿着统计
                </h2>
                <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">
                  累计穿着{' '}
                  <span className="font-bold text-primary">{clothing.wearStats.total}</span> 次 ·{' '}
                  本月 <span className="font-bold text-primary">{clothing.wearStats.monthly}</span> 次 ·{' '}
                  最后一次 <span className="font-bold text-primary">{clothing.wearStats.lastWorn}</span>
                </p>
              </div>

              <div className="mt-6">
                <h2 className="text-lg font-semibold mb-2 text-text-primary-light dark:text-text-primary-dark">
                  耐久度
                </h2>
                <div className="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2.5">
                  <div
                    className="bg-success h-2.5 rounded-full"
                    style={{ width: `${clothing.durability.percentage}%` }}
                  ></div>
                </div>
                <p className="text-xs text-right mt-1 text-text-secondary-light dark:text-text-secondary-dark">
                  预计剩余寿命: {clothing.durability.estimatedLife}
                </p>
              </div>
            </Card>
          </div>
        </div>

        {/* 穿着历史记录 */}
        <div className="mt-12">
          <h2 className="text-xl font-bold mb-6 text-text-primary-light dark:text-text-primary-dark">
            穿着历史
          </h2>
          <Card className="p-6">
            <div className="space-y-4">
              {clothing.wearHistory.map((record, index) => (
                <div
                  key={record.id}
                  className={`flex items-start space-x-4 pb-4 ${
                    index < clothing.wearHistory.length - 1
                      ? 'border-b border-border-light dark:border-border-dark'
                      : 'border-0'
                  }`}
                >
                  <div className="flex-shrink-0 text-center">
                    <div className="text-sm font-semibold text-text-primary-light dark:text-text-primary-dark">
                      {record.date.split('月')[0]}月
                    </div>
                    <div className="text-2xl font-bold text-primary">
                      {record.date.split('月')[1].split('日')[0]}
                    </div>
                  </div>
                  <div className="flex-grow">
                    <div className="flex items-center justify-between mb-2">
                      <span className="text-sm font-medium text-text-primary-light dark:text-text-primary-dark">
                        {record.occasion}
                      </span>
                      <span className="text-xs text-text-secondary-light dark:text-text-secondary-dark">
                        {record.daysAgo}
                      </span>
                    </div>
                    <div className="flex items-center space-x-2 text-xs text-text-secondary-light dark:text-text-secondary-dark">
                      <span className="material-icons-outlined text-sm">{record.weather.icon}</span>
                      <span>
                        {record.weather.condition} • {record.weather.temperature}
                      </span>
                      <span className="ml-2">穿着时长: {record.duration}</span>
                    </div>
                  </div>
                  <img
                    src={record.image}
                    alt="穿搭照片"
                    className="w-16 h-16 rounded-lg object-cover"
                  />
                </div>
              ))}
            </div>
            <button className="w-full mt-6 py-2 text-sm font-medium text-text-secondary-light dark:text-text-secondary-dark hover:text-text-primary-light dark:hover:text-text-primary-dark border border-border-light dark:border-border-dark rounded-lg hover:bg-background-light dark:hover:bg-background-dark transition-colors">
              查看全部历史记录
            </button>
          </Card>
        </div>

        {/* 相关穿搭推荐 */}
        <div className="mt-12">
          <div className="flex items-center justify-between mb-6">
            <h2 className="text-xl font-bold text-text-primary-light dark:text-text-primary-dark">
              相关穿搭
            </h2>
            <button className="text-sm font-medium text-primary hover:text-primary-hover transition-colors">
              查看更多
            </button>
          </div>
          <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {clothing.relatedOutfits.map((outfit) => (
              <div key={outfit.id} className="group cursor-pointer">
                <div className="relative aspect-square rounded-lg overflow-hidden bg-card-light dark:bg-card-dark shadow-sm">
                  <img
                    src={outfit.image}
                    alt={outfit.title}
                    className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-black/60 to-transparent"></div>
                  <div className="absolute bottom-3 left-3 right-3">
                    <p className="text-white text-sm font-medium">{outfit.title}</p>
                    <p className="text-gray-300 text-xs mt-1">{outfit.date}</p>
                  </div>
                </div>
              </div>
            ))}
            <div className="group cursor-pointer">
              <div className="relative aspect-square rounded-lg overflow-hidden bg-card-light dark:bg-card-dark shadow-sm border-2 border-dashed border-border-light dark:border-border-dark flex items-center justify-center hover:border-primary transition-colors">
                <div className="text-center">
                  <span className="material-icons-outlined text-4xl text-text-secondary-light dark:text-text-secondary-dark mb-2">
                    add_circle_outline
                  </span>
                  <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark">
                    创建新搭配
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* 全屏图片查看器 */}
      {showFullscreen && (
        <>
          <div
            onClick={() => setShowFullscreen(false)}
            className="fixed inset-0 bg-black z-50 flex items-center justify-center"
          />
          <div className="fixed inset-0 z-50 flex items-center justify-center">
            <button
              onClick={() => setShowFullscreen(false)}
              className="absolute top-4 right-4 text-white p-2 rounded-full hover:bg-white hover:bg-opacity-20"
            >
              <span className="material-icons-outlined text-3xl">close</span>
            </button>
            <button
              onClick={handlePreviousImage}
              className="absolute left-4 text-white p-2 rounded-full hover:bg-white hover:bg-opacity-20"
            >
              <span className="material-icons text-3xl">chevron_left</span>
            </button>
            <button
              onClick={handleNextImage}
              className="absolute right-4 text-white p-2 rounded-full hover:bg-white hover:bg-opacity-20"
            >
              <span className="material-icons text-3xl">chevron_right</span>
            </button>
            <img
              src={clothing.images[currentImage]}
              alt=""
              className="max-h-[90vh] max-w-[90vw] object-contain"
            />
          </div>
        </>
      )}

      {/* 删除确认模态框 */}
      {showDeleteModal && (
        <>
          <div
            onClick={() => setShowDeleteModal(false)}
            className="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center p-4"
          />
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <div className="bg-card-light dark:bg-card-dark rounded-lg shadow-xl max-w-md w-full p-6">
              <h3 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">
                确认删除
              </h3>
              <p className="text-sm text-text-secondary-light dark:text-text-secondary-dark mb-6">
                确定要删除这件衣物吗？此操作无法撤销。
              </p>
              <div className="flex justify-end space-x-3">
                <Button
                  variant="secondary"
                  onClick={() => setShowDeleteModal(false)}
                >
                  取消
                </Button>
                <Button
                  onClick={handleDelete}
                  className="bg-red-500 hover:bg-red-600 text-white"
                >
                  确认删除
                </Button>
              </div>
            </div>
          </div>
        </>
      )}

      <Footer />
    </div>
  );
}