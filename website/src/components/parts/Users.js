import React from "react";
import "animate.css";

export default function Hero() {
  return (
    <section className="relative !tracking-wide flex flex-col home-wrapper items-center overflow-hidden">
            <div class="container md:mt-24 mt-16">
                <div class="grid grid-cols-1 pb-8 text-center wow animate__animated animate__fadeInUp" data-wow-delay=".1s">
                    <h3 class="mb-4 md:text-3xl md:leading-normal text-2xl leading-normal font-semibold">What Our Users Say</h3>
                </div>

                <div class="grid grid-cols-1 mt-8">
                    <div class="tiny-three-item wow animate__animated animate__fadeInUp" data-wow-delay=".3s">
                        <div class="tiny-slide text-center">
                            <div class="customer-testi">
                                <div class="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                                    <i class="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                                    <p class="text-slate-400">" It seems that only fragments of the original text remain in the Lorem Ipsum texts used today. "</p>
                                    <ul class="list-none mb-0 text-amber-400 mt-3">
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                    </ul>
                                </div>
                                
                                <div class="text-center mt-5">
                                    <img src="assets/images/client/01.jpg" class="h-14 w-14 rounded-full shadow-md mx-auto" alt="" />
                                    <h6 class="mt-2 font-semibold">Calvin Carlo</h6>
                                    <span class="text-slate-400 text-sm">Manager</span>
                                </div>
                            </div>
                        </div>

                        <div class="tiny-slide text-center">
                            <div class="customer-testi">
                                <div class="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                                    <i class="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                                    <p class="text-slate-400">" The most well-known dummy text is the 'Lorem Ipsum', which is said to have originated in the 16th century. "</p>
                                    <ul class="list-none mb-0 text-amber-400 mt-3">
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                    </ul>
                                </div>
                                
                                <div class="text-center mt-5">
                                    <img src="assets/images/client/02.jpg" class="h-14 w-14 rounded-full shadow-md mx-auto" alt="" />
                                    <h6 class="mt-2 font-semibold">Christa Smith</h6>
                                    <span class="text-slate-400 text-sm">Manager</span>
                                </div>
                            </div>
                        </div>

                        <div class="tiny-slide text-center">
                            <div class="customer-testi">
                                <div class="content relative rounded shadow dark:shadow-gray-800 m-2 p-6 bg-white dark:bg-slate-900">
                                    <i class="mdi mdi-format-quote-open mdi-48px text-indigo-600"></i>
                                    <p class="text-slate-400">" According to most sources, Lorum Ipsum can be traced back to a text composed by Cicero. "</p>
                                    <ul class="list-none mb-0 text-amber-400 mt-3">
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                        <li class="inline"><i class="mdi mdi-star"></i></li>
                                    </ul>
                                </div>
                                
                                <div class="text-center mt-5">
                                    <img src="assets/images/client/06.jpg" class="h-14 w-14 rounded-full shadow-md mx-auto" alt="" />
                                    <h6 class="mt-2 font-semibold">Cristino Murfi</h6>
                                    <span class="text-slate-400 text-sm">Manager</span>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
    </section>
  );
}
