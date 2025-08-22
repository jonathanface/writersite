import { Route, Routes } from 'react-router-dom';
import { BooksList } from './BooksList';
import { News } from './News';
import { About } from './About';
import { Contact } from './Contact';



export const Body = () => {

  return (
    <div className="overflow-auto h-screen grow-[999] basis-0 p-6">
      <Routes>
        <Route
          path="/news"
          element={
            <News />
          }
        />
        <Route
          path="/books"
          element={
            <BooksList />
          }
        />
        <Route
          path="/about"
          element={<About />}
        />
        <Route
          path="/contact"
          element={<Contact />}
        />
        <Route
          path="/"
          element={
            <News />
          }
        />
      </Routes>
    </div>
  )
};

