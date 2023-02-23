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
    <section className="relative !tracking-wide flex flex-col home-wrapper items-center overflow-hidden">
      <div className="container md:mt-24 mt-16">
        <div
          className="grid grid-cols-1 pb-8 text-center wow animate__animated animate__fadeInUp"
          data-wow-delay=".1s"
        >
          <h3 className="mb-4 md:text-3xl md:leading-normal text-2xl leading-normal font-semibold">
            What Our Users Say
          </h3>

          <p className="text-slate-400 max-w-xl mx-auto">
            Start working with Tailwind CSS that can provide everything you need
            to generate awareness, drive traffic, connect.
          </p>
        </div>

        <div className="grid grid-cols-1 mt-8">
          <div
            className="tiny-three-item wow animate__animated animate__fadeInUp"
            data-wow-delay=".3s"
          >
            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " It seems that only fragments of the original text remain
                    in the Lorem Ipsum texts used today. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/01.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Calvin Carlo</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " The most well-known dummy text is the 'Lorem Ipsum', which
                    is said to have originated in the 16th century. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/02.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Christa Smith</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " One disadvantage of Lorum Ipsum is that in Latin certain
                    letters appear more frequently than others. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/03.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Jemina CLone</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " Thus, Lorem Ipsum has only limited suitability as a visual
                    filler for German texts. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/04.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Smith Vodka</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " There is now an abundance of readable dummy texts. These
                    are usually used when a text is required. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/05.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Cristino Murfi</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>

            <div className="tiny-slide text-center">
              <div className="customer-testi">
                <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                  <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                  <p className="text-slate-400">
                    " According to most sources, Lorum Ipsum can be traced back
                    to a text composed by Cicero. "
                  </p>
                  <ul className="list-none mb-0 text-amber-400 mt-3">
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                    <li className="inline">
                      <i className="mdi mdi-star"></i>
                    </li>
                  </ul>
                </div>

                <div className="text-center mt-5">
                  <img
                    src="assets/images/client/06.jpg"
                    className="h-14 w-14 rounded-full shadow-md mx-auto"
                    alt=""
                  />
                  <h6 className="mt-2 font-semibold">Cristino Murfi</h6>
                  <span className="text-slate-400 text-sm">Manager</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
