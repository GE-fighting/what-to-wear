'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button, Card, CardHeader, CardContent, Input } from '@/components/ui';
import { Badge, Avatar } from '@/components/ui';
import { useToast } from '@/components/ToastProvider';

// 模拟图片数据
const mockImages = [
  {
    id: '1',
    url: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBfXUvHl_ZjPaWqqe58_4LD2RTGeU0QaJQXmf-JXkC1CUK2IM0-qXg_PxwTTiB49YHlMZBJ-CGpgLaBdOD6iqe7UX6l8iKOk58ZGfvHD1lrETWOcx1VGe444V67KwXbT8aSX-nGa7FSGUCfsSJB97IHcyea9ChEmdS0bvu5L3fWa6fc6LI71zJbm60d5mcBC9sne0S2Lvm6Q2mlpokzTOaFnOx2QbwwLvtBdFq00Mp-4vdF9ET2xPk0P8pNU6ae8LWWGmsCkO0AX26T',
    alt: '黑色T恤',
  },
  {
    id: '2',
    url: 'https://lh3.googleusercontent.com/aida-public/AB6AXuCykmBlONpG_g9Jsqom6b9elDQa4cHB4aMEoUiaPGLskaNa-kiYQx1p-egTA78yaoqGw11kHmzrqgnsWOPmtmahIjcvvv8kU5GJpziLdv4Jhb2kIMjH5_p3jklc86CXYZefBTWu1bDT6oOGpYG8OYjrA4BFcTygNB_1VoXEQUxQqClgVkzvIh201YOX5EDHVT09aO7xy-EptNCLtE9KFqjK0jyIJy1VY6XIaUygg4n7Nba8KNjcFsw4j-LwbmxWG7jPEx9zr4Po738g',
    alt: '黑色T恤背面',
  },
  {
    id: '3',
    url: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBhPzMq0rTN8DW9qnYZMrborfQ_SOLL6LmC5Wa6IB0qQMdVNbCQoC_iw60TZI0zgYxRWeOzM-Ge1g53uLCVyUN_e4XKslWpOEp4u22TAYrCSiySlKBPsJAWztC0_5XrqImr_Je-ScWbW0t0nftQsSK8LVoJEBCqpY1y-HRwE9Bs8HEfC9aI_9YmuuZq95sqLmbzZgT1ZQdU7tZQwOTS0fZ4-X41tZpx2CBty7_jZ2Lov8JWKxZEZ1NEB8Thkl5nAXhmio63LKkd4h1K',
    alt: '黑色T恤细节',
  },
];

// AI建议标签
const aiSuggestedTags = [
  { id: '1', label: '颜色: 黑色', type: 'color' },
  { id: '2', label: '款式: T恤', type: 'style' },
  { id: '3', label: '材质: 棉', type: 'material' },
  { id: '4', label: '风格: 简约', type: 'style' },
];

export default function AddClothingPage() {
  const router = useRouter();
  const { showToast } = useToast();
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState({
    name: '黑色T恤',
    category: '上装',
    brand: '',
    color: '黑色',
    material: '棉',
    style: '简约',
    size: '',
    season: '',
    occasion: '',
    price: '',
    store: '',
    purchaseDate: '',
    notes: '',
  });
  const [images, setImages] = useState(mockImages);
  const [suggestedTags, setSuggestedTags] = useState(aiSuggestedTags);
  const [appliedTags, setAppliedTags] = useState<string[]>([]);

  const handleInputChange = (name: string, value: string) => {
    setFormData(prev => ({
      ...prev,
      [name]: value
    }));
  };

  const handleRemoveImage = (imageId: string) => {
    setImages(prev => prev.filter(img => img.id !== imageId));
    showToast('图片已删除', 'info');
  };

  const handleAddImage = () => {
    showToast('图片上传功能开发中', 'info');
  };

  const handleRemoveTag = (tagId: string) => {
    setSuggestedTags(prev => prev.filter(tag => tag.id !== tagId));
  };

  const handleApplyAllTags = () => {
    const newTags = suggestedTags.map(tag => tag.label);
    setAppliedTags(prev => [...prev, ...newTags]);
    setSuggestedTags([]);
    showToast('AI标签已应用', 'success');
  };

  const handleSave = async () => {
    setLoading(true);
    try {
      // 模拟保存过程
      await new Promise(resolve => setTimeout(resolve, 1500));
      showToast('衣物添加成功！', 'success');
      router.push('/main');
    } catch (error) {
      showToast('保存失败，请重试', 'error');
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    router.push('/main');
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

      <main className="flex-grow container mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="max-w-4xl mx-auto">
          <div className="flex items-center justify-between mb-8">
            <h1 className="text-3xl font-bold text-text-primary-light dark:text-text-primary-dark">添加新衣物</h1>
            <div className="flex space-x-2">
              <Button variant="secondary" onClick={handleCancel}>
                取消
              </Button>
              <Button onClick={handleSave} loading={loading} icon="save">
                保存衣物
              </Button>
            </div>
          </div>

          <div className="space-y-10">
            {/* 衣物图片 */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">衣物图片</h2>
              <div className="grid grid-cols-3 sm:grid-cols-4 md:grid-cols-5 gap-4">
                {images.map((image) => (
                  <div key={image.id} className="relative group aspect-square">
                    <img
                      alt={image.alt}
                      className="w-full h-full object-cover rounded-lg"
                      src={image.url}
                    />
                    <div className="absolute inset-0 bg-black bg-opacity-40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity rounded-lg">
                      <button
                        onClick={() => handleRemoveImage(image.id)}
                        className="text-white hover:text-red-400 transition-colors"
                      >
                        <span className="material-icons-outlined">delete</span>
                      </button>
                    </div>
                  </div>
                ))}
                <div
                  onClick={handleAddImage}
                  className="aspect-square flex items-center justify-center border-2 border-dashed border-border-light dark:border-border-dark rounded-lg text-text-secondary-light dark:text-text-secondary-dark hover:border-primary hover:text-primary cursor-pointer transition-colors"
                >
                  <div className="text-center">
                    <span className="material-icons-outlined text-4xl">add_photo_alternate</span>
                    <p className="text-xs mt-1">添加图片</p>
                  </div>
                </div>
              </div>
            </Card>

            {/* AI 标签建议 */}
            {suggestedTags.length > 0 && (
              <Card className="p-4 bg-success-light/10 dark:bg-success-dark/20 border border-success/30">
                <div className="flex items-start space-x-3">
                  <div className="flex-shrink-0">
                    <span className="material-icons-outlined text-success">auto_awesome</span>
                  </div>
                  <div className="flex-grow">
                    <h3 className="text-base font-semibold text-success-dark dark:text-success-light">AI 标签建议</h3>
                    <p className="text-sm text-success/80 dark:text-success/70 mt-1">
                      根据您上传的图片，我们为您推荐了以下标签。您可以直接使用或进行修改。
                    </p>
                    <div className="mt-4">
                      <div className="flex flex-wrap gap-2">
                        {suggestedTags.map((tag) => (
                          <Badge
                            key={tag.id}
                            variant="secondary"
                            className="flex items-center gap-1 pr-1"
                          >
                            {tag.label}
                            <button
                              onClick={() => handleRemoveTag(tag.id)}
                              className="ml-1 text-text-secondary-light dark:text-text-secondary-dark hover:text-error"
                            >
                              <span className="material-icons-outlined text-sm">close</span>
                            </button>
                          </Badge>
                        ))}
                      </div>
                    </div>
                    <div className="mt-4 flex items-center space-x-3">
                      <Button size="sm" onClick={handleApplyAllTags} icon="done">
                        全部应用
                      </Button>
                      <Button size="sm" variant="secondary" icon="edit">
                        编辑
                      </Button>
                    </div>
                  </div>
                </div>
              </Card>
            )}

            {/* 基本信息 */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">基本信息</h2>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-4">
                <Input
                  label="衣物名称"
                  value={formData.name}
                  onChange={(e) => handleInputChange('name', e.target.value)}
                  placeholder="例如：纯棉白色T恤"
                />

                <div>
                  <label className="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark mb-2">
                    分类
                  </label>
                  <select
                    value={formData.category}
                    onChange={(e) => handleInputChange('category', e.target.value)}
                    className="w-full px-3 py-2 bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-lg focus:ring-2 focus:ring-primary focus:border-primary text-text-primary-light dark:text-text-primary-dark"
                  >
                    <option value="上装">上装</option>
                    <option value="下装">下装</option>
                    <option value="鞋履">鞋履</option>
                    <option value="配饰">配饰</option>
                  </select>
                </div>

                <Input
                  label="品牌"
                  value={formData.brand}
                  onChange={(e) => handleInputChange('brand', e.target.value)}
                  placeholder="例如：Uniqlo"
                />

                <Input
                  label="颜色"
                  value={formData.color}
                  onChange={(e) => handleInputChange('color', e.target.value)}
                  placeholder="例如：白色"
                />

                <Input
                  label="材质"
                  value={formData.material}
                  onChange={(e) => handleInputChange('material', e.target.value)}
                  placeholder="例如：100%纯棉"
                />

                <Input
                  label="风格"
                  value={formData.style}
                  onChange={(e) => handleInputChange('style', e.target.value)}
                  placeholder="例如：简约、休闲"
                />

                <Input
                  label="尺码"
                  value={formData.size}
                  onChange={(e) => handleInputChange('size', e.target.value)}
                  placeholder="例如：M / 170"
                />
              </div>
            </Card>

            {/* 适用场景 */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">适用场景</h2>
              <div className="grid grid-cols-1 md:grid-cols-2 gap-x-6 gap-y-4">
                <Input
                  label="季节"
                  value={formData.season}
                  onChange={(e) => handleInputChange('season', e.target.value)}
                  placeholder="例如：春夏"
                />

                <Input
                  label="场合"
                  value={formData.occasion}
                  onChange={(e) => handleInputChange('occasion', e.target.value)}
                  placeholder="例如：日常通勤、周末出游"
                />
              </div>
            </Card>

            {/* 购买信息 */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">购买信息</h2>
              <div className="grid grid-cols-1 md:grid-cols-3 gap-x-6 gap-y-4">
                <div>
                  <label className="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark mb-2">
                    价格
                  </label>
                  <div className="relative">
                    <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                      <span className="text-text-secondary-light dark:text-text-secondary-dark sm:text-sm">¥</span>
                    </div>
                    <input
                      type="text"
                      value={formData.price}
                      onChange={(e) => handleInputChange('price', e.target.value)}
                      placeholder="0.00"
                      className="block w-full pl-7 pr-12 px-3 py-2 bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-lg focus:ring-2 focus:ring-primary focus:border-primary text-text-primary-light dark:text-text-primary-dark"
                    />
                  </div>
                </div>

                <Input
                  label="商店"
                  value={formData.store}
                  onChange={(e) => handleInputChange('store', e.target.value)}
                  placeholder="例如：淘宝官方旗舰店"
                />

                <Input
                  label="购买日期"
                  value={formData.purchaseDate}
                  onChange={(e) => handleInputChange('purchaseDate', e.target.value)}
                  type="date"
                />
              </div>
            </Card>

            {/* 其他备注 */}
            <Card className="p-6">
              <h2 className="text-lg font-semibold mb-4 text-text-primary-light dark:text-text-primary-dark">其他备注</h2>
              <div>
                <label className="block text-sm font-medium text-text-primary-light dark:text-text-primary-dark mb-2">
                  备注
                </label>
                <textarea
                  value={formData.notes}
                  onChange={(e) => handleInputChange('notes', e.target.value)}
                  placeholder="可以记录洗护方式、搭配建议等..."
                  rows={4}
                  className="block w-full px-3 py-2 bg-card-light dark:bg-card-dark border border-border-light dark:border-border-dark rounded-lg focus:ring-2 focus:ring-primary focus:border-primary text-text-primary-light dark:text-text-primary-dark placeholder-text-secondary-light dark:placeholder-text-secondary-dark"
                />
              </div>
            </Card>
          </div>

          <div className="flex justify-end mt-8 space-x-2">
            <Button variant="secondary" onClick={handleCancel}>
              取消
            </Button>
            <Button onClick={handleSave} loading={loading} icon="save">
              保存衣物
            </Button>
          </div>
        </div>
      </main>

      <footer className="bg-card-light dark:bg-card-dark mt-auto border-t border-border-light dark:border-border-dark">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <p className="text-center text-sm text-text-secondary-light dark:text-text-secondary-dark">
            ©2024 StyleSense. All rights reserved.
          </p>
        </div>
      </footer>
    </div>
  );
}