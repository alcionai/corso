import React from "react";
import "animate.css";
import loadable from "@loadable/component";
import Hero from "./Hero";
import Demo from "./Demo";
import CTA from "./CTA";
import Cookies from "./Cookies";
import KeyLoveFAQ from "./KeyLoveFAQ";
import UsersTestimonials from "./UsersTestimonials";

const BackToTopComp = loadable(() => import("./BackToTop"));

export function MainComp() {
  return (
    <>
      <Hero />
      <Demo />
      <KeyLoveFAQ />
      <CTA />
      <BackToTopComp />
      <UsersTestimonials />
      <Cookies />
    </>
  );
}
