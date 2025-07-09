import { Calendar, CheckSquare, Filter, Inbox, Star } from "lucide-react"

export const mainNavItems = [
  {
    title: "Inbox",
    url: "/inbox",
    icon: Inbox,
  },
  {
    title: "Today",
    url: "/today",
    icon: Calendar,
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
]
