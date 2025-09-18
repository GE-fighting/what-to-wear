"use client";
import React from "react";
import Link from "next/link";
import WardrobeLayout from "@/components/WardrobeLayout";

export default function WardrobeOverviewPage() {
  return (
    <WardrobeLayout
      title="我的衣橱 · 概览"
      actions={
        <div className="header-actions">
          <Link href="/main/wardrobe/add" className="btn">
            ➕ 添加衣物
          </Link>
          <Link href="/main/wardrobe/list" className="btn secondary">
            📚 查看全部
          </Link>
        </div>
      }
    >
      <div className="grid">
        {/* 欢迎区域 */}
        <div style={{ gridColumn: "span 12" }}>
          <div className="grid" style={{ gridTemplateColumns: "1.3fr 1fr", gap: "20px" }}>
            <div className="wardrobe-card" style={{ padding: "0", overflow: "hidden" }}>
              <div
                style={{
                  background: "linear-gradient(135deg, #667eea 0%, #764ba2 100%)",
                  color: "white",
                  borderRadius: "14px",
                  padding: "24px",
                  display: "flex",
                  flexDirection: "column",
                  justifyContent: "space-between",
                  minHeight: "180px",
                }}
              >
                <div>
                  <h2 style={{ fontSize: "22px", marginBottom: "8px" }}>我的衣橱总览</h2>
                  <p style={{ color: "#eef2ff" }}>精心管理你的每一件衣物，让搭配变得更简单。</p>
                </div>
                <div
                  style={{
                    display: "grid",
                    gridTemplateColumns: "repeat(4, 1fr)",
                    gap: "12px",
                    marginTop: "14px",
                  }}
                >
                  <div
                    style={{
                      background: "rgba(255,255,255,0.15)",
                      border: "1px solid rgba(255,255,255,0.25)",
                      borderRadius: "12px",
                      padding: "12px",
                    }}
                  >
                    <div style={{ fontSize: "22px", fontWeight: 700 }}>483</div>
                    <div style={{ fontSize: "12px", opacity: 0.9 }}>件衣物</div>
                  </div>
                  <div
                    style={{
                      background: "rgba(255,255,255,0.15)",
                      border: "1px solid rgba(255,255,255,0.25)",
                      borderRadius: "12px",
                      padding: "12px",
                    }}
                  >
                    <div style={{ fontSize: "22px", fontWeight: 700 }}>127</div>
                    <div style={{ fontSize: "12px", opacity: 0.9 }}>常穿单品</div>
                  </div>
                  <div
                    style={{
                      background: "rgba(255,255,255,0.15)",
                      border: "1px solid rgba(255,255,255,0.25)",
                      borderRadius: "12px",
                      padding: "12px",
                    }}
                  >
                    <div style={{ fontSize: "22px", fontWeight: 700 }}>8</div>
                    <div style={{ fontSize: "12px", opacity: 0.9 }}>新增衣物</div>
                  </div>
                  <div
                    style={{
                      background: "rgba(255,255,255,0.15)",
                      border: "1px solid rgba(255,255,255,0.25)",
                      borderRadius: "12px",
                      padding: "12px",
                    }}
                  >
                    <div style={{ fontSize: "22px", fontWeight: 700 }}>94%</div>
                    <div style={{ fontSize: "12px", opacity: 0.9 }}>利用率</div>
                  </div>
                </div>
              </div>
            </div>

            <div className="wardrobe-card" style={{ padding: "20px" }}>
              <div className="section-title">快速筛选</div>
              <div
                style={{
                  display: "flex",
                  gap: "8px",
                  flexWrap: "wrap",
                }}
              >
                {["春秋", "夏季", "冬季", "上装", "下装", "鞋履", "中性色", "近7天未穿", "最常穿", "需要清洗", "待整理"].map(
                  (filter) => (
                    <div key={filter} className="chip">
                      {filter}
                    </div>
                  ),
                )}
              </div>
            </div>
          </div>
        </div>

        {/* 分类概览 */}
        <div className="wardrobe-card" style={{ gridColumn: "span 12", padding: "20px" }}>
          <div className="section-title">分类概览</div>
          <div
            style={{
              display: "grid",
              gridTemplateColumns: "repeat(6, 1fr)",
              gap: "12px",
            }}
          >
            {[
              { icon: "👕", name: "上装", count: 128 },
              { icon: "👖", name: "下装", count: 96 },
              { icon: "🧥", name: "外套", count: 54 },
              { icon: "👗", name: "裙装", count: 36 },
              { icon: "👟", name: "鞋履", count: 42 },
              { icon: "🎒", name: "配饰", count: 127 },
            ].map((category) => (
              <Link
                key={category.name}
                href={`/main/wardrobe/categories?filter=${category.name}`}
                style={{
                  display: "flex",
                  alignItems: "center",
                  gap: "10px",
                  padding: "14px",
                  border: "1px solid var(--line)",
                  borderRadius: "12px",
                  background: "#f8fafc",
                  textDecoration: "none",
                  color: "inherit",
                  transition: "all 0.2s ease",
                }}
                onMouseEnter={(e) => {
                  e.currentTarget.style.background = "#e2e8f0";
                }}
                onMouseLeave={(e) => {
                  e.currentTarget.style.background = "#f8fafc";
                }}
              >
                <div style={{ fontSize: "18px" }}>{category.icon}</div>
                <div style={{ fontWeight: 600 }}>{category.name}</div>
                <div style={{ color: "var(--muted)", fontSize: "12px", marginLeft: "auto" }}>
                  {category.count}
                </div>
              </Link>
            ))}
          </div>
        </div>

        {/* 双列布局 */}
        <div style={{ gridColumn: "span 12", display: "grid", gridTemplateColumns: "2fr 1fr", gap: "20px" }}>
          {/* 衣橱状态 */}
          <div className="wardrobe-card" style={{ padding: "20px" }}>
            <div className="section-title">衣橱状态</div>
            <div
              style={{
                display: "grid",
                gridTemplateColumns: "repeat(2, 1fr)",
                gap: "16px",
                marginBottom: "16px",
              }}
            >
              {[
                { value: "12", label: "需要清洗", color: "#ef4444" },
                { value: "5", label: "需要整理", color: "#f59e0b" },
                { value: "43", label: "30天未穿", color: "#64748b" },
                { value: "8", label: "本周新购", color: "#10b981" },
              ].map((status, index) => (
                <div
                  key={index}
                  style={{
                    padding: "12px",
                    background: "#f1f5f9",
                    borderRadius: "8px",
                    textAlign: "center",
                  }}
                >
                  <div style={{ fontSize: "20px", fontWeight: 600, color: status.color }}>
                    {status.value}
                  </div>
                  <div style={{ fontSize: "12px", color: "#64748b" }}>
                    {status.label}
                  </div>
                </div>
              ))}
            </div>
            <div
              style={{
                display: "grid",
                gridTemplateColumns: "repeat(4, 1fr)",
                gap: "12px",
              }}
            >
              {[
                { icon: "🧥", name: "机能风防风外套", meta: "外套 · 春秋 · 深灰" },
                { icon: "👟", name: "白色运动鞋", meta: "鞋履 · 全季节" },
                { icon: "👖", name: "直筒牛仔裤", meta: "下装 · 春秋" },
                { icon: "👕", name: "简约白T", meta: "上装 · 夏季" },
              ].map((item, index) => (
                <Link
                  key={index}
                  href="/main/wardrobe/detail/1"
                  className="item-card"
                  style={{ textDecoration: "none", color: "inherit" }}
                >
                  <div className="item-thumb">{item.icon}</div>
                  <div className="item-body">
                    <div className="item-name">{item.name}</div>
                    <div className="item-meta">
                      <span>{item.meta}</span>
                    </div>
                  </div>
                </Link>
              ))}
            </div>
          </div>

          {/* 衣橱管理 */}
          <div className="wardrobe-card" style={{ padding: "20px" }}>
            <div className="section-title">衣橱管理</div>
            <div style={{ display: "flex", flexDirection: "column", gap: "12px", marginBottom: "20px" }}>
              {[
                { icon: "🧹", text: "清理长期未穿衣物" },
                { icon: "📊", text: "生成穿搭报告" },
                { icon: "🏷️", text: "批量标签管理" },
              ].map((action, index) => (
                <button
                  key={index}
                  style={{
                    display: "flex",
                    alignItems: "center",
                    gap: "8px",
                    padding: "12px",
                    background: "#f8fafc",
                    border: "1px solid #e2e8f0",
                    borderRadius: "8px",
                    cursor: "pointer",
                    fontSize: "14px",
                    color: "#1e293b",
                    width: "100%",
                    textAlign: "left",
                  }}
                >
                  <span>{action.icon}</span>
                  <span>{action.text}</span>
                </button>
              ))}
            </div>
            <div className="section-title" style={{ fontSize: "14px", marginBottom: "8px" }}>管理建议</div>
            <ul style={{ listStyle: "none", padding: 0 }}>
              {[
                "12件衣物需要清洗，建议及时处理。",
                "43件衣物超过30天未穿，考虑整理或捐赠。",
                "春季衣物使用率较低，可适当调整摆放位置。",
              ].map((tip, index) => (
                <li key={index} style={{ color: "var(--muted)", fontSize: "14px", marginBottom: "8px" }}>
                  {tip}
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </WardrobeLayout>
  );
}
