import Header from "./Header";
import Footer from "./Footer";
import React, {useEffect, useState} from "react";
import {toast, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import axios from "axios";
import {NavigateFunction, useLocation, useNavigate} from "react-router-dom";

interface LayoutProps {
    children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({children}) => {
    const [isLogged, setIsLogged] = useState<boolean>(false);
    const navigate: NavigateFunction = useNavigate();
    const location = useLocation()

    useEffect(() => {
        if (location.pathname !== "/login" && location.pathname !== "/register") {
            axios
                .get("http://localhost:8000/api/users/", {
                    withCredentials: true,
                })
                .then((res) => {
                    setIsLogged(res.status === 200);
                })
                .catch((error) => {
                    if (error.response && error.response.status === 401) {
                        toast.warning("Unauthorized access. Please log in.");
                        setTimeout(() => {
                            navigate("/login");
                        }, 2000);
                    }
                });
        }
    }, [navigate, location]);

    return (
        <div className="layout">
            <Header isUserLoggedIn={isLogged}/>
            <ToastContainer/>
            {children}
            <Footer/>
        </div>
    );
};

export default Layout;
