"use client";
import React from "react";
import WardrobeLayout from "@/components/WardrobeLayout";

export default function WardrobeStatsPage() {
  const kpiData = [
    { value: "483", label: "总件数" },
    { value: "23", label: "本月穿搭" },
    { value: "8", label: "本月新增" },
    { value: "¥126", label: "平均成本/次" },
  ];

  const topItems = [
    { name: "白色运动鞋", category: "鞋履", monthly: 8, total: 22, costPerWear: "¥18" },
    { name: "简约白T", category: "上装", monthly: 6, total: 18, costPerWear: "¥4" },
    { name: "直筒牛仔裤", category: "下装", monthly: 5, total: 12, costPerWear: "¥33" },
  ];

  return (
    <WardrobeLayout title="数据统计">
      <div className="grid" style={{ margin: "0 32px" }}>
        {/* KPI 指标 */}
        {kpiData.map((kpi, index) => (
          <div
            key={index}
            className="wardrobe-card"
            style={{
              gridColumn: "span 3",
              padding: "16px",
              textAlign: "center",
            }}
          >
            <div style={{ fontSize: "26px", fontWeight: 800, color: "var(--text)", marginBottom: "4px" }}>
              {kpi.value}
            </div>
            <div style={{ color: "var(--muted)", fontSize: "12px" }}>
              {kpi.label}
            </div>
          </div>
        ))}

        {/* 图表区域 */}
        <div
          className="wardrobe-card"
          style={{
            gridColumn: "span 8",
            padding: "16px",
          }}
        >
          <div style={{ fontWeight: 700, marginBottom: "8px" }}>近90天穿着趋势</div>
          <div
            style={{
              height: "220px",
              border: "1px dashed var(--line)",
              borderRadius: "12px",
              background: "#f8fafc",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              color: "var(--muted)",
            }}
          >
            折线图占位
          </div>
        </div>

        <div
          className="wardrobe-card"
          style={{
            gridColumn: "span 4",
            padding: "16px",
          }}
        >
          <div style={{ fontWeight: 700, marginBottom: "8px" }}>品类占比</div>
          <div
            style={{
              height: "220px",
              border: "1px dashed var(--line)",
              borderRadius: "12px",
              background: "#f8fafc",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              color: "var(--muted)",
            }}
          >
            饼图占位
          </div>
        </div>

        {/* 利用率排行榜 */}
        <div
          className="wardrobe-card"
          style={{
            gridColumn: "span 12",
            padding: "16px",
            overflow: "auto",
          }}
        >
          <div style={{ fontWeight: 700, marginBottom: "8px" }}>利用率 Top 10</div>
          <table style={{ width: "100%", borderCollapse: "collapse" }}>
            <thead>
              <tr>
                <th
                  style={{
                    borderBottom: "1px solid var(--line)",
                    padding: "10px 12px",
                    textAlign: "left",
                    fontSize: "14px",
                    color: "var(--muted)",
                    fontWeight: 700,
                  }}
                >
                  名称
                </th>
                <th
                  style={{
                    borderBottom: "1px solid var(--line)",
                    padding: "10px 12px",
                    textAlign: "left",
                    fontSize: "14px",
                    color: "var(--muted)",
                    fontWeight: 700,
                  }}
                >
                  分类
                </th>
                <th
                  style={{
                    borderBottom: "1px solid var(--line)",
                    padding: "10px 12px",
                    textAlign: "left",
                    fontSize: "14px",
                    color: "var(--muted)",
                    fontWeight: 700,
                  }}
                >
                  近30天
                </th>
                <th
                  style={{
                    borderBottom: "1px solid var(--line)",
                    padding: "10px 12px",
                    textAlign: "left",
                    fontSize: "14px",
                    color: "var(--muted)",
                    fontWeight: 700,
                  }}
                >
                  累计
                </th>
                <th
                  style={{
                    borderBottom: "1px solid var(--line)",
                    padding: "10px 12px",
                    textAlign: "left",
                    fontSize: "14px",
                    color: "var(--muted)",
                    fontWeight: 700,
                  }}
                >
                  成本/次
                </th>
              </tr>
            </thead>
            <tbody>
              {topItems.map((item, index) => (
                <tr key={index}>
                  <td
                    style={{
                      borderBottom: "1px solid var(--line)",
                      padding: "10px 12px",
                      fontSize: "14px",
                    }}
                  >
                    {item.name}
                  </td>
                  <td
                    style={{
                      borderBottom: "1px solid var(--line)",
                      padding: "10px 12px",
                      fontSize: "14px",
                    }}
                  >
                    {item.category}
                  </td>
                  <td
                    style={{
                      borderBottom: "1px solid var(--line)",
                      padding: "10px 12px",
                      fontSize: "14px",
                    }}
                  >
                    {item.monthly}
                  </td>
                  <td
                    style={{
                      borderBottom: "1px solid var(--line)",
                      padding: "10px 12px",
                      fontSize: "14px",
                    }}
                  >
                    {item.total}
                  </td>
                  <td
                    style={{
                      borderBottom: "1px solid var(--line)",
                      padding: "10px 12px",
                      fontSize: "14px",
                    }}
                  >
                    {item.costPerWear}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </WardrobeLayout>
  );
}
