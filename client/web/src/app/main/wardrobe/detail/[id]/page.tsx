"use client";
import React from "react";
import Link from "next/link";
import { useParams } from "next/navigation";
import WardrobeLayout from "@/components/WardrobeLayout";

export default function WardrobeDetailPage() {
  const params = useParams();
  const id = params.id;

  // 模拟数据 - 在实际项目中应该根据 id 从 API 获取
  const item = {
    id,
    name: "机能风防风外套",
    category: "外套",
    brand: "Arc'teryx",
    material: "尼龙混纺",
    price: "¥1999",
    purchaseDate: "2024/03/12",
    wearCount: 6,
    monthlyWearCount: 2,
    lastWorn: "3天前",
    costPerWear: "¥333",
    tags: ["户外", "春秋", "深灰"],
    careInstructions: ["30℃ 轻柔洗，悬挂阴干", "避免高温暴晒与熨烫"],
  };

  return (
    <WardrobeLayout
      title={item.name}
      showSubNav={true}
      actions={
        <div style={{ margin: "8px 0 12px" }}>
          <Link href="/main/wardrobe/list" className="btn ghost">
            ← 返回列表
          </Link>
        </div>
      }
    >
      <div style={{ display: "grid", gridTemplateColumns: "1.2fr 1fr", gap: "20px", margin: "0 32px" }}>
        {/* 图片展示区 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          <div style={{ display: "grid", gridTemplateColumns: "1fr", gap: "12px" }}>
            {/* 主图 */}
            <div
              style={{
                height: "340px",
                borderRadius: "12px",
                background: "linear-gradient(135deg, #f1f5f9, #e2e8f0)",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                fontSize: "42px",
                color: "#64748b",
              }}
            >
              🧥
            </div>
            {/* 缩略图 */}
            <div style={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", gap: "10px" }}>
              {[1, 2, 3, 4].map((index) => (
                <div
                  key={index}
                  style={{
                    height: "70px",
                    borderRadius: "10px",
                    background: "#f1f5f9",
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                    fontSize: "14px",
                    color: "#64748b",
                    cursor: "pointer",
                  }}
                >
                  {index}
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* 信息详情区 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          {/* 标题和标签 */}
          <div style={{ display: "flex", alignItems: "center", justifyContent: "space-between", marginBottom: "16px" }}>
            <div style={{ fontSize: "22px", fontWeight: 700 }}>{item.name}</div>
            <div style={{ display: "flex", gap: "8px" }}>
              {item.tags.map((tag) => (
                <span key={tag} className="badge">
                  {tag}
                </span>
              ))}
            </div>
          </div>

          {/* 基本信息 */}
          <div style={{ display: "grid", gridTemplateColumns: "repeat(2, 1fr)", gap: "12px", marginBottom: "16px" }}>
            {[
              { label: "分类", value: item.category },
              { label: "品牌", value: item.brand },
              { label: "材质", value: item.material },
              { label: "购入", value: `${item.price} · ${item.purchaseDate}` },
            ].map((info, index) => (
              <div
                key={index}
                style={{
                  border: "1px solid var(--line)",
                  borderRadius: "12px",
                  padding: "12px",
                }}
              >
                <div style={{ color: "var(--muted)", fontSize: "12px", marginBottom: "4px" }}>{info.label}</div>
                <div style={{ fontWeight: 600 }}>{info.value}</div>
              </div>
            ))}
          </div>

          {/* 穿着统计 */}
          <div style={{ marginBottom: "16px" }}>
            <h3 style={{ fontSize: "14px", color: "var(--muted)", fontWeight: 700, marginBottom: "8px" }}>
              穿着统计
            </h3>
            <ul style={{ listStyle: "none", padding: 0 }}>
              <li style={{ color: "var(--muted)", fontSize: "14px", marginBottom: "6px" }}>
                累计穿着 {item.wearCount} 次 · 本月 {item.monthlyWearCount} 次 · 最后一次 {item.lastWorn}
              </li>
              <li style={{ color: "var(--muted)", fontSize: "14px", marginBottom: "6px" }}>
                估算成本 / 次：{item.costPerWear}
              </li>
            </ul>
          </div>

          {/* 护理与备注 */}
          <div style={{ marginBottom: "16px" }}>
            <h3 style={{ fontSize: "14px", color: "var(--muted)", fontWeight: 700, marginBottom: "8px" }}>
              护理与备注
            </h3>
            <ul style={{ listStyle: "none", padding: 0 }}>
              {item.careInstructions.map((instruction, index) => (
                <li key={index} style={{ color: "var(--muted)", fontSize: "14px", marginBottom: "6px" }}>
                  {instruction}
                </li>
              ))}
            </ul>
          </div>

          {/* 操作按钮 */}
          <div style={{ display: "flex", gap: "10px" }}>
            <button className="btn">✏️ 编辑</button>
            <button className="btn ghost">📦 存档</button>
            <button className="btn ghost">🗑️ 删除</button>
          </div>
        </div>
      </div>
    </WardrobeLayout>
  );
}
