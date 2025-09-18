"use client";
import React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

interface SubNavItem {
  href: string;
  label: string;
}

const subNavItems: SubNavItem[] = [
  { href: "/main/wardrobe", label: "概览" },
  { href: "/main/wardrobe/list", label: "列表" },
  { href: "/main/wardrobe/categories", label: "分类与标签" },
  { href: "/main/wardrobe/stats", label: "数据统计" },
  { href: "/main/wardrobe/import", label: "批量导入" },
];

export default function WardrobeSubNav() {
  const pathname = usePathname();

  return (
    <nav className="subnav">
      {subNavItems.map((item) => (
        <Link
          key={item.href}
          href={item.href}
          className={pathname === item.href ? "active" : ""}
        >
          {item.label}
        </Link>
      ))}
    </nav>
  );
}
