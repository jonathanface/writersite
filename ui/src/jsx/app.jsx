import React from 'react';
import ReactDOM from 'react-dom/client';
import {Sidebar} from './sidebar.jsx';
import {Body} from './body.jsx';

const App = () => {
  return (
    <div className="root-child flex flex-wrap">
      <Sidebar />
      <Body />
    </div>);
};

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
);
