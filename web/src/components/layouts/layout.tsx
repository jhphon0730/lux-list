import { useEffect } from "react"
import { Outlet, Navigate, useLocation } from "react-router-dom"

import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar"
import Sidebar from "@/components/layouts/sidebar"
import Navbar from "@/components/layouts/navbar"

import { ping } from "@/lib/api/auth";

const Layout = () => {
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
            <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  )
}

export default Layout