import React, { useEffect } from "react";
import "animate.css";
import { tns } from "tiny-slider/src/tiny-slider";

export default function Users() {
  useEffect(function () {
    if (typeof window !== "undefined") {
      window.tns = tns;
    }
  }, []);

  return (
    <section className="relative !tracking-wide flex flex-col items-center overflow-hidden">
      <div className="container md:mt-24 mt-16">
        <div className="grid grid-cols-1 mt-8">
          <h3 className="text-center mb-4 md:text-3xl md:leading-normal text-2xl leading-normal font-semibold">
            What Corso Users Say
          </h3>
          <div
            className="tiny-three-item wow animate__animated animate__fadeInUp"
            data-wow-delay=".3s"
          >
            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    “Documentation is great... initial steps for setup are
                    really useful.”
                  </p>
                  <h6 className="mt-2 font-semibold">
                    Microsoft 365 Administrator
                  </h6>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    "I liked the tool a lot. But once I connected with the team
                    on Discord, I could see that this team really knows what
                    they're doing. That made me a lot more confident."
                  </p>
                  <h6 className="mt-2 font-semibold">
                    Kias Hanifa, CTO, Fonicom
                  </h6>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    "Corso is a fantastic tool, especially the backend logic
                    with Kopia that lets me run incremental backups"
                  </p>
                  <h6 className="mt-2 font-semibold">Backup Administrator</h6>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div></div>
      </div>
    </section>
  );
}
