import React from "react";
import "../styles/NotFound.css";

const NotFound: React.FC = () => {
  return (
    <div className="not-found">
      <h1>404 - Page Not Found</h1>
      <p>The page you requested could not be found.</p>
    </div>
  );
};

export default NotFound;
