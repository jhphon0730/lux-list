import { useEffect } from "react"
import { Outlet, Navigate, useLocation } from "react-router-dom"

import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar"
import Sidebar from "@/components/layouts/sidebar"
import Navbar from "@/components/layouts/navbar"

import { ping } from "@/lib/api/auth";
import { useAuthStore } from "@/store/useAuthStore"


type ProtectedRouteProps = {
  children: React.ReactNode
}

const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const { user } = useAuthStore()

  if (!user) {
    return <Navigate to="/login" />
  }

  return <>{children}</>
}

export default function Layout() {
  const location = useLocation()

  useEffect(() => {
    ping()
  }, [location.pathname])

  return (
    <SidebarProvider>
      <Sidebar />
      <SidebarInset>
        <Navbar />
        <main className="flex-1 overflow-auto">
          <ProtectedRoute>
            <Outlet />
          </ProtectedRoute>
        </main>
      </SidebarInset>
    </SidebarProvider>
  )
}
