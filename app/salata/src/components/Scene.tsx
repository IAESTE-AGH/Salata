"use client";

import React, { use, useRef } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/dist/ScrollTrigger";
import { useGSAP } from "@gsap/react";
import Image from "next/image";
import { text } from "stream/consumers";

if (typeof window !== "undefined") {
  gsap.registerPlugin(ScrollTrigger);
  gsap.registerPlugin(useGSAP);
}

export default function Section() {
  const mainRef = useRef<HTMLDivElement>(null);
  const imageRef = useRef<HTMLDivElement>(null);
  const placeholderRef = useRef<HTMLDivElement>(null);
  const aboutRef = useRef<HTMLDivElement>(null);
  const heroRef = useRef<HTMLDivElement>(null);
  const text1Ref = useRef<HTMLHeadingElement>(null);
  const text2Ref = useRef<HTMLHeadingElement>(null);
  const footerRef = useRef<HTMLDivElement>(null);

  useGSAP(
    () => {
      const mm = gsap.matchMedia();

      mm.add(
        {
          isDesktop: "(min-width: 768px)",
          isMobile: "(max-width: 767px)",
        },
        (context) => {
          const { isDesktop } = context.conditions!;

          const setupPosition = () => {
            const target = placeholderRef.current?.getBoundingClientRect();
            if (target && imageRef.current) {
              gsap.set(imageRef.current, {
                top: target.top,
                left: target.left,
                width: target.width,
                height: target.height,
                x: 0,
                y: 0,
                autoAlpha: 0,
                scale: 0.7,
                transformOrigin: "center center",
              });
            }
          };

          setupPosition();

          const introTl = gsap.timeline();

          introTl
            .to(imageRef.current, {
              autoAlpha: 1,
              scale: 1,
              duration: 1.2,
              ease: "back.out(1.7)",
            })
            .from(
              ".hero-text",
              {
                y: 40,
                opacity: 0,
                stagger: 0.15,
                duration: 1,
                ease: "power4.out",
              },
              "-=0.8"
            );

          const imageTL = gsap.timeline({
            scrollTrigger: {
              trigger: heroRef.current,
              start: "top 32",
              end: "700%",
              scrub: 1.5,
              markers: true,
            },
          });

          imageTL
            .to(
              imageRef.current,
              {
                scale: 0.8,
                rotate: 360,
                duration: 1,
                x: isDesktop ? "40vw" : 0,
                y: isDesktop ? 0 : "40vh",
              },
              "start"
            )
            .addLabel("left1")
            .to(
              imageRef.current,
              {
                x: isDesktop ? "0" : 0,
                y: isDesktop ? 0 : "20vh",
                rotate: "-=360",
                duration: 2,
              },
              "left1"
            )
            .to(
              text1Ref.current,
              {
                x: "-=40vw",
                opacity: 0,
                duration: 2,
                ease: "none",
              },
              "left1"
            )
            .from(
              text2Ref.current,
              {
                opacity: 0,
                x: "+=40vw",
                duration: 2,
              },
              "left1+=0.1"
            )
            .addLabel("right1")
            .to(
              imageRef.current,
              {
                x: isDesktop ? "60vw" : 0,
                y: isDesktop ? 0 : "40vh",
                rotate: "+=360",
                duration: 2,
              },
              "right1"
            )
            .to(
              text2Ref.current,
              {
                opacity: 0,
                x: "+=40vw",
                duration: 1,
              },
              "right1"
            )
            .to(imageRef.current, {
              x: isDesktop ? "0" : 0,
              y: isDesktop ? 0 : "20vh",
              rotate: "-=360",
              duration: 2,
            });

          const aboutTl = gsap.timeline({
            scrollTrigger: {
              trigger: aboutRef.current,
              start: "top 32",
              end: "600%",
              scrub: true,
              pin: true,
              markers: true,
            },
          });
        }
      );
    },
    { scope: mainRef }
  );

  return (
    <section ref={mainRef} className="bg-yellow min-h-screen overflow-hidden">
      <div
        ref={imageRef}
        className="fixed opacity-0 invisible z-50 pointer-events-none will-change-transform"
      >
        <Image
          src="/salata.png"
          alt="Animated Hero"
          width={1095}
          height={1041}
          className="object-cover w-full "
          priority
        />
      </div>

      <div
        ref={heroRef}
        className="relative min-h-screen -mt-30 md:-mt-20 w-full flex flex-col items-center justify-center md:flex-row md:justify-end md:pr-[10vw] px-4 sm:px-6 gap-8 md:gap-10"
      >
        <div className="hero-text text-center md:text-left z-20 w-full md:w-auto">
          <h1 className="text-6xl sm:text-6xl md:text-8xl text-(--foreground) font-bold leading-tight">
            Sprawdź co dziś <br /> na{" "}
            <span className="text-(--primary)">straganach</span>
          </h1>
          <h2 className="text-(--foreground) opacity-70 text-lg sm:text-xl md:text-4xl mt-4">
            Wszystkie dostępne stoiska w jednym miejscu
          </h2>
        </div>

        {/* PLACEHOLDER */}
        <div
          ref={placeholderRef}
          className="relative md:absolute md:top-1/2 md:-translate-y-1/2 md:left-[10vw] w-64 h-64 md:w-lg md:h-128 shadow-xl rounded-2xl border-4 border-(--primary) bg-(--background) "
        >
          {/* Empty  box */}
        </div>
      </div>

      <div
        className="h-screen bg-gray-100 flex items-center justify-center  w-full"
        ref={aboutRef}
      >
        <div className="max-w-7xl items-center  flex">
          <h2
            ref={text1Ref}
            className="relative  text-8xl italic font-extrabold  "
          >
            Poznaj moliwości <br />{" "}
            <span className="text-(--primary)">Sałaty</span>
          </h2>
          <h2
            ref={text2Ref}
            className="relative  text-8xl italic font-extrabold "
          >
            Przeglądaj aktywne{" "}
            <span className="text-(--primary)">stanowiska</span>
          </h2>
        </div>
      </div>
      <div ref={footerRef} className="h-screen bg-green-200 "></div>
    </section>
  );
}
