import { Routes, Route } from "react-router-dom";

import NotFound from "./pages/NotFound";
import { MovieContextProvider } from "./context/MovieContext";
import ShowTvInfo from "./helpers/ShowTvInfo";

import Search from "./pages/Search";
import Home from "./pages/Home";
import Movies from "./pages/Movies";
import TvSeries from "./pages/TvSeries";
import ShowInfo from "./pages/ShowInfo";

export default function App() {
  return (
    <>
      <MovieContextProvider>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/movies" element={<Movies />} />
          <Route path="/tv-series" element={<TvSeries />} />
          <Route path="/search" element={<Search />} />
          <Route path="/:id" element={<ShowInfo />} />
          <Route path="/tv/:id" element={<ShowTvInfo />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </MovieContextProvider>
    </>
  );
}
