import { withAuth } from "next-auth/middleware"
import { NextResponse } from "next/server"
import type { NextRequestWithAuth } from "next-auth/middleware"

// This function can be marked `async` if using `await` inside
export default withAuth(
  // `withAuth` augments your `Request` with the user's token.
  function middleware(request: NextRequestWithAuth) {
    return NextResponse.next()
  },
  {
    callbacks: {
      authorized: ({ token }) => token != null,
    },
    pages: {
      signIn: "/login",
    },
    theme: {
      colorScheme: "dark",
    },
  }
)

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api/auth (auth endpoints)
     * - login (login page)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    "/((?!api/auth|login|_next/static|_next/image|favicon.ico).*)",
  ],
} 