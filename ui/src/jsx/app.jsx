import React from 'react';
import ReactDOM from 'react-dom/client';
import {Sidebar} from './sidebar.jsx';
import {Body} from './body.jsx';
import {useState, useEffect} from 'react';

const locations = ['news', 'books', 'about', 'contact'];

const validateOrDefaultLocation = (location) => {
  if (locations.includes(location)) {
    return location;
  }
  window.location.pathname = '/' + locations[0];
  return locations[0];
};

const updatePage = (setLocation) => {
  const path = window.location.pathname;
  const section = path.substring(path.lastIndexOf('/') + 1);
  const location = validateOrDefaultLocation(section);
  setLocation(location);
  document.title = 'Author Jonathan Face' + ' - ' + location;
};

const App = () => {
  const path = window.location.pathname;
  const initialLocation = validateOrDefaultLocation(path.substring(path.lastIndexOf('/') + 1));
  document.title = 'Author Jonathan Face' + ' - ' + initialLocation;
  const [location, setLocation] = useState(initialLocation);

  useEffect(() => {
    window.addEventListener('locationchange', () => updatePage(setLocation));
    return () => window.removeEventListener('locationchange', () => updatePage(setLocation));
  });
  return (
    <div className="root-child flex flex-wrap">
      <Sidebar />
      <Body section={location}/>
    </div>);
};

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
);
