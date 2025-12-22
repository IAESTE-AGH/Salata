import Link from "next/link";
import MobileMenu from "./MobileMenu";
import Image from "next/image";

function Navbar() {
  return (
    <nav className="w-full bg-white border-b border-gray-200 sticky top-0 z-50 h-16 overflow-x-hidden">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-full">
        <div className="flex justify-between items-center h-full">
          {/* Logo */}
          <div className="shrink-0 flex items-center">
            <Link href="/" className="text-xl font-bold text-gray-800">
              <span className="text-3xl bold text-(--primary)">SAŁATA</span>
            </Link>
          </div>

          <div className="hidden md:flex space-x-8 items-center">
            <Link
              href="/rejestracja"
              className="text-gray-600 hover:text-gray-900 transition-colors whitespace-nowrap"
            >
              Zarejestruj się
            </Link>
            <Link
              href="/login"
              className="bg-(--primary) text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors whitespace-nowrap"
            >
              Zaloguj się
            </Link>
          </div>

          <MobileMenu />
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
