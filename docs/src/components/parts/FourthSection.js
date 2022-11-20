import React from "react";
import "animate.css";
import { Icon } from "@iconify/react";

export default function FourthSection() {
  return (
    <section className="relative !tracking-wide  md:py-16 py-12 md:pt-0 pt-0">
      <div className="absolute bottom-0 left-0 !z-0 right-0 sm:h-2/3 h-4/5 bg-gradient-to-b from-indigo-500 to-indigo-600"></div>

      <div className="container !z-50">
        <div
          className="grid grid-cols-1 justify-center wow animate__animated animate__fadeInUp"
          data-wow-delay=".1s"
        >
          <div className="relative  flex flex-col items-center justify-center z-1">
            <div className="grid grid-cols-1 md:text-left text-center justify-center">
              <div className="relative">
                <img
                  src="assets/images/laptop-macbook.png"
                  className="mx-auto"
                  alt="Laptop image showing Microsoft 365 icons"
                />
              </div>
            </div>
            <div className="content md:mt-0">
              <div className="grid lg:grid-cols-12 grid-cols-1 md:text-left text-center justify-center">
                <div className="lg:col-start-2 lg:col-span-10">
                  <div className="grid md:grid-cols-2 grid-cols-1 items-center">
                    <div className="mt-8">
                      <div className="section-title text-md-start">
                        <h3 className="md:text-3xl text-2xl md:leading-normal leading-normal font-semibold text-white mt-2">
                          Start Protecting Your
                          <br /> Microsoft 365 Data!
                        </h3>
                        <h6 className="text-white/50 text-lg font-semibold">
                          Corso is Free and Open Source
                        </h6>
                      </div>
                    </div>

                    <div className="mt-8">
                      <div className="section-title text-md-start">
                        <p className="text-white/50 max-w-xl mx-auto mb-2">
                          Follow our quick-start guide to start protecting your
                          business-critical Microsoft 365 data in just a few
                          minutes.
                        </p>
                        <a
                          href="https://docs.corsobackup.io/quickstart"
                          className="!text-white !no-underline flex flex-row items-center !hover:text-white"
                        >
                          Get Started{" "}
                          <Icon
                            icon="uim:angle-right-b"
                            className="align-middle"
                          />
                        </a>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
