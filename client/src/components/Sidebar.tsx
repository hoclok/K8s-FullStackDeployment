'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';
import AddIcon from '@mui/icons-material/Add';
import RemoveIcon from '@mui/icons-material/Remove';

interface SidebarProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function Sidebar({ isOpen, onClose }: SidebarProps) {
  const pathname = usePathname();

  if (!isOpen) return null;

  return (
    <>
      {/* Overlay */}
      <div
        className="fixed inset-0 bg-black bg-opacity-50 z-40"
        onClick={onClose}
      />

      {/* Sidebar */}
      <div className="fixed left-0 top-0 h-full w-64 bg-white shadow-lg z-50 p-4">
        <div className="space-y-4">
          <Link
            href="/products/add"
            className={`flex items-center space-x-2 p-2 rounded-lg ${
              pathname === '/products/add'
                ? 'bg-blue-100 text-blue-600'
                : 'hover:bg-gray-100'
            }`}
          >
            <AddIcon />
            <span>Add Product</span>
          </Link>
          <Link
            href="/products/remove"
            className={`flex items-center space-x-2 p-2 rounded-lg ${
              pathname === '/products/remove'
                ? 'bg-blue-100 text-blue-600'
                : 'hover:bg-gray-100'
            }`}
          >
            <RemoveIcon />
            <span>Remove Product</span>
          </Link>
        </div>
      </div>
    </>
  );
} 