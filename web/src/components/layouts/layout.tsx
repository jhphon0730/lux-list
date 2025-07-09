import { Outlet } from "react-router-dom"
import {
  Calendar,
  CheckSquare,
  Filter,
  Inbox,
  Plus,
  Search,
  Settings,
  Star,
  Tag,
  User,
  Bell,
  MoreHorizontal,
} from "lucide-react"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Badge } from "@/components/ui/badge"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInset,
  SidebarMenu,
  SidebarMenuAction,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
  SidebarRail,
  SidebarTrigger,
} from "@/components/ui/sidebar"
import { Separator } from "@/components/ui/separator"

const mainNavItems = [
  {
    title: "Inbox",
    url: "/inbox",
    icon: Inbox,
    badge: 12,
  },
  {
    title: "Today",
    url: "/today",
    icon: Calendar,
    badge: 5,
  },
  {
    title: "Upcoming",
    url: "/upcoming",
    icon: Star,
  },
  {
    title: "Completed",
    url: "/completed",
    icon: CheckSquare,
  },
]

const projectItems = [
  {
    title: "Personal",
    url: "/projects/personal",
    color: "bg-blue-500",
    count: 8,
  },
  {
    title: "Work",
    url: "/projects/work",
    color: "bg-green-500",
    count: 15,
  },
  {
    title: "Shopping",
    url: "/projects/shopping",
    color: "bg-purple-500",
    count: 3,
  },
]

const labelItems = [
  {
    title: "Priority",
    url: "/labels/priority",
    color: "bg-red-500",
  },
  {
    title: "Important",
    url: "/labels/important",
    color: "bg-orange-500",
  },
  {
    title: "Later",
    url: "/labels/later",
    color: "bg-gray-500",
  },
]

export default function Layout() {
  return (
    <SidebarProvider>
      <Sidebar className="border-r">
        <SidebarHeader className="border-b px-6 py-4">
          <div className="flex items-center gap-2">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-primary-foreground">
              <CheckSquare className="h-4 w-4" />
            </div>
            <div className="flex flex-col">
              <span className="text-sm font-semibold">TodoMaster</span>
              <span className="text-xs text-muted-foreground">Organize your life</span>
            </div>
          </div>
        </SidebarHeader>

        <SidebarContent>
          {/* Main Navigation */}
          <SidebarGroup>
            <SidebarGroupContent>
              <SidebarMenu>
                {mainNavItems.map((item) => (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton asChild>
                      <a href={item.url} className="flex items-center gap-3">
                        <item.icon className="h-4 w-4" />
                        <span>{item.title}</span>
                        {item.badge && (
                          <Badge variant="secondary" className="ml-auto h-5 px-1.5 text-xs">
                            {item.badge}
                          </Badge>
                        )}
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>

          <Separator className="my-2" />

          {/* Projects */}
          <SidebarGroup>
            <SidebarGroupLabel className="flex items-center justify-between">
              Projects
              <SidebarMenuAction>
                <Plus className="h-4 w-4" />
              </SidebarMenuAction>
            </SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {projectItems.map((project) => (
                  <SidebarMenuItem key={project.title}>
                    <SidebarMenuButton asChild>
                      <a href={project.url} className="flex items-center gap-3">
                        <div className={`h-3 w-3 rounded-full ${project.color}`} />
                        <span>{project.title}</span>
                        <span className="ml-auto text-xs text-muted-foreground">{project.count}</span>
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>

          <Separator className="my-2" />

          {/* Labels */}
          <SidebarGroup>
            <SidebarGroupLabel className="flex items-center justify-between">
              Labels
              <SidebarMenuAction>
                <Plus className="h-4 w-4" />
              </SidebarMenuAction>
            </SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {labelItems.map((label) => (
                  <SidebarMenuItem key={label.title}>
                    <SidebarMenuButton asChild>
                      <a href={label.url} className="flex items-center gap-3">
                        <Tag className="h-4 w-4" />
                        <span>{label.title}</span>
                        <div className={`ml-auto h-2 w-2 rounded-full ${label.color}`} />
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>

          <Separator className="my-2" />

          {/* Filters & Settings */}
          <SidebarGroup>
            <SidebarGroupContent>
              <SidebarMenu>
                <SidebarMenuItem>
                  <SidebarMenuButton asChild>
                    <a href="/filters" className="flex items-center gap-3">
                      <Filter className="h-4 w-4" />
                      <span>Filters & Labels</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
                <SidebarMenuItem>
                  <SidebarMenuButton asChild>
                    <a href="/settings" className="flex items-center gap-3">
                      <Settings className="h-4 w-4" />
                      <span>Settings</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </SidebarContent>

        <SidebarFooter className="border-t p-4">
          <SidebarMenu>
            <SidebarMenuItem>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <SidebarMenuButton className="w-full">
                    <Avatar className="h-6 w-6">
                      <AvatarImage src="/placeholder-user.jpg" />
                      <AvatarFallback>JD</AvatarFallback>
                    </Avatar>
                    <div className="flex flex-col items-start text-left">
                      <span className="text-sm font-medium">John Doe</span>
                      <span className="text-xs text-muted-foreground">john@example.com</span>
                    </div>
                    <MoreHorizontal className="ml-auto h-4 w-4" />
                  </SidebarMenuButton>
                </DropdownMenuTrigger>
                <DropdownMenuContent side="top" className="w-56">
                  <DropdownMenuItem>
                    <User className="mr-2 h-4 w-4" />
                    Profile
                  </DropdownMenuItem>
                  <DropdownMenuItem>
                    <Settings className="mr-2 h-4 w-4" />
                    Settings
                  </DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Log out</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </SidebarMenuItem>
          </SidebarMenu>
        </SidebarFooter>

        <SidebarRail />
      </Sidebar>

      <SidebarInset>
        {/* Navbar */}
        <header className="sticky top-0 z-10 flex h-16 shrink-0 items-center gap-4 border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 px-6">
          <SidebarTrigger className="-ml-1" />
          <Separator orientation="vertical" className="h-6" />

          {/* Search */}
          <div className="flex flex-1 items-center gap-4">
            <div className="relative max-w-md flex-1">
              <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
              <Input placeholder="Search tasks, projects..." className="pl-9 bg-muted/50" />
            </div>
          </div>

          {/* Right side actions */}
          <div className="flex items-center gap-2">
            <Button variant="ghost" size="icon">
              <Bell className="h-4 w-4" />
            </Button>

            <Button size="sm" className="gap-2">
              <Plus className="h-4 w-4" />
              Add Task
            </Button>

            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" size="icon">
                  <Avatar className="h-7 w-7">
                    <AvatarImage src="/placeholder-user.jpg" />
                    <AvatarFallback>JD</AvatarFallback>
                  </Avatar>
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent align="end" className="w-56">
                <DropdownMenuItem>
                  <User className="mr-2 h-4 w-4" />
                  Profile
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <Settings className="mr-2 h-4 w-4" />
                  Settings
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem>Log out</DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </header>

        {/* Main Content Area - Outlet */}
        <main className="flex-1 overflow-auto">
          <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  )
}
