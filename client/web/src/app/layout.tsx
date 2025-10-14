import type { Metadata } from "next";
import { Noto_Sans_SC } from "next/font/google";
import "./globals.css";
import { ToastProvider } from "@/components/ToastProvider";

const notoSansSC = Noto_Sans_SC({
  subsets: ["latin"],
  weight: ["400", "500", "700"],
  variable: "--font-noto-sans-sc",
});

export const metadata: Metadata = {
  title: "StyleSense | 智能服饰管理平台",
  description: "基于黑白极简主义设计的智能服饰管理平台",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="zh-CN">
      <head>
        <link
          href="https://fonts.googleapis.com/icon?family=Material+Icons"
          rel="stylesheet"
        />
        <link
          href="https://fonts.googleapis.com/icon?family=Material+Icons+Outlined"
          rel="stylesheet"
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;700;900&display=swap"
          rel="stylesheet"
        />
      </head>
      <body
        className={`${notoSansSC.variable} font-display antialiased bg-background-light text-text-light-primary dark:bg-background-dark dark:text-dark-primary`}
      >
        <ToastProvider>
          {children}
        </ToastProvider>
      </body>
    </html>
  );
}
