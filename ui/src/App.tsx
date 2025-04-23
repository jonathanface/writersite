import { Sidebar } from './components/Sidebar';
import { Body } from './components/Body';
import { BrowserRouter } from 'react-router-dom';

export const App = () => {
  return (
    <BrowserRouter>
      <div className="root-child flex flex-wrap">
        <Sidebar />
        <Body />
      </div>
    </BrowserRouter>
  );
};
