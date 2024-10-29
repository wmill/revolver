// import { getServerSession } from 'next-auth';
// import { redirect } from 'next/navigation';

export default async function ProtectedPage() {
  return (
    <div>
      <h1>Protected Content</h1>
      <p>You can only see this if you're logged in</p>
    </div>
  );
} 