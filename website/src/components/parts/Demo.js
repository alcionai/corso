import React from "react";

export default function Demo() {
  return (
    <section className="relative flex !tracking-wide  flex-col items-center overflow-hidden">
      <div className="!container relative">
        <div className="flex flex-col content-center items-center justify-start relative md:mt-24 mt-16 text-center">
          <div className="wow w-[95%] sm:w-[80%] animate__animated relative  animate__fadeIn">
            <div className="flex flex-row items-center bg-gray-200 rounded-t-lg h-6">
              <div className="align-middle flex flex-col items-center justify-center">
                <img className="h-4 px-2" src="assets/images/powershell.svg" />
              </div>
            </div>
            <div
              className="!p-2 relative rounded-b-lg overflow-clip"
              style={{
                borderLeft: "2px solid #e5e7eb",
                borderRight: "2px solid #e5e7eb",
                borderBottom: "2px solid #e5e7eb",
              }}
            >
              <video
                className="w-full"
                poster="assets/images/corso_demo_thumbnail.png"
                muted
                loop
                autoPlay
                playsInline
              >
                <source src="assets/images/corso_demo.mp4" type="video/mp4" />
              </video>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
