import { BrowserRouter as Router, Routes, Route } from "react-router-dom"

import Layout from "@/components/layouts/layout"
import LoginPage from "@/pages/login"


// 예시 페이지 컴포넌트들
function InboxPage() {
  return (
    <div className="p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold">Inbox</h1>
        <p className="text-muted-foreground">All your tasks in one place</p>
      </div>
      <div className="space-y-4">
        <div className="rounded-lg border p-4">
          <h3 className="font-medium">Sample Task 1</h3>
          <p className="text-sm text-muted-foreground">This is a sample task description</p>
        </div>
        <div className="rounded-lg border p-4">
          <h3 className="font-medium">Sample Task 2</h3>
          <p className="text-sm text-muted-foreground">Another sample task</p>
        </div>
      </div>
    </div>
  )
}

function TodayPage() {
  return (
    <div className="p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold">Today</h1>
        <p className="text-muted-foreground">Focus on what matters today</p>
      </div>
      <div className="space-y-4">
        <div className="rounded-lg border p-4">
          <h3 className="font-medium">Important Task</h3>
          <p className="text-sm text-muted-foreground">Due today</p>
        </div>
      </div>
    </div>
  )
}

function UpcomingPage() {
  return (
    <div className="p-6">
      <div className="mb-6">
        <h1 className="text-2xl font-bold">Upcoming</h1>
        <p className="text-muted-foreground">Plan ahead with upcoming tasks</p>
      </div>
    </div>
  )
}

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        {/* 기본 레이아웃을 사용하는 라우트들 */}
        <Route path="/" element={<Layout />}>
          <Route index element={<InboxPage />} />
          <Route path="inbox" element={<InboxPage />} />
          <Route path="today" element={<TodayPage />} />
          <Route path="upcoming" element={<UpcomingPage />} />
          {/* 여기에 더 많은 라우트를 추가할 수 있습니다 */}
        </Route>
      </Routes>
    </Router>
  )
}
