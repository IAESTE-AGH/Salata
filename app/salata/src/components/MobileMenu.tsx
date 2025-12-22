"use client";

import { useRef, useState } from "react";
import Link from "next/link";
import gsap from "gsap";
import { useGSAP } from "@gsap/react";

gsap.registerPlugin(useGSAP);

export default function MobileMenu() {
  const [isOpen, setIsOpen] = useState(false);

  const containerRef = useRef<HTMLDivElement>(null);
  const menuRef = useRef<HTMLDivElement>(null);
  const linksRef = useRef<(HTMLAnchorElement | null)[]>([]);

  const line1Ref = useRef<HTMLSpanElement>(null);
  const line2Ref = useRef<HTMLSpanElement>(null);
  const line3Ref = useRef<HTMLSpanElement>(null);

  const tl = useRef<gsap.core.Timeline | null>(null);

  useGSAP(
    () => {
      gsap.set(menuRef.current, {
        yPercent: -100,
        display: "none",
        visibility: "hidden",
        opacity: 0,
      });

      tl.current = gsap
        .timeline({ paused: true })

        // --- HAMBURGER ANIMATION ---
        .to(
          line1Ref.current,
          { y: 8, rotate: 45, duration: 0.2, ease: "power1.inOut" },
          0
        )
        .to(
          line2Ref.current,
          { opacity: 0, duration: 0.2, ease: "power1.inOut" },
          0
        )
        .to(
          line3Ref.current,
          { y: -8, rotate: -45, duration: 0.2, ease: "power1.inOut" },
          0
        )

        // --- MENU SLIDE  ---
        .set(
          menuRef.current,
          {
            display: "block",
            visibility: "visible",
          },
          0
        )

        .to(
          menuRef.current,
          {
            yPercent: 0,
            opacity: 1,
            duration: 0.5,
            ease: "power3.out",
          },
          0
        )

        // --- LINKS ANIMATION ---
        .from(
          linksRef.current,
          {
            y: 20,
            opacity: 0,
            stagger: 0.1,
            duration: 0.3,
            ease: "back.out(1.7)",
          },
          "-=0.2"
        );
    },
    { scope: containerRef }
  );

  const toggleMenu = () => {
    if (!tl.current) return;
    if (isOpen) {
      tl.current.reverse();
    } else {
      tl.current.play();
    }
    setIsOpen(!isOpen);
  };

  const addToRefs = (el: HTMLAnchorElement | null) => {
    if (el && !linksRef.current.includes(el)) {
      linksRef.current.push(el);
    }
  };

  return (
    <div ref={containerRef} className="md:hidden">
      {/* HAMBURGER BUTTON */}
      <button
        onClick={toggleMenu}
        className="relative z-50 p-2 text-gray-600 focus:outline-none"
        aria-label="Toggle menu"
      >
        <div className="w-6 h-6 flex flex-col justify-between items-center py-1">
          <span
            ref={line1Ref}
            className="w-6 h-0.5 bg-gray-600 rounded-full block"
          ></span>
          <span
            ref={line2Ref}
            className="w-6 h-0.5 bg-gray-600 rounded-full block"
          ></span>
          <span
            ref={line3Ref}
            className="w-6 h-0.5 bg-gray-600 rounded-full block"
          ></span>
        </div>
      </button>

      {/* MOBILE MENU */}
      <div
        ref={menuRef}
        className="fixed top-16 left-0 right-0 w-full bg-gray-50 border-b border-gray-200 shadow-xl origin-top invisible opacity-0"
        style={{ display: "none" }}
      >
        <div className="flex flex-col px-4 sm:px-6 py-8 space-y-4 max-w-7xl mx-auto">
          <Link
            ref={addToRefs}
            href="/rejestracja"
            onClick={toggleMenu}
            className="text-xl font-medium text-gray-800  transition-colors"
          >
            Zarejestruj się
          </Link>
          <Link
            ref={addToRefs}
            href="/login"
            onClick={toggleMenu}
            className="text-xl font-medium text-(--primary) transition-colors"
          >
            Zaloguj się
          </Link>
        </div>
      </div>
    </div>
  );
}
