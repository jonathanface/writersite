
import React from 'react';

const Sidebar = () => {
  return (
    <ul className="grow md:basis-80 md:pt-6 border-r-4 md:border-r-8 border-r-black md:pl-6 pr-6 md:pr-0 text-[2.25rem] md:text-[3.75rem] leading-loose">
      <li>
        <h2><a onClick={() => {
          window.history.pushState({'html': 'news', 'pageTitle': 'news'}, 'news', './news');
        }}>News</a></h2>
      </li>
      <li>
        <h2><a onClick={() => {
          window.history.pushState({'html': 'books', 'pageTitle': 'books'}, 'books', './books');
        }}>Books</a></h2>
      </li>
      <li>
        <h2><a onClick={() => {
          window.history.pushState({'html': 'about', 'pageTitle': 'about'}, 'about', './about');
        }}>About</a></h2>
      </li>
      <li>
        <h2><a onClick={() => {
          window.history.pushState({'html': 'contact', 'pageTitle': 'contact'}, 'contact', './contact');
        }}>Contact</a></h2>
      </li>
    </ul>
  );
};
export {Sidebar};
