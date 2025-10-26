import { http } from "@/lib/api/http";
import type { LoginRequest, RegisterRequest } from "@/types/auth";

export function login(body: LoginRequest) {
  return http.post<string>("/api/auth/login", body);
}

export function register(body: RegisterRequest) {
  return http.post<{ message: string }>("/api/auth/register", body);
}
