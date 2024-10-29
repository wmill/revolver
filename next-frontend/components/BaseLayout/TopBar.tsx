import { Button } from '@mantine/core';
import { signOut } from 'next-auth/react';

export function TopBar({ togglePanel }: { togglePanel: (e: React.MouseEvent) => void }) {
  return (
    <div className="flex items-center justify-between w-full px-4 py-2 border-b border-gray-200 ">
      <div>
        <Button onClick={togglePanel}>(*)</Button>
      </div>
      <div className="flex-1 text-center">
        Top Bar
      </div>
      <div>
        <Button onClick={signOut}>Sign out</Button>
      </div>
    </div>
  );
}