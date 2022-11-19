import React from "react";
import "animate.css";

export default function SecondSection() {
  return (
    <section className="relative items-center overflow-hidden">
      <div className="container relative">
        <div className="grid content-center relative w-full md:mt-24 mt-16 text-center">
          <div className="wow animate__animated relative !w-full animate__fadeIn">
            <div className="flex items-center bg-gray-200 rounded-t-lg h-6">
              <div className="align-middle">
                <img className="h-4 px-2" src="assets/images/powershell.svg" />
              </div>
            </div>
            <div
              className="!p-2 relative rounded-b-lg overflow-clip"
              style={{
                borderLeft: "1px solid #e5e7eb",
                borderRight: "1px solid #e5e7eb",
                borderBottom: "1px solid #e5e7eb",
              }}
            >
              <video className="w-full" muted loop autoPlay playsInline>
                <source src="assets/images/corso_demo.mp4" type="video/mp4" />
              </video>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
