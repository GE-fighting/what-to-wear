'use client';

import { useState, useEffect, useCallback } from 'react';
import { useToast } from '@/components/ToastProvider';

// 模拟衣物数据类型
export interface ClothingItem {
  id: string;
  name: string;
  category: string;
  season: string;
  image: string;
  tags: string[];
  wearCount: number;
  brand?: string;
  color?: string;
  material?: string;
  size?: string;
  price?: string;
  purchaseDate?: string;
  notes?: string;
}

export interface WardrobeStats {
  totalItems: number;
  totalValue: string;
  toClean: number;
  weeklyNew: number;
  unworn30Days: number;
  avgPrice: string;
}

export interface WardrobeFilters {
  category: string;
  season: string;
  searchTerm: string;
  sortBy: 'name' | 'category' | 'wearCount' | 'dateAdded';
  sortOrder: 'asc' | 'desc';
}

// 模拟数据
const mockClothingItems: ClothingItem[] = [
  {
    id: '1',
    name: '简约白T',
    category: '上装',
    season: '夏',
    image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBpKRSkI0k_rv3-f3pkIHx8gSRFlXMKWkz5pMMTOl91whfAGmxZ7W_Wvp_51sCXqPS05Ghwswh09y79iYU2YbqzgVI4qFgBJ2Mczj3QLPKX7k1sxRZ3sMniVjXctOedrgdGbFGiCcgRUuiEp4TgBZ2UcKsMvm9ii-cWu6FEeSZ_3WqYfeCv_IZ8HWbm70HZy2d4gHO0jqVJ3ZoyD639pZGAXsiMRunOgs_pq5LK-DtbU-il8n_jKcAe8zuVzcacFavxvNxlkCyNSpRJ',
    tags: ['通勤', '百搭'],
    wearCount: 8,
  },
  {
    id: '2',
    name: '直筒牛仔裤',
    category: '下装',
    season: '春秋',
    image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuC1CcPAuYZDAzOoS09aTYninYVkA0F2bZ6MQS1XvIVykKnhdjXc3ILR5f1bddH6AKyZii9O-OWLbeYZdr5JF5c0jgotXh3R0Hduw2ZjFcjneACgQHUmJPTsB-w-bezs3EzNP9D6lDQIA3d5eCo1LWojxp7-mkYLjJme5Xrah723APDeA0vX2eD6n2vS4l3r4FQBpYvXVRPkp0ZzRO0f70ezn9BaACBMIfvjgDn-THgQ-sLvJjcAijZoDaqZ2xj2qpx16du5mQTaOANO',
    tags: ['休闲', '耐穿'],
    wearCount: 12,
  },
  {
    id: '3',
    name: '机能风冲锋衣',
    category: '外套',
    season: '春秋',
    image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuAkHfG8nBe7u5ky9SU6swqaWeW1-vlQh4NGnzLV5wcO1HGZ8UrMgwB1I8etuA48UQjqQTaC7FyDzurm8Yg_tjUXEXX0K9HqsQk4J16OZbuegSNbXf35h2CzqYrbzyReOOQQeD5llB_Kxkk_FzpBnnFtvBm6D9IfD9JV4vIepeh28KsFUbiffkw7PAya8M5RvMmByPuoK1cGq4zL6mSQqiZaJVIU67cpamB-8RDI3B3pv7kb0_HOsBXODmEfF0y4IEcft4h388KEs9Yd',
    tags: ['户外', '防风'],
    wearCount: 6,
  },
  {
    id: '4',
    name: '白色运动鞋',
    category: '鞋履',
    season: '全季',
    image: 'https://lh3.googleusercontent.com/aida-public/AB6AXuBDGHZOG1AVpehpXck7NDvsGZfMo9PMolJN4mPrh3Qw9f47KWQniwwoRyFEOOoIZ-MmJzuHCGjaW1eUC5ix2wdCz5t2taSiXbrfHTGnRrEBKZaX4zL3vYa5ttzh0MA_FOYAiW9kDO6s8SRWVfCwAlTjeHcu5SoAsMlbcOjgkizcEMVl4Q8PPe7wfF3rl0UPRmpa23iHIaKr8jPVC3OjxAjqodqcwh5slG7xQLzUgbvkYNLhHp5IAvuDlHHHdAySC0o0vN0sxYWXQ5gf',
    tags: ['百搭', '舒适'],
    wearCount: 22,
  },
];

const mockStats: WardrobeStats = {
  totalItems: 250,
  totalValue: '¥15,000',
  toClean: 25,
  weeklyNew: 5,
  unworn30Days: 50,
  avgPrice: '¥60',
};

export const useWardrobe = () => {
  const { showToast } = useToast();
  const [loading, setLoading] = useState(false);
  const [items, setItems] = useState<ClothingItem[]>([]);
  const [stats, setStats] = useState<WardrobeStats>(mockStats);
  const [filters, setFilters] = useState<WardrobeFilters>({
    category: '所有类别',
    season: '所有季节',
    searchTerm: '',
    sortBy: 'name',
    sortOrder: 'asc',
  });

  // 获取衣物列表
  const fetchItems = useCallback(async () => {
    setLoading(true);
    try {
      // 模拟API调用
      await new Promise(resolve => setTimeout(resolve, 500));
      setItems(mockClothingItems);
    } catch (error) {
      showToast('获取衣物列表失败', 'error');
    } finally {
      setLoading(false);
    }
  }, [showToast]);

  // 获取统计数据
  const fetchStats = useCallback(async () => {
    try {
      // 模拟API调用
      setStats(mockStats);
    } catch (error) {
      showToast('获取统计数据失败', 'error');
    }
  }, [showToast]);

  // 添加衣物
  const addItem = useCallback(async (item: Omit<ClothingItem, 'id' | 'wearCount'>) => {
    try {
      const newItem: ClothingItem = {
        ...item,
        id: Date.now().toString(),
        wearCount: 0,
      };

      setItems(prev => [...prev, newItem]);
      showToast('衣物添加成功', 'success');
      return newItem;
    } catch (error) {
      showToast('添加衣物失败', 'error');
      throw error;
    }
  }, [showToast]);

  // 更新衣物
  const updateItem = useCallback(async (id: string, updates: Partial<ClothingItem>) => {
    try {
      setItems(prev =>
        prev.map(item =>
          item.id === id ? { ...item, ...updates } : item
        )
      );
      showToast('衣物更新成功', 'success');
    } catch (error) {
      showToast('更新衣物失败', 'error');
      throw error;
    }
  }, [showToast]);

  // 删除衣物
  const deleteItem = useCallback(async (id: string) => {
    try {
      setItems(prev => prev.filter(item => item.id !== id));
      showToast('衣物删除成功', 'success');
    } catch (error) {
      showToast('删除衣物失败', 'error');
      throw error;
    }
  }, [showToast]);

  // 筛选衣物
  const filteredItems = items.filter(item => {
    const matchesCategory = filters.category === '所有类别' || item.category === filters.category;
    const matchesSeason = filters.season === '所有季节' || item.season === filters.season;
    const matchesSearch = item.name.toLowerCase().includes(filters.searchTerm.toLowerCase()) ||
                         item.tags.some(tag => tag.toLowerCase().includes(filters.searchTerm.toLowerCase()));

    return matchesCategory && matchesSeason && matchesSearch;
  });

  // 排序衣物
  const sortedItems = [...filteredItems].sort((a, b) => {
    let comparison = 0;

    switch (filters.sortBy) {
      case 'name':
        comparison = a.name.localeCompare(b.name);
        break;
      case 'category':
        comparison = a.category.localeCompare(b.category);
        break;
      case 'wearCount':
        comparison = a.wearCount - b.wearCount;
        break;
      case 'dateAdded':
        comparison = a.id.localeCompare(b.id); // 使用ID作为简单的日期排序替代
        break;
    }

    return filters.sortOrder === 'asc' ? comparison : -comparison;
  });

  // 更新筛选条件
  const updateFilters = useCallback((newFilters: Partial<WardrobeFilters>) => {
    setFilters(prev => ({ ...prev, ...newFilters }));
  }, []);

  // 重置筛选条件
  const resetFilters = useCallback(() => {
    setFilters({
      category: '所有类别',
      season: '所有季节',
      searchTerm: '',
      sortBy: 'name',
      sortOrder: 'asc',
    });
  }, []);

  // 初始化数据
  useEffect(() => {
    fetchItems();
    fetchStats();
  }, [fetchItems, fetchStats]);

  return {
    // 数据
    items: sortedItems,
    rawItems: items,
    stats,
    loading,
    filters,

    // 操作方法
    fetchItems,
    fetchStats,
    addItem,
    updateItem,
    deleteItem,
    updateFilters,
    resetFilters,

    // 计算属性
    filteredCount: filteredItems.length,
    hasItems: items.length > 0,
    isFiltered: filters.category !== '所有类别' ||
                filters.season !== '所有季节' ||
                filters.searchTerm !== '',
  };
};