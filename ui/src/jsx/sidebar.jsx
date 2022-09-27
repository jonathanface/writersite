
import React from 'react';
import {useState} from 'react';

const NavItem = (props) => {
  return (
    <h2>
      <button onClick={(event) => {
        props.setActive(props.label);
        window.history.pushState({
          'html': props.label,
          'pageTitle': props.label,
        }, props.label, './' + props.label);
      }} className={props.activeSection === props.label ? 'sidebarActive' : ''}>{props.label.charAt(0).toUpperCase() + props.label.substring(1)}
      </button>
    </h2>
  );
};

const Sidebar = () => {
  const [activeSection, setActiveSection] = useState(window.location.pathname.length > 1 ? window.location.pathname.substring(1) : 'news');
  return (
    <ul className="bg-[#F1F1F1] sidebar h-screen grow md:basis-80 md:pt-6 border-r-4 md:border-r-8 border-r-black md:pl-6 pr-6 md:pr-0 text-[2.25rem] md:text-[3.75rem] leading-loose">
      <li key="news" className="sm:ml-8 ml-4"><NavItem label="news" activeSection={activeSection} setActive={setActiveSection}/></li>
      <li key="books" className="sm:ml-8 ml-4"><NavItem label="books" activeSection={activeSection} setActive={setActiveSection}/></li>
      <li key="about" className="sm:ml-8 ml-4"><NavItem label="about" activeSection={activeSection} setActive={setActiveSection}/></li>
      <li key="contact" className="sm:ml-8 ml-4"><NavItem label="contact" activeSection={activeSection} setActive={setActiveSection}/></li>
    </ul>
  );
};
export {Sidebar};
