export interface ClothingItemData {
  name: string;
  category_id: number;
  category_name: string;
  brand: string;
  color: string;
  size: string;
  material: string;
  season: string[];
  occasion: string[];
  style: string;
  description: string;
  tags: number[];
  tag_names: string[];
  status: ClothingStatus;
  is_favorite: boolean;
  purchase_info?: PurchaseInfo | null;
  specific_attributes: Record<string, any>;
}

export interface PurchaseInfo {
  price: number;
  store: string;
  purchase_date: string;
  notes: string;
}

export type ClothingStatus = 'active' | 'inactive' | 'donated' | 'sold' | 'lost' | 'damaged';

export interface Tag {
  id: number;
  name: string;
  type: string;
  color: string;
  description: string;
}

export interface ClothingCategory{
  id: number;
  name: string;
  description: string;
  parent_id?: number | null;
  parent_name?: string;
  icon?: string;
  sort_order: number;
  is_active: boolean;
  item_count?: number;
  created_at: string;
  updated_at: string;
  children?: ClothingCategory[];
}

