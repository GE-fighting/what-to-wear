import { http } from "@/lib/api/http";
import type { Weather } from "@/types/weather";

export function getCurrentWeather() {
  return http<Weather>("/api/weather/current");
}
