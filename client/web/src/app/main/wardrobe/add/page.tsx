'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { useToast } from '@/components/ToastProvider';
import { Footer } from '@/components/Footer';

type ImageItem = {
  id: string;
  url: string;
  alt: string;
};

type TagItem = {
  id: string;
  label: string;
  type: string;
};

const mockImages: ImageItem[] = [
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

const aiSuggestedTags: TagItem[] = [
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
  const [images, setImages] = useState<ImageItem[]>(mockImages);
  const [suggestedTags, setSuggestedTags] = useState<TagItem[]>(aiSuggestedTags);
  const [appliedTags, setAppliedTags] = useState<string[]>([]);

  const handleInputChange = (name: string, value: string) => {
    setFormData(prev => ({
      ...prev,
      [name]: value,
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
    showToast('AI 标签已应用', 'success');
  };

  const handleSave = async () => {
    setLoading(true);
    try {
      await new Promise(resolve => setTimeout(resolve, 1500));
      showToast('衣物添加成功', 'success');
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
    <div className="flex min-h-screen flex-col bg-background-light font-display dark:bg-background-dark">
      <header className="sticky top-0 z-20 bg-card-light shadow-sm dark:bg-card-dark">
        <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div className="flex h-16 items-center justify-between">
            <div className="flex items-center space-x-8">
              <h1 className="text-xl font-bold text-text-light-primary dark:text-text-dark-primary">
                StyleSense
              </h1>
              <nav className="hidden space-x-8 md:flex">
                <button className="text-sm font-medium text-text-light-secondary transition-colors hover:text-primary dark:text-text-dark-secondary">
                  记录穿搭
                </button>
                <button className="border-b-2 border-primary text-sm font-medium text-primary">
                  我的衣橱
                </button>
                <button className="text-sm font-medium text-text-light-secondary transition-colors hover:text-primary dark:text-text-dark-secondary">
                  风格灵感
                </button>
                <button className="text-sm font-medium text-text-light-secondary transition-colors hover:text-primary dark:text-text-dark-secondary">
                  穿搭分析
                </button>
              </nav>
            </div>
            <div className="flex items-center space-x-4">
              <div className="relative hidden md:block">
                <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-text-light-secondary dark:text-text-dark-secondary">
                  search
                </span>
                <input
                  type="text"
                  placeholder="搜索衣物"
                  className="w-48 rounded-full border border-border-light bg-background-light py-2 pl-10 pr-4 text-sm text-text-light-primary focus:border-primary focus:ring-primary dark:border-border-dark dark:bg-background-dark dark:text-text-dark-primary"
                />
              </div>
              <button className="rounded-full p-2 text-text-light-secondary transition-colors hover:bg-gray-100 dark:text-text-dark-secondary dark:hover:bg-gray-800">
                <span className="material-icons-outlined">notifications_none</span>
              </button>
              <button className="h-8 w-8 overflow-hidden rounded-full">
                <img
                  src="https://lh3.googleusercontent.com/aida-public/AB6AXuA_qyQReAQqPnp4kLTe4O7X0yHidcXaj2eUnJyZ1VtAgg8kmw6iegfQgMEs2lGGO6nBxZos29reVGCsOSVzLa_LURh9KLREupnhZs9zsi-1CrBrM6Bbf25eIXlGpazovZNI0Xg8J37PVviC-gd0qT2Uj-SQJkN1ihlAb4-fBjEAKwYzFHMXFyHzL6MYZ1pI67jUsw6c5uJ1qqU3-_RZfPsnAu6JWFwAPVgKBP68lF2jyBKS-XrxKrL-9AEOHcFatvQEO6sOzYQxggKL"
                  alt="用户头像"
                  className="h-full w-full object-cover"
                />
              </button>
            </div>
          </div>
        </div>
      </header>

      <main className="container mx-auto flex-grow px-4 py-8 sm:px-6 lg:px-8">
        <div className="mx-auto max-w-4xl">
          <div className="mb-8 flex items-center justify-between">
            <h2 className="text-3xl font-bold text-gray-900 dark:text-white">添加新衣物</h2>
            <div className="flex space-x-2">
              <button
                onClick={handleCancel}
                className="rounded-lg border border-gray-300 bg-background-light px-4 py-2 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-gray-600"
              >
                取消
              </button>
              <button
                onClick={handleSave}
                disabled={loading}
                className="flex items-center space-x-2 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-gray-700 disabled:cursor-not-allowed disabled:opacity-70 dark:hover:bg-gray-600"
              >
                {loading ? (
                  <span className="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent" />
                ) : (
                  <span className="material-icons-outlined text-base">save</span>
                )}
                <span>保存衣物</span>
              </button>
            </div>
          </div>

          <div className="space-y-10">
            <section className="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-700 dark:bg-gray-800">
              <h3 className="mb-4 text-lg font-semibold text-gray-900 dark:text-white">衣物图片</h3>
              <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-5">
                {images.map(image => (
                  <div key={image.id} className="group relative aspect-square">
                    <img
                      src={image.url}
                      alt={image.alt}
                      className="h-full w-full rounded-lg object-cover"
                    />
                    <div className="absolute inset-0 flex items-center justify-center rounded-lg bg-black/40 opacity-0 transition-opacity group-hover:opacity-100">
                      <button
                        onClick={() => handleRemoveImage(image.id)}
                        className="text-white"
                      >
                        <span className="material-icons-outlined">delete</span>
                      </button>
                    </div>
                  </div>
                ))}
                <button
                  onClick={handleAddImage}
                  className="flex aspect-square flex-col items-center justify-center rounded-lg border-2 border-dashed border-gray-300 text-gray-400 transition-colors hover:border-blue-500 hover:text-blue-500 dark:border-gray-600"
                >
                  <span className="material-icons-outlined text-4xl">add_photo_alternate</span>
                  <p className="mt-1 text-xs">添加图片</p>
                </button>
              </div>
            </section>

            <section className="rounded-lg border border-green-500/30 bg-green-100/50 p-4 dark:border-green-500/30 dark:bg-green-900/20">
              <div className="flex items-start space-x-3">
                <span className="material-icons-outlined flex h-10 w-10 items-center justify-center rounded-full bg-white text-green-500 shadow-sm dark:bg-green-900">
                  auto_awesome
                </span>
                <div className="flex-1">
                  <div className="flex flex-col gap-1 md:flex-row md:items-center md:justify-between">
                    <div>
                      <h4 className="text-base font-semibold text-green-900 dark:text-green-100">
                        AI 标签建议
                      </h4>
                      <p className="text-sm text-green-700/80 dark:text-green-200/70">
                        根据您上传的图片，我们为您推荐了以下标签。您可以直接使用或进行修改。
                      </p>
                    </div>
                    <div className="mt-2 flex space-x-2 md:mt-0">
                      <button
                        onClick={handleApplyAllTags}
                        disabled={suggestedTags.length === 0}
                        className="rounded-lg bg-green-500 px-3 py-2 text-sm font-medium text-white transition-colors hover:bg-green-600 disabled:cursor-not-allowed disabled:opacity-60"
                      >
                        全部应用
                      </button>
                      <button
                        className="flex items-center space-x-1 rounded-lg border border-green-500 px-3 py-2 text-sm font-medium text-green-600 transition-colors hover:bg-green-500/10 dark:text-green-300"
                      >
                        <span className="material-icons-outlined text-base">edit</span>
                        <span>编辑</span>
                      </button>
                    </div>
                  </div>

                  <div className="mt-4 flex flex-wrap gap-2">
                    {suggestedTags.map(tag => (
                      <span
                        key={tag.id}
                        className="flex items-center space-x-1 rounded-full bg-white px-3 py-1 text-sm text-green-700 shadow-sm dark:bg-green-800 dark:text-green-200"
                      >
                        <span>{tag.label}</span>
                        <button
                          onClick={() => handleRemoveTag(tag.id)}
                          className="text-green-500 transition-colors hover:text-green-700 dark:text-green-300"
                        >
                          <span className="material-icons-outlined text-sm leading-none">close</span>
                        </button>
                      </span>
                    ))}
                    {suggestedTags.length === 0 && (
                      <span className="rounded-full bg-white px-3 py-1 text-sm text-green-700 shadow-sm dark:bg-green-800 dark:text-green-200">
                        所有标签已应用
                      </span>
                    )}
                  </div>

                  {appliedTags.length > 0 && (
                    <div className="mt-4 border-t border-green-500/20 pt-4">
                      <p className="mb-2 text-sm font-medium text-green-800 dark:text-green-100">
                        已应用标签
                      </p>
                      <div className="flex flex-wrap gap-2">
                        {appliedTags.map(tag => (
                          <span
                            key={tag}
                            className="rounded-full bg-green-500/20 px-3 py-1 text-sm text-green-700 dark:bg-green-700/40 dark:text-green-200"
                          >
                            {tag}
                          </span>
                        ))}
                      </div>
                    </div>
                  )}
                </div>
              </div>
            </section>

            <section className="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-700 dark:bg-gray-800">
              <h3 className="mb-6 text-lg font-semibold text-gray-900 dark:text-white">基本信息</h3>
              <div className="grid grid-cols-1 gap-x-6 gap-y-4 md:grid-cols-2">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    衣物名称
                  </label>
                  <input
                    value={formData.name}
                    onChange={event => handleInputChange('name', event.target.value)}
                    placeholder="例如：黑色T恤"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    分类
                  </label>
                  <select
                    value={formData.category}
                    onChange={event => handleInputChange('category', event.target.value)}
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                  >
                    <option value="上装">上装</option>
                    <option value="下装">下装</option>
                    <option value="鞋履">鞋履</option>
                    <option value="配饰">配饰</option>
                  </select>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    品牌
                  </label>
                  <input
                    value={formData.brand}
                    onChange={event => handleInputChange('brand', event.target.value)}
                    placeholder="例如：Uniqlo"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    颜色
                  </label>
                  <input
                    value={formData.color}
                    onChange={event => handleInputChange('color', event.target.value)}
                    placeholder="例如：黑色"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    材质
                  </label>
                  <input
                    value={formData.material}
                    onChange={event => handleInputChange('material', event.target.value)}
                    placeholder="例如：棉"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    风格
                  </label>
                  <input
                    value={formData.style}
                    onChange={event => handleInputChange('style', event.target.value)}
                    placeholder="例如：简约、休闲"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    尺码
                  </label>
                  <input
                    value={formData.size}
                    onChange={event => handleInputChange('size', event.target.value)}
                    placeholder="例如：M / 170"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
              </div>
            </section>

            <section className="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-700 dark:bg-gray-800">
              <h3 className="mb-6 text-lg font-semibold text-gray-900 dark:text-white">适用场景</h3>
              <div className="grid grid-cols-1 gap-x-6 gap-y-4 md:grid-cols-2">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    季节
                  </label>
                  <input
                    value={formData.season}
                    onChange={event => handleInputChange('season', event.target.value)}
                    placeholder="例如：春季"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    场合
                  </label>
                  <input
                    value={formData.occasion}
                    onChange={event => handleInputChange('occasion', event.target.value)}
                    placeholder="例如：日常通勤、周末出游"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
              </div>
            </section>

            <section className="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-700 dark:bg-gray-800">
              <h3 className="mb-6 text-lg font-semibold text-gray-900 dark:text-white">购买信息</h3>
              <div className="grid grid-cols-1 gap-x-6 gap-y-4 md:grid-cols-3">
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    价格
                  </label>
                  <div className="relative mt-1">
                    <div className="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
                      <span className="text-gray-500 sm:text-sm">¥</span>
                    </div>
                    <input
                      value={formData.price}
                      onChange={event => handleInputChange('price', event.target.value)}
                      placeholder="0.00"
                      className="block w-full rounded-lg border border-gray-300 bg-gray-50 pl-7 pr-3 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                      type="text"
                    />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    商店
                  </label>
                  <input
                    value={formData.store}
                    onChange={event => handleInputChange('store', event.target.value)}
                    placeholder="例如：淘宝官方旗舰店"
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="text"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                    购买日期
                  </label>
                  <input
                    value={formData.purchaseDate}
                    onChange={event => handleInputChange('purchaseDate', event.target.value)}
                    className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                    type="date"
                  />
                </div>
              </div>
            </section>

            <section className="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-700 dark:bg-gray-800">
              <h3 className="mb-6 text-lg font-semibold text-gray-900 dark:text-white">其他备注</h3>
              <div>
                <label className="block text-sm font-medium text-gray-700 dark:text-gray-300">
                  备注
                </label>
                <textarea
                  value={formData.notes}
                  onChange={event => handleInputChange('notes', event.target.value)}
                  placeholder="可以记录洗护方式、搭配建议等..."
                  rows={4}
                  className="mt-1 block w-full rounded-lg border border-gray-300 bg-gray-50 px-3 py-2 text-sm shadow-sm focus:border-primary focus:ring-primary dark:border-gray-600 dark:bg-gray-700 dark:text-white"
                />
              </div>
            </section>
          </div>

          <div className="mt-8 flex justify-end space-x-2">
            <button
              onClick={handleCancel}
              className="rounded-lg border border-gray-300 bg-background-light px-4 py-2 text-sm font-medium text-gray-700 transition-colors hover:bg-gray-50 dark:border-gray-600 dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-gray-600"
            >
              取消
            </button>
            <button
              onClick={handleSave}
              disabled={loading}
              className="flex items-center space-x-2 rounded-lg bg-primary px-4 py-2 text-sm font-medium text-white transition-colors hover:bg-gray-700 disabled:cursor-not-allowed disabled:opacity-70 dark:hover:bg-gray-600"
            >
              {loading ? (
                <span className="h-4 w-4 animate-spin rounded-full border-2 border-white border-t-transparent" />
              ) : (
                <span className="material-icons-outlined text-base">save</span>
              )}
              <span>保存衣物</span>
            </button>
          </div>
        </div>
      </main>

      <Footer variant="card" />
    </div>
  );
}
