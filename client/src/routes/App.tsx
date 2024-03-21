import React from "react";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import NotFound from "../component/NotFound";
import Layout from "../component/Layout";
import Register from "../component/Register";
import Login from "../component/Login";
import Home from "../component/Home";

const App: React.FC = () => {
  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/register" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  );
};

export default App;
