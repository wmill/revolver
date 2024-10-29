import { MainPanel } from "./MainPanel";
import { SideBar } from "./SideBar";
import { useSession } from "next-auth/react"

export function BaseLayout({ children }: { children: React.ReactNode }) {
  const { status } = useSession();
  if (status !== 'authenticated') {
    return (<div>{children}</div>)
  }

  return (
    <div className="w-screen min-h-screen">
      <SideBar />
      <MainPanel>{children}</MainPanel>
    </div>
  )
}