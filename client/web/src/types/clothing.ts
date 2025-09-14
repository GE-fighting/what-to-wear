export interface ClothingItemData {
  name: string;
  categoryId: string;
  brand: string;
  color: string;
  size: string;
  sizeSystem: string;
  material: string;
  price: string;
  purchaseDate: string;
  notes: string;
  isFavorite: boolean;
  specificAttributes: {
    sleeveLength?: string;
    neckline?: string;
    fit?: string;
    length?: string;
    waistType?: string;
    closure?: string;
    heelHeight?: string;
    shoeType?: string;
    occasion?: string;
    accessoryType?: string;
    style?: string;
    season?: string;
  };
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

