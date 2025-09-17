import { http } from "@/lib/api/http";
import type { ClothingCategory, ClothingItemData, Tag } from "@/types/clothing";

export function createClothingItem(body: ClothingItemData) {
  return http<{ id: string }>("/api/clothing/item", { method: "POST", body: JSON.stringify(body) });
}

export function getClothingCategories() {
  return http<ClothingCategory[]>("/api/public/clothing/categories/tree", {method: "GET"});
}

export function getSystemTagEnums(type: string) {
  return http<Tag[]>(`/api/public/clothing/tags/enums/${type}`, { method: "GET" });
}