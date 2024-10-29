import { TopBar } from './TopBar';

export function MainPanel({ children }: { children: React.ReactNode }) {
  return (
    <div className="w-screen min-h-screen absolute">
      Main Panel
      <TopBar />
      <div>
        {children}
      </div>
    </div>
  )
}