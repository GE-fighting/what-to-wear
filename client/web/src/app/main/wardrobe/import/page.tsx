"use client";
import React, { useState } from "react";
import WardrobeLayout from "@/components/WardrobeLayout";

export default function WardrobeImportPage() {
  const [dragOver, setDragOver] = useState(false);

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    setDragOver(true);
  };

  const handleDragLeave = (e: React.DragEvent) => {
    e.preventDefault();
    setDragOver(false);
  };

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    setDragOver(false);
    // 处理文件上传
    const files = Array.from(e.dataTransfer.files);
    console.log("上传的文件:", files);
  };

  const sampleData = [
    { name: "简约白T", category: "上装", season: "夏季", color: "白", brand: "Uniqlo", price: "79" },
    { name: "直筒牛仔裤", category: "下装", season: "春秋", color: "蓝", brand: "Levi's", price: "399" },
  ];

  return (
    <WardrobeLayout title="批量导入">
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: "20px", margin: "0 32px" }}>
        {/* 上传区域 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          <div style={{ fontWeight: 700, marginBottom: "8px" }}>上传图片/文件</div>
          <div
            style={{
              border: `2px dashed ${dragOver ? "var(--primary)" : "var(--line)"}`,
              borderRadius: "12px",
              minHeight: "200px",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              background: dragOver ? "rgba(102, 126, 234, 0.05)" : "#f8fafc",
              color: "var(--muted)",
              transition: "all 0.2s ease",
              cursor: "pointer",
            }}
            onDragOver={handleDragOver}
            onDragLeave={handleDragLeave}
            onDrop={handleDrop}
            onClick={() => {
              // 触发文件选择
              const input = document.createElement("input");
              input.type = "file";
              input.multiple = true;
              input.accept = "image/*,.csv";
              input.onchange = (e) => {
                const files = Array.from((e.target as HTMLInputElement).files || []);
                console.log("选择的文件:", files);
              };
              input.click();
            }}
          >
            {dragOver ? "释放文件到此处" : "拖拽图片或CSV到此处"}
          </div>
          <div style={{ display: "flex", alignItems: "center", gap: "10px", marginTop: "10px" }}>
            <button className="btn">选择文件</button>
            <button className="btn ghost">下载CSV模板</button>
          </div>
          <ul style={{ color: "var(--muted)", fontSize: "14px", marginTop: "10px", paddingLeft: "0", listStyle: "none" }}>
            <li style={{ marginBottom: "4px" }}>• 支持 JPG/PNG，单张 ≤ 5MB；CSV 使用 UTF-8 编码</li>
            <li>• 可从图片中提取颜色与品类作为初始属性</li>
          </ul>
        </div>

        {/* CSV 字段说明 */}
        <div className="wardrobe-card" style={{ padding: "16px" }}>
          <div style={{ fontWeight: 700, marginBottom: "8px" }}>CSV 字段说明</div>
          <div
            style={{
              width: "100%",
              border: "1px solid var(--line)",
              borderRadius: "10px",
              overflow: "auto",
            }}
          >
            <table style={{ width: "100%", borderCollapse: "collapse" }}>
              <thead>
                <tr>
                  {["name", "category", "season", "color", "brand", "price"].map((header) => (
                    <th
                      key={header}
                      style={{
                        borderBottom: "1px solid var(--line)",
                        padding: "10px 12px",
                        textAlign: "left",
                        fontSize: "14px",
                        color: "var(--muted)",
                        fontWeight: 700,
                        background: "#f8fafc",
                      }}
                    >
                      {header}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {sampleData.map((row, index) => (
                  <tr key={index}>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.name}
                    </td>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.category}
                    </td>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.season}
                    </td>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.color}
                    </td>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.brand}
                    </td>
                    <td style={{ borderBottom: "1px solid var(--line)", padding: "10px 12px", fontSize: "14px" }}>
                      {row.price}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {/* 导入预览 */}
      <div className="wardrobe-card" style={{ marginTop: "20px", margin: "20px 32px 0", padding: "16px" }}>
        <div style={{ fontWeight: 700, marginBottom: "8px" }}>导入预览</div>
        <div style={{ color: "var(--muted)", marginBottom: "12px" }}>
          选择文件后将在此处显示解析结果与冲突提示。
        </div>
        <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
          <button className="btn">开始导入</button>
          <button className="btn ghost">仅创建草稿</button>
        </div>
      </div>
    </WardrobeLayout>
  );
}
