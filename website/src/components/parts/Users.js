import React, { useEffect } from "react";
import "animate.css";
import BrowserOnly from "@docusaurus/BrowserOnly";

export default function Users() {
  useEffect(() => {
    if (typeof window !== "undefined") {
      const tns = require("tiny-slider").tns;

      if (document.getElementsByClassName("tiny-single-item").length > 0) {
        var slider = tns({
          container: ".tiny-single-item",
          items: 1,
          controls: false,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          speed: 400,
          gutter: 16,
        });
      }

      if (document.getElementsByClassName("tiny-one-item").length > 0) {
        var slider = tns({
          container: ".tiny-one-item",
          items: 1,
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
        });
      }

      if (document.getElementsByClassName("tiny-two-item").length > 0) {
        var slider = tns({
          container: ".tiny-two-item",
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
          responsive: {
            768: {
              items: 2,
            },
          },
        });
      }

      if (document.getElementsByClassName("tiny-three-item").length > 0) {
        var slider = tns({
          container: ".tiny-three-item",
          controls: false,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          speed: 400,
          gutter: 12,
          responsive: {
            992: {
              items: 3,
            },

            767: {
              items: 2,
            },

            320: {
              items: 1,
            },
          },
        });
      }

      if (document.getElementsByClassName("tiny-six-item").length > 0) {
        var slider = tns({
          container: ".tiny-six-item",
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
          responsive: {
            1025: {
              items: 6,
            },

            992: {
              items: 4,
            },

            767: {
              items: 3,
            },

            320: {
              items: 1,
            },
          },
        });
      }

      if (document.getElementsByClassName("tiny-twelve-item").length > 0) {
        var slider = tns({
          container: ".tiny-twelve-item",
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
          responsive: {
            1025: {
              items: 12,
            },

            992: {
              items: 8,
            },

            767: {
              items: 6,
            },

            320: {
              items: 2,
            },
          },
        });
      }

      if (document.getElementsByClassName("tiny-five-item").length > 0) {
        var slider = tns({
          container: ".tiny-five-item",
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
          responsive: {
            1025: {
              items: 5,
            },

            992: {
              items: 4,
            },

            767: {
              items: 3,
            },

            425: {
              items: 1,
            },
          },
        });
      }

      if (document.getElementsByClassName("tiny-home-slide-four").length > 0) {
        var slider = tns({
          container: ".tiny-home-slide-four",
          controls: true,
          mouseDrag: true,
          loop: true,
          rewind: true,
          autoplay: true,
          autoplayButtonOutput: false,
          autoplayTimeout: 3000,
          navPosition: "bottom",
          controlsText: [
            '<i class="mdi mdi-chevron-left "></i>',
            '<i class="mdi mdi-chevron-right"></i>',
          ],
          nav: false,
          speed: 400,
          gutter: 0,
          responsive: {
            1025: {
              items: 4,
            },

            992: {
              items: 3,
            },

            767: {
              items: 2,
            },

            320: {
              items: 1,
            },
          },
        });
      }
    }
  }, []);

  return (
    <section className="relative !tracking-wide flex flex-col items-center overflow-hidden">
      <div className="container md:mt-24 mt-16">
        <div className="grid grid-cols-1 mt-2">
          <h3 className="mb-6 mt-2 md:text-4xl text-white text-3xl md:leading-normal leading-normal font-bold text-center">
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
