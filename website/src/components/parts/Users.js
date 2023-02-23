import React from "react";
import "animate.css";

export default function Users() {
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
        </div>

        <div className="grid grid-cols-1 mt-8">
          <div className="tns-outer" id="tns1-ow">
            <div
              className="tns-liveregion tns-visually-hidden"
              aria-live="polite"
              aria-atomic="true"
            >
              slide <span className="current">2 to 4</span> of 6
            </div>
            <div id="tns1-mw" className="tns-ovh">
              <div className="tns-inner" id="tns1-iw">
                <div
                  className="tiny-three-item wow animate__animated animate__fadeInUp"
                  data-wow-delay=".3s"
                >
                  <div className="tiny-slide text-center">
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          "I liked the tool [a lot]. But once I connected with
                          the team on Discord, I could see [that] this team
                          really knows what they're doing. That made me a lot
                          more confident."
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
                        <h6 className="mt-2 font-semibold">Kias Hanifa</h6>
                        <span className="text-slate-400 text-sm">
                          CTO Fonicom
                        </span>
                      </div>
                    </div>
                  </div>

                  <div className="tiny-slide text-center">
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          "Documentation is great... initial steps for setup are
                          really useful.""
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
                          "Corso is a fantastic tool, especially the backend
                          logic with Kopia that lets me run incremental backups"
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
          </div>
        </div>
      </div>
    </section>
  );
}
