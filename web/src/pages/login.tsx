import type React from "react"

import Swal from "sweetalert2"
import { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { CheckSquare, Loader2 } from "lucide-react"

import { login } from "@/lib/api/auth"

export default function LoginPage() {
  const navigate = useNavigate()

  const [name, setName] = useState("")
  const [error, setError] = useState("")
  const [isLoading, setIsLoading] = useState(false)

  // 로그인 페이지 접근 시에 세션 만료 여부 확인
  useEffect(() => {
    const queryParams = new URLSearchParams(location.search)
    const isExpired = queryParams.get("expired") === "true"
    if (isExpired) {
      Swal.fire({
        title: "세션 만료",
        text: "로그인 세션이 만료되었습니다. 다시 로그인해 주세요.",
        icon: "warning",
        confirmButtonText: "확인",
        confirmButtonColor: "#3182F6",
      })
    }
  }, [location.search])

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!name.trim()) {
      setError("Please enter your name")
      return
    }

    setIsLoading(true)
    setError("")

    try {
      const res = await login(name.trim())

      if (res.error) {
        setError(res.error)
      } else {
        // 로그인 성공 시 리디렉션
        const redirectPath = sessionStorage.getItem("redirectAfterLogin") || "/"
        sessionStorage.removeItem("redirectAfterLogin")
        navigate(redirectPath)
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-50 to-indigo-100 px-4">
      <Card className="w-full max-w-md shadow-lg">
        <CardHeader className="text-center space-y-4">
          <div className="flex justify-center">
            <div className="flex h-16 w-16 items-center justify-center rounded-2xl bg-primary text-primary-foreground shadow-lg">
              <CheckSquare className="h-8 w-8" />
            </div>
          </div>
          <div className="space-y-2">
            <CardTitle className="text-3xl font-bold">Welcome Back</CardTitle>
            <CardDescription className="text-base">Enter your name to access LuxList</CardDescription>
          </div>
        </CardHeader>

        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="name" className="text-sm font-medium">
                Your Name
              </Label>
              <Input
                id="name"
                type="text"
                placeholder="Enter your name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                disabled={isLoading}
                className="h-12 text-base"
                autoFocus
              />
            </div>

            {error && (
              <Alert variant="destructive">
                <AlertDescription>{error}</AlertDescription>
              </Alert>
            )}

            <Button type="submit" className="w-full h-12 text-base font-medium" disabled={isLoading || !name.trim()}>
              {isLoading ? (
                <>
                  <Loader2 className="mr-2 h-5 w-5 animate-spin" />
                  Signing in...
                </>
              ) : (
                "Sign In"
              )}
            </Button>
          </form>

          <div className="mt-6 text-center">
            <p className="text-sm text-muted-foreground">Simple and secure access to your tasks</p>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
