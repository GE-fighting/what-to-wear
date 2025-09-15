"use client";
import React, { useState, useRef, useEffect } from 'react';
import '@/styles/AddClothingItem.css';
import type { ClothingItemData } from "@/types/clothing";
import type { ClothingCategory } from "@/types/clothing";
import { getClothingCategories } from "@/lib/api/clothing";

interface AddClothingItemProps {
  onSubmit: (data: ClothingItemData) => Promise<void> | void;
  onCancel: () => void;
}

export function AddClothingItem({ onSubmit, onCancel }: AddClothingItemProps) {
  const [formData, setFormData] = useState<ClothingItemData>({
    name: '',
    categoryId: '',
    brand: '',
    color: '',
    size: '',
    sizeSystem: 'CN',
    material: '',
    price: '',
    purchaseDate: '',
    notes: '',
    isFavorite: false,
    specificAttributes: {}
  });

  const [images, setImages] = useState<File[]>([]);
  const [imagePreviews, setImagePreviews] = useState<string[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const [categoryTree, setCategoryTree] = useState<ClothingCategory[]>([]);
  const [categoriesLoading, setCategoriesLoading] = useState(false);
  const [categoriesError, setCategoriesError] = useState<string | null>(null);

  const fetchCategories = async () => {
    setCategoriesLoading(true);
    setCategoriesError(null);
    try {
      const data = await getClothingCategories();
      setCategoryTree(Array.isArray(data) ? data : []);
    } catch (err) {
      const message = err instanceof Error ? err.message : '未知错误';
      setCategoriesError(message);
    } finally {
      setCategoriesLoading(false);
    }
  };

  useEffect(() => {
    fetchCategories();
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    
    if (type === 'checkbox') {
      const checked = (e.target as HTMLInputElement).checked;
      setFormData(prev => ({ ...prev, [name]: checked }));
    } else if (name.startsWith('attr_')) {
      const attrName = name.replace('attr_', '');
      setFormData(prev => ({
        ...prev,
        specificAttributes: {
          ...prev.specificAttributes,
          [attrName]: value
        }
      }));
    } else {
      setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleImageUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = Array.from(e.target.files || []);
    if (files.length === 0) return;

    const newImages = [...images, ...files].slice(0, 5);
    setImages(newImages);

    const newPreviews: string[] = [];
    newImages.forEach(file => {
      const reader = new FileReader();
      reader.onload = (e) => {
        newPreviews.push(e.target?.result as string);
        if (newPreviews.length === newImages.length) {
          setImagePreviews(newPreviews);
        }
      };
      reader.readAsDataURL(file);
    });
  };

  const removeImage = (index: number) => {
    const newImages = images.filter((_, i) => i !== index);
    const newPreviews = imagePreviews.filter((_, i) => i !== index);
    setImages(newImages);
    setImagePreviews(newPreviews);
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    e.currentTarget.classList.add('dragover');
  };

  const handleDragLeave = (e: React.DragEvent) => {
    e.preventDefault();
    e.currentTarget.classList.remove('dragover');
  };

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    e.currentTarget.classList.remove('dragover');
    
    const files = Array.from(e.dataTransfer.files);
    const imageFiles = files.filter(file => file.type.startsWith('image/'));
    
    if (imageFiles.length > 0) {
      const newImages = [...images, ...imageFiles].slice(0, 5);
      setImages(newImages);

      const newPreviews: string[] = [];
      newImages.forEach(file => {
        const reader = new FileReader();
        reader.onload = (e) => {
          newPreviews.push(e.target?.result as string);
          if (newPreviews.length === newImages.length) {
            setImagePreviews(newPreviews);
          }
        };
        reader.readAsDataURL(file);
      });
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsSubmitting(true);

    try {
      await onSubmit(formData);
    } catch (error) {
      console.error('提交失败:', error);
    } finally {
      setIsSubmitting(false);
    }
  };

  const renderCategoryOptions = (categories: ClothingCategory[], depth = 0): React.ReactNode[] => {
    const indent = depth > 0 ? '—'.repeat(depth) + ' ' : '';
    return categories.flatMap((cat) => {
      const label = `${cat.icon ? cat.icon + ' ' : ''}${indent}${cat.name}`;
      const current = (
        <option key={cat.id} value={String(cat.id)}>
          {label}
        </option>
      );
      const children = Array.isArray(cat.children) && cat.children.length > 0
        ? renderCategoryOptions(cat.children, depth + 1)
        : [];
      return [current, ...children];
    });
  };

  const renderDynamicAttributes = () => {
    const categoryId = formData.categoryId;
    
    if (!categoryId) return null;

    switch (categoryId) {
      case '1': // 上衣
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">上衣属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">袖长</label>
                <select 
                  className="form-select" 
                  name="attr_sleeveLength" 
                  value={formData.specificAttributes.sleeveLength || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="短袖">短袖</option>
                  <option value="长袖">长袖</option>
                  <option value="七分袖">七分袖</option>
                  <option value="无袖">无袖</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">领型</label>
                <select 
                  className="form-select" 
                  name="attr_neckline" 
                  value={formData.specificAttributes.neckline || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="圆领">圆领</option>
                  <option value="V领">V领</option>
                  <option value="高领">高领</option>
                  <option value="翻领">翻领</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">版型</label>
                <select 
                  className="form-select" 
                  name="attr_fit" 
                  value={formData.specificAttributes.fit || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="修身">修身</option>
                  <option value="宽松">宽松</option>
                  <option value="标准">标准</option>
                  <option value="紧身">紧身</option>
                </select>
              </div>
            </div>
          </div>
        );
      
      case '2': // 下装
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">下装属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">长度</label>
                <select 
                  className="form-select" 
                  name="attr_length" 
                  value={formData.specificAttributes.length || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="长裤">长裤</option>
                  <option value="短裤">短裤</option>
                  <option value="七分裤">七分裤</option>
                  <option value="九分裤">九分裤</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">腰型</label>
                <select 
                  className="form-select" 
                  name="attr_waistType" 
                  value={formData.specificAttributes.waistType || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="高腰">高腰</option>
                  <option value="中腰">中腰</option>
                  <option value="低腰">低腰</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">开合方式</label>
                <select 
                  className="form-select" 
                  name="attr_closure" 
                  value={formData.specificAttributes.closure || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="拉链">拉链</option>
                  <option value="纽扣">纽扣</option>
                  <option value="松紧带">松紧带</option>
                  <option value="系带">系带</option>
                </select>
              </div>
            </div>
          </div>
        );
      
      case '4': // 鞋子
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">鞋子属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">跟高</label>
                <select 
                  className="form-select" 
                  name="attr_heelHeight" 
                  value={formData.specificAttributes.heelHeight || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="平底">平底</option>
                  <option value="低跟">低跟</option>
                  <option value="中跟">中跟</option>
                  <option value="高跟">高跟</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">鞋型</label>
                <select 
                  className="form-select" 
                  name="attr_shoeType" 
                  value={formData.specificAttributes.shoeType || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="运动鞋">运动鞋</option>
                  <option value="皮鞋">皮鞋</option>
                  <option value="靴子">靴子</option>
                  <option value="凉鞋">凉鞋</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">适用场合</label>
                <select 
                  className="form-select" 
                  name="attr_occasion" 
                  value={formData.specificAttributes.occasion || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="日常">日常</option>
                  <option value="运动">运动</option>
                  <option value="正式">正式</option>
                  <option value="休闲">休闲</option>
                </select>
              </div>
            </div>
          </div>
        );
      
      case '5': // 配饰
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">配饰属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">配饰类型</label>
                <select 
                  className="form-select" 
                  name="attr_accessoryType" 
                  value={formData.specificAttributes.accessoryType || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="帽子">帽子</option>
                  <option value="包包">包包</option>
                  <option value="首饰">首饰</option>
                  <option value="围巾">围巾</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">风格</label>
                <select 
                  className="form-select" 
                  name="attr_style" 
                  value={formData.specificAttributes.style || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="简约">简约</option>
                  <option value="复古">复古</option>
                  <option value="时尚">时尚</option>
                  <option value="优雅">优雅</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">适用季节</label>
                <select 
                  className="form-select" 
                  name="attr_season" 
                  value={formData.specificAttributes.season || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="春季">春季</option>
                  <option value="夏季">夏季</option>
                  <option value="秋季">秋季</option>
                  <option value="冬季">冬季</option>
                  <option value="四季">四季</option>
                </select>
              </div>
            </div>
          </div>
        );
      
      default:
        return null;
    }
  };

  return (
    <form className="add-clothing-form" onSubmit={handleSubmit}>
      {/* 基本信息 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">📝</span>
          基本信息
        </h2>
        <div className="form-grid">
          <div className="form-group">
            <label className="form-label">
              衣物名称 <span className="required">*</span>
            </label>
            <input 
              type="text" 
              className="form-input" 
              name="name" 
              required 
              placeholder="例如：白色棉质T恤"
              value={formData.name}
              onChange={handleInputChange}
            />
            <div className="form-hint">给您的衣物起个容易识别的名字</div>
          </div>
          <div className="form-group">
            <label className="form-label">
              分类 <span className="required">*</span>
            </label>
            <select 
              className="form-select" 
              name="categoryId" 
              required
              value={formData.categoryId}
              onChange={handleInputChange}
              disabled={categoriesLoading}
            >
              <option value="">{categoriesLoading ? '分类加载中...' : '请选择分类'}</option>
              {renderCategoryOptions(categoryTree)}
            </select>
            {categoriesError && (
              <div className="form-hint" style={{ color: '#d33' }}>
                分类加载失败：{categoriesError}
                <button type="button" className="btn btn-secondary" style={{ marginLeft: 8 }} onClick={fetchCategories}>重试</button>
              </div>
            )}
          </div>
          <div className="form-group">
            <label className="form-label">品牌</label>
            <input 
              type="text" 
              className="form-input" 
              name="brand" 
              placeholder="例如：Uniqlo"
              value={formData.brand}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-group">
            <label className="form-label">
              颜色 <span className="required">*</span>
            </label>
            <input 
              type="text" 
              className="form-input" 
              name="color" 
              required 
              placeholder="例如：白色"
              value={formData.color}
              onChange={handleInputChange}
            />
          </div>
        </div>
      </div>

      {/* 尺码信息 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">📏</span>
          尺码信息
        </h2>
        <div className="form-grid">
          <div className="form-group">
            <label className="form-label">尺码</label>
            <div className="size-group">
              <input 
                type="text" 
                className="form-input" 
                name="size" 
                placeholder="例如：M, L, 38, 40"
                value={formData.size}
                onChange={handleInputChange}
              />
              <select 
                className="form-select" 
                name="sizeSystem"
                value={formData.sizeSystem}
                onChange={handleInputChange}
              >
                <option value="CN">中国码</option>
                <option value="US">美国码</option>
                <option value="EU">欧洲码</option>
                <option value="UK">英国码</option>
              </select>
            </div>
          </div>
          <div className="form-group">
            <label className="form-label">材质</label>
            <select 
              className="form-select" 
              name="material"
              value={formData.material}
              onChange={handleInputChange}
            >
              <option value="">请选择材质</option>
              <option value="棉">棉</option>
              <option value="真皮">真皮</option>
              <option value="羊毛">羊毛</option>
              <option value="聚酯纤维">聚酯纤维</option>
              <option value="尼龙">尼龙</option>
              <option value="丝绸">丝绸</option>
              <option value="亚麻">亚麻</option>
              <option value="混纺">混纺</option>
            </select>
          </div>
        </div>
      </div>

      {/* 动态属性 */}
      {renderDynamicAttributes()}

      {/* 图片上传 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">📷</span>
          衣物图片
        </h2>
        <div 
          className="image-upload"
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
          onDrop={handleDrop}
          onClick={() => fileInputRef.current?.click()}
        >
          <div className="upload-icon">📷</div>
          <div className="upload-text">点击上传或拖拽图片到此处</div>
          <div className="upload-hint">支持 JPG、PNG 格式，最多5张图片</div>
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            multiple
            style={{ display: 'none' }}
            onChange={handleImageUpload}
          />
        </div>
        
        {imagePreviews.length > 0 && (
          <div className="image-preview">
            {imagePreviews.map((preview, index) => (
              <div key={index} className="preview-item">
                <img src={preview} alt={`预览 ${index + 1}`} className="preview-image" />
                <button 
                  type="button"
                  className="remove-image" 
                  onClick={() => removeImage(index)}
                >
                  ×
                </button>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* 其他信息 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">📋</span>
          其他信息
        </h2>
        <div className="form-grid">
          <div className="form-group full-width">
            <label className="form-label">备注</label>
            <textarea 
              className="form-textarea" 
              name="notes" 
              placeholder="记录关于这件衣物的其他信息..."
              value={formData.notes}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-group">
            <div className="checkbox-group">
              <input 
                type="checkbox" 
                className="checkbox" 
                name="isFavorite"
                checked={formData.isFavorite}
                onChange={handleInputChange}
              />
              <label className="checkbox-label">标记为收藏</label>
            </div>
          </div>
        </div>
      </div>

      {/* 表单操作 */}
      <div className="form-actions">
        <button type="button" className="btn btn-secondary" onClick={onCancel}>
          取消
        </button>
        <button type="submit" className="btn btn-primary" disabled={isSubmitting}>
          {isSubmitting ? (
            <div className="loading">
              <div className="spinner"></div>
              <span>保存中...</span>
            </div>
          ) : (
            <>
              <span>💾</span>
              <span>保存衣物</span>
            </>
          )}
        </button>
      </div>
    </form>
  );
}
