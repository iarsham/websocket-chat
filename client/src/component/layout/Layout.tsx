import Header from "../header/Header";
import Footer from "../footer/Footer";
import React from "react";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <div className="layout">
      <Header />
      <ToastContainer />
      <main>{children}</main>
      <Footer />
    </div>
  );
};

export default Layout;
