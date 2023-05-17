import './globals.css';
import React from 'react';
import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'Mohammad Mahdi Afshar',
  description: 'Mohammad Mahdi Afshar, another developer.',
};

export default function RootLayout({ children }: { children: React.ReactNode; }) {
  return (
    <html>
    <head>
      <title>Mohammad Mahdi Afshar</title>
      <link rel='icon' href='/favicon.ico' />
    </head>
    <body>
    {children}
    </body>
    </html>
  );
}
