import React, { useEffect } from "react";
import { Icon } from "@iconify/react";

export default function Cookies() {
  function acceptCookies() {
    document.cookie = "cookies=accepted; expires=Fri, 31 Dec 9999 23:59:59 GMT";
    document.getElementById("cookies").style.display = "none";
  }

  return (
    <div
      id="cookies"
      className="cookie-popup !tracking-wide fixed max-w-lg bottom-3 right-3 left-3 sm:left-0 sm:right-0 mx-auto bg-white dark:bg-slate-900 shadow dark:shadow-gray-800 rounded-md pt-6 pb-2 px-6 z-50"
    >
      <p className="text-slate-400">
        This website uses cookies to provide you with a great user experience.
        By using it, you accept our{" "}
        <a
          href="cookies.html"
          target="_blank"
          className="text-emerald-600 dark:text-emerald-500 font-semibold"
        >
          use of cookies
        </a>
        .
      </p>
      <div className="cookie-popup-actions text-right">
        <button
          onClick={() => acceptCookies()}
          className="absolute border-none !bg-transparent p-0 cursor-pointer font-semibold top-2 right-2"
        >
          <Icon
            className="text-dark dark:text-slate-200 text-2xl"
            icon="humbleicons:times"
          />
        </button>
      </div>
    </div>
  );
}
