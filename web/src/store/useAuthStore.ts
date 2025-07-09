import { create } from "zustand"
import { persist } from "zustand/middleware"

import { type User } from "@/types/auth"

interface AuthStore {
  user: User | null
  setUser: (user: User | null) => void
  clearUser: () => void
}

export const useAuthStore = create<AuthStore>()(
  persist(
    (set) => ({
      user: null,
      setUser: (user) => set({ user }),
      clearUser: () => set({ user: null }),
    }),
    {
      name: "auth-storage", // unique name
    }
  )
)