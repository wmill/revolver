import { useState } from 'react';
import { TopBar } from './TopBar';

export function MainPanel({ children }: { children: React.ReactNode }) {
  const [panelOpen, setPanelOpen] = useState(false);
  const togglePanel = (e : React.MouseEvent) => {
    e.preventDefault();
    setPanelOpen(!panelOpen);
  };

  return (
    <div className={`w-screen min-h-screen absolute bg-[var(--mantine-color-body)] drop-shadow-md transition-transform duration-5000 ease-in-out ${panelOpen ? 'translate-x-64' : ''}`}>
      <TopBar togglePanel={togglePanel} />
      <div>
        {children}
      </div>
    </div>
  )
}
