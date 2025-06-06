'use client';

import { useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { Menu as MenuIcon } from '@mui/icons-material';
import Sidebar from './Sidebar';

export default function Navbar() {
  const pathname = usePathname();
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  return (
    <>
      <nav className="bg-white shadow-lg">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <button
                onClick={() => setIsSidebarOpen(true)}
                className="p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100 focus:outline-none"
              >
                <MenuIcon />
              </button>
              <div className="ml-4 flex-shrink-0 flex items-center">
                <Link href="/" className="text-xl font-bold text-gray-800">
                  GoMicroBackend
                </Link>
              </div>
            </div>
            <div className="flex items-center">
              <Link
                href="/login"
                className={`px-4 py-2 rounded-md text-sm font-medium ${
                  pathname === '/login'
                    ? 'bg-gray-900 text-white'
                    : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Login
              </Link>
              <Link
                href="/signup"
                className={`ml-4 px-4 py-2 rounded-md text-sm font-medium ${
                  pathname === '/signup'
                    ? 'bg-gray-900 text-white'
                    : 'text-gray-700 hover:bg-gray-100'
                }`}
              >
                Sign Up
              </Link>
            </div>
          </div>
        </div>
      </nav>

      <Sidebar isOpen={isSidebarOpen} onClose={() => setIsSidebarOpen(false)} />
    </>
  );
} 