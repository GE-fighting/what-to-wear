"use client";
import { useEffect } from "react";

export function ToastProvider() {
  useEffect(() => {
    // Create toast container if it doesn't exist
    if (!document.getElementById("toast-container")) {
      const container = document.createElement("div");
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

  return null; // This component doesn't render anything visible
}
