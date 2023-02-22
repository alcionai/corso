import React from "react";
import "animate.css";

export default function Hero() {
  return (
    <section className="relative !tracking-wide flex flex-col home-wrapper items-center overflow-hidden">
      <div
        className="bg-[#151C3D] absolute"
        style={{
          left: "-20rem",
          right: 0,
          zIndex: 1,
          top: "-30%",
          height: "62rem",
          width: "140rem",
          transform: "rotate(-12deg)",
        }}
      ></div>
      <div
        style={{
          zIndex: "1 !important",
        }}
        className="!container relative !z-10"
      >
        <div className="grid !z-10 grid-cols-1 mt-28 text-center">
          <div className="wow !z-10 animate__animated animate__fadeIn">
            <h4 className="font-bold !text-white !z-10 !leading-normal text-4xl lg:text-5xl mb-5">
              Free, Secure, and Open-Source
              <br /> Backup for Microsoft 365
            </h4>
            <p className="text-slate-300 !z-10 text-xl max-w-xl mx-auto">
              The #1 open-source backup tool for Microsoft 365
            </p>
          </div>

          <div className="mt-12 !z-10 mb-6 flex flex-col 2xs:flex-row items-center justify-center 2xs:space-y-0 space-y-4 2xs:space-x-4">
            <a
              href="../docs/quickstart/"
              className="text-2xl !z-10 !no-underline hover:text-white py-2 px-6 font-bold btn bg-indigo-800 hover:bg-indigo-900 border-indigo-800 hover:border-indigo-900 text-white rounded-md"
            >
              Quickstart
            </a>

          </div>

          <div
            className="home-dashboard mt-8 !z-10 wow animate__ animate__fadeIn animated"
            style={{
              visibility: "visible",
              animationName: "fadeIn",
            }}
          >

          </div>

        </div>

        <div className="bg-indigo-600 w-8 h-16 !z-10 absolute left-8 lg:bottom-28 md:bottom-36 sm:bottom-40 bottom-16"></div>
        <div className="bg-indigo-600/20 w-8 h-16 !z-10 absolute left-20 lg:bottom-32 md:bottom-40 sm:bottom-44 bottom-20"></div>

        <div className="bg-indigo-600/20 !z-10 w-8 h-16 absolute right-20 xl:bottom-[420px] lg:bottom-[315px] md:bottom-[285px] sm:bottom-80 bottom-32"></div>
        <div className="bg-indigo-600 w-8 h-16 !z-10 absolute right-8 xl:bottom-[440px] lg:bottom-[335px] md:bottom-[305px] sm:bottom-[340px] bottom-36"></div>
      </div>
      
    </section>
  );
}
