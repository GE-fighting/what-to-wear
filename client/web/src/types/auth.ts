export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  message?: string;
}

export interface RegisterRequest {
  username: string;
  password: string;
  email: string;
  nickname?: string;
  gender?: string;
  birth_date?: string;
  height?: number | null;
  weight?: number | null;
}

export interface ApiError {
  error: string;
}
