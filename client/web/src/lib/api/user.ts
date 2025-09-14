import { http } from "@/lib/api/http";
import type { UserProfile } from "@/types/user";

export function getUserProfile() {
  return http<UserProfile>("/api/user/profile");
}
