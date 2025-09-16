import { http } from "@/lib/api/http";
import type { LoginRequest, RegisterRequest } from "@/types/auth";

export function login(body: LoginRequest) {
  return http<string>("/api/auth/login", { method: "POST", body: JSON.stringify(body) });
}

export function register(body: RegisterRequest) {
  return http<{ message: string }>("/api/auth/register", { method: "POST", body: JSON.stringify(body) });
}
