import { FetchWithOutAuth, type Response } from "@/lib/api";
import { type User } from "@/types/auth";

import { useAuthStore } from "@/store/useAuthStore";

// 사용자 로그인 요청
export const login = async (name: string): Promise<Response<{ user: User }>> => {
  const res = await FetchWithOutAuth("/auth/login", {
    method: "POST",
    body: JSON.stringify({ name }),
  });

  return {
    data: res.data,
    error: res.error,
  }
}

// 사용자 로그아웃 요청
export const logout = async (): Promise<Response<void>> => {
  useAuthStore.getState().clearUser();

  const res = await FetchWithOutAuth("/auth/logout", {
    method: "GET",
  });

  return {
    data: res.data,
    error: res.error,
  }
}