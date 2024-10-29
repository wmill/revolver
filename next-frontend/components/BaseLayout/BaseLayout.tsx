import { MainPanel } from "./MainPanel";
import { SideBar } from "./SideBar";


export function BaseLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="w-screen min-h-screen">
      <SideBar />
      <MainPanel>{children}</MainPanel>
    </div>
  )
}