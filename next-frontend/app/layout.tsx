'use client'

import '@mantine/core/styles.css';
import './globals.css'

import React from 'react';
import { ColorSchemeScript, MantineProvider } from '@mantine/core';
import { SessionProvider } from 'next-auth/react';
import { theme } from '../theme';
import { BaseLayout } from '@/components/BaseLayout/BaseLayout';

// export const metadata = {
//   title: 'Mantine Next.js template',
//   description: 'I am using Mantine with Next.js!',
// };

export default function RootLayout({ children }: { children: any }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
      </head>
      <body>
        <SessionProvider>
          <MantineProvider theme={theme}>
            <BaseLayout>{children}</BaseLayout>
          </MantineProvider>
        </SessionProvider>
      </body>
    </html>
  );
}
