import React from "react";
import "animate.css";
import { Icon } from "@iconify/react";


export default function Demo() {
  return (
    <section className="relative flex !tracking-wide  flex-col items-center overflow-hidden">
      <h3 className="mb-6 mt-8 md:text-4xl text-white text-3xl md:leading-normal leading-normal font-bold">
        What Corso users are saying          </h3>
      <div className="!container relative">

        <div className="grid md:grid-cols-2 grid-cols-1 items-center gap-[30px]">
          <div
            className="relative wow animate__animated animate__fadeInLeft"
            data-wow-delay=".3s"
          >
            <p className="text-slate-400 text-xl">
              “I liked the tool [a lot]. But once I connected with the team on Discord,
              I could see [that] this team really knows what they're doing.
              That made me a lot more confident.”
            </p>
            <p className="text-slate-500 italic indent-10">
              Kias Hanifa, CTO Fonicom

            </p>
            <div className="overflow-hidden absolute lg:h-[400px] h-[320px] lg:w-[400px] w-[320px] bg-indigo-600/5 bottom-0 left-0 rotate-45 -z-1 rounded-3xl"></div>
          </div>

          <div
            className="lg:ml-8 wow animate__animated animate__fadeInRight"
            data-wow-delay=".3s"
          >


            <ul className="list-none text-slate-400 mt-4">
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="ooui:speech-bubble-rtl"
                />{" "}
                “Documentation is great... initial steps for setup are really useful.”
              </li>
              <li className="mb-1">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="ooui:speech-bubble-rtl"
                />{" "}
                "Corso is a fantastic tool,
                especially the backend logic with Kopia that lets me run incremental backups"
              </li>
              <li className="mb-1 flex">
                <Icon
                  className="text-indigo-600 text-xl mr-2"
                  icon="ooui:speech-bubble-rtl"
                />{" "}
                “Yep, boring backups, righteous restores...”
              </li>
            </ul>


          </div>
        </div>
      </div>
      <div className="!container relative">
        <div className="flex flex-col content-center items-center justify-start relative md:mt-24 mt-16 text-center">
          <div className="wow w-[95%] sm:w-[80%] animate__animated relative  animate__fadeIn">
            <div className="flex flex-row items-center bg-gray-200 rounded-t-lg h-6">

              <div className="align-middle flex flex-col items-center justify-center">
                <img className="h-4 px-2" src="assets/images/powershell.svg" alt="Powershell logo" />
              </div>
            </div>
            <div
              className="!p-2 relative rounded-b-lg overflow-clip"
              style={{
                borderLeft: "2px solid #e5e7eb",
                borderRight: "2px solid #e5e7eb",
                borderBottom: "2px solid #e5e7eb",
                backgroundColor: "#121831",
              }}
            >
              <video className="w-full" poster="assets/images/corso_demo_thumbnail.png" muted loop autoPlay playsInline>
                <source src="assets/images/corso_demo.mp4" type="video/mp4" />
              </video>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
