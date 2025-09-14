"use client";
import React from "react";
import "@/styles/sidebar-layout.css";

export default function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="app-layout">
      {children}
    </div>
  );
}
