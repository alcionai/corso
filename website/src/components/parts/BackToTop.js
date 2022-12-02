import { Icon } from "@iconify/react";
import React, { useEffect } from "react";

export default function BackToTop() {
  function scroll() {
    window.scrollTo({ top: 0, left: 0, behavior: "smooth" });
  }
  function scrollFunction() {
    var mybutton = document.getElementById("back-to-top");
    if (mybutton != null) {
      if (
        document.body.scrollTop > 500 ||
        document.documentElement.scrollTop > 500
      ) {
        mybutton.classList.add("flex");
        mybutton.classList.remove("hidden");
      } else {
        mybutton.classList.add("hidden");
        mybutton.classList.remove("flex");
      }
    }
  }
  useEffect(() => {
    window.onscroll = function () {
      scrollFunction();
    };
  }, []);
  return (
    <a
      href="#"
      onClick={() => scroll()}
      id="back-to-top"
      className="back-to-top flex-col justify-center items-center fixed hidden text-lg rounded-full z-10 bottom-5 right-5 h-9 w-9 text-center bg-indigo-600 text-white leading-9"
    >
      <Icon icon="mdi:arrow-up" color="#fff" />
    </a>
  );
}
