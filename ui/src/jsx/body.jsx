import React from 'react';
import {useState, useEffect} from 'react';

const locations = ['news', 'books', 'about', 'contact'];
const validateLocation = (location) => {
  if (locations.includes(location)) {
    return true;
  }
  return false;
};

const Body = (section) => {
  if (!validateLocation(section.section)) {
    return <div className="grow-[999] basis-0 p-6"><b>NOT FOUND</b>></div>;
  }
  const [content, setContent] = useState('content');

  const fetchData = () => {
    fetch('/api/' + section.section).then((response) => {
      if (response.ok) {
        return response.json();
      }
      setContent("something went wrong");
      throw new Error("something went wrong");
    }).then((data) => {
      setContent(data.body);
    }).catch((error) => {
      console.error(error)
    });
  };

  useEffect(() => {
    fetchData();
  }, [section.section]);

  return <div className="grow-[999] basis-0 p-6">{content}</div>;
};

export {Body};
