import React from "react";
import "animate.css";
import loadable from "@loadable/component";
import Hero from "./Hero";
import Demo from "./Demo";
import FourthSection from "./FourthSection";
import Cookies from "./Cookies";
import KeyLoveFAQ from "./KeyLoveFAQ";

const BackToTopComp = loadable(() => import("./BackToTop"));

export function MainComp() {
  return (
    <>
      <Hero />
      <Demo />
      <KeyLoveFAQ />
      <FourthSection />
      <BackToTopComp />
      <Cookies />
    </>
  );
}
