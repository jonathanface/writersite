import React from 'react';

const Body = (section) => {
  switch (section.section) {
    case 'news':
      return <div className="grow-[999] basis-0 p-6">news</div>;
      break;
    case 'books':
      return <div className="grow-[999] basis-0 p-6">books</div>;
      break;
    case 'about':
      return <div className="grow-[999] basis-0 p-6">about</div>;
      break;
    case 'contact':
      return <div className="grow-[999] basis-0 p-6">contact</div>;
      break;
    default:
      return <div className="grow-[999] basis-0 p-6">contents</div>;
  }
};

export {Body};
