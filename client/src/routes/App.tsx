import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import NotFound from "../component/notfound/NotFound";
import Layout from "../component/layout/Layout";
import Register from "../component/register/Register";
import Login from "../component/login/Login";

const App: React.FC = () => {
  return (
    <>
      <BrowserRouter>
        <Layout>
          <Routes>
            <Route path="/" element={<Register />} />
            <Route path="/login" element={<Login />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Layout>
      </BrowserRouter>
    </>
  );
};

export default App;
