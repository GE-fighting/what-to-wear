"use client";
import { createContext, useContext, useEffect, useState, useCallback } from "react";

// Toast 类型定义
export type ToastType = "success" | "error" | "warning" | "info";

interface ToastMessage {
  id: string;
  message: string;
  type: ToastType;
  duration: number;
}

interface ToastContextValue {
  showToast: (message: string, type?: ToastType, duration?: number) => void;
}

// 创建 Context
const ToastContext = createContext<ToastContextValue | undefined>(undefined);

// Hook to use toast
export function useToast() {
  const context = useContext(ToastContext);
  if (context === undefined) {
    throw new Error("useToast must be used within a ToastProvider");
  }
  return context;
}

// Toast Provider 组件
export function ToastProvider({ children }: { children: React.ReactNode }) {
  const [toasts, setToasts] = useState<ToastMessage[]>([]);

  // 创建或获取 toast 容器
  useEffect(() => {
    let container = document.getElementById("toast-container");
    if (!container) {
      container = document.createElement("div");
      container.id = "toast-container";
      container.className = "toast-container";
      document.body.appendChild(container);
    }

    return () => {
      // Cleanup on unmount
      const container = document.getElementById("toast-container");
      if (container) {
        container.remove();
      }
    };
  }, []);

  // 显示 toast 的函数
  const showToast = useCallback((message: string, type: ToastType = "info", duration: number = 3000) => {
    const id = Date.now().toString();
    const newToast: ToastMessage = { id, message, type, duration };

    setToasts(prev => [...prev, newToast]);

    // 自动移除 toast
    setTimeout(() => {
      setToasts(prev => prev.filter(toast => toast.id !== id));
    }, duration);
  }, []);

  // 渲染 toast 元素
  useEffect(() => {
    const container = document.getElementById("toast-container");
    if (!container) return;

    // 清空容器
    container.innerHTML = "";

    // 渲染所有 toast
    toasts.forEach(toast => {
      const toastElement = document.createElement("div");
      toastElement.className = `toast toast-${toast.type}`;
      toastElement.id = `toast-${toast.id}`;

      const iconMap = {
        success: "✓",
        error: "✕",
        warning: "⚠",
        info: "ℹ"
      };

      toastElement.innerHTML = `
        <div class="toast-icon">${iconMap[toast.type]}</div>
        <div class="toast-message">${toast.message}</div>
        <button class="toast-close">×</button>
      `;

      // 添加关闭功能
      const closeButton = toastElement.querySelector(".toast-close");
      const closeToast = () => {
        toastElement.classList.remove("toast-visible");
        toastElement.classList.add("toast-hidden");
        setTimeout(() => {
          if (toastElement.parentNode) {
            toastElement.parentNode.removeChild(toastElement);
          }
        }, 300);
        setToasts(prev => prev.filter(t => t.id !== toast.id));
      };

      closeButton?.addEventListener("click", closeToast);

      container.appendChild(toastElement);

      // 触发动画
      setTimeout(() => {
        toastElement.classList.add("toast-visible");
      }, 10);
    });
  }, [toasts]);

  return (
    <ToastContext.Provider value={{ showToast }}>
      {children}
    </ToastContext.Provider>
  );
}
