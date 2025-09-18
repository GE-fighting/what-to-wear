"use client";
import React, { useState } from "react";
import WardrobeLayout from "@/components/WardrobeLayout";

interface ChipInputProps {
  id: string;
  placeholder: string;
  onAdd: (value: string) => void;
}

function ChipInput({ id, placeholder, onAdd }: ChipInputProps) {
  const [value, setValue] = useState("");

  const handleAdd = () => {
    if (value.trim()) {
      onAdd(value.trim());
      setValue("");
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === "Enter") {
      handleAdd();
    }
  };

  return (
    <div style={{ display: "flex", gap: "8px", marginTop: "10px" }}>
      <input
        id={id}
        type="text"
        placeholder={placeholder}
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onKeyPress={handleKeyPress}
        style={{
          flex: 1,
          padding: "8px 10px",
          border: "1px solid var(--line)",
          borderRadius: "8px",
          fontSize: "14px",
        }}
      />
      <button
        className="btn"
        onClick={handleAdd}
        disabled={!value.trim()}
      >
        添加
      </button>
    </div>
  );
}

export default function WardrobeCategoriesPage() {
  const [categories, setCategories] = useState([
    { name: "👕 上装", count: 128 },
    { name: "👖 下装", count: 96 },
    { name: "🧥 外套", count: 54 },
    { name: "👟 鞋履", count: 42 },
  ]);

  const [tags, setTags] = useState({
    styles: ["通勤", "休闲", "户外"],
    scenes: ["会议", "约会", "旅行"],
    colors: ["黑", "白", "灰"],
    materials: ["棉", "羊毛", "尼龙"],
  });

  const addCategory = (name: string) => {
    setCategories([...categories, { name, count: 0 }]);
  };

  const addTag = (type: keyof typeof tags, value: string) => {
    setTags({
      ...tags,
      [type]: [...tags[type], value],
    });
  };

  return (
    <WardrobeLayout title="分类与标签">
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "20px", margin: "0 32px" }}>
        {/* 一级分类 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          <div className="filter-title">一级分类</div>
          <div style={{ display: "grid", gridTemplateColumns: "1fr", gap: "8px" }}>
            {categories.map((category, index) => (
              <div
                key={index}
                style={{
                  display: "flex",
                  alignItems: "center",
                  gap: "10px",
                  border: "1px solid var(--line)",
                  borderRadius: "12px",
                  padding: "10px 12px",
                  background: "#f8fafc",
                }}
              >
                <span style={{ flex: 1 }}>{category.name}</span>
                <span className="badge">{category.count}</span>
              </div>
            ))}
          </div>
          <ChipInput
            id="categories-input"
            placeholder="新增分类名称"
            onAdd={addCategory}
          />
        </div>

        {/* 标签体系 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          <div className="filter-title">标签体系（风格 / 场景 / 颜色）</div>
          <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "20px" }}>
            {/* 风格标签 */}
            <div>
              <div className="filter-title">风格</div>
              <div style={{ display: "flex", gap: "8px", flexWrap: "wrap" }}>
                {tags.styles.map((tag) => (
                  <span key={tag} className="badge">
                    {tag}
                  </span>
                ))}
              </div>
              <ChipInput
                id="styles-input"
                placeholder="新增风格"
                onAdd={(value) => addTag("styles", value)}
              />
            </div>

            {/* 场景标签 */}
            <div>
              <div className="filter-title">场景</div>
              <div style={{ display: "flex", gap: "8px", flexWrap: "wrap" }}>
                {tags.scenes.map((tag) => (
                  <span key={tag} className="badge">
                    {tag}
                  </span>
                ))}
              </div>
              <ChipInput
                id="scenes-input"
                placeholder="新增场景"
                onAdd={(value) => addTag("scenes", value)}
              />
            </div>

            {/* 颜色标签 */}
            <div>
              <div className="filter-title">颜色</div>
              <div style={{ display: "flex", gap: "8px", flexWrap: "wrap" }}>
                {tags.colors.map((tag) => (
                  <span key={tag} className="badge">
                    {tag}
                  </span>
                ))}
              </div>
              <ChipInput
                id="colors-input"
                placeholder="新增颜色"
                onAdd={(value) => addTag("colors", value)}
              />
            </div>

            {/* 材质标签 */}
            <div>
              <div className="filter-title">材质</div>
              <div style={{ display: "flex", gap: "8px", flexWrap: "wrap" }}>
                {tags.materials.map((tag) => (
                  <span key={tag} className="badge">
                    {tag}
                  </span>
                ))}
              </div>
              <ChipInput
                id="materials-input"
                placeholder="新增材质"
                onAdd={(value) => addTag("materials", value)}
              />
            </div>
          </div>
        </div>
      </div>
    </WardrobeLayout>
  );
}
