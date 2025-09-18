"use client";
import React, { useState } from "react";
import Link from "next/link";
import WardrobeLayout from "@/components/WardrobeLayout";

// 模拟数据
const mockItems = [
  {
    id: 1,
    name: "简约白T",
    category: "上装",
    season: "夏",
    icon: "👕",
    wearCount: 18,
    tags: ["通勤", "百搭"],
  },
  {
    id: 2,
    name: "直筒牛仔裤",
    category: "下装",
    season: "春秋",
    icon: "👖",
    wearCount: 12,
    tags: ["休闲", "耐穿"],
  },
  {
    id: 3,
    name: "机能风防风外套",
    category: "外套",
    season: "春秋",
    icon: "🧥",
    wearCount: 6,
    tags: ["户外", "防风"],
  },
  {
    id: 4,
    name: "白色运动鞋",
    category: "鞋履",
    season: "全季",
    icon: "👟",
    wearCount: 22,
    tags: ["百搭", "舒适"],
  },
];

export default function WardrobeListPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [sortBy, setSortBy] = useState("recent");

  const filteredItems = mockItems.filter((item) =>
    item.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
    item.category.toLowerCase().includes(searchTerm.toLowerCase()) ||
    item.tags.some((tag) => tag.toLowerCase().includes(searchTerm.toLowerCase()))
  );

  return (
    <WardrobeLayout
      title="衣物列表"
      actions={
        <div className="header-actions">
          <Link href="/main/wardrobe/add" className="btn">
            ➕ 添加衣物
          </Link>
        </div>
      }
    >
      {/* 搜索工具栏 */}
      <div className="toolbar">
        <div className="search">
          <span className="icon">🔎</span>
          <input
            type="text"
            placeholder="搜索名称、品牌、标签..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <select
          className="select"
          value={sortBy}
          onChange={(e) => setSortBy(e.target.value)}
        >
          <option value="recent">按最近添加</option>
          <option value="frequent">按最常穿</option>
          <option value="price">按价格</option>
        </select>
        <Link href="/main/wardrobe" className="btn secondary">
          概览
        </Link>
      </div>

      {/* 主布局 */}
      <div className="layout">
        {/* 侧边栏筛选 */}
        <aside className="wardrobe-card sidebar-content">
          <div className="filter-group">
            <div className="filter-title">季节</div>
            <div>
              {["春秋", "夏季", "冬季"].map((season) => (
                <span key={season} className="chip">
                  {season}
                </span>
              ))}
            </div>
          </div>

          <div className="filter-group">
            <div className="filter-title">分类</div>
            <div>
              {["上装", "下装", "外套", "鞋履"].map((category) => (
                <span key={category} className="chip">
                  {category}
                </span>
              ))}
            </div>
          </div>

          <div className="filter-group">
            <div className="filter-title">颜色</div>
            <div>
              {["白色", "黑色", "蓝色", "灰色"].map((color) => (
                <span key={color} className="chip">
                  {color}
                </span>
              ))}
            </div>
          </div>

          <div className="filter-group">
            <div className="filter-title">标签</div>
            <div>
              {["通勤", "运动", "休闲", "户外"].map((tag) => (
                <span key={tag} className="chip">
                  {tag}
                </span>
              ))}
            </div>
          </div>
        </aside>

        {/* 主内容区 */}
        <main className="wardrobe-card" style={{ padding: "0" }}>
          {/* 列表工具栏 */}
          <div
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              padding: "14px 16px",
              borderBottom: "1px solid var(--line)",
            }}
          >
            <div style={{ color: "var(--muted)", fontSize: "14px" }}>
              共 {filteredItems.length} 件
            </div>
            <div>
              <button className="btn ghost">⬜ 网格</button>
              <button className="btn ghost">≣ 列表</button>
            </div>
          </div>

          {/* 物品网格 */}
          <div className="items-grid">
            {filteredItems.map((item) => (
              <Link
                key={item.id}
                href={`/main/wardrobe/detail/${item.id}`}
                className="item-card"
              >
                <div className="item-thumb">{item.icon}</div>
                <div className="item-body">
                  <div className="item-name">{item.name}</div>
                  <div className="item-meta">
                    <span>
                      {item.category} · {item.season}
                    </span>
                    <span>穿着 {item.wearCount} 次</span>
                  </div>
                  <div className="item-tags">
                    {item.tags.map((tag) => (
                      <span key={tag} className="tag">
                        {tag}
                      </span>
                    ))}
                  </div>
                </div>
              </Link>
            ))}
          </div>
        </main>
      </div>
    </WardrobeLayout>
  );
}
