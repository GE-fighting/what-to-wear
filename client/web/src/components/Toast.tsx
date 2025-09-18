"use client";
import React, { useEffect, useState } from "react";
import "@/styles/Toast.css";

interface ToastProps {
  message: string;
  type?: "success" | "error" | "warning" | "info";
  duration?: number;
  onClose?: () => void;
}

export function Toast({ message, type = "info", duration = 3000, onClose }: ToastProps) {
  const [isVisible, setIsVisible] = useState(true);

  useEffect(() => {
    const timer = setTimeout(() => {
      setIsVisible(false);
      setTimeout(() => {
        onClose?.();
      }, 300); // Wait for animation to complete
    }, duration);

    return () => clearTimeout(timer);
  }, [duration, onClose]);

  const getIcon = () => {
    switch (type) {
      case "success":
        return "✓";
      case "error":
        return "✕";
      case "warning":
        return "⚠";
      case "info":
      default:
        return "ℹ";
    }
  };

  return (
    <div className={`toast toast-${type} ${isVisible ? "toast-visible" : "toast-hidden"}`}>
      <div className="toast-icon">
        {getIcon()}
      </div>
      <div className="toast-message">
        {message}
      </div>
      <button 
        className="toast-close" 
        onClick={() => {
          setIsVisible(false);
          setTimeout(() => onClose?.(), 300);
        }}
      >
        ×
      </button>
    </div>
  );
}

// Toast container component for managing multiple toasts
export function ToastContainer() {
  return (
    <div className="toast-container" id="toast-container">
      {/* Toasts will be dynamically added here */}
    </div>
  );
}

// Utility function to show toast notifications
export function showToast(message: string, type: "success" | "error" | "warning" | "info" = "info", duration: number = 3000) {
  // Only run on client side
  if (typeof window === "undefined") return;
  
  // Toast styles will be automatically loaded via CSS imports
  
  // Create toast element
  const toastElement = document.createElement("div");
  toastElement.className = `toast toast-${type}`;
  
  // Create toast content
  const getIcon = () => {
    switch (type) {
      case "success":
        return "✓";
      case "error":
        return "✕";
      case "warning":
        return "⚠";
      case "info":
      default:
        return "ℹ";
    }
  };

  toastElement.innerHTML = `
    <div class="toast-icon">${getIcon()}</div>
    <div class="toast-message">${message}</div>
    <button class="toast-close">×</button>
  `;

  // Add close functionality
  const closeButton = toastElement.querySelector(".toast-close");
  const closeToast = () => {
    toastElement.classList.remove("toast-visible");
    toastElement.classList.add("toast-hidden");
    setTimeout(() => {
      if (toastElement.parentNode) {
        toastElement.parentNode.removeChild(toastElement);
      }
    }, 300);
  };

  closeButton?.addEventListener("click", closeToast);

  // Get or create container
  let container = document.getElementById("toast-container");
  if (!container) {
    container = document.createElement("div");
    container.id = "toast-container";
    container.className = "toast-container";
    document.body.appendChild(container);
  }

  // Add toast to container
  container.appendChild(toastElement);
  
  // Trigger animation
  setTimeout(() => {
    toastElement.classList.add("toast-visible");
  }, 10);

  // Auto-hide after duration
  setTimeout(closeToast, duration);
}
