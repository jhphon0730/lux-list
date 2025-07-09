"use client"

import { Outlet } from "react-router-dom"
import { SidebarProvider, SidebarInset } from "@/components/ui/sidebar"
import Sidebar from "@/components/layouts/sidebar"
import Navbar from "@/components/layouts/navbar"

export default function Layout() {
  return (
    <SidebarProvider>
      <Sidebar />
      <SidebarInset>
        <Navbar />
        <main className="flex-1 overflow-auto">
          {/** Outlet component to render child routes */}
          <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  )
}
