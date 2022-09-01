/*!-----------------------------------------------------------------------------
 * easy_background
 * v2.0 - built 2017-10-30
 * Licensed under the MIT License.
 * http://www.testersite.it/github/easy_background/v3/
 * ----------------------------------------------------------------------------
 * Copyright (C) 2017 Eugenio Segala
 * --------------------------------------------------------------------------*/

function easy_background(selector, sld_args) {

  function empty_img(x) {
    if (x) {
      return "<img src='" + x + "'>";
    } else {
      return "";
    }
  }

  //use object same as arrays in php {nameofindex:variable} inside object you can use arrays [value1,val2] (variable in object can be as array
  //var sld_args={i:["img/555.jpg","img/44.jpg","img/33.jpg","img/22.jpg","img/11.jpg","img/1.jpg","img/2.jpg","img/3.jpg","img/4.jpg","img/5.jpg"],d:[3000,3000,3000,3000,3000] };

  //if delay is empty or forgotten then use this default value
  var def_del = 3000;

  var p = document.createElement("div");
  p.innerHTML = " ";
  p.classList.add("easy_slider");

  document.body.insertBefore(p, document.body.firstChild);
  //switch all values in object -- objectname.index in you case sld_args is object and i is index of array which keep images (i). We use this function for fill div with img tags
  //and for insert delays into empty or forgotten places in object
  sld_args.slide.forEach(function(v, i) {
    if (v) {
      document.querySelector(".easy_slider").innerHTML += empty_img(v);
      if (typeof sld_args.delay[i] == 'undefined' || typeof sld_args.delay[i] == '' || sld_args.delay[i] == 0) {
        sld_args.delay[i] = def_del;
      }
    }

  });

  //add various style on selector
  document.querySelector(".easy_slider").style.display = "none";


  //add various style on selector
  document.querySelector(selector).style.backgroundSize = "cover";
  document.querySelector(selector).style.backgroundRepeat = "no-repeat";
  document.querySelector(selector).style.backgroundPosition = "center center";


  setTimeout(function() {
    //add various style on selector
    if (typeof sld_args.transition_timing === 'undefined') {
      sld_args.transition_timing = "ease-in";
    }
    if (typeof sld_args.transition_duration === 'undefined') {
      sld_args.transition_duration = 500;
    }
    var transition = "all " + sld_args.transition_duration + 'ms ' + sld_args.transition_timing;
    document.querySelector(selector).style.WebkitTransition = transition;
    document.querySelector(selector).style.MozTransition = transition;
    document.querySelector(selector).style.MsTransition = transition;
    document.querySelector(selector).style.OTransition = transition;
    document.querySelector(selector).style.transition = transition;
  }, 100);


  //this n is number of row  in object - if first row one function if more than 1 then other
  var n = 1;
  //li collection previous delays from previous slides
  var li = 0;

  function slider() {
    //switching all images one by one
    sld_args.slide.forEach(function(vvv, iii) {
      //here go all slides except first
      if (n > 1) {
        //set delay from collected number from previous slides
        var delay = li;
        setTimeout(function() {

          document.querySelector(selector).style.backgroundImage = "url('" + vvv + "')";

        }, delay); // >1
        //collecting delays from curent
        li = li + sld_args.delay[iii];

      } else { //this function for only  first slide

        //next row
        n++;
        //collect delay first time
        li = sld_args.delay[iii];

        document.querySelector(selector).style.backgroundImage = "url('" + vvv + "')";

      }

    });

  };

  slider();

  setInterval(function() { // REPEAT

    slider();
    //here used length of array of delays in object instead you tot_time variable
  }, sld_args.delay.length);

}
