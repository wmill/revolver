import { Button } from '@mantine/core';

export function TopBar({ togglePanel }: { togglePanel: (e: React.MouseEvent) => void }) {
  return (
    <div><Button onClick={togglePanel}>(*)</Button>Top Bar</div>
  );
}