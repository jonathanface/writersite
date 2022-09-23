import React from 'react';
import {useState, useEffect} from 'react';

const locations = ['news', 'books', 'about', 'contact'];
const validateLocation = (location) => {
  if (locations.includes(location)) {
    return true;
  }
  return false;
};

const formatListItems = (data, type) => {
  let response;
  switch (type) {
    case "news":
      response = data.map((item) => {
        return (
            <div className="sm:m-16 sm:max-w-[500px] news">
              <div className="flex justify-between">
                <h2 className="text-[2rem]">{item.title}</h2>
                <span className="text-[2rem]">{new Date(item.posted_on).toDateString()}</span>
              </div>
              <div className="p-8 pt-4 text-base font-roboto">{item.post}</div>
            </div>);
      });
      break;
    case "books":
      response = data.map((item) => {
        return (
            <div className="sm:m-16 book">
              <h2 className="text-[1.5rem]"><a href={item.link}>{item.title}</a></h2>
              <div>{new Date(item.released_on).toLocaleString('default', { month: 'long', year: 'numeric'})}</div>
              <div className="sm:max-w-[400px]"><a href={item.link}><img src={"img/covers/"+item.img}/></a></div>
              <div className="p-8 pl-0 pr-0 pt-4 text-base font-roboto sm:max-w-[400px] text-justify">{item.description}</div>
            </div>);
      });
      break;
  }
  return <div>{response}</div>
}

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
      throw new Error("something went wrong: " + response.body);
    }).then((data) => {
      if (Array.isArray(data)) {
        setContent(formatListItems(data, section.section))
      } else {
        setContent(data.body);
      }
    }).catch((error) => {
      console.error(error)
    });
  };

  useEffect(() => {
    fetchData();
  }, [section.section]);

  return <div className="overflow-auto h-screen grow-[999] basis-0 p-6">{content}</div>;
};

export {Body};
