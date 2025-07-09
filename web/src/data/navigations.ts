import { Calendar, CheckSquare, Filter, Inbox, Settings, Star } from "lucide-react"

export const mainNavItems = [
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

export const projectItems = [
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

export const labelItems = [
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

export const utilityNavItems = [
  {
    title: "Filters & Labels",
    url: "/filters",
    icon: Filter,
  },
  {
    title: "Settings",
    url: "/settings",
    icon: Settings,
  },
]
