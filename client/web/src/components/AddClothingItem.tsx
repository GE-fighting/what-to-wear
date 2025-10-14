"use client";
import React, { useState, useRef, useEffect } from 'react';
import '@/styles/AddClothingItem.css';
import type { ClothingItemData, ClothingCategory, ClothingStatus, Tag } from "@/types/clothing";
import { getClothingCategories, getSystemTagEnums, createClothingItem } from "@/lib/api/clothing";
import { showToast } from "@/components/Toast";

interface AddClothingItemProps {
  onSubmit: (data: ClothingItemData) => Promise<void> | void;
  onCancel: () => void;
}

export function AddClothingItem({ onSubmit, onCancel }: AddClothingItemProps) {
  const [formData, setFormData] = useState<ClothingItemData>({
    name: '',
    category_id: 0,
    category_name: '',
    brand: '',
    color: '',
    size: '',
    material: '',
    season: [],
    occasion: [],
    style: '',
    description: '',
    tags: [],
    tag_names: [],
    status: 'active' as ClothingStatus,
    is_favorite: false,
    purchase_info: null,
    specific_attributes: {}
  });

  const [images, setImages] = useState<File[]>([]);
  const [imagePreviews, setImagePreviews] = useState<string[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const categoryDropdownRef = useRef<HTMLDivElement>(null);

  const [categoryTree, setCategoryTree] = useState<ClothingCategory[]>([]);
  const [categoriesLoading, setCategoriesLoading] = useState(false);
  const [categoriesError, setCategoriesError] = useState<string | null>(null);
  const [selectedParentId, setSelectedParentId] = useState<string>('');
  const [categoryError, setCategoryError] = useState<string | null>(null);
  const [categoryDropdownOpen, setCategoryDropdownOpen] = useState(false);
  const [expandedParents, setExpandedParents] = useState<Set<string>>(new Set());
  const [showPurchaseInfo, setShowPurchaseInfo] = useState(false);
  
  // 动态加载的选项
  const [seasonOptions, setSeasonOptions] = useState<Tag[]>([]);
  const [occasionOptions, setOccasionOptions] = useState<Tag[]>([]);
  const [styleOptions, setStyleOptions] = useState<Tag[]>([]);
  const [tagEnumsLoading, setTagEnumsLoading] = useState(false);
  const [tagEnumsError, setTagEnumsError] = useState<string | null>(null);
  const statusOptions = [
    { value: 'active', label: '在用', icon: '✅' },
    { value: 'inactive', label: '闲置', icon: '⏸️' },
    { value: 'donated', label: '已捐赠', icon: '💝' },
    { value: 'sold', label: '已出售', icon: '💰' },
    { value: 'lost', label: '丢失', icon: '❌' },
    { value: 'damaged', label: '损坏', icon: '🔧' }
  ];

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

  const fetchTagEnums = async () => {
    setTagEnumsLoading(true);
    setTagEnumsError(null);
    try {
      const [seasonData, occasionData, styleData] = await Promise.all([
        getSystemTagEnums('season'),
        getSystemTagEnums('occasion'),
        getSystemTagEnums('style')
      ]);
      setSeasonOptions(Array.isArray(seasonData) ? seasonData : []);
      setOccasionOptions(Array.isArray(occasionData) ? occasionData : []);
      setStyleOptions(Array.isArray(styleData) ? styleData : []);
    } catch (err) {
      const message = err instanceof Error ? err.message : '未知错误';
      setTagEnumsError(message);
    } finally {
      setTagEnumsLoading(false);
    }
  };

  useEffect(() => {
    fetchCategories();
    fetchTagEnums();
  }, []);

  // 点击外部关闭下拉框
  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (categoryDropdownRef.current && !categoryDropdownRef.current.contains(event.target as Node)) {
        setCategoryDropdownOpen(false);
      }
    };

    if (categoryDropdownOpen) {
      document.addEventListener('mousedown', handleClickOutside);
      return () => document.removeEventListener('mousedown', handleClickOutside);
    }
  }, [categoryDropdownOpen]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    
    if (type === 'checkbox') {
      const checked = (e.target as HTMLInputElement).checked;
      setFormData(prev => ({ ...prev, [name]: checked }));
    } else if (name.startsWith('attr_')) {
      const attrName = name.replace('attr_', '');
      setFormData(prev => ({
        ...prev,
        specific_attributes: {
          ...prev.specific_attributes,
          [attrName]: value
        }
      }));
    } else if (name.startsWith('purchase_')) {
      const purchaseField = name.replace('purchase_', '');
      setFormData(prev => ({
        ...prev,
        purchase_info: {
          ...(prev.purchase_info || { price: 0, store: '', purchase_date: '', notes: '' }),
          [purchaseField]: purchaseField === 'price' ? parseFloat(value) || 0 : value
        }
      }));
    } else {
      setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleMultiSelect = (name: string, value: string) => {
    setFormData(prev => {
      const currentArray = prev[name as keyof ClothingItemData] as string[];
      const newArray = currentArray.includes(value) 
        ? currentArray.filter(item => item !== value)
        : [...currentArray, value];
      return { ...prev, [name]: newArray };
    });
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
    if (!formData.category_id) {
      setCategoryError('请选择分类');
      return;
    }
    setIsSubmitting(true);

    try {
      // 处理购买信息：如果都是空的，设为 null
      const cleanedFormData = { ...formData };
      if (cleanedFormData.purchase_info) {
        const { price, store, purchase_date, notes } = cleanedFormData.purchase_info;
        const isEmpty = price === 0 && !store.trim() && !purchase_date.trim() && !notes.trim();
        if (isEmpty) {
          cleanedFormData.purchase_info = null;
        }
      }

      // 收集季节和场合的 ID，放入 tags 字段
      const seasonIds = seasonOptions
        .filter(season => formData.season.includes(season.name))
        .map(season => season.id);
      
      const occasionIds = occasionOptions
        .filter(occasion => formData.occasion.includes(occasion.name))
        .map(occasion => occasion.id);

      // 合并所有标签 ID
      cleanedFormData.tags = [...seasonIds, ...occasionIds];

      // 收集标签名称
      const seasonNames = formData.season;
      const occasionNames = formData.occasion;
      cleanedFormData.tag_names = [...seasonNames, ...occasionNames];

      // 调用 API 创建衣物
      const result = await createClothingItem(cleanedFormData);
      console.log('衣物创建成功:', result);

      // 显示成功提示
      showToast('✅ 衣物保存成功！', 'success', 3000);

      // 调用父组件的 onSubmit 回调
      await onSubmit(cleanedFormData);
    } catch (error) {
      console.error('提交失败:', error);
      // 显示错误提示
      showToast('保存失败，请重试', 'error', 4000);
    } finally {
      setIsSubmitting(false);
    }
  };

  const getTopLevelCategoryId = React.useCallback((id: number): string | null => {
    if (!id) return null;
    for (const parent of categoryTree) {
      if (parent.id === id) return String(parent.id);
      const children = Array.isArray(parent.children) ? parent.children : [];
      if (children.some(child => child.id === id)) {
        return String(parent.id);
      }
    }
    return null;
  }, [categoryTree]);

  const handleCategorySelect = (categoryId: string, parentId?: string) => {
    // 查找分类名称
    let categoryName = '';
    for (const parent of categoryTree) {
      if (String(parent.id) === categoryId) {
        categoryName = parent.name;
        break;
      }
      const children = Array.isArray(parent.children) ? parent.children : [];
      for (const child of children) {
        if (String(child.id) === categoryId) {
          categoryName = child.name;
          break;
        }
      }
      if (categoryName) break;
    }
    
    setFormData(prev => ({ ...prev, category_id: parseInt(categoryId), category_name: categoryName }));
    if (parentId) {
      setSelectedParentId(parentId);
    } else {
      setSelectedParentId(categoryId);
    }
    setCategoryError(null);
    setCategoryDropdownOpen(false);
  };

  const handleParentClick = (parent: ClothingCategory) => {
    const parentId = String(parent.id);
    const hasChildren = Array.isArray(parent.children) && parent.children.length > 0;
    
    if (!hasChildren) {
      // 没有子分类，直接选中
      handleCategorySelect(parentId);
    } else {
      // 有子分类，切换展开/收起状态
      setExpandedParents(prev => {
        const newSet = new Set(prev);
        if (newSet.has(parentId)) {
          newSet.delete(parentId);
        } else {
          newSet.add(parentId);
        }
        return newSet;
      });
    }
  };

  const getSelectedCategoryDisplay = () => {
    if (!formData.category_id) return '请选择分类';
    
    for (const parent of categoryTree) {
      if (parent.id === formData.category_id) {
        return `${parent.icon ? parent.icon + ' ' : ''}${parent.name}`;
      }
      const children = Array.isArray(parent.children) ? parent.children : [];
      for (const child of children) {
        if (child.id === formData.category_id) {
          return `${parent.icon ? parent.icon + ' ' : ''}${parent.name} > ${child.icon ? child.icon + ' ' : ''}${child.name}`;
        }
      }
    }
    return '请选择分类';
  };

  const renderDynamicAttributes = () => {
    const topLevelId = getTopLevelCategoryId(formData.category_id);
    if (!topLevelId) return null;

    switch (topLevelId) {
      case '1': // 上衣
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">上衣属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">袖长</label>
                <select 
                  className="form-select" 
                  name="attr_sleeve_length" 
                  value={(formData.specific_attributes as any)?.sleeve_length || ''}
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
                  value={(formData.specific_attributes as any)?.neckline || ''}
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
                  value={(formData.specific_attributes as any)?.fit || ''}
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
                  value={(formData.specific_attributes as any)?.length || ''}
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
                  name="attr_waist_type" 
                  value={(formData.specific_attributes as any)?.waist_type || ''}
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
                  value={(formData.specific_attributes as any)?.closure || ''}
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
      
      case '3': // 鞋子
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">鞋子属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">跟高</label>
                <select 
                  className="form-select" 
                  name="attr_heel_height" 
                  value={(formData.specific_attributes as any)?.heel_height || ''}
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
                  name="attr_shoe_type" 
                  value={(formData.specific_attributes as any)?.shoe_type || ''}
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
                <label className="form-label">鞋头形状</label>
                <select 
                  className="form-select" 
                  name="attr_toe_shape" 
                  value={(formData.specific_attributes as any)?.toe_shape || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="圆头">圆头</option>
                  <option value="尖头">尖头</option>
                  <option value="方头">方头</option>
                  <option value="杏仁头">杏仁头</option>
                </select>
              </div>
            </div>
          </div>
        );
      
      case '4': // 配饰
        return (
          <div className="dynamic-attributes show">
            <h3 className="attribute-title">配饰属性</h3>
            <div className="attribute-group">
              <div className="form-group">
                <label className="form-label">配饰材质</label>
                <select 
                  className="form-select" 
                  name="attr_accessory_material" 
                  value={(formData.specific_attributes as any)?.accessory_material || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="金属">金属</option>
                  <option value="皮革">皮革</option>
                  <option value="布料">布料</option>
                  <option value="塑料">塑料</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">配饰尺寸</label>
                <select 
                  className="form-select" 
                  name="attr_accessory_size" 
                  value={(formData.specific_attributes as any)?.accessory_size || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="小号">小号</option>
                  <option value="中号">中号</option>
                  <option value="大号">大号</option>
                  <option value="特大">特大</option>
                </select>
              </div>
              <div className="form-group">
                <label className="form-label">配饰功能</label>
                <select 
                  className="form-select" 
                  name="attr_accessory_function" 
                  value={(formData.specific_attributes as any)?.accessory_function || ''}
                  onChange={handleInputChange}
                >
                  <option value="">请选择</option>
                  <option value="装饰">装饰</option>
                  <option value="实用">实用</option>
                  <option value="保暖">保暖</option>
                  <option value="防晒">防晒</option>
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
            <div className="category-selector" ref={categoryDropdownRef}>
              <div 
                className={`category-trigger ${categoryDropdownOpen ? 'open' : ''}`}
                onClick={() => !categoriesLoading && setCategoryDropdownOpen(!categoryDropdownOpen)}
              >
                <span className="category-display">
                  {categoriesLoading ? '分类加载中...' : getSelectedCategoryDisplay()}
                </span>
                <span className="dropdown-arrow">
                  {categoryDropdownOpen ? '▲' : '▼'}
                </span>
              </div>
              
              {categoryDropdownOpen && !categoriesLoading && (
                <div className="category-dropdown">
                  <div className="category-grid">
                    {categoryTree.map((parent) => {
                      const parentId = String(parent.id);
                      const hasChildren = Array.isArray(parent.children) && parent.children.length > 0;
                      const isExpanded = expandedParents.has(parentId);
                      
                      return (
                        <div key={parent.id} className="category-group">
                          <div 
                            className={`category-parent ${parent.id === formData.category_id ? 'selected' : ''} ${hasChildren ? 'expandable' : ''}`}
                            onClick={() => handleParentClick(parent)}
                          >
                            <span className="category-icon">{parent.icon}</span>
                            <span className="category-name">{parent.name}</span>
                            {hasChildren && (
                              <span className={`expand-arrow ${isExpanded ? 'expanded' : ''}`}>▼</span>
                            )}
                          </div>
                          {hasChildren && isExpanded && (
                            <div className="category-children">
                              {parent.children!.map((child) => (
                                <div 
                                  key={child.id}
                                  className={`category-child ${child.id === formData.category_id ? 'selected' : ''}`}
                                  onClick={() => handleCategorySelect(String(child.id), String(parent.id))}
                                >
                                  <span className="category-icon">{child.icon}</span>
                                  <span className="category-name">{child.name}</span>
                                </div>
                              ))}
                            </div>
                          )}
                        </div>
                      );
                    })}
                  </div>
                </div>
              )}
            </div>
            {categoriesError && (
              <div className="form-hint" style={{ color: '#d33' }}>
                分类加载失败：{categoriesError}
                <button type="button" className="btn btn-secondary" style={{ marginLeft: 8 }} onClick={fetchCategories}>重试</button>
              </div>
            )}
            {categoryError && (
              <div className="form-hint" style={{ color: '#d33' }}>{categoryError}</div>
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
          <div className="form-group">
            <label className="form-label">风格</label>
            <select 
              className="form-select" 
              name="style"
              value={formData.style}
              onChange={handleInputChange}
            >
              <option value="">请选择风格</option>
              {styleOptions.map(style => (
                <option key={style.id} value={style.name}>{style.name}</option>
              ))}
            </select>
          </div>
        </div>
      </div>

      {/* 尺码与状态 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">📏</span>
          尺码与状态
        </h2>
        <div className="form-grid">
          <div className="form-group">
            <label className="form-label">尺码</label>
            <input 
              type="text" 
              className="form-input" 
              name="size" 
              placeholder="例如：M, L, 38, 40"
              value={formData.size}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-group">
            <label className="form-label">状态</label>
            <select 
              className="form-select" 
              name="status"
              value={formData.status}
              onChange={handleInputChange}
            >
              {statusOptions.map(option => (
                <option key={option.value} value={option.value}>
                  {option.icon} {option.label}
                </option>
              ))}
            </select>
            <div className="form-hint">设置衣物的当前状态</div>
          </div>
        </div>
        {tagEnumsError && (
          <div className="form-hint" style={{ color: '#d33' }}>
            标签加载失败：{tagEnumsError}
            <button type="button" className="btn btn-secondary" style={{ marginLeft: 8 }} onClick={fetchTagEnums}>重试</button>
          </div>
        )}
      </div>

      {/* 动态属性 */}
      {renderDynamicAttributes()}

      {/* 适用场景 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">🌟</span>
          适用场景
        </h2>
        <div className="form-grid">
          <div className="form-group">
            <label className="form-label">适用季节</label>
            <div className="multi-select-container">
              {tagEnumsLoading ? (
                <div className="loading-text">加载中...</div>
              ) : seasonOptions.length > 0 ? (
                seasonOptions.map(season => (
                  <label key={season.id} className="multi-select-item">
                    <input
                      type="checkbox"
                      checked={formData.season.includes(season.name)}
                      onChange={() => handleMultiSelect('season', season.name)}
                    />
                    <span className="multi-select-label">{season.name}</span>
                  </label>
                ))
              ) : (
                <div className="form-hint" style={{ color: '#666' }}>暂无季节选项</div>
              )}
            </div>
            <div className="form-hint">可以选择多个季节</div>
          </div>
          <div className="form-group">
            <label className="form-label">适用场合</label>
            <div className="multi-select-container">
              {tagEnumsLoading ? (
                <div className="loading-text">加载中...</div>
              ) : occasionOptions.length > 0 ? (
                occasionOptions.map(occasion => (
                  <label key={occasion.id} className="multi-select-item">
                    <input
                      type="checkbox"
                      checked={formData.occasion.includes(occasion.name)}
                      onChange={() => handleMultiSelect('occasion', occasion.name)}
                    />
                    <span className="multi-select-label">{occasion.name}</span>
                  </label>
                ))
              ) : (
                <div className="form-hint" style={{ color: '#666' }}>暂无场合选项</div>
              )}
            </div>
            <div className="form-hint">可以选择多个场合</div>
          </div>
        </div>
      </div>

      {/* 购买信息 */}
      <div className="form-section">
        <h2 className="section-title">
          <span className="section-icon">💰</span>
          购买信息
          <button 
            type="button" 
            className="btn btn-link"
            onClick={() => setShowPurchaseInfo(!showPurchaseInfo)}
          >
            {showPurchaseInfo ? '隐藏' : '显示'}
          </button>
        </h2>
        {showPurchaseInfo && (
          <div className="form-grid">
            <div className="form-group">
              <label className="form-label">购买价格</label>
              <div className="price-input">
                <input 
                  type="number" 
                  className="form-input" 
                  name="purchase_price" 
                  placeholder="0.00"
                  step="0.01"
                  min="0"
                  value={formData.purchase_info?.price || ''}
                  onChange={handleInputChange}
                />
              </div>
            </div>
            <div className="form-group">
              <label className="form-label">购买商店</label>
              <input 
                type="text" 
                className="form-input" 
                name="purchase_store" 
                placeholder="例如：Uniqlo官网"
                value={formData.purchase_info?.store || ''}
                onChange={handleInputChange}
              />
            </div>
            <div className="form-group">
              <label className="form-label">购买日期</label>
              <input 
                type="date" 
                className="form-input" 
                name="purchase_purchase_date" 
                value={formData.purchase_info?.purchase_date || ''}
                onChange={handleInputChange}
              />
            </div>
            <div className="form-group full-width">
              <label className="form-label">购买备注</label>
              <textarea 
                className="form-textarea" 
                name="purchase_notes" 
                placeholder="记录折扣、促销信息等..."
                value={formData.purchase_info?.notes || ''}
                onChange={handleInputChange}
              />
            </div>
          </div>
        )}
      </div>

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
            <label className="form-label">描述</label>
            <textarea 
              className="form-textarea" 
              name="description" 
              placeholder="记录关于这件衣物的其他信息..."
              value={formData.description}
              onChange={handleInputChange}
            />
          </div>
          <div className="form-group">
            <div className="checkbox-group">
              <input 
                type="checkbox" 
                className="checkbox" 
                name="is_favorite"
                checked={formData.is_favorite}
                onChange={handleInputChange}
              />
              <label className="checkbox-label">⭐ 标记为收藏</label>
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
