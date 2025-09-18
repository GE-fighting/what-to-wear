"use client";
import React from "react";
import WardrobeSubNav from "./WardrobeSubNav";

interface WardrobeLayoutProps {
  children: React.ReactNode;
  title: string;
  subtitle?: string;
  actions?: React.ReactNode;
  showSubNav?: boolean;
}

export default function WardrobeLayout({
  children,
  title,
  subtitle,
  actions,
  showSubNav = true,
}: WardrobeLayoutProps) {
  return (
    <div className="wardrobe-layout">
      <header className="content-header">
        <div className="header-top">
          <div>
            <div className="breadcrumbs">What to Wear › 我的衣橱</div>
            <h1 className="page-title">{title}</h1>
            {subtitle && <p className="page-subtitle">{subtitle}</p>}
          </div>
          {actions && <div className="header-actions">{actions}</div>}
        </div>
      </header>

      <div className="content-body">
        {showSubNav && <WardrobeSubNav />}
        {children}
      </div>
    </div>
  );
}
