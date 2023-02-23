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
              slide <span className="current">5 to 6</span> of 6
            </div>
            <div id="tns1-mw" className="tns-ovh">
              <div className="tns-inner" id="tns1-iw">
                <div
                  className="tiny-three-item wow animate__ animate__fadeInUp  tns-slider tns-carousel tns-subpixel tns-calc tns-horizontal animated"
                  data-wow-delay=".3s"
                  id="tns1"
                  style={{
                    transform: "translate3d(-66.6667%, 0px, 0px)",
                    visibility: "visible",
                    animationDelay: "0.3s",
                    animationName: "fadeInUp",
                  }}
                >
                  <div
                    className="tiny-slide text-center tns-item"
                    id="tns1-item0"
                    aria-hidden="true"
                    tabIndex="-1"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " It seems that only fragments of the original text
                          remain in the Lorem Ipsum texts used today. "
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

                  <div
                    className="tiny-slide text-center tns-item"
                    id="tns1-item1"
                    aria-hidden="true"
                    tabIndex="-1"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " The most well-known dummy text is the 'Lorem Ipsum',
                          which is said to have originated in the 16th century.
                          "
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

                  <div
                    className="tiny-slide text-center tns-item"
                    id="tns1-item2"
                    aria-hidden="true"
                    tabIndex="-1"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " One disadvantage of Lorum Ipsum is that in Latin
                          certain letters appear more frequently than others. "
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

                  <div
                    className="tiny-slide text-center tns-item"
                    id="tns1-item3"
                    aria-hidden="true"
                    tabIndex="-1"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " Thus, Lorem Ipsum has only limited suitability as a
                          visual filler for German texts. "
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

                  <div
                    className="tiny-slide text-center tns-item tns-slide-active"
                    id="tns1-item4"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " There is now an abundance of readable dummy texts.
                          These are usually used when a text is required. "
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

                  <div
                    className="tiny-slide text-center tns-item tns-slide-active"
                    id="tns1-item5"
                  >
                    <div className="customer-testi">
                      <div className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                        <i className="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                        <p className="text-slate-400">
                          " According to most sources, Lorum Ipsum can be traced
                          back to a text composed by Cicero. "
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
            <div className="tns-nav" aria-label="Carousel Pagination">
              <button
                type="button"
                data-nav="0"
                aria-controls="tns1"
                aria-label="Carousel Page 1"
                className=""
                tabIndex="-1"
              ></button>
              <button
                type="button"
                data-nav="1"
                aria-controls="tns1"
                aria-label="Carousel Page 2"
                className=""
                tabIndex="-1"
              ></button>
              <button
                type="button"
                data-nav="2"
                aria-controls="tns1"
                aria-label="Carousel Page 3 (Current Slide)"
                className="tns-nav-active"
              ></button>
              <button
                type="button"
                data-nav="3"
                tabIndex="-1"
                aria-controls="tns1"
                style={{ display: "none" }}
                aria-label="Carousel Page 4"
              ></button>
              <button
                type="button"
                data-nav="4"
                tabIndex="-1"
                aria-controls="tns1"
                style={{ display: "none" }}
                aria-label="Carousel Page 5"
              ></button>
              <button
                type="button"
                data-nav="5"
                tabIndex="-1"
                aria-controls="tns1"
                style={{ display: "none" }}
                aria-label="Carousel Page 6"
              ></button>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
