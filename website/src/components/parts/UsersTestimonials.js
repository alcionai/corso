import React from "react";
import "animate.css";
import TinySlider from "tiny-slider-react";
import "tiny-slider/dist/tiny-slider.css";
const settings = {
  loop: true,
  items: 1,
  autoplay: true,
  autoplayTimeout: 4000,
  autoplayHoverPause: false,
  autoplayButton: false,
  lazyload: true,
  autoplayText: [],
  nav: false,
  mouseDrag: true,
  controls: false,
  controlsContainer: document.querySelector(".slider-controls"),

  responsive: {
    600: {
      items: 3,
    },
  },
};

const data = [
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
  {
    text: " Thus, Lorem Ipsum has only limited suitability as a visual filler for German texts. ",
    avatar:
      "https://shreethemes.in/techwind/layouts/assets/images/client/02.jpg",
    name: "Christa Smith",
    role: "Manager",
  },
];

export default function UsersTestimonials() {
  const slideStyles = {
    background: "blue",
    color: "white",
    padding: "20px",
  };
  return (
    <section className="relative md:py-24 !tracking-wide py-16 overflow-hidden">
      <div className="container">
        <div
          className="grid grid-cols-1 pb-8 text-center wow animate__animated animate__fadeInUp"
          data-wow-delay=".1s">
          <h3 className="mb-6 mt-8 md:text-4xl text-white text-3xl md:leading-normal leading-normal font-bold">
            What Our Users Say
          </h3>

          <p className="text-slate-400 max-w-xl mx-auto">
            Start working with Tailwind CSS that can provide everything you need
            to generate awareness, drive traffic, connect
          </p>
        </div>
        <div className="slider-container">
          <TinySlider settings={settings}>
            {data.map((item, index) => (
              <div key={index} className="slide-item">
                <UserCard {...item} />
              </div>
            ))}
          </TinySlider>
          <div className="slider-controls"></div>
        </div>
      </div>
    </section>
  );
}

const UserCard = ({ avatar, role, name, text }) => {
  return (
    <div className="customer-testi">
      <div
        className="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900 
      before:content-[''] before:absolute before:-bottom-[10px] before:left-[calc(50%_-_10px)] before:bg-inherit before:h-5 before:w-5 before:transform before:rotate-45 before:shadow before:dark:shadow-gray-800
      ">
        <p className="text-slate-400">{text}</p>
      </div>
      <div class="text-center mt-5">
        <img
          src={avatar}
          class="h-14 w-14 rounded-full shadow-md mx-auto"
          alt=""
        />
        <h6 class="mt-2 font-semibold">{name}</h6>
        <span class="text-slate-400 text-sm">{role}</span>
      </div>
    </div>
  );
};
