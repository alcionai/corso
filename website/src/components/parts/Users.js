import React, { useEffect } from "react";
import "animate.css";
import { tns } from "tiny-slider/src/tiny-slider";

export default function Users() {
  useEffect(function () {
    if (typeof window !== "undefined") {
      window.tns = tns;
    }
  });

  return (
    <section className="relative !tracking-wide flex flex-col items-center overflow-hidden">
      <div className="container md:mt-24 mt-16">
        <div className="grid grid-cols-1 mt-8">
        <h3 className="text-center mb-4 md:text-3xl md:leading-normal text-2xl leading-normal font-semibold">
            What Corso Users Say</h3>
          <div
            className="tiny-three-item wow animate__animated animate__fadeInUp"
            data-wow-delay=".3s"
          >
            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    "It seems that only fragments of the original text remain
                    in the Lorem Ipsum texts used today. "
                  </p>
                  <h6 className="mt-2 font-semibold">Calvin Carlo, Manager</h6>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    "The most well-known dummy text is the 'Lorem Ipsum', which
                    is said to have originated in the 16th century."
                  </p>
                  <h6 className="mt-2 font-semibold">Christa Smith, Manager</h6>
                </div>
              </div>
            </div>


            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    "According to most sources, Lorum Ipsum can be traced back
                    to a text composed by Cicero."
                  </p>
                  <h6 className="mt-2 font-semibold">Cristino Murfi, Manager</h6>
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
