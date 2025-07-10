import { FetchWithOutAuth, type Response } from "@/lib/api";
import { type User } from "@/types/auth";

import { useAuthStore } from "@/store/useAuthStore";

// 사용자 로그인 요청
export const login = async (name: string): Promise<Response<{ user: User }>> => {
  const res = await FetchWithOutAuth("auth/login", {
    method: "POST",
    body: JSON.stringify({ name }),
  });

  // 로그인 성공 시 사용자 정보를 상태에 저장
  if (res && res.user) {
    useAuthStore.getState().setUser(res.user);
  }

  return {
    data: res,
    error: res.error,
  }
}

// 사용자 로그아웃 요청
export const logout = async (): Promise<Response<void>> => {
  const res = await FetchWithOutAuth("auth/logout", {
    method: "GET",
  });

  // 로그아웃 후 상태 초기화
  useAuthStore.getState().clearUser();

  return {
    data: res.data,
    error: res.error,
  }
}

// 사용자 정보 요청 (프로필)
